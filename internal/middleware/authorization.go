package middleware

import (
	"errors"
	"net/http"

	"github.com/navaneet-rao/Go-API/api"
	"github.com/navaneet-rao/Go-API/internal/tools"
	log "github.com/sirupsen/logrus"
)

var UnAuthorzedError = errors.New("Invalid user Auth token.")

func Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var username string = r.URL.Query().Get("username")
		var token = r.Header.Get("Authorization")
		var err error

		if username == "" || token == "" {
			log.Error(UnAuthorzedError)
			api.RequestErrorHandler(w, UnAuthorzedError)
			return
		}

		var database *tools.DatabaseInterface
		database, err = tools.NewDatabase()

		if err != nil {
			api.InternalErrorHandler(w)
			return
		}

		var loginDetails *tools.LoginDetails
		loginDetails = (*database).GetUserLoginDetails(username)

		if loginDetails == nil || (token != (*loginDetails).AuthToken) {
			log.Error(UnAuthorzedError)
			api.RequestErrorHandler(w, UnAuthorzedError)
			return
		}
		next.ServeHTTP(w, r)

	})
}
