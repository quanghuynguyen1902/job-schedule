package utils

import (
	"fmt"
	"regexp"
	"strconv"
	"time"
)

func ParseISO8601Duration(isoDuration string) (time.Duration, error) {
	re := regexp.MustCompile(`P(?:(\d+)D)?T?(?:(\d+)H)?(?:(\d+)M)?(?:(\d+)S)?`)
	matches := re.FindStringSubmatch(isoDuration)

	if len(matches) == 0 {
		return 0, fmt.Errorf("invalid ISO 8601 duration: %s", isoDuration)
	}

	var duration time.Duration

	// Days
	if matches[1] != "" {
		days, _ := strconv.Atoi(matches[1])
		duration += time.Duration(days) * 24 * time.Hour
	}

	// Hours
	if matches[2] != "" {
		hours, _ := strconv.Atoi(matches[2])
		duration += time.Duration(hours) * time.Hour
	}

	// Minutes
	if matches[3] != "" {
		minutes, _ := strconv.Atoi(matches[3])
		duration += time.Duration(minutes) * time.Minute
	}

	// Seconds
	if matches[4] != "" {
		seconds, _ := strconv.Atoi(matches[4])
		duration += time.Duration(seconds) * time.Second
	}

	return duration, nil
}
