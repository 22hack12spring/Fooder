package model

type QuestionsRepository interface {
	CreateQuestions(args QuestionArgs) (Questions, error)
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
func (repo *SqlxRepository) CreateQuestions(args QuestionArgs) (Questions, error) {
	return Questions{}, nil
}