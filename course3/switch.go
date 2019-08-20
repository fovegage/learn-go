package main

func main() {
	println(score1(99))
	println(score2(60))
}

func score1(s int) string {
	switch s / 10 {
	case 9:
		return "优秀"
	default:
		return "不及格"
	}
}

func score2(s int) string {
	switch {
	case s >= 60:
		return "及格"
	default:
		return "不及格"

	}
}
