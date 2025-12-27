package main

import "fmt"

type Animal struct {
	name      string
	voice     string
	ableToFly bool
}

func main() {

	var input string
	fmt.Print("Masukkan nama hewan: ")
	fmt.Scan(&input)

	var animal Animal

	switch input {
	case "anjing":
		animal = Animal{
			name:      "anjing",
			voice:     "Gung-gung",
			ableToFly: false,
		}
	case "burung":
		animal = Animal{
			name:      "burung",
			voice:     "cuit",
			ableToFly: true,
		}
	case "ayam":
		animal = Animal{
			name:      "ayam",
			voice:     "kukuruyuk",
			ableToFly: false,
		}
		case "elang"
			animal = Animal{
				name:	"elang"
				voice: "hiii-hiiiw"
				ableToFly: true,
		}
	default:
		fmt.Println("Hewan tidak dikenal")
		return
	}

	fmt.Println("\nName: ", animal.name)
	fmt.Println("Suara: ", animal.voice)

	switch animal.ableToFly {
	case true:
		fmt.Println("Hewan ini bisa terbang")
	default:
		fmt.Println("Hewan ini tidak bisa terbang")
	}

}


