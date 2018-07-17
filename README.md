MessageBird's REST API for Go
=============================
This repository contains the open source Go client for MessageBird's REST API. Documentation can be found at: https://developers.messagebird.com.

[![Build Status](https://travis-ci.org/messagebird/go-rest-api.svg?branch=master)](https://travis-ci.org/messagebird/go-rest-api) [![GoDoc](https://godoc.org/github.com/messagebird/go-rest-api?status.svg)](https://godoc.org/github.com/messagebird/go-rest-api)

Requirements
------------
- [Sign up](https://www.messagebird.com/en/signup) for a free MessageBird account
- Create a new access key in the developers sections
- An application written in Go to make use of this API

Installation
------------
The easiest way to use the MessageBird API in your Go project is to install it using *go get*:

```
$ go get github.com/messagebird/go-rest-api
```

Examples
--------
Here is a quick example on how to get started. Assuming the **go get** installation worked, you can import the messagebird package like this:

```go
import "github.com/messagebird/go-rest-api"
```

Then, create an instance of **messagebird.Client**:

```go
client := messagebird.New("test_gshuPaZoeEG6ovbc8M79w0QyM")
```

Now you can query the API for information or send data. For example, if we want to request our balance information you'd do something like this:

```go
// Request the balance information, returned as a Balance object.
balance, err := balance.Read(client)
if err != nil {
	switch errResp := err.(type) {
	case messagebird.ErrorResponse:
		for _, mbError := range errResp.Errors {
			fmt.Printf("Error: %#v\n", mbError)
		}
	}

	return
}

fmt.Println("  payment :", balance.Payment)
fmt.Println("  type    :", balance.Type)
fmt.Println("  amount  :", balance.Amount)
```

This will give you something like:
```shell
$ go run example.go
  payment : prepaid
  type    : credits
  amount  : 9
```

Please see the other examples for a complete overview of all the available API calls.

Documentation
-------------
Complete documentation, instructions, and examples are available at:
[https://developers.messagebird.com](https://developers.messagebird.com).

License
-------
The MessageBird REST Client for Go is licensed under [The BSD 2-Clause License](http://opensource.org/licenses/BSD-2-Clause). Copyright (c) 2014, 2015, MessageBird
