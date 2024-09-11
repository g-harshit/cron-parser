package expressionparser

import "github.com/cron-parser/stratergy"

const (
	startDayOfMonthLimit = 1
	endDayOfMonthLimit   = 31
)

type dayOfMonth struct {
	start           int
	end             int
	parseExpression func(field string, startLimit, endLimit int) ([]int, error)
}

func NewDayOfMonthExpressionParser() *dayOfMonth {
	st := stratergy.NewSimpleServiceStatergy()
	return &dayOfMonth{
		start:           startDayOfMonthLimit,
		end:             endDayOfMonthLimit,
		parseExpression: st.Parse,
	}
}

func (h *dayOfMonth) Parse(ip string) ([]int, error) {
	return h.parseExpression(ip, h.start, h.end)
}
