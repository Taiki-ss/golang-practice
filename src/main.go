package main

import (
	"fmt"
	"runtime"
)

type User struct {
	name string
}

func (user User) cal(weight float32, height float32) (bmi float32) {
	bmi = weight / height / height * 10000
	return
}

func main() {
	// BMI出力
	user := User{name:"taiki"}
	fmt.Printf("%sのBIMは%gです。", user.name, user.cal(68,177))

	// 使用中のOS出力
	os := runtime.GOOS;
	fmt.Println(os)
}

