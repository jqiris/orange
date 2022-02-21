# orange

## 第一步部署环境

1-部署etcd服务器注册中心
cd deploy/etcd && docker-compose up -d
2-部署nats消息中心
cd deploy/nats && docker-compose up -d
3-部署redis
docker run -d -p 6379:6379 --name redis redis
4-部署mysql
docker或者自己安装-暂时不需要

## 启动服务端服务器
说明-configs配置文件里frames.servers里面配置服务器,is_lauch表明是需要在当前节点启动的服务
go run main.go启动所有当前节点的服务

## 启动测试环境测试
说明-第一次进入请执行npm install进行nodejs服务安装
cd client客户端文件夹
node app.js启动客户端服务
打开http://127.0.0.1:3001/index.html点击测试按钮进行测试