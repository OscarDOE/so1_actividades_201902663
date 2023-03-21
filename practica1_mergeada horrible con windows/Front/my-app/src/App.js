import React, {useState, useEffect} from 'react';
import { BrowserRouter, Route, Routes } from 'react-router-dom';

import logo from './logo.svg';
import './App.css';

import {Home} from './pages/home';
import {Calculator} from './pages/calculator';
import {History} from './pages/history';

function App() {
  return (
    <div className="App">
      <header className="App-header">
        <BrowserRouter>
          <Routes>
            <Route path="/" element={<Home></Home>}/>
            <Route path="/calculate" element={<Calculator></Calculator>}/>
            <Route path="/history" element={<History></History>}/>
          </Routes>
        </BrowserRouter>
      </header>
    </div>
  );
}

export default App;
