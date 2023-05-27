const router = {
    path: '/im',
    name: 'im',
    icon: 'wechat', // 显示图标
    routes: [
      {
        path: '/im', // 访问路由地址
        component: '../layouts/ChatLayout', // 聊天界面使用聊天布局
        routes: [
          {
            path: '/im', // 初始化界面
            component: './Im',
          },
          {
            path: '/im/chat', // 具体聊天界面显示
            name: 'chat',
            component: './Im/chat', // 对应组件
          },
        ],
      },

      {
        path: '/blog',
        component: '../layouts/Blog'
      }
    ],
  }
export default router
  