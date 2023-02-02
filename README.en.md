# bluebell

#### Description
基于一套通用的web脚手架搭建的博客社区

#### Software Architecture
Software architecture description

#### Installation

1.  使用分布式ID生成器生成用户ID
    使用雪花算法  生成一个64为的整数 最高位不用 默认为0 41位为时间戳 10位为工作机器id 12位为序列号
    github.com/bwmarrin/snowflake
2.  添加注册业务
    1、参数获取和参数校验
    使用第三方库对参数进行校验
    validator
3.  将用户保存至数据库
    使用md5加盐的方式保存用户的密码

4. 基于Token的无状态会话管理方式
   使用jwt-go 这个库进行实现
   这里还有两个问题：1、refresh Token 双token的认证 这样就可以实现长时间的无状态会话
   2、限制用户在同一时间只能在一个设备登录
5. 为项目编写makefile
   我们可以把makefile简单的理解为它定义了一个项目文件的编译规则
6. 使用AIR 实现热重载
7. 添加查看社区的内容
8. 增加发帖功能
9. 增加常看帖子的详细内容的功能 不过这里只显示了作者的id
10. 增加获取帖子的详情的时候 可以显示作者的名称
11. 获取帖子列表的信息 实现结果分页展示
12. 有一个小问题 前端js number 类型能够表示的最大值 2^53 但是后端go int64能表示
的最大值是2^64-1 所有有可能后端传给前端的数据会出现失真的情况，它的解决办法，后端传递过去的数据
序列化成字符串 前端传过来的数据 反序列化  这里有个小技巧 在 结构体tag里 json:"id,string" 
关于go语言操作json的其他技巧
13. 帖子投票功能  如果在结构体里 我们定义的字段本就是string类型 这个时候千万不要在 json tag 里面加上string 会报错
14. 获取帖子列表 按照时间排序
15. 查询帖子 并且 获取到帖子所得的分数
16. 根据社区去查询帖子列表
17. 使用swagger 生成接口文档
#### Instructions

1.  xxxx
2.  xxxx
3.  xxxx

#### Contribution

1.  Fork the repository
2.  Create Feat_xxx branch
3.  Commit your code
4.  Create Pull Request


#### Gitee Feature

1.  You can use Readme\_XXX.md to support different languages, such as Readme\_en.md, Readme\_zh.md
2.  Gitee blog [blog.gitee.com](https://blog.gitee.com)
3.  Explore open source project [https://gitee.com/explore](https://gitee.com/explore)
4.  The most valuable open source project [GVP](https://gitee.com/gvp)
5.  The manual of Gitee [https://gitee.com/help](https://gitee.com/help)
6.  The most popular members  [https://gitee.com/gitee-stars/](https://gitee.com/gitee-stars/)
