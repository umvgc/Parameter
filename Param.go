package parameter

func Check(num int) string {
	if num%2 == 0 {
		return "GENAP"
	} else {
		return "GANJIL"
	}
}
