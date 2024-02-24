# Simple GO HTMX

## Introdution

This e-commerce platform uses standard Go libraries for its backend. 
It's a simple htmx project designed for a smooth user experience without full page reloads.

To run the project, simply type the command below, and both the database and the project will start up. 

```bash
    docker compose up
```

To access the pages, you just need to connect to localhost:8080.



## MIGRATIONS

for runnning migrations waiting for container start and run: 
```bash
docker exec -it go-htmx migrate -path /usr/src/app/migrations -database "postgresql://usernamepg:password123@postgresgo:5432/db?sslmode=disable" -verbose up
```
