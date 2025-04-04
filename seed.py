# seed.py
from sqlalchemy.orm import Session
from database import engine, SessionLocal
from models import Profile

def create_default_profile():
    db = SessionLocal()
    
    default_profile = Profile(
        full_name="Your Name",
        address="123 Main St, City",
        email="your.email@example.com",
        phone="(123) 456-7890",
        linkedin="https://linkedin.com/in/yourprofile",
        github="https://github.com/yourusername",
        summary="Experienced software developer specializing in..."
    )
    
    try:
        db.add(default_profile)
        db.commit()
        print("✅ Default profile created successfully!")
    except Exception as e:
        db.rollback()
        print(f"❌ Error creating profile: {e}")
    finally:
        db.close()

if __name__ == "__main__":
    create_default_profile()