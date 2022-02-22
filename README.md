# golang学习笔记

## Go语言概述

​		Go（又称Golang）是Google开发的一种静态强类型、编译型、并发型，并具有垃圾回收功能的编程语言。罗伯特·格瑞史莫（Robert Griesemer），罗勃·派克（Rob Pike）及肯·汤普逊（Ken Thompson）于2007年9月开始设计Go，稍后Ian Lance Taylor、Russ Cox加入项目。Go是基于Inferno操作系统所开发的。Go于2009年11月正式宣布推出，成为开放源代码项目，并在Linux及Mac OS X平台上进行了实现，后来追加了Windows系统下的实现。在2016年，Go被软件评价公司TIOBE 选为“TIOBE 2016 年最佳语言”。 目前，Go每半年发布一个二级版本（即从a.x升级到a.y）。

​		参考[go（计算机编程语言）](https://baike.baidu.com/item/go/953521?fromtitle=Go%E8%AF%AD%E8%A8%80) 

## 安装Go语言环境

参考文档：[安装Go语言及搭建Go语言开发环境](https://www.cnblogs.com/aaronthon/p/10587487.html)

### 下载
Go官网下载地址：[Downloads - The Go Programming Language](https://golang.google.cn/dl/)

选择适合自己操作系统的版本，这里以win10操作系统为例
![image-20220222233418906](https://s2.loli.net/2022/02/22/NXwiCbEJGTFqan1.png)

### 安装
双击可执行文件

![image-20220222233759398](https://s2.loli.net/2022/02/22/fYSLg8cyPvlMxiT.png)
![image-20220222233928479](https://s2.loli.net/2022/02/22/f4T3zy62m7tUuiq.png)
![image-20220222233958337](https://s2.loli.net/2022/02/22/cJRhl1X2WIA3Lar.png)

选择安装路径-可自定义C:\Program Files\Go\（和JDK类似)，需要记住该安装路径

![image-20220222234051118](https://s2.loli.net/2022/02/22/ARVTraJeFsEDUP5.png)
![image-20220222234227766](https://s2.loli.net/2022/02/22/f1CMKRHpgFXIxiN.png)

### 配置Go环境变量

#### GOROOT（GO安装路径）

GOROOT配置代表GO的安装路径，本例为C:\Program Files\Go\

> 编辑系统环境变量

![image-20220223000005872](https://s2.loli.net/2022/02/23/IQDW5lCXFkdp2KG.png)
![image-20220223000304242](https://s2.loli.net/2022/02/23/bOEYMHglywWrNPf.png)

> 添加GOROOT系统变量

![image-20220223001833224](https://s2.loli.net/2022/02/23/svGMIdURT6hFlVP.png)

> 添加bin到Path中

![image-20220223002441876](https://s2.loli.net/2022/02/23/TQIfDVACxeScw9l.png)

#### GOPATH（项目工程路径）

> 创建目录

GOPATH表示后续开发存放代码/二进制文件/缓存文件的工作空间（建议新建一个空目录，如：D:\Program Files (x86)\GoProjects）

并在GOPATH路径下创建三个目录，src、bin、pkg

> `$GOPATH/src`：第三方包源代码文件
> `$GOPATH/bin`：产生的二进制可执行文件
> `$GOPATH/pkg`：生成的中间缓存文件

![image-20220223003913830](https://s2.loli.net/2022/02/23/rh5ot6xAFsHwgYv.png)

> 配置GOPATH

![image-20220223004124934](https://s2.loli.net/2022/02/23/yuTbqNdsj64hftQ.png)

#### 验证环境变量生效

打开cmd命令行，输入go version指令测试

![image-20220223010030906](https://s2.loli.net/2022/02/23/6frd2EnTl9GHyJN.png)

