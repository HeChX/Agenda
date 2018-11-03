package entity

import (
	"strconv"
	"strings"
)

type Date struct {
	Year   int
	Month  int
	Day    int
	Hour   int
	Minute int
}

func IsLeapYear(year int) bool {
	if (year%4 == 0 && year%100 != 0) || year%400 == 0 {
		return true
	}
	return false
}

func IsValid(date Date) bool {
	if date.Year > 9999 || date.Year < 1000 ||
		date.Month < 1 || date.Month > 12 ||
		date.Day < 1 ||
		date.Hour < 0 || date.Hour > 23 ||
		date.Minute < 0 || date.Minute > 59 {
		return false
	}
	if date.Month == 1 || date.Month == 3 || date.Month == 5 || date.Month == 7 || date.Month == 8 || date.Month == 10 || date.Month == 12 {
		if date.Day > 31 {
			return false
		}
	} else if date.Month == 4 || date.Month == 6 || date.Month == 9 || date.Month == 11 {
		if date.Day > 30 {
			return false
		}
	} else {
		if IsLeapYear(date.Year) {
			if date.Day > 29 {
				return false
			}
		} else {
			if date.Day > 28 {
				return false
			}
		}
	}
	return true
}

func StringtoDate(str string) (Date, bool) {
	var date = Date{
		Year:   0,
		Month:  0,
		Day:    0,
		Hour:   0,
		Minute: 0,
	}
	count := 0
	if len(str) == 16 {
		if str[4] == '-' && str[7] == '-' && str[10] == '/' && str[13] == ':' {
			for i := 0; i < 16; i++ {
				if i == 4 || i == 7 || i == 10 || i == 13 {
					continue
				}
				if str[i] >= '0' && str[i] <= '9' {
					count++
				}
			}
		}
	}
	if count == 12 {
		date.Year, _ = strconv.Atoi(str[0:4])
		date.Month, _ = strconv.Atoi(str[5:7])
		date.Day, _ = strconv.Atoi(str[8:10])
		date.Hour, _ = strconv.Atoi(str[11:13])
		date.Minute, _ = strconv.Atoi(str[14:16])
	}
	if IsValid(date) {
		return date, true
	}
	return date, false
}

func DatetoString(date Date) string {
	if IsValid(date) {
		str := ""
		Year := strconv.Itoa(date.Year)
		Month := strconv.Itoa(date.Month)
		Day := strconv.Itoa(date.Day)
		Hour := strconv.Itoa(date.Hour)
		Minute := strconv.Itoa(date.Minute)
		if len(Month) == 1 {
			Month = "0" + Month
		}
		if len(Day) == 1 {
			Day = "0" + Day
		}
		if len(Hour) == 1 {
			Hour = "0" + Hour
		}
		if len(Minute) == 1 {
			Minute = "0" + Minute
		}
		str = Year + "-" + Month + "-" + Day + "/" + Hour + ":" + Minute
		return str
	}
	return "0000-00-00/00:00"
}

func DateEqual(d1 Date, d2 Date) bool {
	return DatetoString(d1) == DatetoString(d2)
}

func Dateless(d1 Date, d2 Date) bool {
	if strings.Compare(DatetoString(d1), DatetoString(d2)) == -1 {
		return true
	}
	return false
}

func Datemore(d1 Date, d2 Date) bool {
	if strings.Compare(DatetoString(d1), DatetoString(d2)) == 1 {
		return true
	}
	return false
}

func DateOverlap(d1_start Date, d1_end Date, d2_start Date, d2_end Date) bool {
	if (Datemore(d1_end, d2_start) && Dateless(d1_end, d2_end)) ||
		(Datemore(d1_start, d2_start) && Dateless(d1_end, d2_end)) ||
		(Datemore(d1_start, d2_start) && Dateless(d1_start, d2_end)) ||
		(Dateless(d1_start, d2_start) && Datemore(d1_end, d2_end)) ||
		(DateEqual(d1_start, d2_start) && DateEqual(d1_end, d2_end)) {
		return true
	}
	return false
}
