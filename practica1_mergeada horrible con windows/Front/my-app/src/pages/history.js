import { useEffect, useState } from 'react';
import { useNavigate } from 'react-router';
import '../App.css'

import 'bootstrap/dist/css/bootstrap.css';
import 'bootstrap/dist/js/bootstrap.js';

export const History = () =>{
    const [operations,setOperations] = useState([]);

    const navigate  = useNavigate();
    const calculadora = () =>{
        navigate('/calculate')
    };

    useEffect(()=>{
        fetch('http://localhost:5005/history')
        .then(res => res.json())
        .then(res =>setOperations(res))
            .catch(err => console.log(err))
    },[])


    return (
        <div>
            <h1>HISTORIAL</h1>
            <div>
                <div className='grid-history'>
                    <div className='header-history'>Número 1</div>
                    <div className='header-history'>Operación</div>
                    <div className='header-history'>Número 2</div>
                    <div className='header-history'>Resultado</div>
                    <div className='header-history'>Fecha</div>
                </div>
                {operations.map(function(operacion,index){
                    return(
                        <div className='grid-history'>
                            <div>{operacion.num1}</div>
                            <div>{operacion.operation}</div>
                            <div>{operacion.num2}</div>
                            <div>{operacion.resultado}</div>
                            <div>{operacion.fecha}</div>
                        </div>
                    )
                })}

            </div>

            <button onClick={calculadora} className='btn btn-primary'>Calculadora</button>
        </div>
    )
}