package ctw

import "strings"

var (
	white      string = "\033[38;2;255;255;255m"
	green      string = "\033[38;2;055;183;021m"
	brown      string = "\033[38;2;192;154;107m"
	brailEmpty string = "\u0020"
)

func DisplayColor(b bool) {
	if b == false {
		white = ""
		green = ""
		brown = ""
	}
}

func DisplayBrailEmpty(b bool) {
	if b == false {
		brailEmpty = " "
	}
}

func getGitColor(gitStatus string) string {
	switch strings.Trim(gitStatus, " ") {
	case "":
		return white
	case "U":
		return green
	default:
		return brown
	}
}

func widthsSum(w [][4]int, p int) int {
	s := 0
	for _, v := range w {
		s += v[0] + v[1] + v[2] + v[3] + p
	}
	s -= p
	return s
}
