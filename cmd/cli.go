package main

import (
	"fmt"
	"tsoobame/tsooctl/pkg/pods"
)

func main() {

	command := pods.GetPods{}

	result := command.Execute()

	fmt.Println(result.Message)
}
