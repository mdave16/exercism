// Leap is a package to work with leap years.
package leap

// IsLeapYear calculates if a given year is a Leap Year.
func IsLeapYear(year int) bool {
	return year%400 == 0 || (year%100 != 0 && year%4 == 0)
}
