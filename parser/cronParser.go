package parser

import (
	"fmt"

	expressionparser "github.com/cron-parser/expressionParser"
	"github.com/cron-parser/lib"
)

type cronParser struct {
	input                                      string
	minute, hour, dayOfMonth, month, dayOfWeek []int
	command                                    string
	values                                     []string
	parseHours                                 func(ip string) ([]int, error)
	parseMinute                                func(ip string) ([]int, error)
	parseDayOfMonth                            func(ip string) ([]int, error)
	parseMonth                                 func(ip string) ([]int, error)
	parseDayOfWeek                             func(ip string) ([]int, error)
}

func NewCronParser(expression string) (*cronParser, error) {
	p := &cronParser{
		input:           expression,
		parseHours:      expressionparser.NewHourExpressionParser().Parse,
		parseMinute:     expressionparser.NewMinuteExpressionParser().Parse,
		parseDayOfMonth: expressionparser.NewDayOfMonthExpressionParser().Parse,
		parseMonth:      expressionparser.NewMonthExpressionParser().Parse,
		parseDayOfWeek:  expressionparser.NewDayOfWeekExpressionParser().Parse,
	}
	err := p.validateCronPraserExpression()
	if err != nil {
		return nil, err
	}
	return p, nil
}

func (p *cronParser) Parse() error {
	var err error
	p.minute, err = p.parseMinute(p.values[0])
	if err != nil {
		return err
	}
	p.hour, err = p.parseHours(p.values[1])
	if err != nil {
		return err
	}
	p.dayOfMonth, err = p.parseDayOfMonth(p.values[2])
	if err != nil {
		return err
	}
	p.month, err = p.parseMonth(p.values[3])
	if err != nil {
		return err
	}
	p.dayOfWeek, err = p.parseDayOfWeek(p.values[4])
	if err != nil {
		return err
	}
	p.command = p.values[5]
	return nil
}
func (p *cronParser) Print() {
	fmt.Println("minute -> ", lib.ConvertIntsToSingleString(p.minute))
	fmt.Println("hour -> ", lib.ConvertIntsToSingleString(p.hour))
	fmt.Println("day of month -> ", lib.ConvertIntsToSingleString(p.dayOfMonth))
	fmt.Println("month -> ", lib.ConvertIntsToSingleString(p.month))
	fmt.Println("day of week -> ", lib.ConvertIntsToSingleString(p.dayOfWeek))
	fmt.Println("command -> ", p.command)
}
