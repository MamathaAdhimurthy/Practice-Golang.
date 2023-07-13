package main

import (
	myfunc "example.com/hello/myfun"
	_ "github.com/go-sql-driver/mysql"
)

type person struct {
	name   string
	gender string
	age    int
}

// type Car struct {
// 	Name    string
// 	Age     int
// 	ModelNo int
// }

// //	func add() {
// //		var a int = 10
// //		var b int = 20
// //		var c int = a + b
// //		fmt.Println(c)
// //	}

// func main() {
// 	// a := 10
// 	// b := 20
// 	// //myfun.Sum(a, b)
// 	// result := myfunc.Sum(a, b)
// 	// fmt.Println(result)

// 	// //for loop
// 	// for i := 0; i < 5; i++ {
// 	// 	//fmt.Println(i)
// 	// 	if i == 3 {
// 	// 		//continue //skips the value 3 continues the execution
// 	// 		break
// 	// 	}

// 	// 	fmt.Println(i)
// 	// }

// 	//while loop(in go we can not use while loop but we can use it by using for loop)

// 	// i := 0
// 	// for i < 5 {
// 	// 	fmt.Println("It is a while loop")
// 	// 	i++
// 	// }

// 	// Nested for loop

// 	// adj := []string{"big", "tasty"}
// 	// fruits := []string{"apple", "orange", "bannana"}
// 	// for i := 0; i < len(adj); i++ {
// 	// 	for j := 0; j < len(fruits); j++ {
// 	// 		fmt.Println(adj[i], fruits[j])
// 	// 	}
// 	// }

// 	//switch case

// 	//day := 4
// 	// var array1 = [3]int{1, 2, 3}
// 	// switch array1[0] {
// 	// case 1:
// 	// 	fmt.Println("Monday")
// 	// case 2:
// 	// 	fmt.Println("tuesday")
// 	// case 3:
// 	// 	fmt.Println("wednesday")

// 	// case 4:
// 	// 	fmt.Println("Thursday")
// 	// case 5:
// 	// 	fmt.Println("friday")
// 	// default:
// 	// 	fmt.Println("not a week day")
// 	// }

// 	//strings
// 	// 	m1 := "my name"
// 	// 	m2 := "name"
// 	// 	fmt.Println(strings.Contains(m1, m2))
// 	// 	fmt.Println(strings.ReplaceAll(m1, "m", "Mamatha"))
// 	// 	fmt.Println(strings.Split(m1, " "))
// 	// 	fmt.Println(m1 + " " + m2)

// 	// m1, m2 := 2, 3
// 	// fmt.Println(m1, m2)
// 	// myfunc.Swap(&m1, &m2)
// 	// fmt.Println(m1, m2)

// 	//Structure example two ways u can declare a structure
// 	//c := Car{"Mamatha", 23, 45}
// 	c1 := Car{
// 		Name:    "mamatha",
// 		Age:     23,
// 		ModelNo: 90,
// 	}
// 	c1.Print()
// 	c1.Drive()
// 	fmt.Println(c1.GetName())
// 	//var c2 Ccar
// 	// c2.Name = "Dhanuja"
// 	// c2.Age = 23
// 	// c2.ModelNo = 89
// 	// fmt.Println(c2.Name)
// 	// fmt.Println(c)
// 	// fmt.Println(c1)
// 	//myfunc.Print()
// 	// Print(c1)
// 	// Print(c)
// 	// Print(c2)
// 	// fmt.Println(GetName(c2))

// }

// //	func Print(cs Car) {
// //		fmt.Println("c:", cs.Age)
// //	}
// //

//	func (c1 Car) Print() {
//		fmt.Println(c1)
//	}
//
//	func (c1 Car) Drive() {
//		fmt.Println("Driving")
//	}
//
//	func (c1 Car) GetName() string {
//		return c1.Name
//	}
// type Cars interface {
// 	Drive()
// 	Stop()
// }

// type Lambo struct {
// 	LamboModel string
// }

// func NewModel(arg string) Cars {
// 	return &Lambo{arg}
// }

// type Chevy struct {
// 	ChevyModel string
// }

// func (l *Lambo) Drive() {
// 	fmt.Println("lambo is moving")
// 	fmt.Println(l.LamboModel)
// }
// func (l *Lambo) Stop() {
// 	fmt.Println("stopping")
// }
// func (c *Chevy) Drive() {
// 	fmt.Println("Chevy is moving")
// 	fmt.Println(c.ChevyModel)
// }

// func main() {
// 	l := Lambo{"Gallardo"}
// 	c := Chevy{"c345"}
// 	l.Drive()
// 	c.Drive()
// }

//	func Anything(anything interface{}) {
//		fmt.Println("Anything")
//	}

// GO ROUTINES

// func heavy() {
// 	for {
// 		time.Sleep(time.Second * 1)
// 		fmt.Println("Heavy")
// 	}

// }
//
//	func superHeavy() {
//		for {
//			time.Sleep(time.Second * 2)
//			fmt.Println("superHeavy")
//		}
//	}
func main() {
	// go heavy()
	// go superHeavy()
	// fmt.Println("fin")
	// time.Sleep(time.Second * 5)
	// //select {} // run infinite times hello
	// Anything(2.44)
	// Anything("hiii")
	// Anything(struct{}{})
	// mymap := make(map[string]interface{})
	// mymap["name"] = "asd"
	// mymap["age"] = 10
	// fmt.Println(mymap)
	// 	flag := true
	// 	for i := 1; i <= 10; i++ {
	// 		if i%2 == 0 {
	// 			flag = false
	// 			break
	// 		} else if i == 1 {
	// 			continue
	// 		}

	// 	}
	// 	fmt.Println(flag)
	// }

	////REST APIS
	// mux := http.NewServeMux()
	// mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Println("Request received")
	// 	fmt.Println(r.Method)
	// 	w.Write([]byte("Hello World"))
	// })

	// http.ListenAndServe("localhost:3001", mux)

	// cCONNECTING TO DATA BASE
	// mux := controller.Register()
	// db := model.Connect()
	// defer db.Close()
	// fmt.Println("Serving......")

	// log.Fatal(http.ListenAndServe("localhost:3002", mux))

	//STRINGS
	//myfunc.NumberOrNot()
	//myfunc.RemoveWhiteSpace()
	//myfunc.Multiline()
	//myfunc.CompareString()
	//myfunc.StringContainsOrNot()
	//myfunc.SplitString()
	//myfunc.AllWords()
	//myfunc.StrJoin()
	//myfunc.StrPrefix()
	//myfunc.StringLower()
	//myfunc.StringTrim()
	//myfunc.StringReplace()
	//myfunc.Iiff()

	//CLOUSERS
	// modulus := myfunc.GetModulus()
	// modulus(-1)
	// modulus(2)
	// fmt.Println(modulus(-5))

	// myfunc.Print(myfunc.Area, 2, 4)
	// myfunc.Print(myfunc.Sum1, 2, 4)

	//Return a function from function

	// areaF := myfunc.GetAreaFunc()
	// res := areaF(4, 2)
	// fmt.Println(res)

	// PASS VARIABLE NUMBER OF ARGUMENTS TO A FUNCTION
	// Different number of parameters but of the same type
	// fmt.Println(myfunc.Add(1, 2))
	// fmt.Println(myfunc.Add(1, 2, 3))
	// fmt.Println(myfunc.Add(1, 2, 3, 4))

	//Different number of parameters and of different types
	myfunc.Handle(1, "abc")
	myfunc.Handle("abc", "xyz", 3)
	myfunc.Handle(1, 2, 3, 4)
}
