package rankingHandler

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

type RankingListRequest struct {
	StageId int32 `json:"stageId"`
}

type RankingListResponse struct {
	Name     string     `json:"name"`
	Score    int32      `json:"score"`
	RankInfo []RankInfo `json:"rankinfo"`
}

type RankInfo struct {
	Name  string `json:"name"`
	Score int32  `json:"score"`
}

func HandleRankingList() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {

		// リクエストヘッダからx-token(認証トークン)を取得
		token := request.Header.Get("x-token")
		if token == "" {
			log.Println("x-token is empty")
			return
		}

		// リクエストBodyから更新後情報を取得
		var requestBody RankingListRequest
		if err := json.NewDecoder(request.Body).Decode(&requestBody); err != nil {
			log.Println(err)
			response.InternalServerError(writer, "Internal Server Error")
			return
		}

		//SelectName
		// トークンを用いてユーザを特定
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

		// 上位ランキング取得
		userRankings, err := rankingModel.SelectGettingRanking(requestBody.StageId)
		if err != nil {
			log.Println(err)
			response.InternalServerError(writer, "Internal Server Error")
			return
		}

		// 自スコア取得
		myRanking, err := rankingModel.SelectGettingMyRanking(userName.Name, requestBody.StageId)
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
		response.Success(writer, &RankingListResponse{
			Name:     myRanking.Name,
			Score:    myRanking.Score,
			RankInfo: gameRankings,
		})
	}
}
