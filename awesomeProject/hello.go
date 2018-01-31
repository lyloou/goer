/*
Package regexp implements a simple library for regular expressions.

The syntax of the regular expressions accepted is:

    regexp:
        concatenation { '|' concatenation }
    concatenation:
        { closure }
    closure:
        term [ '*' | '+' | '?' ]
    term:
        '^'
        '$'
        '.'
        character
        '[' [ '^' ] character-ranges ']'
        '(' regexp ')'
*/

package main

import (
	"fmt"
	"reflect"
	"regexp"
	"runtime"
	"strconv"
	"time"
	"math/rand"
	"encoding/json"
	"os"
	"io/ioutil"
	"path/filepath"
)

type Point struct {
	Px int
	Py int
}
type Circle struct {
	Radius int
	Point
}

type Rect struct {
	width  int
	height int
	Point
}

func main() {

	sampleS()
}
func sampleS() {
	fmt.Println(os.Getwd())
}
func sampleR() {
	point := new([]Point)
	fmt.Println("point", point)
	p := make([]Point, 0)
	fmt.Println("p,", p)
	var pp Point
	pp.Px = 32
	ppp := interface{}(pp)
	fmt.Println("pp", ppp.(Point))
	fmt.Println(reflect.TypeOf(ppp))
	switch ppp.(type) {
	case Point:
		fmt.Println("Point")
	default:
		fmt.Println("default")
	}
	var okk map[string]int
	fmt.Println("okk", okk["f"])
	var ook2 interface{}
	fmt.Println("ook2", ook2)
}

func sampleP() {
	s, e := filepath.Rel(".", "../web/socket.go")
	fmt.Println(s, e)
	file, err := ioutil.ReadFile("/web/socket.go")
	fmt.Println(string(file), err)
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)
	fmt.Println(exPath)
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(dir)
	var c = 2.01
	j := float64(int(100*c)) / 100
	fmt.Println(j)
	getenv, _ := os.Getwd()
	fmt.Println(getenv)
	file, err = ioutil.ReadFile("../hello/a1.go")
	fmt.Println(string(file), err)
	k := c * 100
	fmt.Println(int64(k))
}

func sampleQ() {
	type o struct {
		a string
		b string
		c int64
	}
	oo := o{
		a: "nihao",
		b: "nikkkk",
		c: 322,
	}
	ooj, err := json.Marshal(&oo)
	if err != nil {
		fmt.Println("ERROR:", err)
	}
	fmt.Println("DATA:", ooj)
	fmt.Println("DATA:", string(ooj))
	data := fmt.Sprintf("%+v", oo)
	fmt.Println("data:", data)
}

func sampleF() {
	var s string = "this is a string"
	fmt.Println(s)
	var b []byte
	b = []byte(s)
	fmt.Println(b)
	for i := range b {
		fmt.Print(string(b[i]))
	}
	s = string(b)
	fmt.Printf("\n%v", s)
	source := rand.NewSource(1)
	fmt.Println("source.Int63()", source.Int63())
}

func sampleE() {
	P := new(person)
	fmt.Println(*P)
	P.name = "LiLou"
	P.age = 19
	fmt.Println(*P)
	fmt.Println()
	var p Provide
	p = &Man{name: "Man's Name "}
	p.Str("wo")
	p.Age(89)
	fmt.Println(p)
	Set(p)
	p.Str("wo")
	p.Age(89)
	fmt.Println(p)
}

func Set(p Provide) string {
	p.Str("wok")
	p.Age(32)
	return "ok"
}

type Man struct {
	name string
}

func (man Man) Str(str string) {
	fmt.Println(man.name, str)
}
func (man Man) Age(age int64) {
	fmt.Println(man.name, age)
}

type Provide interface {
	Str(str string)
	Age(age int64)
}

func sampleD() {
	value := 32 << 20
	fmt.Println(value)
	lou := new(Lou)
	fmt.Println(lou)
	build := lou.Build("什么？")
	fmt.Println(lou, build)
	fmt.Println(lou, build, lou.Lname())
}

func (lou *Lou) Lname() string {
	return lou.name + " dkk"
}

type Lou struct {
	name string
	age  int
}

func (lou *Lou) Build(str string) string {
	lou.name = "盖楼了"
	lou.age = 28
	return str
}
func sampleC() {
	c := make(chan int)
	go func() {
		c <- 3
	}()
	b := <-c
	fmt.Println(b)
}

