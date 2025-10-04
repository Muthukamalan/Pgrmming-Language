// package declaration
package main

// format packge necessary for IO operations like scanf, printf, Println
import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"unicode/utf8"
)

// var stmnt at package level, we can't use implicit construct
// c := true <NOT Usable>
var c, python, java bool = true, false, true

var raiseErrPurposefully = errors.New("Raise error!")

type Rectangle struct {
	width, height float64
}
type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * math.Sqrt(c.radius)
}
func (c Circle) Perimeter() float64 {
	return math.Pi * 2 * c.radius
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}
func (r Rectangle) Perimeter() float64 {
	return (r.width + r.height) * 2
}

type Shape interface {
	Area() float64
	// Volume() float64
}

func CalculateArea(s Shape) float64 {
	return s.Area()
}

type Measurable interface {
	Perimeter() float64
}

// Embedded Interface or Composistion of Interface
type Geometry interface {
	Shape      // Area()
	Measurable // Perimeter()
}

func describeShape(g Geometry) {
	fmt.Println("Area", g.Area())
	fmt.Println("Perimeter", g.Perimeter())
}

// func main must have no arguments and no return values
func main() {
	fmt.Println("muthu")

	/*
		VARIABLE DECLARATION
		1. var language string
		2. var language string = "golang"
		3. var language = "golang"
		4. language := "golang"
	*/

	// Integers
	var i int = -128
	var i8 int8 = 127
	var i16 int16 = -32768
	var i32 int32 = 2147483647
	var i64 int64 = -9223372036854775808

	var u uint = 255
	var u8 uint8 = 255
	var u16 uint16 = 65535
	var u32 uint32 = 4294967295
	var u64 uint64 = 18446744073709551615

	// Float 32
	var f32 float32 = 1234567890.123456789
	var f64 float64 = 1234567890.123456789

	// Boolean
	var active, close bool = true, false

	// complex
	var ij complex64 = complex(12, 32)

	// string
	var language string = "golang"

	// constants: you cannot declare a constant that can only be computed at run-time like you can in JS
	const PI float32 = 3.141
	const (
		DECIMAL     = 255  // decimal with no prefix
		OCTAL       = 0377 // octal with leading 0
		HEXADECIMAL = 0xff // hexa with leading 0x
	)

	fmt.Println(i, i8, i16, i32, i64)
	fmt.Println(u, u8, u16, u32, u64)
	fmt.Println("Float32", f32)
	fmt.Println("Float64", f64)
	fmt.Println("boolean", active, close)
	fmt.Println("complex", ij)
	fmt.Println("language", language)
	fmt.Println("PI", PI)
	fmt.Println("Decimal: ", DECIMAL, "Octal: ", OCTAL, "Hexadecimal: ", HEXADECIMAL)

	// static type
	fmt.Printf("static type declaration, i8 is %T\n", i8)
	// dynamic type
	fmt.Printf("dynamic type declaration, i is %T\n", i)
	// mixed type
	var name, age, marriage_status = "muthu", 24, false
	fmt.Printf("%s of age %v and his/her marriage status %t\n", name, age, marriage_status)

	/*
		%s: string
		%q: quoted string
		%d: int
		%f: float
		%e: float
		%E: float
		%T: type
		%t: boolean
		%v: Struct
		%+v: Struct
		%#v: Struct
		%c: character
	*/

	// string concatenation
	fmt.Println("ba" + strings.Repeat(`na`, 2))

	// Logical Operators
	fmt.Println("A && B: ", active && close)
	fmt.Println("A || B: ", active || close)
	fmt.Println("!A: ", !active)

	// Assignment Operators
	A := 10
	B := 20

	A += B
	fmt.Println("A += B: ", A)

	A -= B
	fmt.Println("A += B: ", A)

	// Bitwise operator

	// Pointer
	/*
		| Pass by Value | Pass by Reference |
		|---------------|-------------------|
		|	int 		|	slices			|
		|	float		|	maps			|
		|	string		|	channels		|
		|	bool		|	pointers		|
		|	struct		|	function		|
	*/
	X := 10
	ptrX := &X // reference
	fmt.Println("Address of A:", ptrX)
	fmt.Println("Value of A:", *ptrX) // dereference

	var l1 int
	var l2 int
	var l3 = &l1 // sharing memory

	l1 = 100
	l2 = l1 // copy the data

	fmt.Println("before changing")
	fmt.Println(l1, l2, *l3, l3)

	l1 = 200
	fmt.Println(`After changing`)
	fmt.Println(l1, l2, *l3, l3)

	// Control Flow :: If Else
	if age > 18 {
		if age > 60 {
			fmt.Println("over aged!!")
		} else {
			fmt.Println("major, legal to major to marry to your love!")
		}
	} else {
		fmt.Println("minor illegal to married!")
	}

	// Control Flow:: For
	for i := 0; i <= 10; i++ {
		for j := 1; j < 10; j++ {
			fmt.Printf("%d * %d = %d\t", i, j, i*j)
		}
		fmt.Println()
	}

	for i := range 10 {
		if i%2 == 0 {
			fmt.Println("continued!!")
			continue
		}
		if i == 5 {
			fmt.Println("breaked!!")
			break
		}
		fmt.Println(i)
	}

	j := 0
start:
	for j < 10 {
		fmt.Println(j)
		j++
		goto start
	}

	/*
	   start_again:
	   	fmt.Println("Infinite loop")
	   	goto start_again
	   	for {
	   		fmt.Println("infinite loop")
	   	}
	*/

	fmt.Printf("%d + %d : = %d\n", A, B, add(A, B))
	fmt.Printf("before swap:: (%s,%s)\n", name, language)
	name, language = swap(name, language)
	fmt.Printf("after swap:: (%s,%s)\n", name, language)

	// Function Arguments
	// Call by value
	fmt.Println(strings.Repeat(`*`, 15) + " Call by Value " + strings.Repeat(`*`, 15))
	fmt.Println("age before covid: ", age)
	increment(age)
	fmt.Println("age after covid: ", age)

	// Call by reference
	fmt.Println(strings.Repeat(`*`, 10) + " Call by Reference " + strings.Repeat(`*`, 15))
	fmt.Println("age before work: ", age)
	decrement(&age)
	fmt.Println("age after work: ", age)
	fmt.Println(strings.Repeat(`*`, 45))

	// Arrays:: type-specific and contiguous memory
	// var balances [5]float32
	// var balances = [5]float32{10.0, 20, 30, 40, 50}
	var balances = []float32{10.0, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	for i := 0; i < len(balances); i++ {
		fmt.Printf("Elements in balances: %f\n", balances[i])
	}

	// Slice: more-dynamic
	var slic []float32
	fmt.Printf("Empty slice: %v\n", slic) // [](nil)
	slic1 := balances[1:5]
	slic2 := balances[6:7]
	slic3 := append(slic2, 23.232)
	fmt.Println(slic1)
	fmt.Println(slic2)
	fmt.Println(slic3)
	slic4 := make([]float32, len(balances), cap(balances)*2) // len: length of elements and cap:maximum size of array can reference
	fmt.Println(slic4)

	// Map:: Not-Ordered, reference type
	var m map[string]int
	fmt.Printf("empty map:: %v\n", m)
	m = map[string]int{"score1": 10, "score2": 20, "score3": 30}
	m["score2"] = 200
	fmt.Println(m)
	fmt.Println(m["score2"])
	v, ok := m["score10"]
	delete(m, "score2")
	fmt.Println(v, ok)

	// Struct:: value type, heterogeneous define type before using it
	var muthu studentDetails
	fmt.Printf("struct:: %#v\n", muthu)
	var kamalan = studentDetails{name: "kamalan", id: 20}
	fmt.Printf("struct:: %#v\n", kamalan)
	fmt.Println(kamalan)

	// Range
	for balance := range balances {
		fmt.Println("balance ", balance)
	}

	// Scopes
	/*
		1. Local scopes
		2. global scopes
	*/
	fmt.Printf("global variables: c: %t, python: %t, java: %t\n", c, python, java)
	python = true
	fmt.Printf("shadowing global variable: python: %t\n", python)

	// strings
	native_language := "tamil"
	fmt.Printf("my favourite language is: \"%s\"\n", native_language)
	fmt.Printf("len of fav language is: %v\n", len(native_language))

	// display hexadecimal byte values of the string
	for i := 0; i < len(native_language); i++ {
		fmt.Printf("hexabyte:%x\n", native_language[i])
	}

	// Recursion
	fmt.Println("Factorial: ", factorial(5))

	// Type casting
	fmt.Printf("format %T and %f\n", float32(age), float32(age))
	wrong_string := "123"
	num, err := strconv.Atoi(wrong_string)
	if err != nil {
		fmt.Println("error convertion: ", err)
		return
	}
	fmt.Println("converted number: ", num)

	// Error handling
	if value, err := sqrt(25); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("value of sqrt: ", value)
	}

	// RUNE
	/*
		strings are just seq of bytes: they can hold arbitary data.
		1. when you need to work with individual charc in a string, you should you `rune` tyoe.
		2. Go will handle emoji, chinese charc natively. (1 emoji  = 1 rune = 4 bytes)
	*/
	fmt.Printf("constant 'name' byte length: %d\n", len(name))
	fmt.Printf("constant 'name' rune length: %d\n", utf8.RuneCountInString(name))
	fmt.Println("=====================================")
	fmt.Printf("Hi %s, \n", name)

	var operating_sys = "linux"
	fmt.Println(getCreator(operating_sys))

	// Anonymous Function
	func() {
		fmt.Println("WELCOME TO ANONYMOUS FUNCTIONS IN GOLang")
	}()

	anonymous := func(x int) int {
		fmt.Println("WELCOME TO ANOTHER ANONYMOUS FUNCTIONS IN GOLang")
		return x
	}
	anonymous(2)

	//
	revathi := person{
		firstname: "muthu",
		lastname:  "revathi",
		contactInfo: contactInfo{
			email:   "muthurevathi@muthu.com",
			zipcode: 123456,
		},
	}
	fmt.Printf("%#v\n", revathi)
	revathi.updateNamevalue("setta-kari")
	fmt.Printf("%#v\n", revathi)

	/*
		revthiAddress := &revathi
		revthiAddress.updateNamevRef("passa-kara")
		fmt.Printf("%#v\n", revathi)
	*/
	revathi.updateNamevRef("passa-kara")
	fmt.Printf("%#v\n", revathi)

	// Closure
	nextInt := intSeq()
	fmt.Println("Closure: ", nextInt())
	nextInt()
	nextInt()
	fmt.Println("Closure: ", nextInt())

	// INTERFACE
	// acts as an abstraction of a class.
	bot1 := englishBot{}
	bot2 := spanishBot{}
	printGreeting(bot1)
	printGreeting(bot2)
	bot1.getGreeting()
	bot2.getGreeting()

	rect := Rectangle{width: 5, height: 5}
	cir := Circle{radius: 3}
	fmt.Printf("%#v\n", rect)
	fmt.Printf("%#v\n", cir)
	fmt.Println(cir.Area())
	fmt.Println(CalculateArea(cir))
	fmt.Println(CalculateArea(rect))

	describeShape(cir) // describeShape is a polymorphism taking both Rectangle,Circle. It's achieved by common Gemetry. to satisfy gemetry interface needs "Area" & "perimeter"
	describeShape(rect)

	fnFoo()
	fnBar()

	// "defer" statement invokes a fn whose execution is deferred to the moment the surrounding fn returns, either because
	// - the surrounding fn executed a return statement
	// - reached the end of its fn body
	// - or because the corresponding goroutines is panicking
	defer fnFoo()
	fnBar()

	if err := raiseErr(); err != nil {
		switch err {
		case raiseErrPurposefully:
			fmt.Println("manual error raised!")
		default:
			fmt.Println("unknown errors!")
		}

	}

	//
	f, err := os.Create("output.txt")
	if err != nil {
		log.Fatalf("error %s", err)
	}
	defer f.Close()

	s := []byte("Helllo gophers!")
	_, err = f.Write(s)
	if err != nil {
		log.Fatalf("error %s", err)
	}

	// functions were 1st class citizen
	fact := factorial
	fmt.Println(fact(12))

	Foobar("muthu")

	// anonymous fn
	// polymorphism
	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")

	// callback function
	fmt.Printf("%T\n", madd)
	fmt.Printf("%T\n", masubstract)
	fmt.Printf("%T\n", doMath)
	fmt.Printf("do Math add:: %d\n", doMath(42, 32, madd))
	fmt.Printf("do Math add:: %d\n", doMath(42, 32, masubstract))

	// closure
	fx := incrementor()
	fmt.Println("increamentor called ", fx())
	fmt.Println("increamentor called ", fx())
	fmt.Println("increamentor called ", fx())
	fmt.Println("increamentor called ", fx())
	fmt.Println("increamentor called ", fx())
	fmt.Println("increamentor called ", fx())

	fmt.Printf("addI (1,2) %T\n", gaddI(1, 2))
	fmt.Printf("addF (1,2)%T\n", gaddF(1.0, 1232.23))
	fmt.Printf("addF (1,2) %T\n", gaddT(10, 32))

	var alias_num muthuInt = 1
	fmt.Printf("addF (1,2) %T\n", gaddT(alias_num, 32.0))

	ints := map[string]int{"first": 34, "second": 12}
	floats := map[string]float64{"first": 34.0, "second": 12.0}

	fmt.Printf("Non-Generic Sums: %v and %v\n", sumInt(ints), sumFloats(floats))
	fmt.Printf("Generic Sums: %v and %v\n", sumIntorFloat[string, int](ints), sumIntorFloat[string, float64](floats))
	fmt.Printf("Generic Sums, type parameters inferred: %T and %T\n", sumIntorFloat(ints), sumIntorFloat(floats))

	groups := ColorGroup{
		ID:     1,
		Name:   "Reds",
		Colors: []string{"Crimsson", "Red", "Ruby", "Maroon"},
	}
	if b, err := json.Marshal(groups); err == nil {
		os.Stdout.Write(b)
		println()

	}

	var jsonBlob = []byte(`[{"code":"golang","ID":1,"Name":"Reds","Colors":["Crimsson","Red","Ruby","Maroon"]},{"ID":2,"Name":"Blues","Colors":["Crimsson","Red","Ruby","Maroon"], "err":1}]`)
	var color_groups []ColorGroup
	err = json.Unmarshal(jsonBlob, &color_groups)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%#v\n", color_groups)

}


