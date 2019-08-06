package mathsUtil

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

var validIntegerRegex *regexp.Regexp

type bigInteger struct {
	value string
}

func InitBigInt(value string) (*bigInteger, error) {
	value = strings.TrimSpace(value)
	if isValid := isValid(value); !isValid {
		return nil, errors.New("cannot parse to a numeric value")
	}
	return &bigInteger{value: value}, nil
}
func isValid(value string) bool {
	if validIntegerRegex == nil {
		validIntegerRegex, _ = regexp.Compile("^[1-9][0-9]*")
	}
	return validIntegerRegex.MatchString(value)
}
func (num1 *bigInteger) Add(num2 *bigInteger) (*bigInteger, error) {
	len1 := len(num1.value)
	len2 := len(num2.value)
	ans := make([]uint8, Max(len1, len2))
	anslen := len(ans)
	if len1 < len2 {
		num1, num2 = num2, num1
		len1, len2 = len2, len1
	}
	carry := uint8(0)
	k := anslen - 1
	for i, j := len1-1, len2-1; i >= 0; i-- {
		c1 := uint8(num1.value[i] - 48)
		c2 := uint8(0)
		if j >= 0 {
			c2 = num2.value[j] - 48
			j--
		}
		sum := c1 + c2 + carry
		ans[k] = sum % 10
		if sum > 9 {
			carry = 1
		} else {
			carry = 0
		}
		k--
	}
	ansToStr := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(ans)), ""), "[]")
	if carry == 1 {
		ansToStr = "1" + ansToStr
	}
	return &bigInteger{ansToStr}, nil
}
