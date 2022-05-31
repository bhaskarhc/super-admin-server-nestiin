package loginhandler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/bhaskarhc/admin-nestiin/db"
	"github.com/bhaskarhc/admin-nestiin/modules"
	token "github.com/bhaskarhc/admin-nestiin/services/jwt"
	apiresponse "github.com/bhaskarhc/admin-nestiin/utils/apiResponse"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Print("\n\t loginHandler triggred ")
	w.Header().Set("Content-Type", "application/json")
	reqBody := modules.LoginUserRequest{}

	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		resBody := apiresponse.Message(http.StatusBadGateway, false, "invalid input format")
		apiresponse.Respond(w, resBody)
		fmt.Print(err)
		return
	}
	user, err := GetUserByEmail(reqBody.Email, reqBody.Password)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		resBody := apiresponse.Message(http.StatusBadGateway, false, "Invalid email ")
		apiresponse.Respond(w, resBody)
		fmt.Print(err)
		return
	}
	tokenString, err := token.CreateJWT(user.ID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if reqBody.Password != user.Password {

		w.WriteHeader(http.StatusUnauthorized)
		resp := apiresponse.Message(http.StatusUnauthorized, false, reqBody.Email)
		apiresponse.Respond(w, resp)
		return

	}
	resp := apiresponse.Message(200, true, user.Email)
	resp["token"] = tokenString
	apiresponse.Respond(w, resp)
}

func GetUserByEmail(email, password string) (*modules.LoginUserResponse, error) {

	account := &modules.LoginUserResponse{}
	fmt.Printf("\n\n \t Email : %s \n\t password: %s \n\n", email, password)

	err := db.GetDB().Table("user_data").Where("email = ?", email).First(account).Error
	if err != nil {
		fmt.Printf("\n Error on query : \t %v", err)
		return account, err
	}
	fmt.Printf("\n \t Account \t %+v", account)
	return account, nil

}
