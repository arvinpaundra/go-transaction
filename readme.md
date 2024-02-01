# Golang Architecture

## Command Line Usage (flags)

    1. First Migrate
        The command is -i=true for migrating the first time (migrations table)
    2. Migrate (like artisan migrate)
        The command is -m={the_file}.sql
    3. Migrate All Migration Files
        If the files already migrated it skips that file
        The command is -m=all
    4. Create Migration File
        The command is -mmf={migration file name}
        The file format is `yyyy-MM-dd_h:mm:ss_filename.sql`
        Use underscore for each word for the filename
    5. Create all app file
        Using for create all app file (router, handler, service, repository)
        The command is -gen=all, then fill in the required data

## Run Unit Test

    1. Run the test

```sh
$ go test -v -cover ./internal/app/...
```

    2. Run test with coverage and exported to html

```sh
$ go test -v -coverprofile=cover.out ./internal/app/... && go tool cover -html=cover.out
```
