package main

import "log"

// this is entry point
func main() {
	s := NewServer()
	err := s.Start()
	if err != nil {
		log.Fatal(err)
	}
}
