// if

package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

// コンストラクタ
func NewPerson(name string, age int) *Person {
	p := new(Person)
	p.Name = name
	p.Age = age
	return p
}

// value receiver
// pのコピーを受け取るので，値は変更されない．構造体は値渡し．
func (p Person) setAge(age int) {
	p.Age = age
}

// pointer receiver
// pのポインタを受け取るので，値が変更される．
func (p *Person) setAgePointer(age int) {
	p.Age = age
}

func (p Person) sayHello() {
	fmt.Println("Hello, I'm", p.Name, ",", p.Age, "years old")
}

func main() {

	p := Person{Name: "Taro", Age: 20}

	// バリューレシーブなので値は変更できない
	p.setAge(100)
	p.sayHello()

	// ポインタレシーブなので値が変更される
	q := Person{Name: "Jiro", Age: 30}
	q.setAgePointer(100)
	q.sayHello()

	r := NewPerson("Saburo", 40)
	r.sayHello()

}

func switch_practice(x int) string {
	switch {
	case x == 1:
		return "1"
	case x == 2:
		return "2"
	case x > 3:
		return "greater than 3"
	default:
		return "default"
	}
}
