package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"reflect"
	"sync"
)

/* ************** example-1 *************** */

type people []struct {
	Name string `json:"name"`
}

type response struct {
	Items people `json:"items"`
}

/* ************** example-2 *************** */

type person struct {
	Name  string  `json:"name"`
	Hobby string  `json:"hobby,omitempty"`
	Email string  `json:"-"`
	Money float64 `json:"money,string"`
}

/* ************** example-3 *************** */

type human struct {
	name string
	age  int
}

func (h *human) setName(n string) {
	h.name = n
}

func (h *human) setAge(a int) {
	h.age = a
}

type student = human

func getInfo(h human) {
	fmt.Printf("\nHi , my name is %v , and my age is %v\n", h.name, h.age)
}

/* ************** example-4 *************** */

// shown in main()

/* ************** example-5 *************** */

var wg sync.WaitGroup

var mu sync.Mutex

func setGrade(grades map[string]int, gradeName string, gradeValue int) {
	mu.Lock()
	defer mu.Unlock()
	grades[gradeName] = gradeValue
}

/* ************** example-6 *************** */

type square struct {
	length float64
}

type circle struct {
	radius float64
}

// square methods

func (s square) area() float64 {
	return s.length * s.length
}

func (s square) circumf() float64 {
	return s.length * 4
}

// circle methods

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) circumf() float64 {
	return 2 * math.Pi * c.radius
}

// ---
// (The empty interface denoted by interface{} can hold values of any type.)
// ---

// generic interface - which works for both the above shapes
// interface groups types together -> based on their methods
type shape interface {

	/*
		interface method signatures
	*/
	area() float64
	circumf() float64
}

