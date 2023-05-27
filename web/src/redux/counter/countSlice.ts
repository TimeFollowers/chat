import { createSlice } from "@reduxjs/toolkit";

// 设置类型
interface counterState {
    value: number
}
//设置初始化state
const initialState: counterState = {
    value: 0
}

export const counterSlice = createSlice({
    name: 'counter',
    initialState,

    reducers: {
        increment: state => {
            //Redux Toolkit 允许我们在reducers写"可变"逻辑。它并不是真正的改变状态值，因为它使用了Immer库
            //可以检测到"草稿状态"的变化并且基于这些变化生产全新的不可变的状态
            state.value += 1
        },
        decrement: state => {
            state.value -= 1
        },
        incrementByAmount:(state, action) => {
            state.value += action.payload
        }
    }
})

// 每个case reducer 函数会产生对应的action creators
export const {increment, decrement, incrementByAmount} = counterSlice.actions
export default counterSlice.reducer
