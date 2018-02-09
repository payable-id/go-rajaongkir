## UnOfficial Raja Ongkir SDK for Go Programming Language

### Raja Ongkir Website
    https://rajaongkir.com

[![Build Status](https://travis-ci.org/Bhinneka/go-rajaongkir.svg?branch=master)](https://travis-ci.org/Bhinneka/go-rajaongkir)

## Bhinneka :blue_heart: Golang

### Install
  ```shell
  go get github.com/Bhinneka/go-rajaongkir
  ```

### Simpe Usage


  ```go
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

  ```

##
### TODO
  - [x] Create Service for Starter(Free Account)
  - [ ] Create Service for Basic and Pro Raja Ongkir Account
  - [ ] Create Unit Testing
  - [ ] Integration with Travis CI

## Authors
  - Lone Wolf (https://github.com/wuriyanto48)
