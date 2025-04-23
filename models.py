# models.py
from sqlalchemy import Column, Integer, String
from database import Base

class Profile(Base):
    __tablename__ = "profiles"
    
    id = Column(Integer, primary_key=True, index=True)
    full_name = Column(String)
    address = Column(String)
    email = Column(String)
    phone = Column(String)
    linkedin = Column(String, nullable=True)
    github = Column(String, nullable=True)
    summary = Column(String, nullable=True)



class Projects(Base):
    __tablename__ = "projects"
    id = Column(Integer, primary_key=True, index=True)
    full_name = Column(String)
    start_date = Column(String)
    end_date = Column(String)
    small_description = Column(String)
    technologies_used = Column(String)
    role_of_project = Column(String)

class Skills(Base):
    __tablename__ = "skills"
    id = Column(Integer,primary_key=True,index=True)
    skill_name = Column(String)

class Certifications(Base):
    __tablename__ = "certifications"
    id = Column(Integer, primary_key=True,index=True)
    full_name = Column(String)
    start_date = Column(String)
    end_date = Column(String)         

  
  
