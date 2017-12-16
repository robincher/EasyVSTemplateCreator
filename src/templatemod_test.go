package main

import (
	"fmt"
	"os"
	"testing"
)

func Test_createRustProject(t *testing.T) {
	os.RemoveAll("Test_RustLangGitClone")
	createRustProject("Test_RustLangGitClone")
	_, err := os.Stat("Test_RustLangGitClone")
	fmt.Println("Testing Rust Cloning")
	if err == nil {
		os.RemoveAll("Test_RustLangGitClone")
	} else {
		t.Error("Failed")
	}

}

func Test_createGoProject(t *testing.T) {
	os.RemoveAll("Test_GoLangGitClone")
	createGoProject("Test_GoLangGitClone")
	_, err := os.Stat("Test_GoLangGitClone")
	fmt.Println("Testing Go Cloning")
	if err == nil {
		os.RemoveAll("Test_GoLangGitClone")
	} else {
		t.Error("Failed")
	}

}
