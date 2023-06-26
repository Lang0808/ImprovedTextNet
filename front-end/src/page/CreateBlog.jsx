import { Editor } from "./Editor";
import "./CreateBlog.css";
import {CREATE_BLOG_API} from "../api/api.js";
import { ToastContainer, toast } from 'react-toastify';
import 'react-toastify/dist/ReactToastify.css';
import axios from 'axios';
import { useState } from "react";
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faEarth, faUserFriends, faArchive, faItalic} from '@fortawesome/free-solid-svg-icons';

function CreateBlog(props){
    // reminder: review editor when have freetime

    const [title, setTitle]=useState("")
    const [privacy, setPrivacy]=useState("0")

    // reminder: try to dynamic privacy modes
    const privacyModes=[
        {
            "name": "PUBLIC",
            "id": 0,
            "icon": faEarth
        },
        {
            "name": "FRIEND",
            "id": 1,
            "icon": faUserFriends
        },
        {
            "name": "PRIVATE",
            "id": 2,
            "icon": faArchive
        }
    ]

    function onSubmit(event){
        // submit create blog
        // step 1: get all nodes in element has id "editor-content"
        var listChildNodes=document.getElementById("editor-content").childNodes

        // step 2: concatenate all child nodes data
        var html=""
        for(var i =0;i<listChildNodes.length;i++){
            html+=listChildNodes[i].outerHTML
        }

        console.log(html)
        
        // step 3: create form data, submit
        var data = new FormData();
        data.append('Content', html);
        data.append("PrivacyMode", privacy);
        data.append("Title", title);
        console.log(CREATE_BLOG_API)
        console.log(data)
        var config = {
            method: 'post',
            url: CREATE_BLOG_API,
            data : data,
            withCredentials: true
        };
        // submit
        axios(config)
        .then(function (response) {
            data=response.data
            if (data.Err!=0){
                // if there is an error, toast data.Message
                toast(data.Message)
            }
            else{
                // otherwise, success
                toast("SUCCESS !!!")
            }
        })
        .catch(function (error) {
            // when an exception is caught, toaast it
            toast(error)
        });
    }

    function changeTitle(e){
        setTitle(e.target.value)
    }

    function changePrivacy(e){
        setPrivacy(e.target.value)
    }

    function changeContent(e){
        console.log("on change call")
        var listChildNodes=document.getElementById("editor-content").childNodes

        // step 2: concatenate all child nodes data
        var html=""
        for(var i =0;i<listChildNodes.length;i++){
            html+=listChildNodes[i].outerHTML
        }

        console.log(html)


        document.getElementById("display-content").innerHTML=html

    }

    // create option for select "privacy"
    // try to dynamic
    var options=
    <select id="select-container" value={privacy} onChange={changePrivacy}>
        <option value="0" className="option-item">
            &#xf0ac;
        </option>
        <option value="1" className="option-item">
            &#xf007;
        </option>
        <option value="2" className="option-item">
            &#xf023;
        </option>
    </select>
    
    return (
        <div>
            
            <div className="title-container">
                <input type="text" value={title} onChange={changeTitle} placeholder="Title"/>
            </div>
            <div className="privacy-container">
                {options}
            </div>
            <div className="content-container">
                <div className="max-w-[800px] mx-4 mt-2 mb editor-container">
                    <Editor onChange={changeContent} />
                </div>
                <div className="display">
                    <div className="display-title">{title}</div>
                    <div className="display-content" id="display-content"></div>
                </div>
            </div>
            
            <div id="button-container">
                <button className="button-89" role="button" onClick={onSubmit} id="register">Create blog</button>
                <ToastContainer/>
            </div>
        </div>
        
    );
}

export default CreateBlog;