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
	ID        int    `db:"id"`
	ShopId    string `db:"shop_id"`
	SearchId  string `db:"search_id"`
	Number    int    `db:"number"`
	CreatedAt string `db:"created_at"`
}

// CreateQuestions  質問データと search_id から Questions テーブルにデータを追加する
func (repo *SqlxRepository) CreateQuestions(ctx context.Context, args QuestionArgs) (questions [7]Questions, err error) {
	sql := "INSERT questions (shop_id, search_id, number) VALUES (?, ?, ?)"

	for i, s := range args.ShopIds {
		questions[i].ShopId = s
		questions[i].SearchId = args.SearchId
		// 1 Index
		questions[i].Number = i + 1

		_, err = repo.db.ExecContext(ctx, sql, questions[i].ShopId, questions[i].SearchId, questions[i].Number)

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

	sql := "SELECT * FROM questions WHERE id = ? AND search_id = ?"
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
