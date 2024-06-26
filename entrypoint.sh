#!/bin/bash

# Function to check if the database is ready
check_database() {
    
    until nc -z -v -w10 $DB_HOST $DB_PORT
    do
        echo "Waiting for database connection..."
        echo "DB HOST: $DB_HOST, DB_PORT:$DB_PORT"
        sleep 1
    done
}

# Main entry point
main() {
    # Wait for the database to be ready
    echo "Waiting for the database to be ready..."
    check_database

     # Run the database initialization.
    echo "Initializing the database..."
    /bin/bun/bun db init
    
    # Run the database migration.
    echo "Migrating the database..."
    /bin/bun/bun db migrate

    # Run the main application
    echo "Starting the application..."
    exec /bin/server
}

# Call the main function
main "$@"
