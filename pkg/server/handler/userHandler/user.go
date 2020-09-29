package userHandler

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"

	"github.com/2009_proto_h_server/pkg/http/response"
	"github.com/2009_proto_h_server/pkg/server/model/userModel"
)

type UserSigninRequest struct {
	Name     string `json:"name"`
	PassWord string `json:"password"`
}

func HandleUserSignin() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {

		// リクエストBodyから更新後情報を取得
		var requestBody UserSigninRequest
		if err := json.NewDecoder(request.Body).Decode(&requestBody); err != nil {
			log.Println(err)
			response.InternalServerError(writer, "Internal Server Error")
			return
		}

		// UUIDで認証トークンを生成する
		authToken, err := uuid.NewRandom()
		if err != nil {
			log.Println(err)
			response.InternalServerError(writer, "Internal Server Error")
			return
		}

		// データベースにユーザデータを登録する
		// ユーザデータの登録クエリを入力する
		err = userModel.InsertUser(&userModel.User{
			Name:      requestBody.Name,
			PassWord:  requestBody.PassWord,
			AuthToken: authToken.String(),
			Coin:      0,
		})
		if err != nil {
			log.Println(err)
			response.InternalServerError(writer, "Internal Server Error")
			return
		}
	}
}

type UserLoginRequest struct {
	Name     string `json:"name"`
	PassWord string `json:"password"`
}

type UserLoginResponse struct {
	Name  string `json:"name"`
	Token string `json:"token"`
	Coin  int32  `json:"coin"`
}

// HandleUserGet ユーザ情報取得処理
func HandleUserLogin() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {

		// リクエストBodyから更新後情報を取得
		var requestBody UserLoginRequest
		if err := json.NewDecoder(request.Body).Decode(&requestBody); err != nil {
			log.Println(err)
			response.InternalServerError(writer, "Internal Server Error")
			return
		}

		// ユーザデータの取得処理を実装
		user, err := userModel.SelectUser(requestBody.Name, requestBody.PassWord)
		if err != nil {
			log.Println(err)
			response.InternalServerError(writer, "Internal Server Error")
			return
		}
		if user == nil {
			log.Println(errors.New("user not found"))
			response.BadRequest(writer, fmt.Sprintf("user not found. name=%s", requestBody.Name))
			return
		}

		// レスポンスに必要な情報を詰めて返却
		response.Success(writer, &UserLoginResponse{
			Name:  user.Name,
			Token: user.AuthToken,
			Coin:  user.Coin,
		})
	}
}
