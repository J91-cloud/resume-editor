import './App.css';
import "../src/styles/styles.scss"

function App() {
  return (
   <div className="container">
    <div className="row col-md-8 col-12">
        <h1 className='header'>Welcome to your Resume Editor</h1>
    </div>
    <div className="row col-md-4 col-12">
      
    </div>

    

    <div className="row mx-n4">
      <div className="col-12">
        <section className="welcome-section">
          <h1>Jobs Applied To</h1>
          <p>List all the jobs you applied to. Make sure you can delete some. And look on with a search.</p>
        </section>
      </div>
      <div className="col-md-6 col-12">
        <p>This is where you display a large number with how many jobs you applied to.</p>
      </div>
      <div className="col-md-6 col-12">
        <p>Show the 5 most recent jobs you applied to this week.  </p>
      </div>
    </div>


    <div className="row mx-n4">
      <div className="col-12">
      <section className="welcome-section">
        <div className="center">
          <button>Build your resume</button>
        </div>
        <p>This button will need to you having to build a route where you can select projects, skills and all other items to select to add to your resume. </p>
        
      </section>
       </div>
    </div>



    </div>
    
    
   
  );
}

export default App;
