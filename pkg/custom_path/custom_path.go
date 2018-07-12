// Copyright (c) 2018 Davsk℠. All Rights Reserved.
// Use of this source code is governed by an ISC License (ISC)
// that can be found in the LICENSE file.

// 'custom_path.go'

// by David Skinner
// on July 11, 2018
// for Davsk℠ Universe 4.0 project gbase

// package custom_path sets the working directory to the directory containing the program
// and returns an optional custom path defined on the command line.
// code derived from: https://www.kaihag.com/external-assets-working-directories-and-go/
// When using a Docker container for example you could point the asset path to a persistent directory
// so that your files survive a restart.
package custom_path

import (
	"flag"
	"log"
	"os"
	"path/filepath"

	"github.com/kardianos/osext"
)

var (
	// Initialization of the working directory. Needed to load asset files.
	filePath = determineWorkingDirectory()
)

// func Join receives a filename string and joins it to the defined custom path
func Join(fileName string) string {
	return filepath.Join(filePath, fileName)
}

func determineWorkingDirectory() string {
	var customPath string
	// Check if a custom path has been provided by the user.
	flag.StringVar(&customPath, "custom-path", "", "Specify a custom path to the asset files. This needs to be an absolute path.")
	flag.Parse()
	// Get the absolute path this executable is located in.
	executablePath, err := osext.ExecutableFolder()
	if err != nil {
		log.Fatal("Error: Couldn't determine working directory: " + err.Error())
	}
	// Set the working directory to the path the executable is located in.
	os.Chdir(executablePath)
	// Return the user-specified path. Empty string if no path was provided.
	return customPath
}
