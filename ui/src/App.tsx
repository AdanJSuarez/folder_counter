import React from 'react';
import logo from './assets/logo_qohash.png';
import './App.css';
import InputFolder from './components/InputFolder'

function App() {
  return (
    <div className="App">
      <header className="App-header">
        <img src={logo} className="App-logo" alt="logo" />
      </header>
      <p>
          Write the folder name, or leave it empty for actual folder:
        </p>
      <InputFolder folderName={""}/>
    </div>
  );
}

export default App;
