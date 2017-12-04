package pack

func Add(args ...int) int {
	sum := 0
	for _, arg := range args {
		sum += arg
	}
	return sum

}
