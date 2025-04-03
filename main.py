# main.py
from fastapi import FastAPI

app = FastAPI(title="Resume CMS API")

@app.get("/")
def read_root():
    return {"message": "Welcome to your Resume CMS!"}