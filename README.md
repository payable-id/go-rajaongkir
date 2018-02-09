## UnOfficial Raja Ongkir SDK for Go Programming Language

### Raja Ongkir's Website (https://rajaongkir.com)

[![Build Status](https://travis-ci.org/Bhinneka/go-rajaongkir.svg?branch=master)](https://travis-ci.org/Bhinneka/go-rajaongkir)

## Bhinneka :blue_heart: Golang

### Install
  ```shell
  go get github.com/Bhinneka/go-rajaongkir
  ```

### Simpe Usage

  - Get Cost Data

    ```go
    package main

    import (
    	"fmt"
    	"time"

    	"github.com/Bhinneka/go-rajaongkir"
    )

    func main() {

      //Raja Ongkir Constructor, simply call ro.New
      //Paramter 1: "Raja Ongkir API KEY"
      //Paramter 2: HTTP Request Timeout
    	raja := ro.New("your-api-key", 10*time.Second)

      // Paramter 1: City Origin ID
      // Paramter 2: City Destination ID
      // Paramter 3: Item's Weight
      // Paramter 4: Courier's Name
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

    ```

##
### TODO
  - [x] Create Service for Starter(Free Account)
  - [ ] Create Service for Basic and Pro Raja Ongkir Account
  - [ ] Create Unit Testing
  - [ ] Integration with Travis CI

## Authors
  - Lone Wolf (https://github.com/wuriyanto48)
