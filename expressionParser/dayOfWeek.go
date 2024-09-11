package expressionparser

import "github.com/cron-parser/stratergy"

const (
	startDayOfWeekLimit = 1
	endDayOfWeekLimit   = 7
)

type dayOfWeek struct {
	start           int
	end             int
	parseExpression func(field string, startLimit, endLimit int) ([]int, error)
}

func NewDayOfWeekExpressionParser() *dayOfWeek {
	st := stratergy.NewSimpleServiceStatergy()
	return &dayOfWeek{
		start:           startDayOfWeekLimit,
		end:             endDayOfWeekLimit,
		parseExpression: st.Parse,
	}
}

func (h *dayOfWeek) Parse(ip string) ([]int, error) {
	return h.parseExpression(ip, h.start, h.end)
}
