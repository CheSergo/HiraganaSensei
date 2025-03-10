package entity

import "errors"

type Word struct {
	ID         string `json:"id"`
	Katakana   string `json:"katakana"`
	Hiragana   string `json:"hiragana"`
	Romaji     string `json:"romaji"`
	Meaning    string `json:"meaning"`
	Difficulty int    `json:"difficulty"`
}

// NewWord создает новый экземпляр слова
func NewWord(katakana, hiragana, romaji, meaning string, difficulty int) (*Word, error) {
	if katakana == "" || hiragana == "" || romaji == "" || meaning == "" {
		return nil, errors.New("обязательные поля не могут быть пустыми")
	}

	return &Word{
		ID:         "",
		Katakana:   katakana,
		Hiragana:   hiragana,
		Romaji:     romaji,
		Meaning:    meaning,
		Difficulty: difficulty,
	}, nil
}
