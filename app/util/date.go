package util

import "time"

const (
	TIME_FORMAT = "2006-01-02"
)

func ParseDate(dateStr string) (time.Time, error) {
	date, err := time.Parse(TIME_FORMAT, dateStr)
	if err != nil {
		date, err = time.ParseInLocation(TIME_FORMAT, dateStr, time.Now().Location())
		if err != nil {
			return time.Now(), err
		}
	}
	return date, nil
}
