#!/bin/bash

if ! command -v sqlite3 &> /dev/null
then
    echo "sqlite3 could not be found."
    exit
fi

if [ -f database.db ]; then
    read -p "Database already exists. Delete? [y/N] " -r -e response
    echo
    if [[ $response =~ ^[Yy]$ ]]; then
        rm database.db
    else
        echo "Cancelling."
        exit
    fi
fi

echo "Creating database..."

sqlite3 database.db < schema.sql

echo "Done."
