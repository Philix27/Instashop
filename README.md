# Instashop App

A simple REST API application for managing an e-commerce platform


### Technologies used

- API Router: [Echo](https://echo.labstack.com/)
- Database Queries: [SQLC](https://github.com/sqlc-dev/sqlc) 
- Database Migration: [Go Migrate](https://github.com/golang-migrate/migrate)
- API Documentation: [Swaggo](https://github.com/swaggo/swag)
- RBAC: [Gorbac](https://github.com/mikespook/gorbac)


### Requirement

- Docker
- SQLC
- Swag
- Migrate

____

### Getting started

- Create a .env file and add credentials
- Add a  database url
- Run app
- Visit the swagger documentation

#### Auth Flow

1. SendOtp endpoint generates an otp and a token 
   - The OTP is returned for demo purpose, but to be hidden during prod 
2. Verify OTP
    - Verify the OTP, ensure to add the auth token while sending
3. Register
   - This creates a new user
4. Login
    - email and password created earlier
    - returns a generated token which can then be passed at the header of every other request. 
    - Add role. Either `USER` or `ADMIN`. Designed this way for testing purpose
  
  ```sh 
Bearer ey.....
  ```

#### Product Flow

- Only admin can 
  - Create
  - Delete
  - Update products
- Both User and admin role can Read products

#### Orders Flow
- Only users can 
  - Create
  - Delete
  - Update products
- Both User and Admin role can read orders

___
### Commands

Most of the commands have been mapped to a Makefile

##### Spin up a database
```sh
docker compose up
```

##### Run database migrations
```sh
make db_push
```

##### Create new database migrations
```sh
make db_new
```

##### Generate database queries
```sh
make db_gen
```

##### Run app in development

The database needs to be running before running the app.
```sh
make rd
```

##### Swagger documentation 

```sh
http://localhost:1323/swagger/index.html
```
