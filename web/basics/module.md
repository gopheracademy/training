name: web-basics
video: 
description: With a basic understand of how the web and HTTP work, let's write a simple "Hello World" app in Go. We'll cover how to start a web server in Go, take in requests, and return responses.
level: Beginner
topic: Go
# Basics
## HTTP Basics In Go

---
name: web-basics-intro
video: 
thumb:
github:
# Introduction
## 

Learn the basics of building web services and applications in Go.

- The standard library has much of what you need to build services and apps.
- The http package provides the building blocks.---
name: web-basics-links
video: 
thumb:
github:
# Useful Links and Articles
## 

## Documentation

- [godoc](https://golang.org/pkg/net/http/)
- [tutorial](https://golang.org/doc/articles/wiki/)

---
name: web-simple-web
video: 
thumb:
github:go get github.com/gophertrain/material/web/basics/example1
# Simple Web Application
## 

Go's http package provides everything you need to serve web applications.

This sample code shows you how to create an http listener and create a function that will respond to http requests.---
name: web-basics-mux
video: 
thumb:
github:go get github.com/gophertrain/material/web/basics/example2
# Creating A New Request Multiplexer
## 

When you want to route requests to different functions based on the URL or other variables, you need to create a multiplexer -- or mux -- to determine the path each request will take.---
name: web-basics-handler
video: 
thumb:
github:go get github.com/gophertrain/material/web/basics/example3
# User Defined Handlers
## 

Of course, the real power in Go is in its interfaces.  The `http.Handler` interface allows any type to become a web server.  It only needs to implement the `ServeHTTP()` function.---
name: web-basics-handler2
video: 
thumb:
github:go get github.com/gophertrain/material/web/basics/example4
# More Sophisticated Handlers
## 

You can even embed a handler in your type.---
name: web-concurrency
video: 
thumb:
github:go get github.com/gophertrain/material/web/basics/example5
# Web Service in a Goroutine
## 

Of course you don't have to let the web application be the only thing that happens in your web application.  In this example, we launch a web server in a goroutine, then call it from the same application.---
name: web-basics-exercise
video: 
thumb:
github:
# Web Exercise
## Hands On Exercise

Write a simple web service that has a set of different routes that return the string "Hello World" in multiple languages. Build the service using an Application context that will own the different handler methods. Then create your own mux, bind the routes and start the service. Validate your routes work in your browser.
