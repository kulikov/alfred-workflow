package workflows

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func ConvertTimestamp(query string) []Item {
	output := make([]Item, 0)

	if query == "" {
		query = time.Now().Format("2006-01-02 15:04:05.000")
	}

	if timestamp, err := strconv.ParseInt(query, 10, 64); err == nil {
		if len(query) < 13 {
			timestamp = timestamp * 1000
		}

		localDate := time.UnixMilli(timestamp).Format("2006-01-02 15:04:05.000")
		output = append(output, makeItem(localDate, "Local"))

		utcDate := time.UnixMilli(timestamp).UTC().Format("2006-01-02 15:04:05.000")
		output = append(output, makeItem(utcDate, "UTC"))

	} else {
		match := regexp.MustCompile(`(\d{4}-\d{2}-\d{2})?[ T]?((\d{2}:\d{2})(:\d{2})?(.(\d{1,6}))?)?`).FindStringSubmatch(query)

		date, _, hourMinutes, seconds, subseconds := match[1], match[2], match[3], match[4], match[6]

		if date != "" {
			if hourMinutes == "" {
				hourMinutes = "00:00"
			}

			if seconds == "" {
				seconds = ":00"
			}

			subsecondsInt, err := strconv.ParseInt(subseconds, 10, 64)
			if err != nil {
				subsecondsInt = 0
			}

			formatted := fmt.Sprintf("%s %s%s.%d%s", date, hourMinutes, seconds, subsecondsInt, strings.Repeat("0", 6-len(strconv.FormatInt(subsecondsInt, 10))))

			if dt, err := time.ParseInLocation("2006-01-02 15:04:05.000000", formatted, time.Now().Location()); err == nil {
				millis := fmt.Sprintf("%d", dt.UnixMilli())
				output = append(output, makeItem(millis, "Unixtime millis"))
			} else {
				output = append(output, makeItem(fmt.Sprintf("Error: %s", err), fmt.Sprintf("Error for %s", query)))
			}
		}
	}

	return output
}
