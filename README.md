# install

使用前需要安装`docker` `godep`

此处不介绍`docker`的安装配置。 `godep`安装：

```
go get github.com/tools/godep
```

当然，使用其他支持`vendor`的包管理工具代替`godep`是完全可以的

将改项目置于`$HOME/go/src/github.com/CodiesTeam/`目录下，或者根据实际情况修改`docker/docker-compose.yml`文件

# 启动

```
# your the gopath where you place this project
export FIRST_GOPATH=$HOME/go
# the absolutely path of this project
export CODIES_DIR=$FIRST_GOPATH/src/github.com/CodiesTeam/codies-server

make server     # 启动服务
make log        # 显示goserver日志
```
其他命令请查看`Makefile`

## mysql配置

* 将`init.sql`复制到`mysql`容器内
* 进入容器,进入`mysql`服务，手动创建数据库`codies`
* 进入`codies`数据库，使用`source`命令执行`init.sql`文件内容

```
# 复制init.sql文件到容器根目录
docker cp server/init.sql 69ec2f3a1884:/init.sql
# 进入容器
docker exec -it 69ec2f3a1884 bash
# 进入mysql服务
mysql -uroot -pcodies-pwd
# 创建codies数据库
mysql> create database codies;
# 进入codies数据库，解析init.sql文件
mysql> use codies;
mysql> source /init.sql
```

## 说明

因为`golang`的编译器在容器内,容器内只能看到`volume`范围的东西, 所以要用到`vedor`把包依赖都整理起来放到`volume`范围内,
这样容器内的程序才能正常编译运行


`skelton`目录用于存放路由框架 数据库底层封装
`server`目录用于存放业务相关代码
