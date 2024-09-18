package expressionparser

import "github.com/cron-parser/stratergy"

const (
	startHour = 0
	endHour   = 23
)

type hour struct {
	start           int
	end             int
	parseExpression func(field string, startLimit, endLimit int) ([]int, error)
}

func NewHourExpressionParser() *hour {
	st := stratergy.NewSimpleServiceStatergy()
	return &hour{
		start:           startHour,
		end:             endHour,
		parseExpression: st.Parse,
	}
}

func (h *hour) Parse(ip string) ([]int, error) {
	return h.parseExpression(ip, h.start, h.end)
}
