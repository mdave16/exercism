// Leap is a package to work with leap years.
package leap

// IsLeapYear calculates if a given year is a Leap Year.
func IsLeapYear(year int) bool {
	if year%4 == 0 {
		if year%100 == 0 {
			if year%400 == 0 {
				return true
			}
			return false
		}
		return true
	}
	return false
}
