package main

import "fmt"

func starpattern1() {
	for i := 1; i <= 5; i++ {
		for i := 1; i <= 5; i++ {
			for j := i; j > 0; j-- {
				fmt.Print("*")
			}
			fmt.Println("")
		}
	}
}

func starpattern2() {
	for i := 5; i >= 0; i-- {
		for j := i; j >= 0; j-- {
			fmt.Print("*")
		}
		fmt.Println("")
	}
}

func starpattern3() {
	for i := 1; i <= 5; i++ {
		for j := i; j > 0; j-- {
			fmt.Print("*")
		}
		fmt.Println("")
	}
	for i := 5; i >= 0; i-- {
		for j := i; j >= 0; j-- {
			fmt.Print("*")
		}
		fmt.Println("")
	}
}

func starpattern4() {
	for i := 5; i >= 0; i-- {
		for j := i; j >= 0; j-- {
			fmt.Print(" ")
		}
		for j := 5 - i; j >= 0; j-- {
			fmt.Print("*")
		}
		fmt.Println("")
	}
}

func starpattern5() {
	for i := 5; i >= 0; i-- {
		for j := i; j >= 0; j-- {
			fmt.Print("*")
		}
		for j := 5 - i; j >= 0; j-- {
			fmt.Print(" ")
		}
		fmt.Println("")
	}
}

func starpattern6() {
	for i := 5; i >= 0; i-- {
		for j := i; j >= 0; j-- {
			fmt.Print(" ")
		}
		for j := 5 - i; j >= 0; j-- {
			fmt.Print("*")
		}
		fmt.Println("")
	}
	for i := 5; i >= 0; i-- {
		for j := i; j >= 0; j-- {
			fmt.Print("*")
		}
		for j := 5 - i; j >= 0; j-- {
			fmt.Print(" ")
		}
		fmt.Println("")
	}
}

func starpattern7() {
	for i := 5; i >= 0; i-- {
		for j := 5 - i; j > 0; j-- {
			fmt.Print(" ")
		}
		for j := i; j >= 0; j-- {
			fmt.Print("*")
		}
		fmt.Println("")
	}
}

func starpattern8() {
	for i := 5; i >= 0; i-- {
		for j := i; j >= 0; j-- {
			fmt.Print(" ")
		}
		for j := 5 - i; j > 0; j-- {
			fmt.Print("*")
		}
		fmt.Println("")
	}
	for i := 5; i >= 0; i-- {
		for j := 5 - i; j > 0; j-- {
			fmt.Print(" ")
		}
		for j := i; j >= 0; j-- {
			fmt.Print("*")
		}
		fmt.Println("")
	}
}

func main() {
	fmt.Println("1")
	starpattern1()
	fmt.Println()
	fmt.Println("2")
	starpattern2()
	fmt.Println()
	fmt.Println("3")
	starpattern3()
	fmt.Println()
	fmt.Println("4")
	starpattern4()
	fmt.Println()
	fmt.Println("5")
	starpattern5()
	fmt.Println()
	fmt.Println("6")
	starpattern6()
	fmt.Println()
	fmt.Println("7")
	starpattern7()
	fmt.Println()
	fmt.Println("8")
	starpattern8()
	fmt.Println()
}
