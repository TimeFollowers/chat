import {req} from "../utils/axios";


export function fetchMessageList() {
    console.log("==========fetchMessageList  ===================")
    return req.get("/u/chat/list?room_id=1")
}

export function fetchUserDetail() {
    return req.get("/u/user/detail")
}