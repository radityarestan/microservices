package main

import "log"

func (app *Log) info(info string) {
	log.Printf("[INFO] %s", info)
}

func (app *Log) error(err string) {
	log.Printf("[ERROR] %s", err)
}

func (app *Log) panic(panic string) {
	log.Printf("[PANIC] %s", panic)
}
