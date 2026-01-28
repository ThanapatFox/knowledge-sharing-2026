# GORM Template for New Project

## Description
An example project for GORM code first migration and CRUD and simple transaction operations.

## Run Steps
Follow these steps to set up and run the project:

1. Prepare your PostgreSQL database for connection.
2. Create your `.env` file from `.env.example`.
3. Run the following command:
   ```bash
   go run ./cmd/example/main.go
   ```
4. see the magic

## API Requests
1. Create `POST` `/example`
    ``` json
    {
        "Code": "001", //cant be duplicate
        "Name": ""
    }
    ```
2. Update `PUT` `/example/:id`
    ``` json
    {
        "ID": 1,
        "Code": "001",
        "Name": "edited value"
    }
    ```
3. Delete `DELETE` `/example/:id`
4. Find by id `GET` `/example/:id`
5. Find all `GET` `/example`

## Apply To Your Project
1. Copy entire structure to your project
2. Edit your project `.env` add config from `.env.example`
3. Add your table domain to `internal/core/domain` (see example.go in the directory)
4. Add your interface to `internal/core/ports` (see example.go in the directory)
5. Implement your repository in `internal/infrastructure/repository` (see existing files for example)
6. If you want to add more database engine you can add them to `internal/infrastructure/db` and implement more repository for them later
7. Implement service for app or api caller in `internal/services` (see example_service.go for example and `cmd/example/main.go` for caller example)
8. For `transaction` implement example please see `internal\infrastructure\repository\example_repository_sqlserver.go`
