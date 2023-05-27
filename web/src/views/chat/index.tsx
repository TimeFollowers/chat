import { Avatar, Layout, List, Menu, Skeleton } from "antd";
import Sider from "antd/es/layout/Sider";
import { Content } from "antd/es/layout/layout";
import React from "react";
import './index.css'
import TextArea from "antd/es/input/TextArea";
import { store, useAppDispatch, useAppSelector } from "../../redux";
import { getUserListAsync } from "../../redux/chat/chatSlice";

interface IMessage {
  sendId: number; // 发送者id
  recvId: number; // 接收者id
  content: string; // 消息内容
  picture: string; // 头像

}
interface IUser {
  Id: number,
  username: string
}
interface IState {
  listMessage: IMessage[];
  websocket: WebSocket;
  userlist: IUser[];
}




export class ChatView2 extends React.Component<any, IState> {
  // dispath = useAppDispatch()
  readonly state:IState = {
    listMessage: [
      {
        sendId:1,
        recvId:2,
        content: "你好，世界",
        picture: "https://c-ssl.duitang.com/uploads/item/202003/16/20200316100626_rqpov.jpeg"
      }
    ],
    websocket: {} as WebSocket,
    userlist: []
  }

  userlist = store.getState().chat.userList
  onPressEnter = (e: React.KeyboardEvent<HTMLTextAreaElement>) => {
    e.preventDefault()
    let message = e.currentTarget.value
    this.state.websocket.send(message)
  }

  componentDidMount(): void {
      if (typeof(WebSocket) == "undefined") {
        console.log("您的浏览器不支持WebSocket")
      } else {
        console.log("您的浏览器支持WebSocket")
      }
      let websocket = new WebSocket("ws://127.0.0.1:8080/ws")
      // 打开事件
      websocket.onopen = function() {
        console.log("websocket已打开")
      }
      //接收消息
      websocket.onmessage = function(msg) {
        console.log("websocket已连接")
        console.log(msg.data) //第一次进入会显示连接成功
      }
      //关闭事件
      websocket.onclose = function() {
        console.log("websocket已关闭")
      }
      // 发生了错误事件
      websocket.onerror = function() {
        console.log("websocket发生了错误")
      }

      this.setState({websocket: websocket})
      // this.props.dispatch(getUserListAsync)
      // this.dispath(getUserListAsync)
      
      
      console.log(this.userlist)
      
      
  }
  render(): React.ReactNode {
    return (
      <Layout>
        <Sider
          breakpoint="lg"
          collapsedWidth="0"
          onBreakpoint={broken => {
            console.log(broken)
          }}
          onCollapse={(collapse, type) => {
            console.log(collapse, type)
          }}
          theme="light"
        >
          <div className="logo"></div>
          <Menu theme="light" mode="inline" defaultSelectedKeys={['4']}>
            {this.userlist.map((item, index) => <Menu.Item key={item.id}>
          <span className="nav-text">{item.user_name}</span>
        </Menu.Item>)}
            
          </Menu>
        </Sider>
        <Layout>
          {/* <Header style={{background: '#fff', padding:0}}/> */}
          <Content style={{margin: '24px 16px 0', position: "relative"}}>
            {/* <div style={{ padding: 24, background: '#fff', minHeight: 360 }}> */}
              <List
                dataSource={this.state.listMessage}
                renderItem={(item) => (
                  <List.Item
                    actions={[<a key="list-loadmore-edit" href="/">edit</a>, <a key="list-loadmore-more" href="/">more</a>]}
                  >
                    <Skeleton avatar title={false} loading= {false} active>
                      <List.Item.Meta 
                        avatar={<Avatar src={item.picture}/>}
                        title = {<a href="/">{item.sendId}</a>}
                      />
                      <div>{item.content}</div>
                    </Skeleton>

                  </List.Item>
                )}
              ></List>
              {/* <div style={{position: "absolute", bottom: "10px"}}> */}
                <TextArea onPressEnter={this.onPressEnter} style={{position: "absolute", bottom: "10px"}} placeholder="Autosize height based on content lines" autoSize />
              {/* </div> */}
              
            {/* </div> */}
          </Content>
          {/* <Footer style={{ textAlign: 'center' }}>Ant Design ©2018 Created by Ant UED</Footer> */}
        </Layout>
      </Layout>
    )
  }
}