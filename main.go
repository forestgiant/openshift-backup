package main

import (
	"flag"
	"fmt"
	// "os"
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
	fmt.Println("Creating directory named", directory)

	//Create the backup directory if it does not exist

	//Define our openshift command

}
