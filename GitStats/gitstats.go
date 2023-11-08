package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func scanGitFolders(folders []string, folder string) []string {
	folder = strings.TrimSuffix(folder, "/")
	f, err := os.Open(folder)
	checkError(err)
	files, err := f.Readdir(-1)
	defer f.Close()
	checkError(err)

	var path string
	for _, file := range files {
		if file.IsDir() {

			if file.Name() == ".git" {
				folders = append(folders, path)
				continue
			}
			if file.Name() == "android" || file.Name() == "assets" || file.Name() == "build" || file.Name() == "ios" || file.Name() == "lib" || file.Name() == "test" || file.Name() == "web" || file.Name() == "node_modules" || file.Name() == "public" || file.Name() == "src" {
				continue
			}
			path = folder + "/" + file.Name()
			folders = scanGitFolders(folders, path)
		}
	}
	return folders
}

func addNewSliceElementsToFile(path string, newRepos []string) {
	f, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0755)
	if err != nil {
		if os.IsNotExist(err) {
			_, err = os.Create(path)
			checkError(err)
		} else {
			checkError(err)
		}
	}
	defer f.Close()
	var existingRepos []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		existingRepos = append(existingRepos, scanner.Text())
	}
	if err := scanner.Err(); err != io.EOF {
		checkError(err)
	}
	//repos := joinSlices(newRepos, existingRepos)
	//dumpStringsSliceToFile(repos, path)
}

func scan(path string) {
	repos := scanGitFolders(make([]string, 0), path)
	fmt.Print(repos)
	//filePath := "./.gogitlocalstats" //Relative to current folder, tutorial has fixed path
	//addNewSliceElementsToFile(filePath, repos)
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
