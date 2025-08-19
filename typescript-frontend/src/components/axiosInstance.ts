import axios from "axios"


const axiosInstance = axios.create({
    baseURL: "http://100.88.110.98:8084/",
    headers: {
        "Content-Type": "application/json",
        "Accept": "application/json",
    },
});

export default axiosInstance;