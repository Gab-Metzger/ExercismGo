package gigasecond

import "time"

const TestVersion = testVersion

func AddGigasecond(t time.Time) time.Time {
	return t.Add(1e9 * time.Second)
}

var Birthday = time.Date(1993, 9, 3, 23, 0, 0, 0, time.UTC)
