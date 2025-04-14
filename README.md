<p align="center">
  <img style="height: 220px;" src="https://external-content.duckduckgo.com/iu/?u=http%3A%2F%2Fclipground.com%2Fimages%2Fresume-clipart-2.jpg&f=1&nofb=1&ipt=74ce2ceb9ee664d6fbb9e83661b96196ec7e97ce93ae27949f5ba2c06eae580a&ipo=images" />
</p>
<p align="center">
<a href="https://www.linkedin.com/in/baptiste-fernandez-%E5%B0%8F%E7%99%BD-0a958630/" target="blank"><img src="https://img.shields.io/badge/LinkedIn-0077B5?style=for-the-badge&logo=linkedin&logoColor=white" alt="https://img.shields.io/badge/LinkedIn-0077B5?style=for-the-badge&logo=linkedin&logoColor=white"  /></a>
  <img src="https://img.shields.io/badge/LICENSE-Apache-green"  />
</p>

<p align="center">
  My Resume Builder 
</p>


Have you ever been so unmotivated to edit your resume, always having to keep the styling the same and trying to make the most space as possible? Here is the ultimate resume builder. Made with `fastAPI` , hosted with `cloudflare tunnels` and stored with an SQLite database. I used my raspberry pi to host it, if you have your own home server or anything able to host a web server, dont hesitate to use, this applucation is bery lightweight. 


## Starting the project
Like always when running with a python project it is always recommended to create all your dependencies in a virtual environement. Activate a venv environment:

` venv\Scripts\activate`

After install the necesaary modules:

`pip install fastapi uvicorn sqlalchemy python-multipart jinja2 `. 

To reload the project, meaning your main method please run it with univcorn:
`uvicorn main:app --reload`