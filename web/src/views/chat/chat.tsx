import React, { useEffect } from "react";
import { useAppDispatch, useAppSelector } from "../../redux";
import { Avatar, Divider, Input, Layout, List, Skeleton, Space } from "antd";
import Sider from "antd/es/layout/Sider";
import { Content } from "antd/es/layout/layout";
import TextArea from "antd/es/input/TextArea";
import { addMessage, getMessageListAsync, getUserDetailAsync, getUserListAsync } from "../../redux/chat/chatSlice";
import { IMessage } from "../../types/chat";
import Avatar2 from './f07c26248feef6f7a08374b607654e09.jpeg';
import InfiniteScroll from "react-infinite-scroll-component";
const token = localStorage.getItem("token")
const websocket = new WebSocket("ws://127.0.0.1:8080/u/ws")
let recvId = 0;
let sendId = 1;





function ChatView() {
    const dispatch = useAppDispatch()
    const userlist = useAppSelector((state) => state.chat.userList)

    const messagelist = useAppSelector((state) => state.chat.messageList)
    const onPressEnter = (e: React.KeyboardEvent<HTMLTextAreaElement>) => {
      e.preventDefault()
      let content = e.currentTarget.value
      let message:IMessage = {
          content,
          recvId,
          sendId,
      }
      // messagelist.push(message)
      dispatch(addMessage(message))
      console.log(message)
      websocket.send(JSON.stringify(message))
      e.currentTarget.value = ""
      console.log("清空消息")
    }
    const chooseUser = (id: number) => {
      recvId = id
      console.log("recvid:"+recvId)
      dispatch(getMessageListAsync())
    }

    useEffect(() => {
        dispatch(getUserDetailAsync())
        dispatch(getUserListAsync())
        // dispatch(getMessageListAsync())
    },[])
    
    if (typeof(WebSocket) == "undefined") {
        console.log("您的浏览器不支持WebSocket")
    } else {
      console.log("您的浏览器支持WebSocket")
    }
    // let websocket = new WebSocket("ws://127.0.0.1:8080/u/ws")
    // 打开事件
    websocket.onopen = function() {
      const message:IMessage = {
        sendId: sendId,
        recvId: recvId,
        content: "权限验证",
        token: token || "",
      }

      websocket.send(JSON.stringify(message))
      console.log("websocket已打开")
    }
    //接收消息
    websocket.onmessage = function(msg) {
      console.log("websocket已连接")

      const message = JSON.parse(msg.data)
      dispatch(addMessage(message))
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


    const loadMoreData = () => {
      dispatch(getUserListAsync())
    }
    const onSearchPressEnter = (e :React.KeyboardEvent<HTMLInputElement>) => {
      console.log(e.currentTarget.value)
    }
    return (
        <Layout>
        <Sider
          breakpoint="lg"
          collapsedWidth="0"
          onBreakpoint={broken => {
            console.log("onBreakpoint")
          }}
          onCollapse={(collapse, type) => {
            console.log(collapse, type)
          }}
          theme="light"
        >
          <div className="logo"></div>

          <div
            id="scrollableDiv"
            style={{
              height: 800,
              overflow: 'auto',
              padding: '0 16px',
              border: '1px solid rgba(140, 140, 140, 0.35)',
            }}
          >
            {/* <Search placeholder="input search text" onSearch={onSearch} style={{ width: 150 }} /> */}
            <Space direction="vertical" size="middle">
              <Space.Compact>
                <Input onPressEnter={onSearchPressEnter} defaultValue="" />
              </Space.Compact>
            </Space>
          <InfiniteScroll
            dataLength={userlist.length}
            next={loadMoreData}
            hasMore={userlist.length < 2}
            loader={<Skeleton avatar paragraph={{ rows: 1 }} active />}
            endMessage={<Divider plain>It is all, nothing more 🤐</Divider>}
            scrollableTarget="scrollableDiv"
          >
            <List
              dataSource={userlist}
              renderItem={(item, index) => (
                <List.Item onClick={() => chooseUser(item.id)} >
                  <List.Item.Meta 
                    avatar={ <Avatar src={Avatar2}></Avatar> }
                    title={item.user_name}
                    />
                </List.Item>
              )}>
              
            </List>
          </InfiniteScroll>
          </div>
        </Sider>
        <Layout>
          <Content style={{margin: '24px 16px 0', position: "relative"}}>
          <div
            id="scrollableDiv"
            style={{
              height: 600,
              overflow: 'auto',
              padding: '0 16px',
              border: '1px solid rgba(140, 140, 140, 0.35)',
            }}
          >
            <InfiniteScroll
              dataLength={userlist.length}
              next={loadMoreData}
              hasMore={userlist.length < 2}
              loader={<Skeleton avatar paragraph={{ rows: 1 }} active />}
              endMessage={<Divider plain>It is all, nothing more 🤐</Divider>}
              scrollableTarget="scrollableDiv"
            >
              <List
                dataSource={messagelist}
                renderItem={(item) => (
                  <List.Item>
                    <Skeleton avatar title={false} loading= {false} active>
                      <List.Item.Meta 
                        avatar={<Avatar src={Avatar2}/>}
                        title = {item.content}
                      />
                    </Skeleton>

                  </List.Item>
                )}
              ></List>
            </InfiniteScroll>
          </div>
          <TextArea onPressEnter={onPressEnter} style={{position: "absolute", bottom: "10px"}} placeholder="Autosize height based on content lines" autoSize />

 
          </Content>
        </Layout>
      </Layout>
    )
}

export default ChatView