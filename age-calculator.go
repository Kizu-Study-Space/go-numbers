package main
import "fmt"
import "time"

func main() {
	var year_of_birth int
	var age int
	current_year := time.Now().Year()

	fmt.Print("Hello what is your year of birth?\n")
	fmt.Scan(&year_of_birth)
	age = current_year - year_of_birth
	fmt.Printf("You are %d years old.\n", age)
}
