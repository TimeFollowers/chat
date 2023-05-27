import React from 'react';
import ReactDOM from 'react-dom/client';
import './index.css';
import reportWebVitals from './reportWebVitals';
import { BrowserRouter, Route, Routes } from 'react-router-dom';
import App from './App';

import './utils/axios'
import { Provider } from 'react-redux';
import { store } from './redux';
import ChatView  from './views/chat/chat';
import Blog from './views/blog/blog';

const root = ReactDOM.createRoot(
  document.getElementById('root') as HTMLElement
);
root.render(
  <Provider store={store}>
    <BrowserRouter>
    <Routes>
      <Route path='/' element={<App/>} />
      <Route path='/chat' element={<ChatView />}></Route>
      <Route path='/blog' element={<Blog />}></Route>
    </Routes>
    
  </BrowserRouter>
  </Provider>
  

);

// If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
reportWebVitals();
