package helper

import (
	"strings"
	"time"

	"github.com/Invan2/invan_order_service/config"
	"github.com/pkg/errors"
)

var (
	ErrInvalidArgument = errors.New("date is invalid. can't be parsed as date")
	ErrParsingDate     = func(err error) error {
		return errors.Wrap(err, "error parsing date")
	}
)

func ParseDate(date string) (string, error) {
	if len(date) < 10 {
		return time.Time{}.String(), ErrInvalidArgument
	}

	var bDate time.Time
	var err error
	switch {
	case len(date) == 10:
		bDate, err = time.Parse(config.DD_MM_YYYY, date)
		if err != nil {
			return time.Time{}.String(), ErrParsingDate(err)
		}
	default:
		date = strings.Replace(date, "+", " ", -1)
		bDate, err = time.Parse(config.DD_MM_YYYY_HH_MM_SS, date)
		if err != nil {
			return time.Time{}.String(), ErrParsingDate(err)
		}
	}
	bDate = bDate.Add(-5 * time.Hour)
	return bDate.UTC().Format(config.DateTimeFormat), nil
}
