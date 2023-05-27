interface IUser {
    Id: number,
    username: string
}


export interface IChatState {
    messageList: IMessage[];
    userList: IUser[];
}

export interface IMessage {
    msgId?: number; // 消息id
    sendId: number; //发送人id
    recvId: number; //接收人id
    content: string //消息内容
    token?: string // token
}