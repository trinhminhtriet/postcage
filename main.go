package main

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/trinhminhtriet/postcage/cmd"
	sendmail "github.com/trinhminhtriet/postcage/sendmail/cmd"
)

func main() {
	exec, err := os.Executable()
	if err != nil {
		panic(err)
	}

	if normalize(filepath.Base(exec)) == normalize(filepath.Base(os.Args[0])) ||
		!strings.Contains(filepath.Base(os.Args[0]), "sendmail") {
		cmd.Execute()
	} else {
		sendmail.Run()
	}
}

func normalize(s string) string {
	s = strings.ToLower(s)

	return strings.TrimSuffix(s, filepath.Ext(s))
}
