# POC: Package Requirements

## Situation

Image you have the following struct `Handler` that has a member of the type `apiInterface`. It is an interface to allow for mocking in unit tests.

```go
package app

import "github.com/mattrx/poc-package-requirements/internal/api"



type apiInterface interface {
	Do(req api.Request) (*api.Response, error)
}



type Handler struct {
	api apiInterface
}

func (h Handler) Do() {
	_, _ = h.api.Do(api.Request{})
}

```

The implementation of the api client is in the package `api` and it needs some kind of credentials:

```go
package api



type Client struct{}

func New(user string, pass string) (*Client, error) {

	// do some configuration and maybe return an error

	return &Client{}, nil
}



type Request struct{}

type Response struct{}

func (b *Client) Do(req Request) (*Response, error) {
	return &Response{}, nil
}

```

## Problem

When you forget to configure the api client your programm would panic when you try to use it.

```go
package main

import (
	"github.com/mattrx/poc-package-requirements/internal/app"
)

func main() {
	handler := app.Handler{}
	handler.Do()
}
```

```bash
$ go run main.go
panic: runtime error: invalid memory address or nil pointer dereference
```

Now image that you only use the api client in some edge cases. Your program would run fine in most cases but sometimes panics.

## Solution

We add a package variable called `configured` that tracks if the `New()` function was called. We also add an `init()` function that registers a "requirement" that checks that `configured` is `true`.

```go
package api

import (
	"errors"

	"github.com/mattrx/poc-package-requirements/internal/requirements"
)



func init() {
    requirements.Register(func() error {
		if !configured {
			return errors.New("Api client not configured")
		}

		return nil
	})
}

var configured = false



type Client struct{}

func New(user string, pass string) (*Client, error) {

	// do some configuration and maybe return an error

	configured = true

	return &Client{}, nil
}



type Request struct{}

type Response struct{}

func (b *Client) Do(req Request) (*Response, error) {
	return &Response{}, nil
}

```

In the main function you can now check if all requirements are met and exit on startup.

```go
package main

import (
	"log"
	"os"

	"github.com/mattrx/poc-package-requirements/internal/app"
	"github.com/mattrx/poc-package-requirements/internal/requirements"
)

func main() {
	handler := app.Handler{}

	errs := requirements.Check()
	if len(errs) > 0 {
		for _, err := range errs {
			log.Printf("Requirement not met: %v", err)
		}
		os.Exit(1)
	}

	handler.Do()
}
```

```bash
go run main.go
2021/03/19 20:39:13 Requirement not met: Api client not configured
exit status 1
```

## Exceptions

This only works if the `api` package is imported in the code. In this example it is done by referencing `api.Request` and `api.Response` in the api interface. If the `Do` function would not have any arguments or returns from the `api` package the `init` function in the `api` package will be never called and therefore the requirement will not be registered.