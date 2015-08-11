package main

import (
	"flag"
	"fmt"
	"os"
	// "os/exec"
	"time"
)

func main() {
	//Check for our command line configuration flags
	var (
		backupPathPtr = flag.String("backupPath", "~/", "The base directory where the openshift backups will be stored.")
	)
	flag.Parse()
	fmt.Println("Running openshift-backup with backup path set to ", *backupPathPtr)

	//Get the name of the directory where we want to save this backup
	weekdays := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
	directory := weekdays[time.Now().Weekday()]

	//Create the backup directory if it does not exist
	fi, err := os.Stat(directory)
	if err != nil {
		fmt.Println("Creating directory named", directory)
		// create folder
		os.Mkdir(directory, 0600)
	} else {
		fmt.Println("Folder exists!: ", fi.Name())
	}

	//Define our openshift command

}
