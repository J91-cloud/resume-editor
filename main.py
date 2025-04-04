# main.py
from fastapi import FastAPI, Depends, HTTPException
from sqlalchemy.orm import Session
from models import Profile
from database import SessionLocal, engine
from pydantic import BaseModel
from fastapi.templating import Jinja2Templates
from fastapi import Request

app = FastAPI()

# Create database tables
Profile.metadata.create_all(bind=engine)

# Dependency
def get_db():
    db = SessionLocal()
    try:
        yield db
    finally:
        db.close()

# Pydantic model for Profile
class ProfileCreate(BaseModel):
    full_name: str
    address: str
    email: str
    phone: str
    linkedin: str = None
    github: str = None
    summary: str = None

# Templates setup
templates = Jinja2Templates(directory="templates")

# Routes
@app.get("/")
async def read_profile(request: Request, db: Session = Depends(get_db)):
    profile = db.query(Profile).first()
    # If no profile exists, return a form with empty fields
    if not profile:
        return templates.TemplateResponse("profile.html", {
            "request": request, 
            "profile": None  # Indicates new profile
        })
    return templates.TemplateResponse("profile.html", {"request": request, "profile": profile})

@app.post("/profile/")
def create_profile(profile: ProfileCreate, db: Session = Depends(get_db)):
    db_profile = Profile(**profile.dict())
    db.add(db_profile)
    db.commit()
    db.refresh(db_profile)
    return db_profile

@app.put("/profile/{profile_id}")
def update_profile(profile_id: int, profile: ProfileCreate, db: Session = Depends(get_db)):
    db_profile = db.query(Profile).filter(Profile.id == profile_id).first()
    if not db_profile:
        raise HTTPException(status_code=404, detail="Profile not found")
    
    for field, value in profile.dict().items():
        setattr(db_profile, field, value)
    
    db.commit()
    db.refresh(db_profile)
    return db_profile