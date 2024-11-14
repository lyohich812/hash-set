package main

import (
	"errors"
	"fmt"
	"mentoring/week01/hash-set/hashset"
	"reflect"
)

func testHashset() {
	mySet := hashset.NewHashSet()
	mySet.Add("test")
	mySet.Add(1)
	mySet.Add("FIVE")
	mySet.Add("to_delete")
	mySet.Remove("to_delete")
	mySet.Remove("to_delete")
	mySet.Remove("to_delete_not_existing")
	fmt.Println(mySet)
	fmt.Println(mySet.Contains("test"))
	fmt.Println(mySet.Contains("badtest"))
	fmt.Println(mySet.Contains(1))
	secSet := hashset.NewHashSet("one", "two", "three", 0, "test", 1)
	fmt.Println("first set", mySet)
	fmt.Println("second set", secSet)
	fmt.Println("")
	fmt.Println("intersection: ", mySet.Intersection(secSet))
	fmt.Println("")
	fmt.Println("union: ", mySet.Union(secSet))
	fmt.Println("")
	fmt.Println("difference: ", mySet.Difference(secSet))
}

func findFirstDuplicate(slice interface{}) interface{} {
	set := hashset.NewHashSet()

	v := reflect.ValueOf(slice)
	if v.Kind() != reflect.Slice {
		return nil
	}

	for i := 0; i < v.Len(); i++ {
		element := v.Index(i).Interface()
		if set.Contains(element) {
			return element
		}
		set.Add(element)
	}
	return nil
}

func findFirstDuplicateGenerics[T any](slice []T) (T, error) {
	set := hashset.NewHashSet()

	for _, el := range slice {
		if set.Contains(el) {
			return el, nil
		}
		set.Add(el)
	}
	var zeroVal T
	return zeroVal, errors.New("no duplicates")
}

func main() {
	// testHashset()

	slice := []int{2, 5, 1, 2, 3, 5, 1, 2, 4}
	//slice := []int{1}
	//slice := []int{1, 2, 3, 4}
	sliceStrings := []string{"jon", "5", "doe", "jon", "3", "5"}

	fmt.Println("first duplicate:", findFirstDuplicate(slice))
	fmt.Println("")
	fmt.Println("first duplicate:", findFirstDuplicate(sliceStrings))

	fmt.Println("")
	{
		res, err := findFirstDuplicateGenerics([]int{1, 2, 3, 4})
		fmt.Printf("first duplicate: %v, error %v\n", res, err)
	}
	{
		res, err := findFirstDuplicateGenerics([]string{"jon", "5", "doe", "jon", "3", "5"})
		fmt.Printf("first duplicate: %v, error %v\n", res, err)
	}

}
