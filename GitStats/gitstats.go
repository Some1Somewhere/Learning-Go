package main

import (
	"bufio"
	"flag"
	"fmt"

	"log"
	"os"
	"strings"
)

func checkError(err error, linenum int) {
	if err != nil {
		log.Fatal("Line Num : ", linenum, " Error : ", err)
	}
}

func scanGitFolders(folders []string, folder string) []string {
	folder = strings.TrimSuffix(folder, "/")
	f, err := os.Open(folder)
	checkError(err, 22)
	files, err := f.Readdir(-1)
	defer f.Close()
	checkError(err, 25)

	var path string
	for _, file := range files {
		if file.IsDir() {

			if file.Name() == ".git" {
				folders = append(folders, path)
				continue
			}
			if file.Name() == ".dart_tool" || file.Name() == ".github" || file.Name() == ".firebase" || file.Name() == "android" || file.Name() == "assets" || file.Name() == "build" || file.Name() == "ios" || file.Name() == "lib" || file.Name() == "test" || file.Name() == "web" || file.Name() == "node_modules" || file.Name() == "public" || file.Name() == "src" {
				continue
			}
			path = folder + "/" + file.Name()
			folders = scanGitFolders(folders, path) //keeps checking internal folders for .git
		}
	}
	return folders
}

func addNewSliceElementsToFile(path string, newRepos []string) {
	f, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0755) //Opens the gogitlocalstats file
	if err != nil {
		if os.IsNotExist(err) {
			_, err = os.Create(path) //creates if it doesn't exist
			checkError(err, 50)
		} else {
			checkError(err, 52)
		}
	}
	defer f.Close()
	var existingRepos []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		existingRepos = append(existingRepos, scanner.Text())
	}
	repos := append(newRepos, existingRepos...) //Makes sure if folder is already scanned, doesn't have to scan again
	contentToWrite := strings.Join(repos, "\n")
	err2 := os.WriteFile(path, []byte(contentToWrite), 0755)
	checkError(err2, 66)
}

func scan(path string) {
	repos := scanGitFolders(make([]string, 0), path)
	filePath := "./.gogitlocalstats" //Relative to current folder, tutorial has fixed path
	addNewSliceElementsToFile(filePath, repos)
}

func stats(email string) {
	fmt.Println(email)
}

func main() {
	var folder string
	var email string
	flag.StringVar(&folder, "add", "", "New folder to scan to gitstats")
	flag.StringVar(&email, "email", "dhruvshetty3@gmail.com", "Email to Scan")
	flag.Parse()

	if folder != "" {
		scan(folder)
		return
	}

	stats(email)

}
