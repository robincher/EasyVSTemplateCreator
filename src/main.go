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
	"strings"
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

func printHelp() {
	fmt.Println("easycreate <project> <golang/rust/ts/js>")
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
