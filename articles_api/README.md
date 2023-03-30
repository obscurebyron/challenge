Articles server
===============

To start locally:

* Have Docker running
* Get the database running using the docker-compose file in postgres_db: `docker compose up` in that directory
* Start the app: `go run main.go`
    
Endpoints:

Add a new article:
-----------------

PUT /article 
content-type: "application/json"

required structure example: 
{
    "title" : "my title",
    "content" : "this is my article's content"
}

Retrieve an article by its title
--------------------------------

GET /article/:title 