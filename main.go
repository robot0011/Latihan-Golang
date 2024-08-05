package main

import (
	"fmt"
)

// init function

func add(a int, b int) int {
	return a + b
}
func sayHello() {
	fmt.Println("Hello")
}

func countRecursive(n int) int {
	if n == 0 {
		return 0
	}
	fmt.Println(n)
	return n + countRecursive(n-1)
}

type User struct {
	Id    int
	Name  string
	Email string
	Pass  string
}

func main() {
	// init var
	var userName string = "Agus"
	var userAge int = 20
	fmt.Println("==================================")
	fmt.Println("initialize var")
	fmt.Println("user name : ", userName)
	fmt.Println("user age : ", userAge)
	fmt.Println("==================================")
	//another way to init var
	kota := "Banyuwangi"
	lama := 10
	fmt.Println("==================================")
	fmt.Println("another way to initialize var")
	fmt.Printf("Agus tinggal di kota %s dan sudah tinggal disana selama %d tahun", kota, lama)
	fmt.Println("==================================")

	//init multiple var
	ayam, bebek := "ayam", "bebek"
	fmt.Println("==================================")
	fmt.Printf("makanan kesukaan agus adalah %s dan %s \n", ayam, bebek)
	fmt.Println("==================================")

	//init array and slice

	arr1 := [2]string{"Anggur", "Jeruk"}
	arr2 := [...]int{1, 2, 3, 4, 5}
	slice1 := []int{1, 2, 3, 4, 5}
	fmt.Println("==================================")
	fmt.Println("Latihan array dan slice")
	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(slice1[0:2])
	fmt.Println("===================================")

	//init map
	var myMap = map[string]string{
		"Name": "Agus",
		"Age":  "20",
	}
	var myMap2 = make(map[string]string)
	myMap2["Name"] = "Bejo"
	myMap2["Age"] = "13"
	fmt.Println(myMap)
	fmt.Println(myMap2)
	//loop test
	fmt.Println("===================================")
	fmt.Println("Looping")
	for i := 0; i < 100000; i++ {
		for j := 0; j < 100000; j++ {
			// fmt.Println(i, j)
		}
	}
	//loop array
	arr12 := [2]string{"Agus", "Bejo"}

	for i, v := range arr12 {
		fmt.Println(i, v)
	}

	// init function
	fmt.Println("===================================")
	fmt.Println("Init Function")
	a := add(1, 2)
	fmt.Print(a)
	sayHello()
	resut := countRecursive(10)
	fmt.Println(resut)
	fmt.Println("===================================")

	// defer, panic, recover
	fmt.Println("Defer, Panic, Recover")
	handlePanic := func() {
		r := recover()
		if r != nil {
			fmt.Println("Recovered from panic", r)
		}
	}
	Division := func(a, b int) int {
		defer handlePanic()
		if b == 0 {
			panic("Panic: Division by zero")
		}
		return a / b
	}
	fmt.Println(Division(3, 2))
	fmt.Println(Division(1, 0))
	fmt.Println("===================================")

	// struct
	var user User
	user.Id = 1
	user.Name = "Ayam geprek"
	user.Email = "ayam_geprek@gmail.com"
	user.Pass = "ayamdigepreksampehalus123"
	fmt.Println("user id : ", user.Id)
	fmt.Println("user name : ", user.Name)
	fmt.Println("user email : ", user.Email)
	fmt.Println("user password : ", user.Pass)

	// fmt.Println(user.Id, user.Name, user.Email, user.Pass)
}
