# Project Title

User API

## Description

Backend data service used to fetch and return user data from the Cassandra user-api db


## Usage

The application requires a DB_HOST environment variable to be set. This should be the IP or FQDN of the Cassandra database.


DB_HOST = IP address or FQDN of the Cassandra database
DB_KEYSPACE = Keyspace name in the target database that stores user information


e.g sudo docker run [...] -e DB_HOST="127.0.0.1"
