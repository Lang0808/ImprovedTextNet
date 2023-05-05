import logo from './logo.svg';
import './App.css';
import {
  BrowserRouter as Router,
  Routes,
  Route,
  Link
} from 'react-router-dom';
import Navbar from "./page/Navbar";
import Login from "./page/Login";
import Home from "./page/Home";

function App() {
  return (
    <Router>
      <div id="page-body">
        <Navbar></Navbar>
        <Routes>
          <Route path='/login' element={<Login />} />
          <Route path="/" element={<Home/>}/>
        </Routes>
      </div>
    </Router>
  )
}

export default App;
