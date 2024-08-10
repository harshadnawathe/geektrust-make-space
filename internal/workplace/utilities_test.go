package workplace

import (
	"fmt"
	"strconv"
)

func PeriodForTest(start, end string) Period {
	return NewPeriod(TimeForTest(start), TimeForTest(end))	
} 

func TimeForTest(time string) Time {
	var hh, mm int
	var err error

	hh, err = strconv.Atoi(time[:2])
	if err != nil {
		panic(fmt.Errorf("cannot parse time `%s`: %w", time, err ))
	}
  mm, err = strconv.Atoi(time[3:])
	if err != nil {
		panic(fmt.Errorf("cannot parse time `%s`: %w", time, err ))
	}

	return NewTime(uint8(hh), uint8(mm))
}
