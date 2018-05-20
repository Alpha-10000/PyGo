package main

type Frame struct {
	data Stack
	block Stack
	lastFrame *Frame
	offset int
}
