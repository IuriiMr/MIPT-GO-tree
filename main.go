package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

func dirTree(out io.Writer, path string, printFiles bool) error {

	err := printTree("", out, path, printFiles)
	return err
}

func printTree(prefix string, out io.Writer, workDir string, printFiles bool) error {

	files, err := ioutil.ReadDir(filepath.Join(workDir))

	if !printFiles {
		var dirList []os.FileInfo = []os.FileInfo{}
		for _, f := range files {
			if f.IsDir() {
				dirList = append(dirList, f)
			}
		}
		files = dirList
	}

	for i, f := range files {

		if f.IsDir() {
			var newPrefix string
			if len(files) > i+1 {
				fmt.Fprintf(out, prefix+"├───"+"%s\n", f.Name())
				newPrefix = prefix + "│\t"
			} else {
				fmt.Fprintf(out, prefix+"└───"+"%s\n", f.Name())
				newPrefix = prefix + "\t"
			}
			newDir := filepath.Join(workDir, f.Name())
			printTree(newPrefix, out, newDir, printFiles)

		} else if printFiles {
			if f.Size() > 0 {
				if len(files) > i+1 {
					fmt.Fprintf(out, prefix+"├───"+"%s (%vb)\n", f.Name(), f.Size())
				} else {
					fmt.Fprintf(out, prefix+"└───"+"%s (%vb)\n", f.Name(), f.Size())
				}
			} else {
				if len(files) > i+1 {
					fmt.Fprintf(out, prefix+"├───"+"%s (empty)\n", f.Name())
				} else {
					fmt.Fprintf(out, prefix+"└───"+"%s (empty)\n", f.Name())
				}
			}
		}
	}
	return err
}

func main() {
	out := os.Stdout
	if !(len(os.Args) == 2 || len(os.Args) == 3) {
		panic("usage go run main.go . [-f]")
	}
	path := os.Args[1]
	printFiles := len(os.Args) == 3 && os.Args[2] == "-f"
	err := dirTree(out, path, printFiles)
	if err != nil {
		panic(err.Error())
	}
}
