package mivideo

import (
	"errors"

	"github.com/bitly/go-simplejson"
)

var (
	ErrFormatInvalid = errors.New("Json format invalid")
)

func ParseError(json *simplejson.Json) error {
	msg, err := json.Get("error").String()
	if err != nil {
		return ErrFormatInvalid
	}
	return errors.New(msg)
}
