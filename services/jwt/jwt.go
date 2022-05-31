package jwt_service

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	apiresponse "github.com/bhaskarhc/admin-nestiin/utils/apiResponse"
	"github.com/dgrijalva/jwt-go"
)

type Token struct {
	UserId int
	jwt.StandardClaims
}

var SECRET = []byte("sample_signed_key")
var api_key = "1234"

func CreateJWT(userID int) (string, error) {

	tokenClaims := &Token{
		UserId: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tokenClaims)

	tokenStr, err := token.SignedString(SECRET)
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}

	return tokenStr, nil
	// token := jwt.New(jwt.SigningMethodHS256)

	// claims := token.Claims.(jwt.MapClaims)

	// claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	// claims["userID"] = userID

	// tokenStr, err := token.SignedString(SECRET)

	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return "", err
	// }
	// fmt.Printf("\n %s \n", tokenStr)

	// return tokenStr, nil
}

func ValidateJWT(next func(w http.ResponseWriter, r *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Header["Token"] != nil {
			token, err := jwt.Parse(r.Header["Token"][0], func(t *jwt.Token) (interface{}, error) {
				_, ok := t.Method.(*jwt.SigningMethodHMAC)
				if !ok {
					w.WriteHeader(http.StatusUnauthorized)
					w.Write([]byte("not authorized"))
				}
				return SECRET, nil
			})

			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte("not authorized: " + err.Error()))
			}

			if token.Valid {
				next(w, r)
			}
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("not authorized"))
		}
	})
}

func GetJwt(w http.ResponseWriter, r *http.Request) {
	if r.Header["Access"] != nil {
		if r.Header["Access"][0] != api_key {
			return
		} else {
			token, err := CreateJWT(0)
			if err != nil {
				return
			}
			fmt.Fprint(w, token)
		}
	}
}

var JwtAuthentication = func(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		tokenHeader := r.Header.Get("Authorization")
		if tokenHeader == "" { //Token is missing, returns with error code 403 Unauthorized
			response := apiresponse.Message(http.StatusUnauthorized, false, "Missing auth token")
			w.WriteHeader(http.StatusUnauthorized)
			apiresponse.Respond(w, response)
			return
		}

		splitted := strings.Split(tokenHeader, " ") //The token normally comes in format `Bearer {token-body}`, we check if the retrieved token matched this requirement
		if len(splitted) != 2 {
			response := apiresponse.Message(http.StatusExpectationFailed, false, "Invalid/Malformed auth token")
			w.WriteHeader(http.StatusExpectationFailed)
			apiresponse.Respond(w, response)
			return
		}
		tokenPart := splitted[1]

		tk := &Token{}

		token, err := jwt.ParseWithClaims(tokenPart, tk, func(token *jwt.Token) (interface{}, error) {
			return SECRET, nil
		})
		if err != nil { //Malformed token, returns with http code 403 as usual
			response := apiresponse.Message(http.StatusUnauthorized, false, err.Error())
			w.WriteHeader(http.StatusUnauthorized)
			apiresponse.Respond(w, response)
			return
		}
		if !token.Valid { //Token is invalid, maybe not signed on this server
			response := apiresponse.Message(http.StatusUnauthorized, false, "Invalid JWT token")
			w.WriteHeader(http.StatusUnauthorized)
			apiresponse.Respond(w, response)
			return
		}
		ctx := context.WithValue(r.Context(), "user", tk.UserId)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)

	})
}
