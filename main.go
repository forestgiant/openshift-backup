package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/user"
	"strings"
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
		appNamePtr   = flag.String("appname", "", appNameUsage)

		// Postgres
		userNameUsage = "*REQUIRED* Username for Postgres DB"
		userNamePtr   = flag.String("username", os.Getenv("PGUSER"), userNameUsage)

		passwordUsage = "*REQUIRED* Username for Postgres DB"
		passwordPtr   = flag.String("password", os.Getenv("PGPASSWORD"), passwordUsage)

		portUsage = "*REQUIRED* Port for Postgres DB"
		portPtr   = flag.String("port", os.Getenv("PGPORT"), portUsage)

		dbNameUsage = "Name of Postgres DB"
		dbNamePtr   = flag.String("dbname", os.Getenv("PGDATABASE"), portUsage)

		backupPathPtr = flag.String("path", user.HomeDir, "The base directory where the openshift backups will be stored.")
		folderNamePtr = flag.String("folder", "OpenShiftBackUps", "Name of folder that backups will be stored in.")
	)

	// Set up short hand flags
	flag.StringVar(appNamePtr, "a", "", appNameUsage+" (shorthand)")
	flag.StringVar(userNamePtr, "u", "", userNameUsage+" (shorthand)")
	flag.StringVar(passwordPtr, "w", "", passwordUsage+" (shorthand)")
	flag.StringVar(portPtr, "p", "", portUsage+" (shorthand)")
	flag.StringVar(dbNamePtr, "d", "", dbNameUsage+" (shorthand)")

	flag.Parse()

	// If an appName isn't set then return
	if *appNamePtr == "" {
		log.Fatalln("Must set --appName (-a) flag")
	} else if *userNamePtr == "" {
		log.Fatalln("Must set --username (-u) flag")
	} else if *passwordPtr == "" {
		log.Fatalln("Must set --password (-w) flag")
	} else if *portPtr == "" {
		log.Fatalln("Must set --port (-p) flag")
	}

	// If the DB Name is black set it to the appNamePtr
	if *dbNamePtr == "" {
		*dbNamePtr = *appNamePtr
	}

	// Set environment variables

	os.Setenv("PGHOST", "127.0.0.1")
	os.Setenv("PGPORT", *portPtr)
	os.Setenv("PGDATABASE", *dbNamePtr)
	os.Setenv("PGUSER", *userNamePtr)
	os.Setenv("PGPASSWORD", *passwordPtr)

	// Set Path
	path := *backupPathPtr + "/" + *folderNamePtr

	// Create OpenShiftBackUps directory
	createDir(path, 0700)

	//Get the name of the directory where we want to save this backup
	weekdays := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
	directory := weekdays[time.Now().Weekday()]

	dirPath := path + "/" + directory

	//Create the backup directory if it does not exist
	createDir(dirPath, 0700)

	// Define commands
	var (
		cmd    *exec.Cmd
		output []byte
	)

	// TODO: Setup port forwarding so it's not blocking
	// cmd = exec.Command("rhc", "port-forward", "-a", *appNamePtr)

	// Change directory to dirPath to save pg_dump
	os.Chdir(dirPath)

	fmt.Println("Running openshift-backup with backup path set to ", dirPath)

	// Call pg_dump -w (don't prompt password)
	cmd = exec.Command("pg_dump", "-w", "-f", *appNamePtr+".sql")
	output, err = cmd.CombinedOutput()
	prettyOutput := strings.Replace(string(output), "\n", "", -1)

	if err != nil {
		fmt.Println(errors.New(err.Error() + ": " + prettyOutput))
	} else {
		fmt.Printf("Backup complete: %v/%v.sql", dirPath, *appNamePtr)
	}

	fmt.Println(prettyOutput)

}

func createDir(name string, perm os.FileMode) error {
	_, err := os.Stat(name)
	if err != nil {
		fmt.Println("Creating directory named", name)

		// Create folder
		err = os.Mkdir(name, perm)
		if err != nil {
			fmt.Println("Couldn't create directory: ", err)

			return err
		}

	}

	return nil
}
