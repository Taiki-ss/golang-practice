package main

import (
	"fmt"
)

type User struct {
	name string
}

func (user User) cal(weight float32, height float32) (bmi float32) {
	bmi = weight / height / height * 10000
	return
}

func main() {
	user := User{name:"taiki"}
	fmt.Printf("%sのBIMは%gです。", user.name, user.cal(68,177))
}
