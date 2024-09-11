package expressionparser

import "github.com/cron-parser/stratergy"

const (
	startMonth = 1
	endMonth   = 12
)

type month struct {
	start           int
	end             int
	parseExpression func(field string, startLimit, endLimit int) ([]int, error)
}

func NewMonthExpressionParser() *month {
	st := stratergy.NewSimpleServiceStatergy()
	return &month{
		start:           startMonth,
		end:             endMonth,
		parseExpression: st.Parse,
	}
}

func (h *month) Parse(ip string) ([]int, error) {
	return h.parseExpression(ip, h.start, h.end)
}
