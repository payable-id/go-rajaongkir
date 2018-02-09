package main

import (
	"fmt"
	"time"

	"github.com/Bhinneka/go-rajaongkir"
)

func main() {

	raja := ro.New("your-api-key", 10*time.Second)

	q := ro.QueryRequest{Origin: "501", Destination: "114", Weight: 1700, Courier: "tiki"}
	result := raja.GetCost(q)

	if result.Error != nil {
		fmt.Println(result.Error.Error())
	}

	cost, ok := result.Result.(ro.Cost)
	if !ok {
		fmt.Println("Result is not Cost")
	}

	fmt.Println(cost)
}
