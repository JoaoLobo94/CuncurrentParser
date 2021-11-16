# Donut test

## Getting started
### .env file
You need a .env file for the application to get it's 

------------

env variables. You can simply rename the `.env.sample` you have out of the box functionality
### Setup
1. Install docker, and start it. Get started at [Docker](http://https://www.docker.com/products/docker-desktop "Docker")
2. There is a make file with all the other commands you need. Run the following to have the app completly setup.  
```
make install_golang_migrate
make ulimit
make install_sqlc
make pullpostgres
make runpostgres
make createdb
make create_test_db
make migratedb
make migrate_test_db
make sqlc

------------
### Running the app
To run the app simply `go run main.go`, you will be prompted for you name, and how many fake random bank statements you will want to round up.
You will be presented with a selection screen. Choose the option you feel like.
**You MUST reset the databse after each batching, with the provided option**
## How it works
### Database
There are only 4 entries. Users that represent the users; batches represent the batching of transacitons; transactions represent a broadcasted amount from batch into the bank, and actions, that simulates user behaviour in their bank
### Logic
App works with cuncurrent workers, with the acceptance criteria from the assigment

## Future improvments
Alow user to pick individual transaction amounts, to simulate this user action. App currently only simulates user choosing an option where all low volume transactions are sent to the batch.
Implement Mock database for tests. This will cleanup much of the current test functionality.
Improve on testing functionality.
