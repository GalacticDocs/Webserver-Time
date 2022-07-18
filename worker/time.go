package worker

import "time"

var now = time.Now()

func amOrPm() string {
	var hour = now.Hour()
	var builder = StringBuilder("{0} {1}")

	if hour < 12 {
		
	}
}