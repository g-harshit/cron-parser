package parser

import (
	"errors"
	"strings"
)

func (p *cronParser) validateCronPraserExpression() error {
	p.values = strings.Split(p.input, " ")

	if len(p.values) != 6 {
		err := errors.New("invalid input")
		return err
	}
	return nil
}
