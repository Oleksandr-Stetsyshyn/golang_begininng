package structs

import "fmt"

func Dereference() {
	postPointer := new(Post)
	postValue := Post{}

	postValue.User.Name = "John Value"

	postPointer.User.Name = "John Pointer"

	(*postPointer).User.Age = 22

	fmt.Printf("postPointer: %+v, %T\n", postPointer, postPointer)
	fmt.Printf("postValue: %+v, %T\n", postValue, postValue)

}
