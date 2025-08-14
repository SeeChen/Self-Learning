package main

import "fmt"

func main() {
	// A Map in Golang is an unordered collection of key-value pairs.
	// A key is like an index that points to a data value
	// In fact, we can think of a map as a data format similar to JSON.

	// Define a Map data
	// Define a map structure using map[KeyType]ValueType
	// We can use the make() function to define a variable of Map structure.
	map1 := make(map[string]int) // For this line, define a map1 variable whose key is string type and value is int type.
	fmt.Println(map1)            // But it's empty data

	// In addition, we define a map variable with an initial value.
	var map2 map[string]string = map[string]string{
		"FirstName":  "SeeChen",
		"FamilyName": "Lee"}

	// To get the mapping value, can use the map[key] to get
	var firstName string = map2["FirstName"]
	fmt.Printf("My First Name is %s.\n", firstName)

	// And can use `key, judge := map[key]` to determine whether a key exists in the map.
	middleName, exists := map2["MiddleName"]
	if exists {
		fmt.Println(middleName)
	} else {
		fmt.Println("Middle name does not exists.")
	}

	// To add a key value to a map variable, can use map[key] = value to do this
	map2["MiddleName"] = "Handsome"
	for k, v := range map2 {
		fmt.Printf("%-10s: %s.\n", k, v)
	}

	fmt.Println()
	// But if the key exists in the map variable, this code will change the value in its mapping key.
	map2["MiddleName"] = "Very Handsome"
	for k, v := range map2 {
		fmt.Printf("%-10s: %s.\n", k, v)
	}

	// Delete function in Golang
	// To delete an item from a Map variable, use the delete() function.
	// The delete() function requires two parameters.
	// The first parameter indicates the map to be deleted.
	// The second parameter indicates the key to be deleted.
	// This operation deletes both the key and value from the Map variable passed to the delete function
	fmt.Println()
	delete(map2, "MiddleName")
	fmt.Println(map2)
}
