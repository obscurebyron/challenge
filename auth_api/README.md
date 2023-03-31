Authentication API
------------------

To start locally:

* Have Docker running
* Start a Python virtual environment: `python3 -m venv ./`'
* Activate the virtual env: `source bin/activate`
* Install the requirements in your virtual env: `python -m pip install -r requirements.txt`
* Set the values needed for the app to connect with the database: `export MONGODB_URL="mongodb://root:example@localhost/?retryWrites=true&w=majority"`
* Get the database running using the docker-compose file in mongo_db: `docker compose up` in that directory
* Run the app: `python -m uvicorn app:app --reload`
* View the Swagger docs at http://127.0.0.1:8000/docs#/

Endpoints:

Register a new user:
--------------------

POST /register

example value:

    {
        "name": "Jane Doe",
        "username": "jdoe@example.com",
        "password": "password123"
    }

Login with user credentials:
---------------------------

POST /login

example value:

    {
        "username": "string",
        "password": "string"
    }



    