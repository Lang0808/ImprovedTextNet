import { Editor } from "./Editor";
import "./CreateBlog.css";
import {CREATE_BLOG_API} from "../api/api.js";
import { ToastContainer, toast } from 'react-toastify';
import 'react-toastify/dist/ReactToastify.css';
import axios from 'axios';


function CreateBlog(props){

    function onSubmit(event){
        console.log("Clicked")
        var element=document.getElementById("editor-content")
        console.log(element)
        var data = new FormData();
        data.append('Content', element.outerHTML);
        data.append("PrivacyMode", "0");
        data.append("Title", "Test title");
        console.log(CREATE_BLOG_API)
        var config = {
            method: 'post',
            url: CREATE_BLOG_API,
            data : data,
            withCredentials: true
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
        // data.append('PrivacyMode', password);
    }
    
    return (
        <div>
            <div className="max-w-[800px] mx-4 mt-2 mb">
                <Editor/>
            </div>
            <div id="button-container">
                <button className="button-89" role="button" onClick={onSubmit} id="register">Create blog</button>
                <ToastContainer/>
            </div>
            <div id="post-container">

            </div>
        </div>
        
    );
}

export default CreateBlog;