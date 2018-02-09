## UnOfficial Raja Ongkir SDK for Go Programming Language

### Raja Ongkir's Website (https://rajaongkir.com)

[![Build Status](https://travis-ci.org/Bhinneka/go-rajaongkir.svg?branch=master)](https://travis-ci.org/Bhinneka/go-rajaongkir)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://github.com/Bhinneka/go-rajaongkir/blob/master/LICENSE)

## Bhinneka :blue_heart: Golang

### Install
  ```shell
  go get github.com/Bhinneka/go-rajaongkir
  ```

### Simple Usage

  - Get Cost Data

    ```go
    package main

    import (
    	"fmt"
    	"time"

    	"github.com/Bhinneka/go-rajaongkir"
    )

    func main() {

      // Raja Ongkir Constructor, simply call ro.New
      // Parameter 1: "Raja Ongkir API KEY"
      // Parameter 2: HTTP Request Timeout
    	raja := ro.New("your-api-key", 10*time.Second)

      // Parameter 1: City Origin ID
      // Parameter 2: City Destination ID
      // Parameter 3: Item's Weight
      // Parameter 4: Courier's Name
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
