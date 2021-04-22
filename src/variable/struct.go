package variable

import "fmt"

/**
 * @Author xgSama
 * @Description:
 * @Date 2021/4/21 14:56
 */
type myInt int

type Book struct {
	title string
	auth  string
}

func PrintBook(book Book) {
	fmt.Println(book)
}

func PrintBook1(book *Book) {
	book.title = "java"
	fmt.Println(book)
}

func (book Book) GetTitle() string {
	return book.title
}

func (book *Book) SetTitle(newName string) {
	book.title = newName
}

type Human struct {
	name string
	sex  string
}

func (this Human) Eat() {
	fmt.Println("Human.Eat()...")
}

func (this Human) Walk() {
	fmt.Println("Human.Walk()...")
}

type SuperMan struct {
	Human
	level int
}

func (this *SuperMan) Eat() {
	fmt.Println("SuperMan.Eat()...")
}

type Animal interface {
	Sleep()
	GetColor() string
	GetType() string
}

type Cat struct {
	name  string
	color string
}

func (cat *Cat) GetColor() string {
	return cat.color
}

func (cat *Cat) GetType() string {
	return "Cat"
}
func (cat *Cat) Sleep() {
	fmt.Println("Cat Sleep ...")
}

func ShowAnimal(animal Animal) {
	fmt.Println(animal)
}
