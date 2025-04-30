// Package actioninfo предоставляет функциональность для обработки и вывода информации о действиях.
package actioninfo

import (
	"log"
)

// Создайте интерфейс DataParser, в котором объявите сигнатуры методов Parse() и ActionInfo().
type DataParser interface {
	Parse(string) error
	ActionInfo() (string, error)
}

func Info(dataset []string, dp DataParser) {
	for _, data := range dataset { //Перебрать все значения слайса dataset в цикле.
		err := dp.Parse(data) //Распарсить каждое значение с помощью метода Parse().
		if err != nil {       //Обработать ошибку парсинга.
			log.Printf("Data parsing error: %v", err) //Залогировать ошибку.
			continue                                  //Перейти к следующей итерации цикла.
		}

		info, err := dp.ActionInfo() //Получить информацию о данных с помощью метода ActionInfo().
		if err != nil {
			log.Printf("Error getting information: %v", err) //При возникновении ошибки ее нужно залогировать.
			continue
		}
		log.Println(info) //Сформировать и вывести строку с информацией об активности с помощью метода ActionInfo().
	}
}
