/**
Entry point for "folder_counter"
Adan J. Suarez
e: adanjsuarez@gmail.com
*/

package main

import (
	"fmt"
	"os"

	"github.com/AdanJSuarez/folder_counter/server/cli"

	"github.com/AdanJSuarez/folder_counter/server/api"
)

func readArgument() string {
	return os.Args[1]
}

func main() {
	api := api.API{}
	cli := cli.CLI{}
	arg := readArgument()
	if arg == "cli" {
		cli.Init()
	} else if arg == "api" {
		api.Init()
	}
	fmt.Printf("'%v' is not a valid argument, you should pass 'api' or 'cli'\n", arg)
	os.Exit(0)

}
