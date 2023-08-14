package encrypt

const times = 4

func CopyStr(str string) string {
	newStr := ""
	for i := 0; i < times; i++ {
		newStr += str
	}
	return newStr
}

func ReduceStr(str string) string {
	return str[0:36]
}
