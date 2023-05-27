import React from 'react';
import { Button, Checkbox, Form, Input } from 'antd';
import { loginAsync } from '../../redux/login/loginSlice';
import { useAppDispatch } from '../../redux';

import './index.css'
export default function Login() {
  const dispatch = useAppDispatch()
  const onFinish = (values: any) => {
      console.log('Success:', values);
      
      dispatch(loginAsync(values))
  };
    
  const onFinishFailed = (errorInfo: any) => {
    console.log('Failed:', errorInfo);
  };
    

  return (
    <div className="login-warp">
      <Form
        name="basic"
        labelCol={{ span: 8 }}
        wrapperCol={{ span: 16 }}
        style={{ maxWidth: 600 }}
        initialValues={{ remember: true }}
        onFinish={onFinish}
        onFinishFailed={onFinishFailed}
        autoComplete="off"
      >
        <Form.Item
          label="Username"
          name="username"
          rules={[{ required: true, message: 'Please input your username!' }]}
        >
          <Input />
        </Form.Item>

        <Form.Item
          label="Password"
          name="password"
          rules={[{ required: true, message: 'Please input your password!' }]}
        >
          <Input.Password />
        </Form.Item>

        <Form.Item name="remember" valuePropName="checked" wrapperCol={{ offset: 8, span: 16 }}>
          <Checkbox>Remember me</Checkbox>
        </Form.Item>

        <Form.Item wrapperCol={{ offset: 8, span: 16 }}>
          <Button type="primary" htmlType="submit">
            Submit
          </Button>
        </Form.Item>
      </Form>
    </div>
    
  )

  
};