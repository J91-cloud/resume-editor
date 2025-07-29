import React from "react";
import "./Home.css"
import suitcase from "../../styles/suitcase.svg"
export default function Home() {
    return(
        <>
            <section className="introduction-section">
                <div className="container">
                    <div className="row">
                        <div className="col-md-6 col-12 extra-padding">

                          
                                 <h1>Jobs Applied</h1>
                                 <p>32</p>
                                 <p className="sub-heading">The goal is for you to apply to as many jobs as possible, do not hesitate to apply as much as you can. </p>
                           
                           
                        </div>
                        <div className="col-md-6 col-12">
                            <img src="https://www.artm.quebec/wp-content/uploads/2023/11/ARTM-CHRONO-ILLU-2.png" alt="" />
                        </div>
                    </div>
                </div>
            </section>

            <section className="importance-section">
                <div className="container">
                    <div className="row">
                        <div className="col-md-6 col-12">
                            <img src={suitcase} alt="suitcase-image" />
                        </div>
                        <div className="col-md-6 col-12">
                            <div className="text-block">
                            <p>The purpose for this web application is that i found myself always needing to create or edit my original resume since i was applying to different types of tech jobs. This resume was for me to select the skills, projects, certifcates i need and then add it to a pdf tmeplate i can download right away. This was a much better experience.</p>

                            </div>
                        </div>
                    </div>
                </div>
            </section>


            <section className="jobs-section">
                <div className="container">
                    <h1>All Jobs Applied in Detail</h1>
                    <div className="row">
                        <div className="col-md-6 col-12">
                            <p>These are all of the locations i have applied to in detail. Please remove the ones that you already have an answer from. Meaning if they rejected you, please delete them from the database.</p>
                            <button className="btn btn-warning">Add A New Application</button>
                        </div>
                        <div className="col-md-6 col-12"></div>
                    </div>
                </div>
            </section>
            
            
            
        </>
    )
}