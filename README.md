# Paystack Golang SDK

🐈 An unofficial Golang SDK for [Paystack](http://paystack.com).

## Installation

To install this package, use the command below

```sh
go get github.com/rxxcc/paystack-go-sdk
```

## Usage/Examples

### Import the package

```go
import "github.com/rxxcc/paystack-go-sdk"
```

_Get your keys from your Paystack dashboard and add it into your environment variables_

`To initialize a new client`

```go
package main

import (
    paystack "github.com/rxxcc/paystack-go-sdk"
)

var (
    sKeys = "{Your secret key}"
)

func main() {
    // Handle error
    newClient, _ := paystack.NewClient(sKeys)
}
```

`To create a new transaction`

```go
// --------------------------------- //
func main() {
    testCase := &paystack.TransactionBody{
		Amount:   "300",
		Email:    "test@test.com",
	}
    transactionData, err := newClient.InitializeTransaction(transactionData)
}
```

## License

MIT

## Author

Inu John O.
