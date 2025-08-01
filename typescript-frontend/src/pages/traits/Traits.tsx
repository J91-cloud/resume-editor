import react, { useEffect, useState } from 'react';
import {project} from "../../models/project";
import {certificate} from "../../models/certificates"
import {skill} from "../../models/skill"
import folderIMG from "../../styles/projects-folder.svg"
import axiosInstance from '../../components/axiosInstance';
import "./Traits.css"

export default function Traits(){

    const [skills, setSkills] = useState<skill[]>([]);
    const [projects, setProjects] = useState<project[]>([]);
    const [certificates,setCertificates] =  useState<certificate[]>([]);


    const FetchProjects = async () => {
        useEffect(() => {
            axiosInstance.get<project[]>('/projects')
            .then(response =>  setProjects(response.data))
            .catch(error => console.error("Error creating this", error));
        },[]);
    }

    const FetchSkills = async () => {
        useEffect(() => {
            axiosInstance.get<skill[]>('/skills')
            .then(response => setSkills(response.data))
            .catch(error => console.error("Error with skills fetching",error));
        },[])
    }

    const FetchCertificates = async () => {
        useEffect(() => {
            axiosInstance.get<certificate[]>('/certicates')
            .then(response => setCertificates(response.data))
            .catch(error => console.error("Error fetching certificates", error));
        }, [])
    }




    
    return (
        <>
        <section className="introduction-section-traits">
            <div className="container">
                <div className="row">
                    <div className="col-12">
                        <h1>My Traits</h1>
                        <p>This page displays all that i have learnt so i do not need to always change my resume.
                            This page will display all of my projects, skills, certificates that i can then add to my builder, so when i construct my resume it will showcase 
                            only the items i selected. 
                        </p>
                        <img src={folderIMG} alt="folder-image"  />
                    </div>
                </div>
            </div>
        </section>

        <section className="project-section-traits">
            <div className="container">
                <div className="row">
                    <div className="col-md-7 col-12">
                        <h1>My Projects</h1>
                        <p>This is a selection of all my projects</p>                        
                    </div>
                    <div className="col-md-5 col-12">
                        {/* Add an image of some sort */}
                    </div>
                </div>
                <div className="row">
                    <ul>
                        {projects.map(project => (
                            <li key={project.id}>{project.name}</li>
                        ))}
                    </ul>
                </div>
            </div>
        </section>

        
    
        
        
        </>
    )
}