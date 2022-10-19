package utils

import (
	"errors"
	"time"

	"github.com/tcc-uniftec-5s/internal/infra/constants"
)

var ERR_INVALID_DATE_FORMAT = errors.New("invalid date format")

const datePattern = "T00:00:00Z"

func ParseDate(date string) (time.Time, error) {
	t, err := time.Parse(time.RFC3339, date+datePattern)
	if err != nil {
		return t, ERR_INVALID_DATE_FORMAT
	}

	return t, nil
}

func FormatStringDate(date time.Time) string {
	return date.Format(constants.ParsingDateLayout)
}
