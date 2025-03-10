package main

import (
	"fmt"
	"log"

	"github.com/CheSergo/HiraganaSensei/internal/domain/entity"
)

func main() {
	// Создание нового иероглифа
	basicA, err := entity.NewHiraganaCharacter(
		"あ",
		"a",
		"あさ - утро",
		"basic",
		"audio/a.mp3",
		"images/a.png",
		[]string{"верхний горизонтальный штрих", "вертикальный штрих"},
	)
	if err != nil {
		log.Fatal(err)
	}

	// Обновление порядка штрихов
	err = basicA.UpdateStrokeOrder([]string{
		"верхний горизонтальный штрих",
		"вертикальный штрих сверху вниз",
		"диагональный штрих слева направо",
	})
	if err != nil {
		log.Fatal(err)
	}

	// Валидация
	if err := basicA.Validate(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(basicA)
}
