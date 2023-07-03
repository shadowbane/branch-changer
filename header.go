package main

import "fmt"

func printHeader() {
	appTitle()
	fmt.Println(`#  `)
	printVersion()
	printMaintainer()
	fmt.Println(``)
}

func appTitle() {
	fmt.Println(`#  ██████  ██████   █████  ███    ██  ██████ ██   ██      ██████ ██   ██  █████  ███    ██  ██████  ███████ ██████`)
	fmt.Println(`#  ██   ██ ██   ██ ██   ██ ████   ██ ██      ██   ██     ██      ██   ██ ██   ██ ████   ██ ██       ██      ██   ██`)
	fmt.Println(`#  ██████  ██████  ███████ ██ ██  ██ ██      ███████     ██      ███████ ███████ ██ ██  ██ ██   ███ █████   ██████`)
	fmt.Println(`#  ██   ██ ██   ██ ██   ██ ██  ██ ██ ██      ██   ██     ██      ██   ██ ██   ██ ██  ██ ██ ██    ██ ██      ██   ██`)
	fmt.Println(`#  ██████  ██   ██ ██   ██ ██   ████  ██████ ██   ██      ██████ ██   ██ ██   ██ ██   ████  ██████  ███████ ██   ██`)
}

func printVersion() {
	fmt.Println("#  Version:\t", Version)
}

func printMaintainer() {
	fmt.Println("#  Maintainer:\t", Maintainer)
}
