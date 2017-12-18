ApiAdmin
====
什么东西？What?
----
API管理工具 golang开发，基于beego，页面基于layUi    
练手使用，尚未用于生产环境

有什么价值？
----
1、RBAC权限完善，多角色管理系统    
2、后台界面完整，多标签页面    
3、API相关页面有比较复杂的使用案例    
所以，可以作为一个基础框架使用，快速开发。初学者还可以作为熟悉beego使用。    

安装方法    
----
1、创建mysql数据库，并将api_admin.sql导入    
2、修改config 配置数据库    
3、运行 go build    
4、运行 ./run.sh start|stop


前台访问：http://your_host:8081    
用户名：admin 密码：123456    



