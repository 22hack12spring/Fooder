package model

import (
	"errors"

	"github.com/google/uuid"
)

type QuestionsRepository interface {
	CreateQuestions(args QuestionArgs) ([7]Questions, error)
	GetQuestionsBySearchId(search_id string) ([7]Questions, error)
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

// GetQuestionsBySearchId  search_id から questions のデータを取得する
func (repo *SqlxRepository) GetQuestionsBySearchId(search_id string) ([7]Questions, error) {
	sql := "SELECT * FROM questions WHERE search_id = ?"
	
	rows, err := repo.db.DB.Query(sql, search_id)

	if err != nil {
		return [7]Questions{}, nil
	}
	
	var res []Questions
	var question Questions

	for rows.Next() {
		rows.Scan(&question.ID, &question.Shop_id, &question.Search_id, &question.Number, &question.CreatedAt)
		res = append(res, question)
	}

	if len(res) != 7 {
		return [7]Questions{}, errors.New("backend: Incorrect number of questions are registered")
	}

	var questions [7]Questions

	for i, r := range(res) {
		questions[i] = r
	}

	return questions, nil
}