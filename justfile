
db := "test.sqlite"

# Run wp24-athletes.
run:
    @go run main.go

create-database:
    rm -f test.sqlite
    sqlite3 {{db}} ".read extras/athletes.sql"
