package main

// condicionais
func main() {
	a := 10
	b := 20
	c := 30
	d := 40

	if a > b {
		println("a > b")
	}

	if a > b {
		println("a > b")
	} else {
		println("a < b")
	}

	if a > b && c > d {
		println("a > b && c > d")
	}

	if a > b || c > d {
		println("a > b || c > d")
	}

	switch a {
	case 10:
		println("a == 10")
	case 20:
		println("a == 20")
	default:
		println("a != 10 && a != 20")
	}
}
