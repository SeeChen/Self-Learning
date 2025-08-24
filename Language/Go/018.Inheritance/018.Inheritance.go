package main

import "fmt"

// MyName defines an interface for types that can return their name.
type MyName interface {
	MyName() string
}

// Animal represents a base struct to be embedded in other types.
type Animal struct {
	Name string
}

// Dog embeds Animal and adds its own behavior.
type Dog struct {
	Animal
	Voice string
}

// MyName implements the MyName interface for Dog.
// This overrides the default behavior from Animal.
func (d *Dog) MyName() string {
	return fmt.Sprintf("I am a Dog, my name is %s.", d.Name)
}

// Speak is Dog's specific method for making sound.
func (d *Dog) Speak() string {
	return fmt.Sprintf("%s barks: %s", d.Name, d.Voice)
}

// Cat embeds Animal and has its own implementation for MyName.
type Cat struct {
	Animal
	Color string
}

// MyName implements the MyName interface for Cat.
// Notice: Different output than Dog -> demonstrates polymorphism.
func (c *Cat) MyName() string {
	return fmt.Sprintf("I am a Cat named %s with %s fur.", c.Name, c.Color)
}

// Speak is Cat's specific method.
func (c *Cat) Speak() string {
	return fmt.Sprintf("%s meows softly.", c.Name)
}

func main() {
	// Instantiate Dog and Cat
	dog := &Dog{Animal: Animal{Name: "Buddy"}, Voice: "Woof Woof"}
	cat := &Cat{Animal: Animal{Name: "Kitty"}, Color: "white"}

	// Interface slice to hold multiple types
	animals := []MyName{dog, cat}

	// Polymorphic behavior: calling MyName on different types
	for _, a := range animals {
		fmt.Println(a.MyName())
	}

	// Specific behaviors via concrete types
	fmt.Println(dog.Speak()) // Buddy barks: Woof Woof
	fmt.Println(cat.Speak()) // Kitty meows softly.
}
