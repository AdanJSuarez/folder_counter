import React from 'react';
import logo from './logo_qohash.png';
import './App.css';
import InputFolder from './components/InputFolder'

function App() {
  return (
    <div className="App">
      <header className="App-header">
        <img src={logo} className="App-logo" alt="logo" />
        <p>
          Write the folder name, or leave it empty for actual folder:
        </p>
          <InputFolder/>
        <a
          className="App-link"
          href="https://qohash.com/"
          target="_blank"
          rel="noopener noreferrer"
        >
          Qohash.com
        </a>
      </header>
    </div>
  );
}

export default App;
