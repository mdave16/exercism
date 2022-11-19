package parsinglogfiles

import (
	"fmt"
	"regexp"
)

// IsValidLine determines if text starts with valid log descriptions.
func IsValidLine(text string) bool {
	var re = regexp.MustCompile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)\]`)
	return re.MatchString(text)
}

// SplitLogLine by separators used by the new team.
func SplitLogLine(text string) []string {
	var re = regexp.MustCompile(`<[~*=-]*>`)
	return re.Split(text, -1)
}

// CountQuotedPasswords counts lines that reference "password" in quotes.
func CountQuotedPasswords(lines []string) int {
	var re = regexp.MustCompile(`(?i)".*password.*"`)
	count := 0
	for _, log := range lines {
		if re.MatchString(log) {
			count++
		}
	}
	return count
}

// RemoveEndOfLineText removes end-of-line<lineNumber> used by some teams.
func RemoveEndOfLineText(text string) string {
	var re = regexp.MustCompile(`end-of-line\d+`)
	return re.ReplaceAllString(text, "")
}

// TagWithUserName will tag lines referencing users to the users mentioned.
func TagWithUserName(lines []string) []string {
	var re = regexp.MustCompile(`User\s+([^\s]+)`)
	for index, log := range lines {
		if match := re.FindStringSubmatch(log); match != nil {
			lines[index] = fmt.Sprintf("[USR] %s %s", match[1], log)
		}
	}
	return lines
}
