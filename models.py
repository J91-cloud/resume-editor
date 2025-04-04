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