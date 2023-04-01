Articles server
===============

To start locally:

* Have Docker running
* Get the database running using the docker-compose file in postgres_db: `docker compose up` in that directory
* Start the app: `go run main.go`
* Site should be running at http://127.0.0.1:4000/
    
Endpoints:

Get all articles:
-----------------

GET /article 

Retrieve an article by its slug
--------------------------------

GET /article/:slug 