import React from 'react';
import Navbar from './Components/Navbar'
import CreateMeeting from './Components/CreateMeeting'
import EditMeeting from './Components/EditMeeting'
import './App.css'
import {
  BrowserRouter as Router,
  Route
} from "react-router-dom";
function App() {
  return (
    <Router>
      <div className ="App">
        <Navbar></Navbar>
        <Route path="/" component={CreateMeeting} exact/>
        <Route path="/meeting" component={EditMeeting} exact/>
      </div>
    </Router>
  );
}

export default App;