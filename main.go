package main

import (
	"io"

	"gosenbay/cmd"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		if err != io.EOF {
			panic(err)
		}
		return
	}
}
