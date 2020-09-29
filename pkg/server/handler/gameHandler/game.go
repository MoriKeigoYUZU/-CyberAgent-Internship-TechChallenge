package gameHandler

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/2009_proto_h_server/pkg/http/response"
	"github.com/2009_proto_h_server/pkg/server/model/rankingModel"
	"github.com/2009_proto_h_server/pkg/server/model/userModel"
)

type GameFinishRequest struct {
	Score   int32 `json:"score"`
	StageId int32 `json:"stageId"`
}

type GameFinishResponse struct {
	Name     string     `json:"name"`
	Score    int32      `json:"score"`
	RankInfo []RankInfo `json:"rankinfo"`
}

type RankInfo struct {
	Name  string `json:"name"`
	Score int32  `json:"score"`
}

func HandleGameFinish() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {

		// リクエストヘッダからx-token(認証トークン)を取得
		token := request.Header.Get("x-token")
		if token == "" {
			log.Println("x-token is empty")
			return
		}

		// トークンを用いてユーザを特定
		//SelectName
		userName, err := userModel.SelectName(token)
		if err != nil {
			log.Println(err)
			response.InternalServerError(writer, "Internal Server Error")
			return
		}
		if userName == nil {
			log.Println(errors.New("user not found"))
			response.BadRequest(writer, fmt.Sprintln("user not found."))
			return
		}

		// リクエストBodyから更新後情報を取得
		var requestBody GameFinishRequest
		if err := json.NewDecoder(request.Body).Decode(&requestBody); err != nil {
			log.Println(err)
			response.InternalServerError(writer, "Internal Server Error")
			return
		}

		// 自スコア取得
		myRanking, err := rankingModel.SelectGettingMyRanking(userName.Name, requestBody.StageId)
		if err != nil {
			log.Println(err)
			response.InternalServerError(writer, "Internal Server Error 1")
			return
		}

		// 入っているランキングのハイスコア < 今回の獲得したスコア
		if myRanking.Score <= requestBody.Score {
			// ランキング更新
			err := rankingModel.UpdateRanking(userName.Name, requestBody.StageId, requestBody.Score)
			if err != nil {
				log.Println(err)
				response.InternalServerError(writer, "Internal Server Error 2")
				return
			}
		}

		// 上位ランキング取得
		userRankings, err := rankingModel.SelectGettingRanking(requestBody.StageId)
		if err != nil {
			log.Println(err)
			response.InternalServerError(writer, "Internal Server Error")
			return
		}

		// 自スコア取得
		myHighRanking, err := rankingModel.SelectGettingMyRanking(userName.Name, requestBody.StageId)
		if err != nil {
			log.Println(err)
			response.InternalServerError(writer, "Internal Server Error")
			return
		}

		// 詰める
		gameRankings := make([]RankInfo, len(userRankings))
		for i, gameRanking := range userRankings {
			gameRankings[i] = RankInfo{
				Name:  gameRanking.Name,
				Score: gameRanking.Score,
			}
		}

		// 返す
		response.Success(writer, &GameFinishResponse{
			Name:     myHighRanking.Name,
			Score:    myHighRanking.Score,
			RankInfo: gameRankings,
		})
	}
}
