package main

import (
	"flag"
	"fmt"
	"os"
	"os/user"
	// "os/exec"
	"log"
	"time"
)

func main() {
	// Get Users home directory
	user, err := user.Current()
	if err != nil {
		log.Fatalln(err)
	}

	//Check for our command line configuration flags
	var (
		appNameUsage = "*REQUIRED* Name of application to snapshot."
		appNamePtr   = flag.String("appName", "", appNameUsage)

		backupPathPtr = flag.String("backupPath", user.HomeDir, "The base directory where the openshift backups will be stored.")
		folderNamePtr = flag.String("folderName", "OpenShiftBackUps", "Name of folder that backups will be stored in.")
	)

	// Set up short hand flags
	flag.StringVar(appNamePtr, "a", "", appNameUsage+" (shorthand)")

	flag.Parse()

	// If an appName isn't set then return
	if *appNamePtr == "" {
		log.Fatalln("Must set --appName (-a) flag")
	}

	fmt.Println("Running openshift-backup with backup path set to ", *backupPathPtr)

	// Set Path
	path := *backupPathPtr + "/" + *folderNamePtr

	// Create OpenShiftBackUps directory
	createDir(path, 0700)

	//Get the name of the directory where we want to save this backup
	weekdays := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
	directory := weekdays[time.Now().Weekday()]

	dirPath := path + "/" + directory

	//Create the backup directory if it does not exist
	createDir(dirPath, 0600)

	//Define our openshift command
	fmt.Println("App name: ", *appNamePtr)

	// rhc snapshot save -a {appName}

}

func createDir(name string, perm os.FileMode) error {
	fi, err := os.Stat(name)
	if err != nil {
		fmt.Println("Creating directory named", name)

		// Create folder
		err = os.Mkdir(name, perm)
		if err != nil {
			fmt.Println("Couldn't create directory: ", err)

			return err
		}

	} else {
		fmt.Println("Folder exists!: ", fi.Name())

	}

	return nil
}
