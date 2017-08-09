package main

import "fmt"
import "sort"

type person struct {
	Name string
	Age  int
}

//"Equivalent of overriding toString method for object in java
func (p person) String() string {
	return fmt.Sprintf("Object  %s: %d", p.Name, p.Age)
}

type byAge []person

func (a byAge) Len() int           { return len(a) }
func (a byAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

func main() {
	people := []person{
		{"Gauri", 30},
		{"Shankar", 40},
		{"Ram", 32},
		{"Kumar", 42},
	}
	// Basically printing object as toString
	fmt.Println(people[2])
	fmt.Println(people)
	sort.Sort(byAge(people))
	fmt.Println(people)
}
