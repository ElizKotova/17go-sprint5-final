package daysteps

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type DaySteps struct {
	Steps                 int           //Количество шагов.
	Duration              time.Duration //Длительность прогулки.
	personaldata.Personal               //Встроенная структура Personal из пакета personaldata
}

func (ds *DaySteps) Parse(datastring string) (err error) {
	parts := strings.Split(datastring, ",")
	if len(parts) != 2 {
		return fmt.Errorf("invalid data format") //Метод возвращает ошибку.
	}

	steps, err := strconv.Atoi(parts[0])
	if err != nil {
		return fmt.Errorf("error parsing number of steps: %v", err)
	}
	ds.Steps = steps

	duration, err := time.ParseDuration(parts[1])
	if err != nil {
		return fmt.Errorf("error parsing duration: %v", err)
	}
	ds.Duration = duration

	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {
	const stepLength = 0.7
	const caloriesPerKm = 50.0

	distance := float64(ds.Steps) * stepLength / spentenergy.MInKm                                 //Вычислите дистанцию.
	calories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration) //Вычислите количество сожжённых калорий.
	if err != nil {                                                                                //При возникновении ошибки верните пустую строку и ошибку.
		return "", fmt.Errorf("error calculating calories: %v", err)
	}
	//Сформируйте и верните строку с информацией.
	info := fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.",
		ds.Steps, distance, calories)

	return info, nil
}
