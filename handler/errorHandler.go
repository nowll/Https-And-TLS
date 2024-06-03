package handler

import "fmt"

func ErrorHand(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