func printShapeInfo(s shape) {

	/* ******************************************************** */

	// (The empty interface denoted by interface{} can hold values of any type.)

	// style-1 : A type switch lets you choose between types
	switch val := s.(type) {
	case square:
		fmt.Printf("\n-- square -- : %v", val)
	case circle:
		fmt.Printf("\n-- circle -- : %v", val)
	default:
		fmt.Printf("\n-- unknown type -- : %v", val)

	}

	// style-2 :using %T for querying type

	sType := fmt.Sprintf("%T", s)
	fmt.Printf("\nsType : %v", sType)

	// style-3 : using reflect

	sTypeReflect := reflect.TypeOf(s)

	/* ******************************************************** */

	fmt.Printf("\nsTypeReflect : %v", sTypeReflect)

	fmt.Printf("\nprintShapeInfo : %T", s)
	fmt.Printf("\narea is : %0.2f", s.area())
	fmt.Printf("\ncircumference is : %0.2f", s.circumf())
}
func main() {
	fmt.Println("main")
	log.Printf("hello...")

	/* ************** example-1 *************** */

	p := people{
		{Name: "John"},
		{Name: "Steve"},
		{Name: "Peter"},
		{Name: "Joseph"},
		{Name: "Danny"},
		{Name: "Giri"},
		{Name: "Aaron"},
		{Name: "Joe"},
	}

	bs1, _ := json.Marshal(response{Items: p})
	bs2, _ := json.Marshal(response{})
	bs3, _ := json.Marshal(response{Items: people(nil)})
	bs4, _ := json.Marshal(response{Items: people{}})

	fmt.Printf("\n")

	fmt.Printf("\nbs1:")
	fmt.Printf("\n%v\n", string(bs1))

	fmt.Printf("\nbs2:")
	fmt.Printf("\n%v\n", string(bs2))

	fmt.Printf("\nbs3:")
	fmt.Printf("\n%v\n", string(bs3))

	fmt.Printf("\nbs4:")
	fmt.Printf("\n%v\n", string(bs4))

	/*
		Output:

		bs1:
		{"items":[{"name":"John"},{"name":"Steve"},{"name":"Peter"},{"name":"Joseph"},{"name":"Danny"},{"name":"Giri"},{"name":"Aaron"},{"name":"Joe"}]}

		bs2:
		{"items":null}

		bs3:
		{"items":null}

		bs4:
		{"items":[]}
	*/

	pRange := p[1:3:4] // [low:high:max] => cap = max-low

	fmt.Printf("\ncapacity : %v\n", cap(pRange))

	bs5, _ := json.Marshal(response{Items: pRange})

	fmt.Printf("\nbs5:")
	fmt.Printf("\n%v\n", string(bs5))

	/*
		Output:

		capacity : 3
		bs5:
		{"items":[{"name":"Steve"},{"name":"Peter"}]}
	*/

	/* ************** example-2 *************** */

	person1 := person{
		Name:  "John",
		Hobby: "Music",
	}
	person1bs, _ := json.Marshal(person1)
	fmt.Printf("\nperson1bs:")
	fmt.Printf("\n%v\n", string(person1bs))

	person2 := person{
		Name:  "John",
		Email: "a@b.com",
	}
	person2bs, _ := json.Marshal(person2)
	fmt.Printf("\nperson2bs:")
	fmt.Printf("\n%v\n", string(person2bs))

	person3 := person{
		Hobby: "Dance",
		Email: "a@b.com",
	}
	person3bs, _ := json.Marshal(person3)
	fmt.Printf("\nperson3bs:")
	fmt.Printf("\n%v\n", string(person3bs))

	person4 := person{
		Hobby: "Dance",
		Money: 55,
	}
	person4bs, _ := json.Marshal(person4)
	fmt.Printf("\nperson4bs:")
	fmt.Printf("\n%v\n", string(person4bs))

	/*
		Output:

		person1bs:
		{"name":"John","hobby":"Music","money":"0"}

		person2bs:
		{"name":"John","money":"0"}

		person3bs:
		{"name":"","hobby":"Dance","money":"0"}

		person4bs:
		{"name":"","hobby":"Dance","money":"55"}
	*/

	/* ************** example-3 *************** */

	var student1 student

	student1.setName("Gary")
	student1.setAge(30)

	getInfo(student1)

	/*
		Output:

		Hi , my name is Gary , and my age is 30
	*/

	/* ************** example-4 *************** */

	// -- Anonymous Structs --

	res := struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}{
		Name: "Steve",
		Age:  25,
	}

	resbs, _ := json.Marshal(res)
	fmt.Printf("\nresbs:\n%v\n", string(resbs))

	/*
		Output:

		resbs:
		{"name":"Steve","age":25}
	*/

	/* ************** example-5 *************** */

	g := map[string]int{
		"English": 9,
		"Math":    8,
	}

	wg.Add(2)

	go func() {
		setGrade(g, "Math", 5)
		wg.Done()
	}()

	go func() {
		setGrade(g, "Math", 6)
		wg.Done()
	}()

	wg.Wait()

	fmt.Printf("\ngrades:\n")
	dataByteArray, _ := json.MarshalIndent(g, "", "    ")
	fmt.Println(string(dataByteArray))

	/*
		output:

		grades:
		{
			"English": 9,
			"Math": 5
		}
	*/

	/* ************** example-6 *************** */

	shapes := []shape{
		square{length: 15.2},
		circle{radius: 7.5},
		circle{radius: 12.3},
		square{length: 4.9},
	}

	for _, shape := range shapes {
		printShapeInfo(shape)
		fmt.Printf("\n---")
	}

	/*
		Output:

		-- square -- : {15.2}
		sType : main.square
		sTypeReflect : main.square
		printShapeInfo : main.square
		area is : 231.04
		circumference is : 60.80
		---
		-- circle -- : {7.5}
		sType : main.circle
		sTypeReflect : main.circle
		printShapeInfo : main.circle
		area is : 176.71
		circumference is : 47.12
		---
		-- circle -- : {12.3}
		sType : main.circle
		sTypeReflect : main.circle
		printShapeInfo : main.circle
		area is : 475.29
		circumference is : 77.28
		---
		-- square -- : {4.9}
		sType : main.square
		sTypeReflect : main.square
		printShapeInfo : main.square
		area is : 24.01
		circumference is : 19.60
		---
	*/

}
