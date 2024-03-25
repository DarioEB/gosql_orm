# Go Sql orm - First steps

## Local tests

### Run local mysql database with docker
```
docker run --name gosql-database -e MYSQL_ROOT_PASSWORD=devpass -e MYSQL_PASSWORD=devpass -e MYSQL_USER=devuser -e MYSQL_DATABASE=gosql -v $(pwd)/database:/var/lib/mysql -dp 3306:3306 mysql:8.0.3
```

### Run local psql database with docker
```
docker run --name gopsql-database -e POSTGRES_PASSWORD=devpass -e POSTGRES_USER=devuser -e POSTGRES_DB=gopsql -v $(pwd)/ps-database:/var/lib/postgresql/data -dp 5432:5432 postgres:16.2
```