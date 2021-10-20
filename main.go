package main

import (
	"fmt"
	"github/vonuki/template_api/api"
	"os"
	"strconv"
)

func main() {
	fmt.Println("Start template API")

	// Read args
	for i := 0; i < len(os.Args); i++ {
		fmt.Printf("input args [%d] = %s \n", i, os.Args[i])
	}

	// convert str arg to int
	value, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Printf("Arguemtn format is not a integer: %s \n", os.Args[1])
	} else {
		fmt.Printf("Get integer value = %d \n", value)
	}

	app := api.GetCurrentTimeAPI()
	resultTime, err := app.GetTime(os.Args[1])
	fmt.Printf("Result time = %s \n", resultTime)

	fmt.Println("Finsh template API")
}
