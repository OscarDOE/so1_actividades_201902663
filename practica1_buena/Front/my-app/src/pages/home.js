import '../App.css'
import logo from '../logo.svg'
import 'bootstrap/dist/css/bootstrap.css';
import 'bootstrap/dist/js/bootstrap.js';

import React, {useEffect, useState} from 'react';
import { useNavigate } from 'react-router-dom';

export const Home = () =>{
    const navigate  = useNavigate();
    const calculadora = () =>{
        navigate('/calculate')
    };

    return (
        <div>
            <img src={logo} className="App-logo" alt="logo" />
            <p>
                Calculadora<br></br> 
                FrontEnd -- <code>React</code> <br></br>
                BackEnd -- <code>Go</code> <br></br>
                DB -- <code>MySQL</code> <br></br>

            </p>
            <a
                className="App-link"
                href="https://reactjs.org"
                target="_blank"
                rel="noopener noreferrer"
            >
                Learn React
            </a>
            <h1></h1>
            <button onClick={calculadora} className='btn btn-primary'>Calculadora</button>
        </div>
    )
}
