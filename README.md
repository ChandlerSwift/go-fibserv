# go-fibserv

[![Go Report Card](https://goreportcard.com/badge/edenprairie.chandlerswift.com/git/chandlerswift/go-fibserv)](https://goreportcard.com/report/edenprairie.chandlerswift.com/git/chandlerswift/go-fibserv)

A fibonnaci number HTTP server written in Go, used for load testing containers
as the load scales quickly with increased fib index. (Note that this is only
effective to scale CPU load; memory, disk, and network usage are not high.)
