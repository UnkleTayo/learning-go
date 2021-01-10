package main

import "fmt"

type employee struct {  
	salary  int
	country string
}

func main() {
	// declaring a map
	// var colors map[string]string
	// colors := make(map[string]string)
	// colors := map[string]string{
	// 	"red":   "#ff0000",
	// 	"green": "#4bf745",
	// 	"white": "#fff",
	// }

	// printMap(colors)
	// printSalary()
	// employeeStruct()
	mapReference()
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex Code for", color, "is", hex)
	}
}

func printSalary() {  
	employeeSalary := make(map[string]int)
	employeeSalary["steve"] = 12000
	employeeSalary["jamie"] = 15000
	employeeSalary["mike"] = 9000

	// Retrieving value for a key from a map
	employee:= "mike"
	salary := employeeSalary[employee]
	fmt.Println("Salary of", employee, "is", salary)


	// checking od a key is present 
	newEmp := "joe"
	value, ok := employeeSalary[newEmp]

	if ok == true {
		fmt.Println("Salary of", newEmp, "is", value)
		return
		}
		fmt.Println(newEmp, "not found")

		fmt.Println("content of map")

		// Iterate over all elements in a map
		for key, value := range employeeSalary{
			fmt.Printf("employeeSalary[%s] = %d\n", key, value)
		}

		// Deleting items from a map
		fmt.Println("map before deletion", employeeSalary)
    delete(employeeSalary, "steve")
    fmt.Println("map after deletion", employeeSalary)
}


func employeeStruct(){
			emp1 := employee{
				salary:  12000,
				country: "USA",
		}
		emp2 := employee{
				salary:  14000,
				country: "Canada",
		}
		emp3 := employee{
				salary:  13000,
				country: "India",
		}

		employeeInfo:= map[string]employee{
			"Steve": emp1,
			"Jamie": emp2,
			"Mike": emp3,
		}

		for name, info := range employeeInfo {
			fmt.Printf("Employee: %s Salary:$%d  Country: %s\n", name, info.salary, info.country)
		}

}

func mapReference(){
	employeeSalary := map[string]int{
		"steve": 12000,
		"jamie": 15000,     
		"mike": 9000,
}
fmt.Println("Original employee salary", employeeSalary)
modified := employeeSalary
modified["mike"] = 18000
fmt.Println("Employee salary changed", employeeSalary)
}