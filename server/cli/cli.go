package cli

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/AdanJSuarez/folder_counter/server/folderreader"
)

// Init is use to initialize CLI
func Init() {
	fr := folderreader.FolderReader{}
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
		fr.Read(folderName)
		fmt.Printf("Number of file: %v - Total size of the folder: %v\n", fr.GetTotalNumberOfFiles(), fr.GetTotalSize())
		fmt.Println("List of folders and files in a beautiful JSON format:\n", fr.GetFilesStatsString())
		fr = folderreader.FolderReader{}
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
