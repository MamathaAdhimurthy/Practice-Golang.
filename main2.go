package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	//creating Map

	// var m_a_p = map[int]string{
	// 	90: "Dog",
	// 	91: "Cat",
	// 	92: "Cow",
	// }

	// a := map[int]string{
	// 	90: "Dog",
	// 	91: "Cat",
	// 	92: "Cow",
	// }

	//USING MAKE FUNCTION CREATE MAP

	// var a = make(map[int]string)
	// a[1] = "sri"
	// a[2] = "sai"
	// a[3] = "ram"

	// fmt.Println("map:", a)

	// b := make(map[int]string)
	// b[1] = "sri"
	// b[2] = "sai"
	// b[3] = "ram"

	// fmt.Println("map:", b)

	//CREATING EMPTY MAP

	// var a map[string]string
	// var b = make(map[string]string)

	// fmt.Println(a == nil)
	// fmt.Println(b == nil)

	//ITERATE OVER AN MAP
	// var m_a_p = map[int]string{
	// 	90: "Dog",
	// 	91: "Cat",
	// 	92: "Cow",
	// }

	// for pet, ok := range m_a_p {
	// 	fmt.Println(pet, ok)
	// }

	//add key value pair in map

	// m_a_p[93] = "Monkey"
	// fmt.Println(m_a_p)

	//update key value pair

	// m_a_p[91] = "Horse"
	// fmt.Println(m_a_p)

	//retrive the value relayedto key
	// val1 := m_a_p[90]
	// fmt.Println(val1)

	// check the existence of the key in the map

	// pet, ok := m_a_p[90]
	// fmt.Println("key value present or not", ok)
	// fmt.Println("value:", pet)

	// _, ok1 := m_a_p[91]
	// fmt.Println(ok1)

	//DELETE KEY FROM THE MAP
	// delete(m_a_p, 90)
	// fmt.Println(m_a_p)

	//MODIFYING THE MAP
	// fmt.Println("before modifying", m_a_p)
	// new_map := m_a_p
	// new_map[90] = "DOG"
	// new_map[93] = "Horse"

	// fmt.Println("after modifying", m_a_p)
	// fmt.Println("after modifying new map", new_map)
	// fmt.Println(len(m_a_p))

	//CONVERTING MAP TO JSON

	a := make(map[int]string)
	a[1] = "Sri"
	//convert map to json
	j, err := json.Marshal(a)
	if err != nil {
		fmt.Println("Error:%s", err.Error())

	} else {
		fmt.Println(string(j))
	}

	//CONVERT JSON TO MAP

	var a1 map[int]string
	json.Unmarshal(j, &a1)
	fmt.Println(a1)

	//convert a map to a JSON where we have a struct for the value in the map. Below is the code for that
	type employee struct {
		Name string
	}

	b := make(map[string]employee)
	b["1"] = employee{Name: "Sai"}

	k, err := json.Marshal(b)
	if err != nil {
		fmt.Println("error :%s", err.Error())

	} else {
		fmt.Println(string(k))
	}

	//CONVERT JSON TO MAP

	var b1 map[int]employee
	json.Unmarshal(k, &b1)
	fmt.Println(b1)

}
