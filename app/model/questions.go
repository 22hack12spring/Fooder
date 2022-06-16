package model

type QuestionsRepository interface {
	CreateQuestions(args QuestionArgs) (Questions, error)
}

type QuestionsData struct {
	Shop_id string
	Number int
}

type QuestionArgs struct {
	QuestionsArr [7]QuestionsData
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