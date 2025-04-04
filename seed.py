# seed.py
from sqlalchemy.orm import Session
from database import engine, SessionLocal
from models import Profile,Projects,Skills,Certifications

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

def create_default_projects():
    db = SessionLocal()
    default_project = Projects(
      full_name="Champlain Pet Clinic",
      start_date="September 2024",
      end_date="December 2024",
      small_description="Fictional Pet Clinic utilised by all students, has all crud operations and more for a tipical client for a pet clinic, worked with a group of 30 students.",
      technologies_used="Used React-Typescript, with MongoDB, SpringBoot Java, Spring Security, and C#",
      role_of_project = "Took the role fo Developer, Porject manager and scrum manaster in a team of 4, following the Scrum Methodology. Used Jira to help with this."  
    )
    try:
        db.add(default_project)
        db.commit()
        print("Default Project has been created")
    except Exception as e:
        db.rollback()
        print("Error Creating default profile")
    finally:
        db.close()   

def create_default_skills():
    db = SessionLocal()
    default_skills1 = Skills(
        skill_name="Spring Boot Java"
    )
    default_skills2 = Skills(
        skill_name="HTML/CSS"
    )
    default_skills3 = Skills(
        skill_name="JavaScript"
    )
    try:
        db.add(default_skills1,default_skills2,default_skills3)
        db.commit()
        print("Skills have now been added")
    except Exception as e:
        db.rollback()
        print("Error creating skills")
    finally:
        db.close()

def create_default_certifications():
    db = SessionLocal()
    default_certification = Certifications(
        full_name= "Advanced React (Meta)",
        start_date= "November 2025",
        end_date= "December 2025"
    )
    try:
        db.add(default_certification)
        db.commit()
        print("Skills have now been added")
    except Exception as e:
        db.rollback()
        print("Error creating skills")
    finally:
        db.close()


if __name__ == "__main__":
    create_default_profile()
    create_default_projects()
    create_default_certifications()
    create_default_skills()
    