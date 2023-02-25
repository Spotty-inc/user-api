# Project Title

User API

## Description

Backend data service used to fetch and return user data from the Cassandra user-api db


## Usage

This application should be run using docker. The following environment variables are required:

| Name      | Description |
| ----------- | ----------- |
| DB_HOST      | IP address or FQDN of the Cassandra database     |
| DB_KEYSPACE   | Keyspace name in the target database that stores user information     |


e.g sudo docker run -e DB_HOST='127.0.0.1' -e DB_KEYSPACE='userapi' -d user-api:latest


## Cassandra

Cassandra DB must have a keyspace, and a table called "users" with the following attributes. This will be scripted later but is manual for now. The following commands will set up a cassandra DB for development purposes.

CREATE KEYSPACE userapi WITH replication = {'class':'SimpleStrategy', 'replication_factor' : 1};

USE userapi;

CREATE TABLE users(
           ... id int PRIMARY KEY,
           ... name text,
           ... high_score int
           ... );
