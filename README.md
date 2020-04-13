# IANA WHOIS Crawler

IANA WHOIS服务器爬虫

## 文件结构介绍

|文件|说明|
|-------------|:-------------:|
|latest.log|用于实时记录当前成功抓取的TLD, 便于程序重启后可以从断点开始工作|
|main.go|HTTP服务器, 与爬虫脚本进行交互|
|main.js|爬虫脚本, 实时将抓取结果发送到HTTP服务器|
|serverlist.txt|WHOIS服务器抓取结果。PS：已经抓好了，拿来即用|

![抓取结果](/assets/serverlist.png "抓取结果")

## 运行爬虫

1. 在运行爬虫之前

    * 假定你具备基础的编程以及调试知识（Javascript、Golang）

    * 假定你具备基础的Golang知识

    * 假定你本地已安装 [Golang二进制包](https://golang.google.cn/dl/)

2. 编译得到HTTP服务器

    打开存储库根目录，运行以下脚本

    ```go
    go build main.go
    ```

3. 运行HTTP服务器

    在 PowerShell 或 Cmd 中运行编译得到的 main.exe

4. 定位采集源

    在浏览器（推荐用Chrome浏览器）中打开 [IANA 根域数据库](https://www.iana.org/domains/root/db)

5. 运行爬虫脚本

    打开浏览器的开发人员工具（F12），定位到Console选项卡

    将 main.js 中所有内容粘贴到光标处，敲击回车键

    ![爬虫脚本(main.js)使用方法演示](/assets/mainjs.gif "爬虫脚本(main.js)使用方法演示")

## 结尾

我只实现了基本的爬虫功能，进度展示没做~

你可以查看 latest.log 文件，当它的内容显示 **.zuerich** 就算是运行结束了
