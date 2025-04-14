# main.py
from fastapi import FastAPI, Depends, HTTPException
from sqlalchemy.orm import Session,sessionmaker
from models import Profile,Projects,Skills, Certifications
from database import SessionLocal, engine
from pydantic import BaseModel
from fastapi.templating import Jinja2Templates
from fastapi import Request
from fastapi.staticfiles import StaticFiles

app = FastAPI()


app.mount("/static", StaticFiles(directory="static"), name="static")
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


class ProjectRequestDTO(BaseModel):
    full_name: str
    start_date: str
    end_date: str
    small_description: str
    technologies_used: str
    role_of_project: str

class SkillsDTO(BaseModel):
    skill_name: str

class CertificationsDTO(BaseModel):
    full_name: str
    start_date: str
    end_date: str


# Templates setup
templates = Jinja2Templates(directory="templates")

# Routes
# @app.get("/")
# async def read_profile(request: Request, db: Session = Depends(get_db)):
#     profile = db.query(Profile).first()
   
#     if not profile:
#         return templates.TemplateResponse("profile.html", {
#             "request": request, 
#             "profile": None  
#         })
#     return templates.TemplateResponse("profile.html", {"request": request, "profile": profile})



########################################################################################
    #                            HTTP GET Requests
########################################################################################
@app.get("/")
async def show_projects(request: Request):
    return templates.TemplateResponse("index.html", {"request": request})

@app.get("/projects")
async def get_projects(request:Request):
    Session = sessionmaker(bind=engine)

    db = Session()
    projects = db.query(Projects).all()
    db.close()

    return templates.TemplateResponse("projects/project.html",{"request":request, "projects": projects})


@app.get("/profile")
async def get_profile(request: Request):

    Session = sessionmaker(hind=engine)
    db = Session()
    profile = db.query(Profile).all
    
    return templates.TemplateResponse("profile.html", {"request" : request, "profile" : profile})

@app.get("/skills")
async def get_skills(request: Request):

    Session = sessionmaker(hind=engine)
    db = Session()
    skills = db.query(Skills).all()
    return templates.TemplateResponse("skills.html", {"request": request, "skills":skills})


@app.get("/certifications")
async def get_certifications(request:Request):

    Session = sessionmaker(hind=engine)
    db = Session()
    certification = db.query(Certifications).all()
    return templates.TemplateResponse("certifications.html", {"request":request, "certifications": certification})


@app.get("/add-project")
async def add_project_form(request:Request):
    return templates.TemplateResponse("projects/addproject.html")



##############################################################################################











@app.get("/resume")
async def read_resume(request: Request, db: Session = Depends(get_db)):
    profile = db.query(Profile).first()
    if not profile:
        return templates.TemplateResponse("resume.html", {
            "request": request, 
            "profile": None  
        })
    return templates.TemplateResponse("resume.html", {"request": request, "profile": profile})


################################################################################################################

@app.post("/profile/")
def create_profile(profile: ProfileCreate, db: Session = Depends(get_db)):
    db_profile = Profile(**profile.dict())
    db.add(db_profile)
    db.commit()
    db.refresh(db_profile)
    return db_profile

# @app.post("/project/add")
# def create_project(projectRequestDTO: ProjectRequestDTO, db: Session ):
#     db_project = Projects(**projectRequestDTO.dict())
#     db.add(db_project)
#     db.commit()
#     db.refresh(db_project)
#     return db_project



################################################################################################################
########################################################



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

