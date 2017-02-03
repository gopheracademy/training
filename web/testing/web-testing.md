name: web-testing
video: 
description: Now that we have some code written, let's start understanding how to test HTTP Go apps. We'll look at two different ways of testing HTTP apps.
level: Beginner
topic: Go
# Testing HTTP in Go
## Ensure Your Web Works

---
name: web-testing-intro
video: 
thumb:
github:
# Testing HTTP In Go
## 

## Testing 

Now that we have some code written, let's start understanding how to test HTTP Go apps. We'll look at two different ways of testing HTTP apps.

- The standard library has a package named `httptest` with good support.
- There are several ways to create unit and integration tests in Go.---
name: web-testing-links
video: 
thumb:
github:
# Useful Links and Articles
## 

## Links and Articles

- [godoc](https://golang.org/pkg/net/http/)
- [Go Wiki](https://golang.org/doc/articles/wiki/)---
name: web-testing-unit
video: 
thumb:
github:go get github.com/gophertrain/material/web/testing/example1
# Basic Unit Testing
## 

## Unit Testing
Using the `httptest` package it's easy to test your web application.  Create a new `httptest.NewRecorder` and record the response for a given request.   Then compare actual vs. expected results for accurate unit tests.---
name: web-testing-handler
video: 
thumb:
github:go get github.com/gophertrain/material/web/testing/example2
# Testing an HTTP Handler
## 

You can also test handlers on types that you create.---
name: web-testing-routes
video: 
thumb:
github:go get github.com/gophertrain/material/web/testing/example3
# Testing Route Matching
## 

You will also want to ensure that your mux is wired correctly to send routes to the correct handler. ---
name: web-testing-mock
video: 
thumb:
github:go get github.com/gophertrain/material/web/testing/example4
# Integration Testing with httptest.NewServer
## 

Just calling your handler functions is fine for simple tests, but often you'll want to test the entire server.  Using an `httptest.NewServer` you can create a web server that uses your handlers.---
name: web-testing-mock2
video: 
thumb:
github:go get github.com/gophertrain/material/web/testing/example5
# Mocking Servers with http.Handler
## 

You can put it all together by creating your own type that implements http.Handler and creating an httptest.NewServer that tests your entire type.  Notice in the example how the test server doesn't specify the port or URL, it's all built into the test library.