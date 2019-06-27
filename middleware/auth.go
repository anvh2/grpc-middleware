package middleware

import (
	"grpc-middleware/models"
	"grpc-middleware/utils"
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

//JwtAuthentication ...
var JwtAuthentication = func(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		//list endpoint does not need authen
		notAuth := []string{"api/user/new", "api/user/login"}

		requestPath := req.URL.Path //current request path

		//serve request does not need authen
		for _, value := range notAuth {
			if value == requestPath {
				next.ServeHTTP(w, req)
				return
			}
		}

		resposne := make(map[string]interface{})
		tokenHeader := req.Header.Get("Authorization")

		if tokenHeader == "" { //token is missing, return error code 403
			resposne = utils.Message(false, "Missing auth token")
			w.WriteHeader(http.StatusForbidden)
			w.Header().Add("Content-Type", "application/json")
			utils.Respond(w, resposne)
			return
		}

		splited := strings.Split(tokenHeader, " ") // the token normally comes i
		if len(splited) != 2 {
			resposne = utils.Message(false, "Invalid/Malformed auth token")
			w.WriteHeader(http.StatusForbidden)
			w.Header().Add("Content-Type", "application/json")
			utils.Respond(w, resposne)
			return
		}

		tokenPart := splited[1] //grab token part
		tk := &models.Token{}

		token, err := jwt.ParseWithClaims(tokenPart, tk, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("token_password")), nil
		})

		if err != nil {
			resposne = utils.Message(false, "Malformed authentication token")
			w.WriteHeader(http.StatusForbidden)
			w.Header().Add("Content-Type", "application/json")
			utils.Respond(w, resposne)
			return
		}

		if !token.Valid {//token is invalid, maybe not signed on this server
			resposne = utils.Message(false, "token is not valid")
			w.WriteHeader(http.StatusForbidden)
			w.Header().Add("Content-Type", "application/json")
			utils.Respond(w, resposne)
			return
		}

		
	})
}
