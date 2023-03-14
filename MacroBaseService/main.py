from typing import Union
from fastapi import FastAPI, Depends
from Services.AnomalyService import AnomalyService
from pydantic import BaseModel
from fastapi.responses import JSONResponse

app = FastAPI()


class Item(BaseModel):
    GoodPart: str
    BadPart: str


@app.get("/")
async def read_root():
    return {"Hello": "World"}


@app.post("/find-outlines")
async def read_item(body: Item, service=Depends(AnomalyService)):
    ret = service.find_outlines(body.GoodPart, body.BadPart)
    return JSONResponse(content=ret)
