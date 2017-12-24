package main

import "os"

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

func renameRepo(templateName string, newName string) bool {
	//if err := os.Rename("VSCodeGoLangStarterTemplate", newName); err != nil {
	if err := os.Rename(templateName, newName); err != nil {
		log("failed to rename")
		log(err.Error())
		return false
	}
	return true
}
