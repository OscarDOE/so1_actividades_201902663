import { useEffect, useState } from 'react';
import { useNavigate } from 'react-router';
import '../App.css'

export const History = () =>{
    const [operations,setOperations] = useState([]);

    const navigate  = useNavigate();
    const calculadora = () =>{
        navigate('/calculate')
    };


    return (
        <div>
            <h1>HISTORIAL</h1>
            <div className='grid-history'>
                <div className='header-history'>Número 1</div>
                <div className='header-history'>Operación</div>
                <div className='header-history'>Número 2</div>
                <div className='header-history'>Resultado</div>
                <div className='header-history'>Fecha</div>
                <div>1</div>
                <div>1</div>
                <div>1</div>
                <div>1</div>
                <div>1</div>
                <div>1</div>
                <div>1</div>
                <div>1</div>
                <div>1</div>
                <div>1</div>
            </div>

            <button onClick={calculadora} className='btn btn-primary'>Calculadora</button>
        </div>
    )
}