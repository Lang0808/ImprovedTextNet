import "./Login.css";
import login_picture from "../resource/login_picture.png"
import { useState } from "react";
import axios from 'axios';
import {REGISTER_API, LOGIN_API} from "../api/api.js";
import { ToastContainer, toast } from 'react-toastify';
import 'react-toastify/dist/ReactToastify.css';

function Login(){

    const [username, setUsername]=useState("")
    const [password, setPassword]=useState("")

    function handleText(e){
        var id=e.target.id
        if(id=="username"){
            setUsername(e.target.value)
        }
        else if(id=="password"){
            setPassword(e.target.value)
        }
    }

    function register(){
        var data = new FormData();
        data.append('Username', username);
        data.append('Password', password);
        console.log(REGISTER_API)
        var config = {
            method: 'post',
            url: REGISTER_API,
            data : data
        };
        axios(config)
        .then(function (response) {
            console.log(JSON.stringify(response.data));
            data=response.data
            if (data.Err!=0){
                toast(data.Message)
            }
            else{

            }
        })
        .catch(function (error) {
            toast(error)
        });
    }

    function login(){
        var data = new FormData();
        data.append('Username', username);
        data.append('Password', password);
        console.log(LOGIN_API)
        var config = {
            method: 'post',
            url: LOGIN_API,
            data : data
        };
        axios(config)
        .then(function (response) {
            console.log(JSON.stringify(response.data));
            data=response.data
            if (data.Err!=0){
                toast(data.Message)
            }
            else{

            }
        })
        .catch(function (error) {
            toast(error)
        });
    }

    function handleClick(e){
        var id=e.target.id
        if(id=="register"){
            register()
        }
        else if(id=="login"){
            login()
        }
    }

    return (
        <div id="container">
            <div id="welcome-pic"><img src={login_picture} /></div>
            <div id="login-section">
                <div id="form-container">
                    <h2>Welcome !</h2>
                    <input type="text" name="username" id="username" placeholder="Your username" value={username} onChange={handleText}/>
                    <input type="text" name="password" id="password" placeholder="Your password" value={password} onChange={handleText}/>
                    <div id="button-container">
                        <button className="button-89" role="button" onClick={handleClick} id="register">Create account</button>
                        <button className="button-89" role="button" onClick={handleClick} id="login">Sign in</button>
                        <ToastContainer />
                    </div>
                    
                </div>
            </div>
        </div>
    )
}

export default Login;