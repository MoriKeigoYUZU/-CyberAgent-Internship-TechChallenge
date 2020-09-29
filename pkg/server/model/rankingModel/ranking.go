package rankingModel

import (
	"database/sql"
	"errors"
	"log"

	"github.com/2009_proto_h_server/pkg/db"
)

type Ranking struct {
	Name    string
	StageId string
	Score   int32
}

// Users 複数のuserデータが必要なため、Userをスライス
type Rankings []*Ranking

func SelectGettingRanking(stage int32) (Rankings, error) {

	// score順に１０件取得
	// todo マジックナンバーを使わない
	rows, err := db.Conn.Query("SELECT name, score FROM ranking WHERE stage_id = ? ORDER BY score DESC LIMIT ?;", stage, 10)
	if err != nil {
		return nil, err
	}
	return convertToRankings(rows)
}

func SelectGettingMyRanking(name string, stage int32) (*Ranking, error) {

	//stmt, _ := db.Conn.Prepare("INSERT INTO ranking (name, stage_id, score) VALUES (?, ?, ?); ")
	stmt, _ := db.Conn.Prepare("INSERT INTO ranking (name, stage_id, score) SELECT ?, ?, ? FROM dual WHERE NOT EXISTS (SELECT name, score FROM ranking WHERE name = ? AND stage_id = ?)")

	//if err != nil {
	//	return nil err
	//}
	_, _ = stmt.Exec(name, stage, 0, name, stage)

	row := db.Conn.QueryRow("SELECT name, score FROM ranking WHERE name = ? AND stage_id = ?;", name, stage)

	return convertToRanking(row)
}

func InsertRanking(name string, stageId int, score int32) error {

	stmt, err := db.Conn.Prepare("INSERT INTO ranking (name, stage_id, score) VALUES (?, ?, ?);")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(name, stageId, score)
	return err
}

func UpdateRanking(name string, stageId int32, score int32) error {

	stmt, err := db.Conn.Prepare("UPDATE ranking set score = ? where name = ? AND stage_id = ?")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(score, name, stageId)
	return err
}

// convertToRanking rowデータをRankingデータへ変換する
func convertToRanking(row *sql.Row) (*Ranking, error) {
	ranking := Ranking{}
	err := row.Scan(&ranking.Name, &ranking.Score)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		log.Println(err)
		return nil, err
	}
	return &ranking, nil
}

// 複数レコード
// convertToRankings rowsデータをrankingsデータへ変換する
func convertToRankings(rows *sql.Rows) (Rankings, error) {
	// users用意
	var rankings Rankings
	for rows.Next() {
		var ranking Ranking
		if err := rows.Scan(&ranking.Name, &ranking.Score); err != nil {
			log.Println(errors.New("scan failed"))
			return nil, err
		}
		//要素を追加
		rankings = append(rankings, &ranking)
	}
	if err := rows.Err(); err != nil {
		log.Println(errors.New("rows scan failed"))
		return nil, err
	}
	return rankings, nil
}
