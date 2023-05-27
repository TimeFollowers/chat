import { Layout, List, Space } from "antd"
import { useState } from "react"

const Blog = () => {
    const {Header, Content, Footer} = Layout
    const [blogList, setlBlogList] = useState([])

    return (

        <Space direction="vertical" style={{ width: '100%' }} size={[0, 48]}>
            <Layout>
                <Header>header</Header>
                <Content>
                    <List>
                        <List.Item></List.Item>
                    </List>
                </Content>
                <Footer>footer</Footer>
            </Layout>
        </Space>
    )
}

export default Blog