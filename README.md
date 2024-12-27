# Instashop App

A simple REST API application for managing an e-commerce platform


### Technologies used

- API Router: [Echo](https://echo.labstack.com/)
- Database Queries: [SQLC](https://github.com/sqlc-dev/sqlc) 
- Database Migration: [Migrate](https://github.com/golang-migrate/migrate)
- Documentation: [Swaggo](https://github.com/swaggo/swag)


### Requirement

- Docker
- SQLC


### Getting started

- Create a .env file and add credentials
- Add a  database url
- Run app
- Visit the swagger documentation

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
