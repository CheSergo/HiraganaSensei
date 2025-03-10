package entity

type KatakanaCharacter struct {
	Symbol       string   `json:"symbol"`
	StrokeOrder  []string `json:"stroke_order"`
	ExampleWords []Word   `json:"example_words"`
}
