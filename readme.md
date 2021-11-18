# Donut test

## Getting started

### .env file

You need a .env file for the application to get it's env variables. You can simply rename the `.env.sample` you have for out of the box functionality

------------
### Setup
1. Install docker, and start it. Get started at [Docker](www.docker.com/products/docker-desktop "Docker")
2. There is a make file with all the commands you need. Run the following to have the app completly setup.  
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
```
------------
### Running the app
To run the app simply `go run main.go`, you will be prompted for you name, and how many fake random bank statements you will want to round up.

You will be presented with a selection screen. Choose the option you feel like.

You must run `make ulimit` before running the app, if you close your terminal window.

When seeding the database bank transactions keep it < 150 for a smooth experience

**You MUST reset the database after each batching, with the provided option**

## How it works
### Database
There are only 4 tables. Users that represent the users; batches represent the batching of transactions; transactions represent a broadcasted amount from batch into the bank, and actions, that simulates user behaviour in their bank. You can see a picture of the schema in the project.

### Logic
App works with concurrent workers, with the acceptance criteria from the assignment.
We currently add to each batch 1 + rounded up amount of user's bank transactin so more batches get dispatched

## Future improvments
Allow user to pick individual transaction amounts, to simulate this user action. App currently only simulates user choosing an option where all low volume transactions are sent to the batch.
Improve concurrent functionality on seeding, to prevent crashing when seeding > 150 at one time.
Implement Mock database for tests. This will clean-up much of the current test functionality.
Improve on testing functionality.
