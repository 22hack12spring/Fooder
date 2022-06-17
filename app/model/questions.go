package model

import (
	"context"
	"errors"
)

type QuestionsRepository interface {
	CreateQuestions(ctx context.Context, args QuestionArgs) ([7]Questions, error)
	GetQuestion(ctx context.Context, questionId int, searchId string) (Questions, error)
	GetQuestionsBySearchId(ctx context.Context, searchId string) ([7]Questions, error)
}

type QuestionArgs struct {
	ShopIds  [7]string
	SearchId string
}

type Questions struct {
	ID        string `db:"id"`
	Shop_id   string `db:"shop_id"`
	Search_id string `db:"search_id"`
	Number    int    `db:"number"`
	CreatedAt string `db:"created_at"`
}

// CreateQuestions  質問データと search_id から Questions テーブルにデータを追加する
func (repo *SqlxRepository) CreateQuestions(ctx context.Context, args QuestionArgs) (questions [7]Questions, err error) {
	sql := "INSERT questions (shop_id, search_id, number) VALUES (?, ?, ?)"

	for i, s := range args.ShopIds {
		questions[i].Shop_id = s
		questions[i].Search_id = args.SearchId
		questions[i].Number = i

		_, err = repo.db.ExecContext(ctx, sql, questions[i].Shop_id, questions[i].Search_id, questions[i].Number)

		if err != nil {
			questions = [7]Questions{}
			return
		}
	}

	return
}

// GetQuestion  queestionId, searchId から Questions のデータを取得する
func (repo *SqlxRepository) GetQuestion(ctx context.Context, questionId int, searchId string) (Questions, error) {
	var question Questions

	sql := "SELECT * FROM questions WHERE question_id = ? AND search_id = ?"
	err := repo.db.GetContext(ctx, &question, sql, questionId, searchId)

	if err != nil {
		return Questions{}, nil
	}

	return question, nil
}

// いつか使うかもしれないので残しておきます
// GetQuestionsBySearchId  searchId から questions のデータを取得する
func (repo *SqlxRepository) GetQuestionsBySearchId(ctx context.Context, searchId string) ([7]Questions, error) {
	sql := "SELECT * FROM questions WHERE search_id = ?"

	var res []Questions

	err := repo.db.SelectContext(ctx, &res, sql, searchId)

	if err != nil {
		return [7]Questions{}, err
	}

	if len(res) != 7 {
		return [7]Questions{}, errors.New("backend: Incorrect number of questions are registered")
	}

	var questions [7]Questions

	for i, r := range res {
		questions[i] = r
	}

	return questions, nil
}
