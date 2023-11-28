# Go Workshop

This workshop will help you learn the [go programming language](https://go.dev/), also known as golang.

Work through the challenges in order. Ensure that you look at resources for each challenge as they include tips and useful resources on solving the challenge.

This workshop is modeled after an OpenHack. The solution is not provided and it's expected that some exploration would take place as part of the learning process.

Scenario:

A development team created a simple sdk to represent dogs called [GoDog](https://github.com/hattan/godog). You have been tasked with improving the GoDog sdk and adjacent projects. This repo includes two supporting projects.

- [GoDog API](./godogapi/) A restful api built in Go using the Gin Framework and contains an in memory dog database.
- [GoDog CLI](./godogcli/) A simple CLI tool written Go using the Cobra framework. 


Note while you will be modifying the web api and cli, it is not expected that you dig deep into Gin or Cobra. In fact all the api and cli hooks are already there, you only need to implement the go logic for specific challenges.

## Challenges

- [Challenge 1 - Refactor the GoDo SDK Demo](./challenge-1.md)


### Resources:

- [The Go Programming Language Website](https://go.dev/)
- [A tour of Go](https://go.dev/tour/welcome/1)
- [Go Package Docs](https://pkg.go.dev/)