func sampleB() {
	fmt.Println(strconv.Atoi("3"))
	fmt.Println(regexp.MatchString("^[0-9]+$", "3"))
	fmt.Println(regexp.MatchString("^\\p{Han}+$", " 中国"))
}

func sampleA() {
	var e *Err
	fmt.Println(e == nil)
	e = new(Err)
	e.err = "221"
	e.Print()
	//e.Printf()
	e.Println()
	var err error
	err = returnErr()
	fmt.Println(err == nil)
}

func returnErr() *Err {
	return nil
}

type Err struct {
	err string
}

func (e *Err) Error() string {
	return e.err
}
func (e *Err) Print() {
	fmt.Println("e.Print run")
}

func (e *Err) Printf() {
	fmt.Println(e.err)
}

func (e Err) Println() {
	fmt.Println("e.Println run")
}

func sample9() {
	runtime.GOMAXPROCS(2)
	println(runtime.NumGoroutine())
	println(runtime.NumCPU())
	c := make(chan int, 3)
	o := make(chan bool)
	c <- 3
	go func() {
		for {
			select {
			case v := <-c:
				println(v)
			case <-time.After(5 * time.Second):
				println("timeout")
				o <- true
				break
			}
		}
	}()
	<-o
}
func sample8() {
	c := make(chan int, 3)
	c <- 2
	c <- 1
	select {
	case <-c:
		fmt.Println(<-c)
	}
	//fmt.Println(<-c)
}
func sample7() {
	a := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int)
	go sum(a[:len(a)/2], c)
	go sum(a[len(a)/2:], c)
	x, y := <-c, <-c
	fmt.Println("x=", x, " y=", y, " x+y=", x+y)
}

func sum(a []int, c chan int) {
	total := 0
	for _, value := range a {
		total += value
	}
	c <- total
}

func sample6() {
	go say("world")
	say("hello")
}

func say(s string) {
	for i := 0; i < 5; i++ {
		runtime.Gosched()
		fmt.Println(s)
	}
}
func sample5() {
	var x float64 = 3.4
	v := reflect.ValueOf(&x)
	fmt.Println(v.Type())
	fmt.Println(v.Kind())
	v.Elem().SetFloat(9.8)
	fmt.Println(x)
	//fmt.Println(v.Elem().Field(0).String())
}

func sample4() {
	list := make(List, 3)
	list[0] = 1
	// an int
	list[1] = "Hello"
	// a string
	list[2] = Person{"Dennis", 70}
	for index, element := range list {
		if value, ok := element.(int); ok {
			fmt.Printf("list[%d] is an int and its value is %d\n", index, value)
		} else if value, ok := element.(string); ok {
			fmt.Printf("list[%d] is a string and its value is %s\n", index, value)
		} else if value, ok := element.(Person); ok {
			fmt.Printf("list[%d] is a Person and its value is %s\n", index, value)
		} else {
			fmt.Printf("list[%d] is of a different type\n", index)
		}
	}
}

type Element interface{}
type List []Element

type Person struct {
	name string
	age  int
}

//定义了String方法，实现了fmt.Stringer
func (p Person) String() string {
	return "(name: " + p.name + " - age: " + strconv.Itoa(p.age) + " years)"
}

func sample3() {
	mark := Student{Human{"Mark", 25, "888,888,888"}, "High School"}
	mark.SayHi()
	jack := Employee{Human{"Jack", 32, "223,223,223"}, "QQ"}
	jack.SayHi()
}

type Human struct {
	name  string
	age   int
	phone string
}

type Student struct {
	Human
	school string
}

type Employee struct {
	Human
	company string
}

func (h *Human) SayHi() {
	fmt.Printf("Human, I'm %s you can call me on %s \n", h.name, h.phone)
}

func (h *Employee) SayHi() {
	fmt.Printf("Employee, I'm %s you can call me on %s \n", h.name, h.phone)
}

func smaple2() {
	P := new(person)
	P.name = "LiLou"
	P.age = 19
	fmt.Println(*P)
}

func sample1() {
	fmt.Println("hello, world")
	// dkfje
	fmt.Println("hello")
	// kehhkj
	fmt.Println("hello, world dkfjkej")
	// kek
	x := 1
	// jek
	y := 2
	// yyy
	iiiii := x<<8 + y<<16
	//kjke
	fmt.Println(iiiii)
}

type person struct {
	name string
	age  int
}
