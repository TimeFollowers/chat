import {req} from "../utils/axios";


export function fetchLogin(params: string) {
    // const token = localStorage.getItem("token")
    return req.post("/login", params, {
        headers: {
            'Content-Type': 'application/x-www-form-urlencoded',
            // 'authorization': token
        }
    })
}

export function fetchUserList() {
    return req.get("/u/user/list")
}
