# go-fibserv

[![Go Report Card](https://goreportcard.com/badge/edenprairie.chandlerswift.com/git/chandlerswift/go-fibserv)](https://goreportcard.com/report/edenprairie.chandlerswift.com/git/chandlerswift/go-fibserv)
[![Build Status](https://drone.blackolivepineapple.pizza/api/badges/chandlerswift/go-fibserv/status.svg)](https://drone.blackolivepineapple.pizza/chandlerswift/go-fibserv)
[![Build Status](https://dev.azure.com/chandlerswift0627/demo-project/_apis/build/status/demo-project?branchName=master)](https://dev.azure.com/chandlerswift0627/demo-project/_build/latest?definitionId=1&branchName=master)

A fibonnaci number HTTP server written in Go, used for load testing containers
as the load scales quickly with increased fib index. (Note that this is only
effective to scale CPU load; memory, disk, and network usage are not high.)
