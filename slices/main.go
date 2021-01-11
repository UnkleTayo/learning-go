package main

import "fmt"

func main() {


	goSlices()
}

func basicArray (){
		// Array declaration 
		var a [3]int //int array with length 3
		fmt.Println(a)
	
		b:=[3]int{3,5,6}
		fmt.Println(b)
	
		countryList := [...]string{"USA", "China", "India", "Germany", "France"}
		countryListCopy := countryList// a copy of a is assigned to b
		countryListCopy[0] = "Singapore"
		fmt.Println("countryList is ", countryList)
		fmt.Println("countryListCopy is ", countryListCopy) 
}


// looping through an array 
// using forLoop 
func forLooping(){
	h := [...]float64{67.8, 89.8, 21, 78}

	for i:=0; i<len(h); i++ {
		fmt.Printf("%d th element of h is %.2f\n", i, h[i])
	}
}

func rangeLooping(){
	a := [...]float64{67.7, 89.8, 21, 78}
	sum := float64(0)
	for i, v := range a {
		// range retursn both index and value
		fmt.Printf("%d the element of a is %.2f\n", i, v)
		sum += v
	}
	
	fmt.Println("\nsum of all elements of a",sum)
}

func multidimensionalArray(){
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

func  goSlices()  {
	a := [5]int{7,98,45,5,7}
	var b []int = a[1:4] //creates slice from index 1 to 3
	fmt.Println(b)
}

func createSlice() {  
	c := []int{6, 7, 8} //creates and array and returns a slice reference
	fmt.Println(c)
}