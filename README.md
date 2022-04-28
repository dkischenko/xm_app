XM Golang Exercise - v21.0.0

For starting app at local machine please do next:
```
docker run -it --rm --name go-postgres -p 5432:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=secret -e PGDATA=/var/lib/postgresql/dat
a/pgdata -v c:/tmp/postgres:/var/lib/postgresql/data postgres:14.0
```

```
go mod download
```

```
go run cmd/main/app.go
```
Or another way
Make a work files for config.yml and .env from theirs copies.
```
docker-compose build
```

```
docker-compose up -d xm_app db
```

Create and delete operation allows for people from Cyprus or authorization needed.