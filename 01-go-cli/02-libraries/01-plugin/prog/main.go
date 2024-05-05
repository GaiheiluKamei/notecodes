package main

import (
	"flag"
	"log"
	"plugin"
)

func main() {
	path := flag.String("plugin", "", "plugin to execute")
	flag.Parse()

	if *path == "" {
		log.Fatal("path to plugin not specified")
	}

	p, err := plugin.Open(*path)
	if err != nil {
		log.Fatal(err)
	}

	symbol, err := p.Lookup("ThingsToDo")
	if err != nil {
		log.Fatal(err)
	}

	thingsToDo, ok := symbol.(func())
	if !ok {
		log.Fatalf("could not find function 'ThingsToDo' in plugin")
	}

	thingsToDo()
	log.Println("Did the things")
}
