package main

import "fmt"
import "os" 

func checkModThree(a int) string{
	if(a % 3 == 0) {
		return "Bar"	
	} else {
		return "Foo"	
	}
}

func numbers(a int, b int) bool{
	if (a > b) {
		return true	
	}
	fmt.Printf("%-3d %s\n",a,checkModThree(a))
	return numbers(a+1, b)
}

func main() {
	var num int

	fmt.Print("Введите число: ")
	fmt.Fscan(os.Stdin, &num)

	if(num < 1) {
		fmt.Println("Ошибка! Введите число больше нуля!")
		os.Exit(0)		
	}
	
	numbers(1, num)
}