// jsonP
type ColorGroup struct {
	ID     int
	Name   string
	Colors []string
}

// tutorials GOlang
func sumIntorFloat[K string, V int | float64](m map[K]V) V {
	var s V
	for _, val := range m {
		s += val
	}
	return s
}

func sumInt(m map[string]int) int {
	s := 0
	for k, v := range m {
		fmt.Println(v, k)
		s += v
	}
	return s
}
func sumFloats(m map[string]float64) float64 {
	s := 0.
	for _, v := range m {
		s += v
	}
	return s
}

// Type constraint ( contraining which type should take )
type muthuInt int
type MuthuNumber interface{ ~int | ~float64 } // include all types of int underlying structure

func gaddT[X MuthuNumber](a, b X) X { return a + b }

// Generics
func gaddI(a, b int32) int32     { return a + b }
func gaddF(a, b float32) float32 { return a + b }

// func gaddT[X int | float64](a, b X) X { return a + b }

// closure
func incrementor() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}

// callBack
func doMath(a, b int, f func(int, int) int) int { return f(a, b) }
func masubstract(a, b int) int                  { return a - b }
func madd(a, b int) int                         { return a + b }

// switch
func Foobar(s string) {
	switch s {
	case "foo":
		fnFoo()
	case "bar":
		fnBar()
	default:
		fmt.Println("fn called FooBar!!")
	}
}

