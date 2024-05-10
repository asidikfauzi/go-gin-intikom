# Test Go Gin Intikom
Test the PT Intikom  Tasks. This application was built using Golang 1.21.6, PostgreSQL latest, and Docker.

## Link Documentation Postman
[Click Link Postman](https://crimson-flare-213787.postman.co/workspace/Team-Workspace~79f6082a-b6ff-401a-8e2d-348a27ef9881/collection/34821541-960f09ce-9bdc-467c-b611-06276471ad24?action=share&creator=34821541&active-environment=34821541-f862945e-f2f1-4218-b80b-2e9625d22140)


## Installation
Use the package manager docker to install app Postgres latest.
```bash
docker-compose up -d --build
```

Create a new Postgres latest database with name `intikom`.

## Mac OS / Linux users
Run the `make migrate` command to populate the database tables.
```bash
make migrate
```

Run the `make seed` command to fill data in several tables.
```bash
make seed
```

Run the `make all` command to run the application.
```bash
make all
```

## Windows users
Run the following command to populate the database table.
``` bash
go mod vendor -v
rm -f cmd/migrate/migrate
go build -o cmd/migrate/migrate cmd/migrate/migrate.go
./cmd/migrate/migrate
```

Run the following command to fill data in several tables.
``` bash
go mod vendor -v
rm -f cmd/seed/seed
go build -o cmd/seed/seed cmd/seed/seed.go
./cmd/seed/seed
```
Run the following command to run the application.
``` bash
go mod vendor -v
rm -f cmd/app/app
go build -o cmd/app/app cmd/app/app.go
./cmd/app/app
```