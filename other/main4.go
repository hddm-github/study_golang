package main

import (
	"fmt"
	"sync"
)

func main() {
	//var arr []int = []int{1, 2, 3, 4, 5}
	//arr2 := [...]int{1,2,3,4,5}
	//for i, value := range arr {
	//	fmt.Printf("%d",i)
	//}
	//type arr3 [3]int
	//var arr arr3 = arr3{2: 3}

	//a := [2]int{1, 2}
	//b := [...]int{1, 2}
	//c := [2]int{1, 3}
	//
	//fmt.Println(a == b, a == c, b == c)
	//
	//d := [...]int{1, 2}
	//fmt.Println(a == d)
	//e := [2][2]int{{1, 2}, {3, 4}}
	//fmt.Printf("%d", e)
	//f := [4][2]int{3: {1: 2}}

	//append()
	//a := make([]int, 2)
	//b := make([]int, 3, 20)
	//fmt.Println(a, b)
	//fmt.Println(len(a), cap(a))
	//fmt.Println(len(b), cap(b))
	//slice1 := []int{1, 2, 3, 4, 5}
	//slice2 := []int{5, 4, 3}
	//copy(slice2, slice1)
	////fmt.Println(slice1, slice2)
	//copy(slice1, slice2)
	//fmt.Println(slice1, slice2)
	//sense := make(map[string][]int, 2)
	//sense["pig"] = []int{1, 2, 3}
	//delete(sense, "pig")

	//m := make(map[int]int)
	m2 := sync.Map{}
	m2.Store(1, 1)
	m2.Store(2, 2)
	m2.Store(3, 3)
	fmt.Println(m2.Load(2))

	m2.Range(func(key, value interface{}) bool {
		fmt.Println(key, value)
		return true
	})
	go func() {
		for {
			m2.Store(1, 2)
		}
	}()
	go func() {
		for {
			m2.Delete(1)
		}
	}()

	select {}

}
