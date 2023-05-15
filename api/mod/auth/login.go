package auth

import (
	"encoding/json"
	"errors"
	"jb/lib"
	"net/http"

	"github.com/teris-io/shortid"
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
	err := db.findUser(&cred).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		lib.NewErr("No Account Found").HandleErr(w)
		return
	}

	if err := authenticateUser(cred, user, db); err != nil {
		if errors.Is(err, ErrAuthFailed) {
			w.WriteHeader(http.StatusUnauthorized)
			lib.NewErr("Authentication failed").HandleErr(w)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		lib.StdErr(err)
		return
	}

	sessionID, err := getSessionID()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		lib.StdErr(err)
		return
	}

	if err := db.addSessionToUser(user.ID, sessionID); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		lib.StdErr(err)
		return
	}

	c := http.Cookie{
		Name:     "SessionID",
		Value:    sessionID,
		HttpOnly: true,
		MaxAge:   15 * 60, // NOTE: 15 minutes
	}
	http.SetCookie(w, &c)
	w.WriteHeader(http.StatusOK)
}

// TODO: test this
func authenticateUser(cred, user Credentials, db *authDatabase) error {
	if user.Session.FailedAttempts >= 5 {
		if err := db.lockUser(user.ID); err != nil {
			return err
		}
		return ErrAuthFailed
	}

	plain := []byte(cred.Password)
	hashed := []byte(user.Password)
	err := bcrypt.CompareHashAndPassword(hashed, plain)
	if err == nil {
		return nil
	}

	if err := db.authFailed(user.ID); err != nil {
		return err
	}

	return ErrAuthFailed
}

func getSessionID() (string, error) {
	sid, err := shortid.New(1, shortid.DefaultABC, 2342)
	if err != nil {
		return "", err
	}

	return sid.Generate()
}
