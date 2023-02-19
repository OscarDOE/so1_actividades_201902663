import { useEffect, useState } from 'react';
import { useNavigate } from 'react-router';
import '../App.css'

export const Calculator = () =>{
    const [presionado, setPresionado] = useState([]);
    const [reserva, setReserva] = useState([]);
    const [lnum, setLnum] = useState([]);
    const [rnum, setRnum] = useState([]);
    const navigate = useNavigate();
    let operation = ""
    let isnumber = false
    let textarea = presionado

    const historial = async () =>{
        navigate('/history')
    }

    const press = async (e,f) =>{
        textarea = presionado
        if(e == "DEL"){
            await setPresionado("")
            await setReserva("")
            await setLnum(0)
            await setRnum(0)

        }else if(e == "="){
            let aux = presionado
            await setRnum(parseInt(aux))
            let res = lnum+rnum
            await setPresionado("")
            await setReserva(res)


        }else if(e == "+" || e == "-" || e == "*" || e == "/"){
            let aux = presionado
            await setLnum(parseInt(aux))
            await setReserva(presionado+e)
            await setPresionado("")

        }else{
            textarea = presionado+e
            await setPresionado(textarea)
        }
        textarea = presionado
    }
    return (
        <div>
            <h1>CALCULADORA</h1>
            <div className='grid-container'>
                <h1 className='area-result'>{textarea}</h1>
                <button className='calc-del' onClick={(e) => press("DEL")}>DEL</button>
                <button onClick={(e) => press("7",0)}>7</button>
                <button onClick={(e) => press("8",0)}>8</button>
                <button onClick={(e) => press("9",0)}>9</button>
                <button onClick={(e) => press("/",0)}>/</button>
                <button onClick={(e) => press("4",0)}>4</button>
                <button onClick={(e) => press("5",0)}>5</button>
                <button onClick={(e) => press("6",0)}>6</button>
                <button onClick={(e) => press("*",0)}>*</button>
                <button onClick={(e) => press("1",0)}>1</button>
                <button onClick={(e) => press("2",0)}>2</button>
                <button onClick={(e) => press("3",0)}>3</button>
                <button onClick={(e) => press("+",0)}>+</button>
                <button onClick={(e) => press("0",0)}>0</button>
                <button onClick={(e) => press(".",0)}>.</button>
                <button onClick={(e) => press("=",0)}>=</button>
                <button onClick={(e) => press("-",0)}>-</button>
                <h1 className='area-result' >{reserva}</h1>
            </div>
            <div style={{paddingTop:"20px"}}>
                <button style={{color:"red"}} onClick={historial} className='btn btn-secondary'> Historial </button>
            </div>
        </div>
    )
};