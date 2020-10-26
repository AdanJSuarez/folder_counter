package cli

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/AdanJSuarez/folder_counter/server/component"
)

//CLI is a structure for handling CLI
type CLI struct {
	folder component.Folder
	reader *bufio.Reader
}

// Init is use to initialize CLI
func (cli *CLI) Init() {
	cli.folder = component.Folder{}
	cli.reader = bufio.NewReader(os.Stdin)
	cli.textHeader()

	for {
		fmt.Print("-> ")
		folderName, _ := cli.reader.ReadString('\n')
		folderName = strings.Replace(folderName, "\n", "", -1)
		if folderName == "EXIT" {
			fmt.Println("Thanks for using our software ;)")
			os.Exit(0)
		}
		cli.folder.New(folderName)
		fmt.Printf("Number of file: %v - Total size of the folder: %v bytes\n", cli.folder.GetTotalFiles(),
			cli.folder.GetSize())
		for _, v := range cli.folder.GetListOfComponent() {
			fmt.Printf("Name: %v - Size: %v - Folder: %v - Last modification: %v\n",
				v.GetName(), v.GetSize(), v.GetIsFolder(), v.GetLastModification())
		}
		cli.folder = component.Folder{}
		fmt.Println("* End of files *")
		cli.textHeader()
	}
}

// textHeader is the header to show when the user start the software as CLI.
func (cli *CLI) textHeader() {
	fmt.Println("")
	fmt.Println("Type folder name or leave empty for ./ folder")
	fmt.Println("Type EXIT to close the program")
	fmt.Println("---------------------")
}
