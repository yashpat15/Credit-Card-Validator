// validator.go

package main

import "strconv"

func ValidateCreditCard(number string) bool {
	if number == "" {
		return false
	}
	nDigits := len(number)
	sum := 0
	isSecond := false

	for i := nDigits - 1; i >= 0; i-- {
		d, err := strconv.Atoi(string(number[i]))
		if err != nil {
			return false
		}

		if isSecond {
			d *= 2
			if d > 9 {
				d -= 9
			}
		}

		sum += d
		isSecond = !isSecond
	}

	return (sum % 10) == 0
}
