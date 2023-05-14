package auth

import (
	"encoding/json"
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

	result := Credentials{}
	err := db.findUser(&cred).Decode(&result)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		lib.NewErr("No Account Found").HandleErr(w)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(cred.Password))

	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		lib.NewErr("Invalid password").HandleErr(w)
		return
	}

	sid, err := shortid.New(1, shortid.DefaultABC, 2342)
	sessionID, err := sid.Generate()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		lib.NewErr("Error generating SID (session)").HandleErr(w)
		return
	}

	if err := db.addSessionToUser(cred.Username, sessionID); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		lib.StdErr(err)
		return
	}


	http.SetCookie(w, &http.Cookie{
		Name: "SessionID",
		Value: sessionID,
		HttpOnly: true,
		MaxAge: 15 * 60,
	})
	w.WriteHeader(http.StatusOK)

}
