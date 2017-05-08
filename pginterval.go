package pginterval

import "time"

func FParse(strInterval string) (*time.Duration, error) {
	var stTemp, stBase time.Time
	var stInterval time.Duration
	var err error
	strIntervalFormat := "15:04:05"
	strTimeFormat := "2006-01-02 15:04:05.000 -07"
	strBaseTime := "0000-01-01 00:00:00.000 +00"
	stTemp, err = time.Parse(strIntervalFormat, strInterval)
	if err != nil { return nil, err }

	stBase, err = time.Parse(strTimeFormat, strBaseTime)
	if err != nil { return nil, err }

	stInterval = stTemp.Sub(stBase)

	return &stInterval, nil
}
