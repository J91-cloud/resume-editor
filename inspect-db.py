# inspect_db.py
from sqlalchemy import create_engine, inspect
from sqlalchemy.orm import sessionmaker
from models import Profile  # Import your models
from database import SQLALCHEMY_DATABASE_URL
from rich.console import Console
from rich.table import Table

def display_profile():
    engine = create_engine(SQLALCHEMY_DATABASE_URL)
    Session = sessionmaker(bind=engine)
    db = Session()
    
    console = Console()
    
    # Check if profiles table exists
    inspector = inspect(engine)
    if "profiles" not in inspector.get_table_names():
        console.print("[red]❌ Profiles table doesn't exist![/red]")
        return
    
    # Get all profiles
    profiles = db.query(Profile).all()
    
    if not profiles:
        console.print("[yellow]⚠️ No profiles found in database[/yellow]")
        return
    
    # Create a rich table
    table = Table(title="Profile Data", show_header=True, header_style="bold magenta")
    
    # Add columns (using the first profile as reference)
    columns = Profile.__table__.columns.keys()
    for col in columns:
        table.add_column(col, style="cyan")
    
    # Add rows
    for profile in profiles:
        table.add_row(*[str(getattr(profile, col)) for col in columns])
    
    console.print(table)
    db.close()

if __name__ == "__main__":
    display_profile()