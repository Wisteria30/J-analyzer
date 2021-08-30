from urllib.parse import unquote

from fastapi import FastAPI, HTTPException
from fastapi.middleware.cors import CORSMiddleware
from fastapi.responses import ORJSONResponse

from extract import extract
from extract_test import extract_test
from load_img import img2url, load_templates
from schemas.paifu import PaifuResponse
from schemas.paifu_test import PaifuTestResponse

app = FastAPI()

app.add_middleware(
    CORSMiddleware,
    allow_origins=['*'],
    allow_credentials=True,
    allow_methods=['*'],
    allow_headers=['*'],
)


hand_templates = load_templates('template/hand/')
bakaze_templates = load_templates('template/bakaze/')
jikaze_templates = load_templates('template/jikaze/')
dora_templates = load_templates('template/dora/')
yohai_templates = load_templates('template/yohai/')


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
    recog_img = img2url(img_url)
    resp = extract(
        recog_img,
        hand_templates,
        bakaze_templates,
        jikaze_templates,
        dora_templates,
        yohai_templates
    )
    return resp


@app.get(
    '/recognition-test',
    response_class=ORJSONResponse,
    response_model=PaifuTestResponse,
    summary='recognition pai test return string',
    tags=['recognition_test']
    )
def recognition_test(img_url: str):
    resp = extract_test(unquote(img_url))
    if type(resp) == str:
        raise HTTPException(status_code=400, detail=resp)
    return resp
