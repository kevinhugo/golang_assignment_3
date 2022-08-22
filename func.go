package main
import "fmt"

func main() {
	names := []string{"Mr","Sir","Bro","Friend","Kevin","Hugo"}
	printNames(names)
}

func printNames(names []string) {
	for index, eachName := range names {
		fmt.Println(index+1,eachName)
	}
}