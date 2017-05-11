package pginterval_test

import (
	"time"
	"testing"
	"github.com/stretchr/testify/require"
	"github.com/rcmonitor/pginterval"
)

//var tAlmostDay = 23 * time.Hour + 59 * time.Minute + 59 * time.Second
var tAlmostDay = 24 * time.Hour - 1 * time.Second

var slsdPGInterval = []struct{
	interval string
	duration time.Duration
	fail bool
}{
	{"00:10:00", 10 * time.Minute, false},
	{"00:00:05", 5 * time.Second, false},
	{"00:00:01", 1 * time.Second, false},
	{"00:05:00", 5 * time.Minute, false},
	{"20 10:40:45", -1, true},
	{"23:59:59", tAlmostDay, false},
	{"01:00:00", 1 * time.Hour, false},
	{"00:00:10", 10 * time.Second, false},
}

func TestFParse(t *testing.T) {
	var err error
	var pstDuration *time.Duration
	require := require.New(t)


	for _, v := range slsdPGInterval {
		pstDuration, err = pginterval.FParse(v.interval)
		if v.fail {
			require.NotNil(err)
			require.Nil(pstDuration)
		}else{
			require.Nil(err)
			require.NotNil(pstDuration)
			require.Equal(v.duration, *pstDuration)
		}
	}
}

func UsageExample() {
	var err error
	var pstDuration *time.Duration
	var strInterval = "10:25:42"
	pstDuration, err = pginterval.FParse(strInterval)
	if err != nil { return }
	pstDuration.String()
}
