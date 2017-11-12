package main

import (
	"testing"
)

func Test_isAlphanumeric(t *testing.T) {
	type args struct {
		foldername string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isAlphanumeric(tt.args.foldername); got != tt.want {
				t.Errorf("isAlphanumeric() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
	// TODO: Add test cases.
	}
	for range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}

func Test_pullRepo(t *testing.T) {
	tests := []struct {
		name string
		want bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pullRepo(); got != tt.want {
				t.Errorf("pullRepo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_renameRepo(t *testing.T) {
	type args struct {
		newName string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := renameRepo(tt.args.newName); got != tt.want {
				t.Errorf("renameRepo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_deleteExistingDirectory(t *testing.T) {
	tests := []struct {
		name string
		want bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := deleteExistingDirectory(); got != tt.want {
				t.Errorf("deleteExistingDirectory() = %v, want %v", got, tt.want)
			}
		})
	}
}
