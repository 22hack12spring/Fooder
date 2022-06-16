package model

import (
	"github.com/google/uuid"
)

type QuestionsRepository interface {
	CreateQuestions(args QuestionArgs) ([7]Questions, error)
}

type QuestionArgs struct {
	Shop_ids [7]string
	Search_id *string
}

type Questions struct {
	ID			string	`db:"id"`
	Shop_id		string	`db:"shop_id"`
	Search_id	string	`db:"search_id"`
	Number		int		`db:"number"`
	CreatedAt	string	`db:"created_at"`
}

// CreateQuestions  質問データと search_id から Questions テーブルにデータを追加する
func (repo *SqlxRepository) CreateQuestions(args QuestionArgs) (questions [7]Questions, err error) {
	var u uuid.UUID
	
	sql := "INSERT questions (id, shop_id, search_id, number) VALUES (?, ?, ?, ?)"

	for i, s := range(args.Shop_ids) {
		u = uuid.New()
		questions[i].ID = u.String()
		questions[i].Shop_id = s
		questions[i].Search_id = *args.Search_id
		questions[i].Number = i

		_, err = repo.db.DB.Exec(sql, &questions[i].ID, &questions[i].Shop_id, &questions[i].Search_id, &questions[i].Number)

		if err != nil {
			questions = [7]Questions{}
			return
		}
	}
	
	return
}