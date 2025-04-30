package trainings

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type Training struct { //Создайте экспортируемую структуру Training.
	Steps                 int           //Количество шагов, проделанных за тренировку.
	TrainingType          string        //Тип тренировки.
	Duration              time.Duration //Длительность тренировки.
	personaldata.Personal               //Встроенная структура Personal из пакета personaldata.
}

func (t *Training) Parse(datastring string) (err error) {
	parts := strings.Split(datastring, ",") //Разделить строку datastring на слайс строк.
	if len(parts) != 3 {                    //Проверить, чтобы длина слайса была равна 3.
		return fmt.Errorf("invalid data format: expected 3 values, got %d", len(parts))
	}

	steps, err := strconv.Atoi(parts[0]) //Преобразовать первый элемент слайса (количество шагов) в тип int.
	if err != nil {                      //Обработать возможные ошибки.
		return fmt.Errorf("error parsing number of steps: %v", err) //При возникновении ошибки вернуть её из метода.
	}
	t.Steps = steps //Сохранить полученное значение в соответствующем поле структуры Training.

	t.TrainingType = parts[1] //Сохранить значение типа тренировки в соответствующем поле структуры Training.

	duration, err := time.ParseDuration(parts[2]) //Преобразовать третий  элемент слайса в time.Duration.
	if err != nil {                               //Обработать возможные ошибки.
		return fmt.Errorf("error parsing duration: %v", err) //При их возникновении вернуть ошибку.
	}
	t.Duration = duration //Сохранить полученное значение в соответствующем поле структуры Training.

	return nil
}

func (t Training) ActionInfo() (string, error) {
	distance := spentenergy.Distance(t.Steps, t.Height)           //Вычислить дистанцию, используя функцию из пакета spentenergy.
	speed := spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration) //Вычислить среднюю скорость, используя функцию из пакета spentenergy.

	var calories float64
	var err error

	switch strings.ToLower(t.TrainingType) { //Проверить, какой вид тренировки содержится в структуре Training.
	case "бег": //Рассчитать калории, используя функцию из пакета spentenergy.
		calories, err = spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	case "ходьба": //Рассчитать калории, используя функцию из пакета spentenergy.
		calories, err = spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	default: //Если был передан неизвестный тип тренировки, верните ошибку с текстом неизвестный тип тренировки.
		return "", fmt.Errorf("неизвестный тип тренировки: %s", t.TrainingType) //Если тип тренировки не соответствует ни одному из известных, вернуть ошибку.
	}

	if err != nil {
		return "", fmt.Errorf("error calculating calories: %v", err) //Описать сообщения об ошибках.
	}

	hours := t.Duration.Hours()
	info := fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f",
		t.TrainingType, hours, distance, speed, calories) //Сформируйте и верните строку, образец которой был выше.

	return info, nil
}
