package main

import "iTypeService/cmd"

func main() {
	defer cmd.Clean()
	cmd.Start()
}
