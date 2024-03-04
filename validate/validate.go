package validate

import (
	"regexp"
)

type Validator interface {
	Validate() error
}

func LengthBetween(str string, min, max int) bool {
	return len(str) >= min && len(str) <= max
}

func allowEmpty(str string, strict bool) bool {
	return !strict && str == ""
}

func MHID(str string, strict bool) bool {
	return allowEmpty(str, strict) || len(str) == 8
}

func RQID(str string, strict bool) bool {
	return allowEmpty(str, strict) || LengthBetween(str, 25, 40)
}

func CorrelationID(str string, strict bool) bool {
	return allowEmpty(str, strict) || LengthBetween(str, 5, 50)
}

func TSID(str string, strict bool) bool {
	return allowEmpty(str, strict) || LengthBetween(str, 25, 35)
}

func MUID(str string, strict bool) bool {
	return allowEmpty(str, strict) || LengthBetween(str, 25, 35)
}

func BPID(str string, strict bool) bool {
	return allowEmpty(str, strict) || LengthBetween(str, 1, 100)
}

func ECID(str string, strict bool) bool {
	return allowEmpty(str, strict) || LengthBetween(str, 25, 35)
}

func LVT(str string, strict bool) bool {
	return allowEmpty(str, strict) || LengthBetween(str, 10, 20)
}

func HVT(str string, strict bool) bool {
	return allowEmpty(str, strict) || LengthBetween(str, 20, 40)
}

func Language(str string, strict bool) bool {
	return allowEmpty(str, strict) || LengthBetween(str, 2, 40)
}

func Timezone(str string, strict bool) bool {
	return allowEmpty(str, strict) || LengthBetween(str, 3, 30)
}

func UserAgent(str string, strict bool) bool {
	return allowEmpty(str, strict) || LengthBetween(str, 4, 255)
}

func CVV(str string, strict bool) bool {
	return allowEmpty(str, strict) || regexp.MustCompile("^[0-9]{3,4}$").MatchString(str)
}

func NameOnCard(str string, strict bool) bool {
	if !strict && str == "" {
		return true
	}
	return LengthBetween(str, 1, 50) && !regexp.MustCompile(`\d{3,}`).MatchString(str)
}

var numberRegex = regexp.MustCompile("[^0-9]")

func CardNumber(str string, strict bool) bool {
	if !strict && str == "" {
		return true
	}
	return LengthBetween(numberRegex.ReplaceAllString(str, ""), 12, 24)
}

func Month(month int32, strict bool) bool {
	return month > 0 && month <= 12 || (!strict && month == 0)
}

func Year(year int32, strict bool) bool {
	return year > 1900 && year < 2100 || (!strict && year == 0)
}
