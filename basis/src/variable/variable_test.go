package variable

import (
	"fmt"
	"testing"
)

/**
 * @Author xgSama
 * @Description:
 * @Date 2021/4/21 14:39
 */

func TestMap(t *testing.T) {
	createMap()
}

func TestSlice(t *testing.T) {
	createSlice()
}

func TestStruct(t *testing.T) {
	//var num myInt = 1
	//var book Book
	//book.title = `GoLang`
	//book.auth = `zhangsan`
	//
	//book.GetTitle()
	//book.SetTitle("book")
	//
	//fmt.Println(book, num)

	var animal Animal
	animal = &Cat{"cat", "red"}
	fmt.Println(animal.GetType())
	ShowAnimal(animal)
}

func TestIo(t *testing.T) {
	testInput()
}

func TestControl(t *testing.T) {
	printShape()
	testGoto()
}

func TestFunc(t *testing.T) {

	testNi()
}

func TestStringFunc(t *testing.T) {

	testStringFunc()
}

func TestTimeFunc(t *testing.T) {

	testTimeFunc()
}

func TestBuiltinFunc(t *testing.T) {

	testBuiltin()
}

func TestError(t *testing.T) {

	defer func() {
		if err := recover(); err!= nil {
			fmt.Println("error tt")
		}
	}()

	testError()
	fmt.Println("test....")

	_, err := testCustomError("c888c")
	if err != nil {
		panic(err)
	}

	fmt.Println("666")

}