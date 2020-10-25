package cli

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/AdanJSuarez/folder_counter/server/component"
)

// Init is use to initialize CLI
func Init() {
	f := component.Folder{}
	reader := bufio.NewReader(os.Stdin)
	textHeader()

	for {
		fmt.Print("-> ")
		folderName, _ := reader.ReadString('\n')
		folderName = strings.Replace(folderName, "\n", "", -1)
		if folderName == "EXIT" {
			fmt.Println("Thanks for using our software ;)")
			os.Exit(0)
		}
		f.New(folderName)
		fmt.Printf("Number of file: %v - Total size of the folder: %v bytes\n", f.GetTotalFiles(), f.GetSize())
		for _, v := range f.GetListOfComponent() {
			fmt.Printf("Name: %v - Size: %v - Folder: %v - Last modification: %v\n",
				v.GetName(), v.GetSize(), v.GetIsFolder(), v.GetLastModification())
		}
		f = component.Folder{}
		fmt.Println("* End of files *")
		textHeader()
	}
}

func textHeader() {
	fmt.Println("")
	fmt.Println("Type folder name or leave empty for actual folder")
	fmt.Println("Type EXIT to close the program")
	fmt.Println("---------------------")
}
