package functions

func Cal(i, j int) (x, y, z int) {
	x = i + j
	y = i - j
	z = i * j
	return
}
func Prime(i int) bool {
	if i < 2 {
		return false
	}
	var flag int = 1
	for k := 2; k <= i/2; k++ {
		if i%k == 0 {
			flag = 0
			break
		}
	}
	if flag == 0 {
		return false
	}
	return true
}
func IsPalindrome(str string) bool {
	reversedStr := ""
	for i := len(str) - 1; i >= 0; i-- {
		reversedStr += string(str[i])
	}

	for i := range str {
		if str[i] != reversedStr[i] {
			return false
		}
	}
	return true
}
