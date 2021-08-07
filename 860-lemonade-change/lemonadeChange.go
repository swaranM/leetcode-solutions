package lemonadechange

func lemonadeChange(bills []int) bool {
	var five, ten int

	for _, v := range bills {
		switch v {

		case 5:
			five++
		case 10:
			if five > 0 {
				five--
				ten++
			} else {
				return false
			}
		case 20:
			if five > 0 && ten > 0 {
				five--
				ten--
			} else if five > 2 {
				five = five - 3
			} else {
				return false
			}

		}
	}

	return true

}
