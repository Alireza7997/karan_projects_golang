package numbers

import "fmt"

// Find Cost of Tile to Cover W x H Floor :
// Calculate the total cost of tile it would take to cover a floor plan of width and height, using a cost entered by the user.

func CostOfTile() {
	var (
		width, height, tileCost float32
	)

	fmt.Print("Enter the cost of tile($): ")
	fmt.Scanln(&tileCost)

	fmt.Print("Enter width : ")
	fmt.Scanln(&width)

	fmt.Print("Enter height : ")
	fmt.Scanln(&height)

	tiles := width * height
	cost := tileCost * tiles
	fmt.Printf("Total cost : $%.2f", cost)
}
