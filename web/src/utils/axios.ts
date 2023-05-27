import axios, { AxiosResponse } from "axios";

var req = axios.create({
    baseURL: "http://127.0.0.1:8080"
})
//响应拦截器
// req.interceptors.response.use(
//     res => res.data, // 拦截器响应对象，将响应对象的的data属性返回给调用的地方
//     err => Promise.reject(err)
// )
// response 拦截器
req.interceptors.response.use(
    (response:AxiosResponse) => {
        // return Promise.resolve(response.data)
        return response
    },
    (error: any) => {

        return Promise.reject(error)
    }
)

const token = localStorage.getItem("token")
console.log("token:"+token)
if (token) {
    req.defaults.headers.common["authorization"] = token;
}

const setAuthToken = (token:string) => {
    if (token) {
        console.log(token)
      // headers 每个请求都需要用到的
      req.defaults.headers.common["authorization"] = token;
    } else {
      delete req.defaults.headers.common["authorization"];
    }
  }

export {req, setAuthToken}