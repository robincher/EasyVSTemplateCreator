package main

/*
Copyrights 2017 Alvin Ng


  This program is free software: you can redistribute it and/or modify
    it under the terms of the GNU General Public License as published by
    the Free Software Foundation, either version 3 of the License, or
    (at your option) any later version.

    This program is distributed in the hope that it will be useful,
    but WITHOUT ANY WARRANTY; without even the implied warranty of
    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
    GNU General Public License for more details.

    You should have received a copy of the GNU General Public License
    along with this program.  If not, see <http://www.gnu.org/licenses/>.

*/
import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
	"unicode"
)

func main() {
	args := os.Args

	if len(args) > 2 {
		var newProjectName = args[1]
		log("deleting any existing template directory...")

		if isAlphanumeric(newProjectName) {
			if strings.ToLower(args[2]) == "go" {
				createGoProject(newProjectName)
			} else if strings.ToLower(args[2]) == "rust" {
				createRustProject(newProjectName)
			} else {
				printHelp()
			}
		} else {
			printHelp()
		}
	} else {
		printHelp()
	}

}

func createRustProject(newProjectName string) {
	deleteExistingRustDirectory()
	log("cloning github template")
	if pullRustRepo() {
		log("cloning done, renaming")
		//allow a pause as I've encountered directory lock in windows
		time.Sleep(time.Millisecond * 200)
		log("removing git")
		purgeExistingGitDirectory()
		renameRepo("RustVSCodeTemplate", newProjectName)
		log("renaming done")
	} else {
		log("error occured")
	}

}

func createGoProject(newProjectName string) {
	deleteExistingGoDirectory()
	log("cloning github template")
	if pullGoRepo() {
		log("cloning done, renaming")
		//allow a pause as I've encountered directory lock in windows
		time.Sleep(time.Millisecond * 200)
		log("removing git")
		purgeExistingGitDirectory()
		renameRepo("VSCodeGoLangStarterTemplate", newProjectName)
		log("renaming done")
	} else {
		log("error occured")
	}

}

func printHelp() {
	fmt.Println("easycreate <project> <go/rust>")
}
func pullGoRepo() bool {
	var cmd = "git"
	var args = []string{"clone", "https://github.com/RoteErde/VSCodeGoLangStarterTemplate"}
	if cmdout, err := exec.Command(cmd, args...).Output(); err != nil {
		log(string(cmdout))
		log(err.Error())
		return false
	}

	return true
}

func pullRustRepo() bool {
	var cmd = "git"
	var args = []string{"clone", "https://github.com/RoteErde/RustVSCodeTemplate"}
	if cmdout, err := exec.Command(cmd, args...).Output(); err != nil {
		log(string(cmdout))
		log(err.Error())
		return false
	}

	return true
}

func renameRepo(templateName string, newName string) bool {

	//if err := os.Rename("VSCodeGoLangStarterTemplate", newName); err != nil {
	if err := os.Rename(templateName, newName); err != nil {
		log("failed to rename")
		log(err.Error())
		return false
	}

	return true
}

/*
	returns true if character is either alphabet or number
	we don't want folders with funny characters
*/
func isAlphanumeric(foldername string) bool {
	for _, character := range foldername {
		if !unicode.IsDigit(character) && !unicode.IsLetter(character) {
			return false
		}
	}
	return true
}

/*
remove current cached directory (in case previous operations failed)
*/
func deleteExistingGoDirectory() bool {
	if err := os.RemoveAll("VSCodeGoLangStarterTemplate"); err != nil {
		log(err.Error())
		return false
	}
	return true
}

func deleteExistingRustDirectory() bool {
	if err := os.RemoveAll("RustVSCodeTemplate"); err != nil {
		log(err.Error())
		return false
	}
	return true
}

func purgeExistingGitDirectory() bool {
	if err := os.RemoveAll("VSCodeGoLangStarterTemplate/.git"); err != nil {
		log(err.Error())
		return false
	}
	return true
}
