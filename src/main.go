/**
Entry point for "folder_counter"
Adan J. Suarez
e: adanjsuarez@gmail.com
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/AdanJSuarez/folder_counter/src/folderreader"
)

func main() {
	fr := folderreader.FolderReader{}

	reader := bufio.NewReader(os.Stdin)
	textHeader()

	for {
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)
		if strings.ToLower(text) == "exit" {
			fmt.Println("Thanks for using our software ;)")
			os.Exit(0)
		}
		fr.Read(text)
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
