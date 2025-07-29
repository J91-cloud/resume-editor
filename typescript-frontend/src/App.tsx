import './App.css';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import "../src/styles/styles.scss"
import Home from "./pages/jobs/Home"

function App() {
  return (
    <>
    <Router>
    <section className="navbar-section">
      <div className="navbar-right">
        <ul className="un-list">
          <li><a href="/home">Home</a></li>
          <li><a href="/traits">Traits</a></li>
          <li><a href="/builder">Build-Resume</a></li>
        </ul>
      </div>

    </section>

    <Routes>
      <Route path='/home' element={<Home/>}/>
    </Routes>
    </Router>

    <section className="footer-section">
      <div className="container">
        <h1>Footer Page</h1>
      </div>
    </section>
    
    </>

    
    
   
  );
}

export default App;
