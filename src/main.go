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
import(
	"fmt"
	"os/exec"
	"os"
	"unicode"
	"time"
)

func main(){
	args := os.Args
	if len(args)>1{
		var newProjectName = args[1]
		log("deleting any existing template directory...")
		deleteExistingDirectory()
		if isAlphanumeric(newProjectName){
			log("cloning github template")
			if(pullRepo()){
				log("cloning done, renaming")
				//allow a pause as I've encountered directory lock in windows
				time.Sleep(time.Millisecond*200)
				log("removing git")
				purgeExistingGitDirectory()
				renameRepo(newProjectName)
				log("renaming done")
			}else{
				log("error occured")
			}
			
		}else{
			fmt.Println("please use alphanumeric characters for project folder")
		}
	}
	
}

func pullRepo() bool{
	var cmd = "git"
	var args = []string{"clone","https://github.com/RoteErde/VSCodeGoLangStarterTemplate"}
	if cmdout, err := exec.Command(cmd, args...).Output(); err !=nil{
		log(string(cmdout))
		log(err.Error())
		return false
		}
	
	return true
}

func renameRepo(newName string) bool{
	
	if err := os.Rename("VSCodeGoLangStarterTemplate", newName); err !=nil{
		log("failed to rename")
		log(err.Error())
		return false;
	}
	
	return true;
}


/*
	returns true if character is either alphabet or number
	we don't want folders with funny characters
*/
func isAlphanumeric(foldername string) bool{
	for _, character := range foldername{
		if !unicode.IsDigit(character) && !unicode.IsLetter(character){
			return false
		}
	}
	return true
}


/*
remove current cached directory (in case previous operations failed)
*/
func deleteExistingDirectory()bool{
	if err := os.RemoveAll("VSCodeGoLangStarterTemplate"); err!=nil{
		log(err.Error())
		return false
	}
	return true
}


func purgeExistingGitDirectory()bool{
	if err := os.RemoveAll("VSCodeGoLangStarterTemplate/.git"); err!=nil{
		log(err.Error())
		return false
	}
	return true
}