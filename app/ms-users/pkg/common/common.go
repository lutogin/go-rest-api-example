package common

import "fmt"

func CriticalErrorHandler(err error) {
	if err != nil {
		fmt.Errorf("Error: %s", err.Error())
		panic(err)
	}
}
