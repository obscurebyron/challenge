import os
from fastapi import FastAPI, Body, HTTPException, status
from fastapi.responses import Response, JSONResponse
from fastapi.encoders import jsonable_encoder
from pydantic import BaseModel, Field
from bson import ObjectId
from typing import Optional
import motor.motor_asyncio
from fastapi.middleware.cors import CORSMiddleware

app = FastAPI()
client = motor.motor_asyncio.AsyncIOMotorClient(os.environ["MONGODB_URL"])
db = client.auth

# This is to handle CORS issues
origins = [
    "http://localhost:3000"
]

app.add_middleware(
    CORSMiddleware,
    allow_origins=origins,
    allow_credentials=True,
    allow_methods=["*"],
    allow_headers=["*"],
)

class PyObjectId(ObjectId):
    @classmethod
    def __get_validators__(cls):
        yield cls.validate

    @classmethod
    def validate(cls, v):
        if not ObjectId.is_valid(v):
            raise ValueError("Invalid objectid")
        return ObjectId(v)

    @classmethod
    def __modify_schema__(cls, field_schema):
        field_schema.update(type="string")


class UserModel(BaseModel):
    id: PyObjectId = Field(default_factory=PyObjectId, alias="_id")
    name: str = Field(...)
    username: str = Field(...)
    password: str = Field(...)

    class Config:
        allow_population_by_field_name = True
        arbitrary_types_allowed = True
        json_encoders = {ObjectId: str}
        schema_extra = {
            "example": {
                "name": "Jane Doe",
                "username": "jdoe@example.com",
                "password": "password123"
            }
        }


class CredentialsModel(BaseModel):
    username: str = Field(...)
    password: str = Field(...)


class UpdateUserModel(BaseModel):
    name: Optional[str]
    username: Optional[str]
    password: Optional[str]

    class Config:
        arbitrary_types_allowed = True
        json_encoders = {ObjectId: str}
        schema_extra = {
            "example": {
                "name": "Jane Doe",
                "username": "jdoe@example.com",
                "password": "password123"
            }
        }


@app.post("/register", response_description="Add new user", response_model=UserModel)
async def register(user: UserModel = Body(...)):
    user = jsonable_encoder(user)
    new_user = await db["users"].insert_one(user)
    created_user = await db["users"].find_one({"_id": new_user.inserted_id})
    return JSONResponse(status_code=status.HTTP_201_CREATED, content=created_user)

@app.post(
    "/login", 
    response_description="verify credentials - i.e. authenticate", 
    response_model=str
)
async def login(response: Response, credentials: CredentialsModel = Body(...)):
    if (await db["users"].
        find_one({"username": credentials.username, "password": credentials.password})
        ) is not None:
        response.set_cookie(key="auth", value="valid")
        return "You have successfully authenticated"

    raise HTTPException(status_code=404, detail="No user found with those credentials")

