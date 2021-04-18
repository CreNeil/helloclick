package util

func GetString() string {
	return "hello world"
}

func GetString2(c string, d string) (a string, b string) {
	a = "hello"
	b = "world"
	var e = "neil"
	b = b + e + " " + c + " " + d
	return a, b
}
