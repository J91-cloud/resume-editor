import './App.css';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import "../src/styles/styles.scss"
import Home from "./pages/home/Home"
import Traits from "./pages/traits/Traits"

function App() {
  return (
    <>
    <Router>
    <section className="navbar-section">
      <div className="navbar-right">
        <ul className="un-list">
          <li><a href="/">Home</a></li>
          <li><a href="/traits">Traits</a></li>
          <li><a href="/builder">Build-Resume</a></li>
        </ul>
      </div>

    </section>

    <Routes>
      <Route path='/' element={<Home/>}/>
      <Route path='/traits' element={<Traits/>}/>
      <Route path='/builder' element={<h1>Builder Page</h1>}/>
    </Routes>
    </Router>

    <section className="footer-section">
      <div className="container">
        <h1>Footer Section of Application</h1>
        <p>Please provide all of your footer functions and applications here.</p>
      </div>
    </section>
    
    </>

    
    
   
  );
}

export default App;
