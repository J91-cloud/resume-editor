# main.py
from fastapi import FastAPI
from fastapi.responses import HTMLResponse
from fastapi import Request
from fastapi.templating import Jinja2Templates

# Configure templates directory
templates = Jinja2Templates(directory="templates")

app = FastAPI(title="Resume CMS API")

@app.get("/", response_class=HTMLResponse)
async def read_root(request: Request):
    return templates.TemplateResponse(
        "main/index.html",  
        {"request": request}  
    )