func raiseErr() error {
	return raiseErrPurposefully //errors.New("Raise error!")
}

// ///////////////////
func fnFoo() {
	fmt.Println("print fn foo!")
}
func fnBar() {
	fmt.Println("print fn Bar!")
}

// ///////////////////////////////////////////////////////////////////////////////////////////////////////
func factorial(n int) int {
	if n == 0 {
		return 1
	} else {
		return n * factorial(n-1)
	}
}

func decrement(nums *int) int {
	*nums--
	fmt.Println("inside function number inc:", *nums)
	return *nums

}

func add(x int, y int) int {
	fmt.Println("given numbers are", x, "and", y)
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func increment(no int) int {
	no++
	fmt.Println("inside function number inc:", no)
	return no
}

type studentDetails struct {
	name string
	id   int
}

type error interface {
	Error() string
}

func sqrt(val float64) (float64, error) {
	if val < 0 {
		return 0, errors.New("negative number passed to sqrt")
	}
	return math.Sqrt(val), nil
}

func getCreator(os string) string {
	var creator string
	switch os {
	case "linux":
		creator = "Linus Torvalds"
	case "windows":
		creator = "Bill Gates"
	case "mac":
		creator = "A Steve"
	default:
		creator = "Unknown"
	}
	return creator
}

//	type person struct {
//		firstname string
//		lastname  string
//		contact   contactInfo
//	}
type person struct {
	firstname string
	lastname  string
	contactInfo
}
type contactInfo struct {
	email   string
	zipcode int
}

func (p person) updateNamevalue(newname string) {
	p.firstname = newname
	fmt.Printf("%#v", p)
}

func (address_of_person *person) updateNamevRef(newname string) {
	(*address_of_person).firstname = newname
}

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

// INTERFACE
// acts as an abstraction of a class.
// collection of method signatures
type englishBot struct{}
type spanishBot struct{}

type Bot interface {
	/*
		if you've fn getGreeting() and returns strings you're consider as bot
	*/
	getGreeting() string
}

func (englishBot) getGreeting() string {
	return "Hello, there!!"
}
func (spanishBot) getGreeting() string {
	return "Hola!!"
}
func printGreeting(b Bot) {
	fmt.Println(b.getGreeting())

}

// func printGreeting(eb englishBot) { fmt.Println(eb.getGreeting()) }
// func printGreeting(sb spanishBot) { fmt.Println(sb.getGreeting()) }

// anoymouse func
// func (parameter(s)) return(s) {}(argmunets(s))
