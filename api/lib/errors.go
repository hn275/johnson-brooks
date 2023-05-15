package lib

import (
	"encoding/json"
	"net/http"
)

type Err struct {
	ErrMessage string `json:"error"`
}

func NewErr(errMessage string) *Err {
	return &Err{errMessage}
}

func (errInfo *Err) HandleErr(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(errInfo); err != nil {
		panic(err)
	}
}
