#Parser for PostgreSQL interval data type into time.Duration

Totally incomplete, unstable and unreliable.

Designed to perform one and only concrete task.

Not aware of EXTRACT_EPOCH.

Not dealing with range difference.

In fact, valid input data range limited to test cases provided.

Usage should be avoided at any cost.

`

	import "github.com/rcmonitor/pginterval"

...

	var err error
	var pstDuration *time.Duration
	var strInterval = "10:25:42"
	pstDuration, err = pginterval.FParse(strInterval)
	if err != nil { return }
	pstDuration.String()
`