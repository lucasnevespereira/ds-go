package main

import (
	"errors"
	"fmt"
)

type Set struct {
	Elements map[string]struct{}
}

func NewSet() *Set {
	var s Set
	s.Elements = make(map[string]struct{})
	return &s
}

func (s *Set) Add(elem string) {
	s.Elements[elem] = struct{}{}
}

func (s *Set) Delete(elem string) error {
	if _, exists := s.Elements[elem]; !exists {
		return errors.New("element not present in set")
	}
	delete(s.Elements, elem)
	return nil
}

func (s *Set) Contains(elem string) bool {
	_, exists := s.Elements[elem]
	return exists
}

func (s *Set) List() {
	for key, _ := range s.Elements {
		fmt.Println(key)
	}
}

func main() {
	// Sets contain only unique values (unordered)
	mySet := NewSet()

	mySet.Add("Earth")
	mySet.Add("Venus")
	mySet.Add("Mars")
	mySet.Add("Earth") // will not show

	mySet.List() // Earth Venus Mars
	fmt.Println("-------")

	mySet.Delete("Venus")

	mySet.List() // Earth Mars
	fmt.Println("-------")

	fmt.Println(mySet.Contains("Mars")) // true
}
