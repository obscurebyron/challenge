Authentication API
------------------

To start locally:

* Have Docker running
* Start a Python virtual environment: `python -m venv ./`
* Set the values needed for the app to connect with the database: `export MONGODB_URL="mongodb://root:example@localhost/?retryWrites=true&w=majority"`
* Get the database running using the docker-compose file in mongo_db: `docker compose up` in that directory
* Run the app: `uvicorn app:app --reload`


    