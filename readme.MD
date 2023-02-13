# Project Title

User API

## Description

Backend data service used to fetch and return user data from the Cassandra user-api db


## Usage

This application should be run using docker. The following environment variables are required:

DB_HOST = IP address or FQDN of the Cassandra database
DB_KEYSPACE = Keyspace name in the target database that stores user information

| Name      | Description |
| ----------- | ----------- |
| DB_HOST      | IP address or FQDN of the Cassandra database     |
| DB_KEYSPACE   | Keyspace name in the target database that stores user information     |


e.g sudo docker run -e DB_HOST='127.0.0.1' -e DB_KEYSPACE='user-api' user-api:latest
