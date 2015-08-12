## OpenShift-Backup
Simiple backup for OpenShift PostgreSQL using `rhc port-forward` and `pg_dump`

## Requirements

* OS X / Linux
* Install rhc tools (https://developers.openshift.com/en/managing-client-tools.html)
* `rhc setup`
* `rhc port-forward "appname"`

## Usage
```
Usage of openshift-backup:
  -a string
    	*REQUIRED* Name of application to snapshot. (shorthand)
  -appname string
    	*REQUIRED* Name of application to snapshot.
  -d string
    	Name of Postgres DB (shorthand)
  -dbname string
    	*REQUIRED* Port for Postgres DB
  -folder string
    	Name of folder that backups will be stored in. (default "OpenShiftBackUps")
  -p string
    	*REQUIRED* Port for Postgres DB (shorthand)
  -password string
    	*REQUIRED* Username for Postgres DB
  -path string
    	The base directory where the openshift backups will be stored. (default "/Users/jesse")
  -port string
    	*REQUIRED* Port for Postgres DB
  -u string
    	*REQUIRED* Username for Postgres DB (shorthand)
  -username string
    	*REQUIRED* Username for Postgres DB
  -w string
    	*REQUIRED* Username for Postgres DB (shorthand)
```
