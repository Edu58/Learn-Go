package main

import (
	"encoding/json"
	"fmt"
)

type cow interface {
	name() string
	run() bool
}

type friesian struct {
	Breed string `json:"breed"`
	Age   int32  `json:"age"`
}

func (f *friesian) name() string {
	return f.Breed
}

func (f *friesian) run() bool {
	return f.Age > 10
}

func NewCow(name string, age int32) (cow, error) {
	result := &friesian{Breed: name, Age: age}
	encoded, _ := json.Marshal(result)
	fmt.Printf("RESULT IS: %v\n", string(encoded))
	return result, nil
}

func main() {
	ff, err := NewCow("gurnsey", 23)
	if err != nil {
		fmt.Println("Error creating cow:", err)
		return
	}

	// Since NewCow returns cow interface, assert to *friesian to access and set breed
	if f, ok := ff.(*friesian); ok {
		f.Breed = "nn"
		fmt.Println(f.name())
	}
}
