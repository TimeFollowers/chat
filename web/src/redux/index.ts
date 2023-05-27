import {Action, ThunkAction, configureStore} from '@reduxjs/toolkit'

// 导入reducer
import counterReducer from './counter/countSlice'

import loginReducer from './login/loginSlice'

import chatReducer  from './chat/chatSlice'

import { TypedUseSelectorHook, useDispatch, useSelector } from 'react-redux'

export const store = configureStore ({
    // 合并reducer
    reducer: {
        counter: counterReducer,
        login: loginReducer,
        chat: chatReducer
    },
})


// 全局定义dispatch 和 sate 的类型，并导出
// 后面使用过程中直接从该文件中引入，而不需要冲react-redux包中引入
export const useAppDispatch: () => AppDispatch = useDispatch
export const useAppSelector: TypedUseSelectorHook<RootState> = useSelector

export type AppDispatch = typeof store.dispatch;
export type RootState = ReturnType<typeof store.getState>;
export type AppThunk<ReturnType = void> = ThunkAction<
  ReturnType,
  RootState,
  unknown,
  Action<string>
>;