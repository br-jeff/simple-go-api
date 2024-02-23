# Simple GO HTMX

## Introdution

This e-commerce platform uses standard Go libraries for its backend. 
It's a simple htmx project designed for a smooth user experience without full page reloads.


## MIGRATIONS

for runnning migrations waiting for container start and run: 
```bash
docker exec -it go-htmx migrate -path /usr/src/app/migrations -database "postgresql://usernamepg:password123@postgres-go:5432/db?sslmode=disable" -verbose up
```
