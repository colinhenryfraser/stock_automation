package main

import "log"

func handle(e error){
	if e != nil {
		log.Fatal(e)
	}
}