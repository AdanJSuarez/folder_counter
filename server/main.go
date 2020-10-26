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

// readArgument read the argument the user pass when running the software.
// It set 'api' if no argument passed.
func readArgument() string {
	if len(os.Args) > 1 {
		return os.Args[1]
	}
	return "api"
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
