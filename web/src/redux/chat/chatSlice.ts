import { createAsyncThunk, createSlice } from "@reduxjs/toolkit";
import { fetchUserList } from "../../api/user";
import { fetchMessageList, fetchUserDetail } from "../../api/message";
import { IMessage } from "../../types/chat";
interface IUser {
    id: number,
    user_name: string
}


interface chatState {
    messageList: IMessage[];
    userList: IUser[];
    user: IUser
}
//设置初始化state
const initialState: chatState = {
    messageList: [],
    userList: [],
    user: {} as IUser
}

export const getUserListAsync = createAsyncThunk(
    'user/fetchUserList',
    async () => {
        const response = await fetchUserList()
        return response.data
    }
)
export const getMessageListAsync = createAsyncThunk(
    "/message/fetchMessageList",
    async () => {
        console.log("========getMessageListAsync============")
        const response = await fetchMessageList()
        return response.data
    }
)

export const getUserDetailAsync = createAsyncThunk(
    "/message/fetchDetail",
    async () => {
        const response = await fetchUserDetail()
        return response.data
    }
)


export const chatSilce = createSlice({
    name: 'chat',
    initialState,

    reducers: {
        addMessage: (state, action) => {
            //Redux Toolkit 允许我们在reducers写"可变"逻辑。它并不是真正的改变状态值，因为它使用了Immer库
            //可以检测到"草稿状态"的变化并且基于这些变化生产全新的不可变的状态
            state.messageList.push(action.payload)
            console.log(action.payload)
            
        },
        
    },

    extraReducers (builder) {
        builder.addCase(getUserListAsync.fulfilled, (state, action) => {
            state.userList = action.payload.data.users
        } )

        builder.addCase(getMessageListAsync.fulfilled, (state, action) => {
            state.messageList = action.payload.data.list
        })

        builder.addCase(getUserDetailAsync.fulfilled, (state, action) => {
            state.user.user_name = action.payload.data.user_name
            state.user.id = action.payload.data.id
        })

    }
})

// 每个case reducer 函数会产生对应的action creators
export const {addMessage} = chatSilce.actions
export default chatSilce.reducer
