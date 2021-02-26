package main

import "fmt"

func basicArray() {
	// Array declaration
	var a [3]int //int array with length 3
	fmt.Println(a)

	b := [3]int{3, 5, 6}
	fmt.Println(b)

	countryList := [...]string{"USA", "China", "India", "Germany", "France"}
	countryListCopy := countryList // a copy of a is assigned to b
	countryListCopy[0] = "Singapore"
	fmt.Println("countryList is ", countryList)
	fmt.Println("countryListCopy is ", countryListCopy)
}

// looping through an array
// using forLoop
func forLooping() {
	h := [...]float64{67.8, 89.8, 21, 78}

	for i := 0; i < len(h); i++ {
		fmt.Printf("%d th element of h is %.2f\n", i, h[i])
	}
}

func rangeLooping() {
	a := [...]float64{67.7, 89.8, 21, 78} // assigns a length of 4 to a
	sum := float64(0)
	for i, v := range a {
		// range returns both index and value
		fmt.Printf("%d the element of a is %.2f\n", i, v)
		sum += v
	}

	fmt.Println("\nsum of all elements of a", sum)
}

func multidimensionalArray() {
	a := [3][2]string{
		{"lion", "tiger"},
		{"cat", "dog"},
		{"pigeon", "peacock"}, //this comma is necessary. The compiler will complain if you omit this comma
	}

	printarray(a)
	var b [3][2]string
	b[0][0] = "apple"
	b[0][1] = "samsung"
	b[1][0] = "microsoft"
	b[1][1] = "google"
	b[2][0] = "AT&T"
	b[2][1] = "T-Mobile"
	fmt.Printf("\n")
	printarray(b)
}

func printarray(a [3][2]string) {
	for _, v1 := range a {
		for _, v2 := range v1 {
			fmt.Printf("%s ", v2)
		}
		fmt.Printf("\n")
	}
}

// muitidmensional array
func muitidmensional(a [3][2]string) {
	for _, v1 := range a {
		for _, v2 := range v1 {
			fmt.Printf("%s, ", v2)
		}
		fmt.Printf("%s, ", v1)

	}
}

func goSlices() {
	a := [5]int{7, 98, 45, 5, 7}
	var b []int = a[1:4] //creates slice from index 1 to 3
	fmt.Println(b)
}

func createSlice() {
	c := []int{6, 7, 8} //creates and array and returns a slice reference
	fmt.Println(c)
}

// slices

func sliceExample() {
	fruitarray := [...]string{"apple", "orange", "grape", "mango", "water melon", "pine apple", "chikoo"}
	fruitslice := fruitarray[1:3]
	fmt.Printf("length of slice %d capacity %d\n", len(fruitslice), cap(fruitslice)) //length of is 2 and capacity is 6
	fruitslice = fruitslice[:cap(fruitslice)]                                        //re-slicing furitslice till its capacity
	fmt.Println("After re-slicing length is", len(fruitslice), "and capacity is", cap(fruitslice))
}

func countries() []string {
	countries := []string{"USA", "Singapore", "Germany", "India", "Australia"}
	neededCountries := countries[:len(countries)-2]
	countriesCpy := make([]string, len(neededCountries))
	copy(countriesCpy, neededCountries) //copies neededCountries to countriesCpy
	return countriesCpy
}

func newExample() {
	slice := []int{10, 20, 30, 40, 50}

	// create  a new slice of length 2 and capacity 4
	// For slice[i:j] with an underlying array of capacity k
	// Length: j - i
	// Capacity:k-i

	newSlice := slice[1:3]

	// Append to new slice

	newSlice = append(newSlice, 60)
	// fmt.Println(newSlice)
	// fmt.Println(slice)

	increaseSliceCap := append(slice, 60)
	fmt.Println(increaseSliceCap)
	fmt.Println(len(increaseSliceCap))

}

func threeIndexSlice() {
	// Create a slice of strings.
	// Contains a length and capacity of 5 elements.
	fruits := []string{"Apple", "Orange", "Plum", "Banana", "Grape"}

	// slice the third element and increase it capacity
	newfruiitSlice := fruits[2:3:4]

	fmt.Println(newfruiitSlice)
}

// This is to show that range provides a copy of each element in a slice and doesnt
//  reference the slice

func rangeCopy() {
	slice := []int{2, 4, 6, 7, 7, 8, 14, 6}

	for index, value := range slice {
		fmt.Printf("Value :%d Value-Addr: %X ElememtAddr: %X\n", value, &value, &slice[index])
	}
}

func main() {

	// countriesNeeded := countries()
	// fmt.Println(countriesNeeded)
	// newExample()
	// threeIndexSlice()
	rangeCopy()
}
