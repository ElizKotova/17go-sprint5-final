package spentenergy

import (
	"fmt"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	MInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе.
)

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 { //Проверить входные параметры на корректность.
		//Если параметры некорректны, вернуть 0 калорий и соответствующую ошибку.
		return 0, fmt.Errorf("invalid input parameters: steps, weight, height, and duration must be positive")
	}
	meanSpeed := MeanSpeed(steps, height, duration) //Рассчитать среднюю скорость с помощью meanSpeed().
	if meanSpeed == 0 {
		return 0, fmt.Errorf("invalid input parameters: cannot calculate mean speed")
	}

	durationInMinutes := duration.Minutes()                       //Преобразовать продолжительность в минуты.
	calories := (weight * meanSpeed * durationInMinutes) / minInH //Рассчитать количество калорий.
	return calories * walkingCaloriesCoefficient, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 { //Проверить входные параметры на корректность.
		//Если параметры некорректны, вернуть 0 калорий и соответствующую ошибку.
		return 0, fmt.Errorf("invalid input parameters: steps, weight, height, and duration must be positive")
	}
	//Рассчитать среднюю скорость с помощью meanSpeed().
	meanSpeed := MeanSpeed(steps, height, duration)
	if meanSpeed == 0 {
		return 0, fmt.Errorf("invalid input parameters: cannot calculate mean speed")
	}

	durationInMinutes := duration.Minutes()
	calories := (weight * meanSpeed * durationInMinutes) / minInH //Рассчитать и вернуть количество калорий.
	return calories, nil
}

// Функция принимает количество шагов steps, рост пользователя height
// и продолжительность активности duration и возвращает среднюю скорость.
func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	if steps <= 0 || duration <= 0 { //Проверить, что продолжительность duration больше 0.
		return 0 //Если это не так, вернуть 0.
	}
	distance := Distance(steps, height) //Вычислить дистанцию с помощью Distance().
	hours := duration.Hours()           //Преобразовать продолжительность в часы.
	if hours == 0 {
		return 0
	}
	return distance / hours //Вычислить и вернуть среднюю скорость.
}

// Функция принимает количество шагов и рост пользователя в метрах, а возвращает дистанцию в километрах.
// Целочисленную переменную steps необходимо будет привести к другому числовому типу.
func Distance(steps int, height float64) float64 {
	stepLength := height * stepLengthCoefficient //Умножьте высоту пользователя на коэффициент длины шага stepLengthCoefficient.
	distance := float64(steps) * stepLength      //Умножьте пройденное количество шагов на длину шага.
	return distance / MInKm                      //Разделите полученное значение на число метров в километре.
}
