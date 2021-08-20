from fastapi import FastAPI, HTTPException
from fastapi.middleware.cors import CORSMiddleware
from fastapi.responses import ORJSONResponse

from extract import extract
from schemas.paifu import PaifuResponse

app = FastAPI()

app.add_middleware(
    CORSMiddleware,
    allow_origins=['*'],
    allow_credentials=True,
    allow_methods=['*'],
    allow_headers=['*'],
)


@app.get('/')
def read_root():
    return {"Hello": "World"}


@app.get(
    '/recognition',
    response_class=ORJSONResponse,
    response_model=PaifuResponse,
    summary='recognition pai',
    tags=['recognition']
    )
def recognition(img_url: str):
    resp = extract(img_url)
    if type(resp) == str:
        raise HTTPException(status_code=400, detail=resp)
    return resp
