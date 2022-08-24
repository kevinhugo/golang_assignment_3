package main
import "fmt"

func main() {
	names := []string{"Mr","Sir","Bro","Friend","Kevin","Hugo"}
	var pointers []*string
	for key, _ := range names {
		pointers = append(pointers, &names[key])
	}
	printNames := func (points []*string) {
		for index, eachPoint := range points {
			fmt.Println(index+1,*eachPoint)
		}
	}
	printNames(pointers)
}

// func printNames(names []string) {
// 	for index, eachName := range names {
// 		fmt.Println(index+1,eachName)
// 	}
// }