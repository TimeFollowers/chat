import { createAsyncThunk, createSlice } from "@reduxjs/toolkit";
import { fetchLogin } from "../../api/user";
import { setAuthToken } from "../../utils/axios";
export interface LoginState {
    username: string;
    password: string;
    status: 'idle' | 'loading' | 'failed'
    token: string
}

const initialState: LoginState = {
    username: "",
    password: "",
    status: 'idle',
    token: ""
}



export const loginAsync = createAsyncThunk(
    'login/fetchLogin',
    async (params:string) => {
        const response = await fetchLogin(params)
        console.log(response.data)
        return response.data
    }
)





const loginSlice = createSlice({
    name: 'login',
    initialState,
    reducers: {
        login: (state) => {
            console.log(state)
        },
        // setuSer: ((state, action) => {
        //     console.log(action)
        //     state.username = action.payload.username
        //     state.password = action.payload.password
        //     console.log(state)
        // })
    },
    extraReducers(builder) {
        builder.addCase(loginAsync.pending, (state) => {
            state.status = 'idle'
        }).addCase(loginAsync.fulfilled, (state, action) => {
            state.status = 'idle'
            console.log(action)
            state.username = action.payload.username
            state.password = action.payload.password
            state.token = action.payload.token
            localStorage.setItem("token", action.payload.data.token)
            window.location.href ="/chat"
            setAuthToken(action.payload.data.token)
            console.log("username:"+state.username)
        }).addCase(loginAsync.rejected,(state) => {
            state.status = 'failed'
        })

    }
})

export const {login} = loginSlice.actions;
export default loginSlice.reducer;