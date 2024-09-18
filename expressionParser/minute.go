package expressionparser

import "github.com/cron-parser/stratergy"

const (
	startMinute = 0
	endMinute   = 59
)

type minute struct {
	start           int
	end             int
	parseExpression func(field string, startLimit, endLimit int) ([]int, error)
}

func NewMinuteExpressionParser() *minute {
	st := stratergy.NewSimpleServiceStatergy()
	return &minute{
		start:           startMinute,
		end:             endMinute,
		parseExpression: st.Parse,
	}
}

func (h *minute) Parse(ip string) ([]int, error) {
	return h.parseExpression(ip, h.start, h.end)
}
