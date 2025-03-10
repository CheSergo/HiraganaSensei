package entity

import (
	"errors"
	// "encoding/json"
	"time"
)

// HiraganaCharacter представляет иероглиф хираганы с дополнительной информацией
type HiraganaCharacter struct {
	ID          string    `json:"id"`
	Symbol      string    `json:"symbol"`
	StrokeOrder []string  `json:"stroke_order"`
	Romaji      string    `json:"romaji"`
	Example     string    `json:"example"`
	Category    string    `json:"category"`
	Audio       string    `json:"audio"`
	Image       string    `json:"image"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// NewHiraganaCharacter создает новый иероглиф
func NewHiraganaCharacter(
	symbol, romaji, example, category, audio, image string,
	strokeOrder []string,
) (*HiraganaCharacter, error) {
	if symbol == "" || romaji == "" {
		return nil, errors.New("символ и ромадзи не могут быть пустыми")
	}

	now := time.Now()
	return &HiraganaCharacter{
		ID:          "", // будет заполнен в базе данных
		Symbol:      symbol,
		Romaji:      romaji,
		Example:     example,
		Category:    category,
		Audio:       audio,
		Image:       image,
		StrokeOrder: strokeOrder,
		CreatedAt:   now,
		UpdatedAt:   now,
	}, nil
}

// UpdateStrokeOrder обновляет порядок штрихов
func (h *HiraganaCharacter) UpdateStrokeOrder(strokeOrder []string) error {
	if len(strokeOrder) == 0 {
		return errors.New("порядок штрихов не может быть пустым")
	}
	h.StrokeOrder = strokeOrder
	h.UpdatedAt = time.Now()
	return nil
}

// Validate проверяет корректность данных иероглифа
func (h *HiraganaCharacter) Validate() error {
	if h.Symbol == "" || h.Romaji == "" {
		return errors.New("символ и ромадзи обязательны")
	}
	if len(h.StrokeOrder) == 0 {
		return errors.New("порядок штрихов обязателен")
	}
	return nil
}
