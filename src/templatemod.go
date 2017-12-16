package main

import (
	"os/exec"
	"time"
)

/**
Alvin Ng

Logic for cloning
*/

// function prototype for map
type sub func(string)

func setupMappedFunctions() map[string]sub {
	mappedFunctions := map[string]sub{
		// add more templates here
		"golang": createGoProject,
		"rust":   createRustProject,
	}

	return mappedFunctions
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

func createTypeScriptProject(newProjectName string) {
	deleteExistingGoDirectory()
	log("cloning github template")
	if pullTSRepo() {
		log("cloning done, renaming")
		//allow a pause as I've encountered directory lock in windows
		time.Sleep(time.Millisecond * 200)
		log("removing git")
		purgeExistingGitDirectory()
		renameRepo("TypeScriptVSCodeTemplate", newProjectName)
		log("renaming done")
	} else {
		log("error occured")
	}

}

func pullTSRepo() bool {
	var cmd = "git"
	var args = []string{"clone", "https://github.com/rebooting/TypeScriptVSCodeTemplate"}
	if cmdout, err := exec.Command(cmd, args...).Output(); err != nil {
		log(string(cmdout))
		log(err.Error())
		return false
	}

	return true
}

func createJavaScriptProject(newProjectName string) {
	deleteExistingGoDirectory()
	log("cloning github template")
	if pullJSRepo() {
		log("cloning done, renaming")
		//allow a pause as I've encountered directory lock in windows
		time.Sleep(time.Millisecond * 200)
		log("removing git")
		purgeExistingGitDirectory()
		renameRepo("JsWebDeVSCodeTemplate", newProjectName)
		log("renaming done")
	} else {
		log("error occured")
	}

}

func pullJSRepo() bool {
	var cmd = "git"
	var args = []string{"clone", "https://github.com/rebooting/JsWebDeVSCodeTemplate"}
	if cmdout, err := exec.Command(cmd, args...).Output(); err != nil {
		log(string(cmdout))
		log(err.Error())
		return false
	}

	return true
}
