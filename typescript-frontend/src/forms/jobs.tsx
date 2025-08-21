import React, { useState } from "react";
import axiosInstance from "../components/axiosInstance";

export default function Jobs() {
    const [formData, setFormData] = useState({
        name: '',
        date_applied: '',
        job_type: ''
    });

    const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
        const { name, value } = e.target;
        setFormData(prevState => ({
            ...prevState,
            [name]: value
        }));
    };

    const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault();
        try {
            const response = await axiosInstance.post("/jobs", formData);
            console.log("Success:", response.data);
            // Optional: clear the form or show a success message
            setFormData({
                name: '',
                date_applied: '',
                job_type: ''
            });
        } catch (error) {
            console.error("Error:", error);
        }
    };

    return (
        <form onSubmit={handleSubmit} method="post">
            <div>
                <label htmlFor="name">Name</label>
                <input 
                    type="text" 
                    id="name" 
                    name="name" 
                    value={formData.name} 
                    onChange={handleChange} 
                />
            </div>
            <div>
                <label htmlFor="date_applied">Date Applied</label>
                <input 
                    type="date" 
                    id="date_applied" 
                    name="date_applied" 
                    value={formData.date_applied} 
                    onChange={handleChange} 
                />
            </div>
            <div>
                <label htmlFor="job_type">Job Type</label>
                <input 
                    type="text" 
                    id="job_type" 
                    name="job_type" 
                    value={formData.job_type} 
                    onChange={handleChange} 
                />
            </div>
            <button type="submit">Submit</button>
        </form>
    );
}