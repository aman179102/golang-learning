package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

//declaration of a map
var m map[string]Vertex


var d = map[string]Vertex{
	"Bell Labs": Vertex{
		40.68433, -74.39967,
	},
	"Google": Vertex{
		37.42202, -122.08408,
	},
}

var l  = map[string]Vertex{
	"Bell Labs": {40.68433, -74.39967},
	"Google":    {37.42202, -122.08408},
}

func main() {
	//giving memory to the map
	m = make(map[string]Vertex)
	// giving value to the map key
	//Here the key is Bell Labs and the value is given to it
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])
	//map literals 
	fmt.Println(d)
	// If the top-level type is just a type name, you can omit it from the elements of the literal. 
	fmt.Println(l)
	//***************
	//mutating maps**
	//***************
	// Create a map named 'store' with string keys and int values
    store := make(map[string]int)

    // Add a key-value pair
    store["Answer"] = 42
    fmt.Println("The value:", store["Answer"])

    // Update the value for the same key
    store["Answer"] = 48
    fmt.Println("The value:", store["Answer"])

    // Delete the key from the map
    delete(store, "Answer")
    fmt.Println("The value:", store["Answer"]) // Key doesn't exist now, so prints: 0

    // Check if key exists using the 'comma ok' idiom
    v, ok := store["Answer"]
    fmt.Println("The value:", v, "Present?", ok)

}
