go mod使用：
    初始化生成.mod文件：  ```go mod init mod文件名```
    安装包依赖：```go mod tidy```
    
    跨文件夹导入自定义mod:    1. 自定义go文件处 go mod init mod文件名 来初始化
                            2. 修改main函数或整个项目的.mod文件：
                                require 自定义mod名 版本号(v0.0.0)
                                replace 自定义mod => 路径