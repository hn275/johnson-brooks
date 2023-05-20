package auth

import (
	"encoding/json"
	"errors"
	"jb/lib"
	"net/http"

	"github.com/teris-io/shortid"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var cred Credentials
	if err := json.NewDecoder(r.Body).Decode(&cred); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		lib.StdErr(err)
		return
	}
	defer r.Body.Close()

	db := newDb()
	defer db.Close()

	var user Credentials
	var err error

	trx := db.findByUsername(&cred)
	err = trx.Err()
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			w.WriteHeader(http.StatusNotFound)
			lib.NewErr("Account not found").HandleErr(w)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		lib.StdErr(err)
		return
	}

	err = trx.Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		lib.StdErr(err)
		return
	}

	err = authenticateUser(&cred, &user, db)
	if err != nil {
		if errors.Is(err, ErrAuthFailed) {
			w.WriteHeader(http.StatusUnauthorized)
			lib.NewErr("Authentication failed").HandleErr(w)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			lib.StdErr(err)
		}
		return
	}

	c := http.Cookie{
		Name:     "SessionID",
		Value:    user.Session.SessionID,
		HttpOnly: true,
		MaxAge:   15 * 60, // NOTE: 15 minutes
	}
	http.SetCookie(w, &c)
	w.WriteHeader(http.StatusOK)
}

func authenticateUser(cred, user *Credentials, db *authDatabase) error {
	if user.Session.FailedAttempts >= 4 {
		if err := db.lockUser(user); err != nil {
			return err
		}
		return ErrAuthFailed
	}

	plain := []byte(cred.Password)
	hashed := []byte(user.Password)
	err := bcrypt.CompareHashAndPassword(hashed, plain)
	if err != nil {
		if err := db.authFailed(user.ID); err != nil {
			return err
		}
		return ErrAuthFailed
	}

	sessionID, err := getSessionID()
	if err != nil {
		return err
	}

	if err := db.addSessionToUser(user, sessionID); err != nil {
		return err
	}

	return nil
}

func getSessionID() (string, error) {
	sid, err := shortid.New(1, shortid.DefaultABC, 2342)
	if err != nil {
		return "", err
	}

	return sid.Generate()
}
