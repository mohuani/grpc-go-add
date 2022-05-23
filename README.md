### 编写 proto 文件

### 生成pb文件
```shell
protoc --go_out=plugins=grpc:. add.proto
```

### 编写 grpc service 服务
- 实现具体的逻辑

### 编写 grpc server
- 监听 tcp xx端口
  - 注意端口的写法不要写错，前面带有冒号
- 实例化 grpc service
- 实例化 grpc Server 
  ```shell
  server := grpc.NewServer()
  ```
  
- 将 grpc service 注册到 grpc Server 
- 启动 grpc Server
    ```shell
    _ = server.Serve(listen)
    ```

