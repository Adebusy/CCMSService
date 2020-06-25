package main

type logger interface {
	Print(v ...interface{})
}
