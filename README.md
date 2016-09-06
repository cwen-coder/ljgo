# ljgo - 静态博客引擎

ljgo 是使用 [GO](https://golang.org) 语言实现的简单静态博客引擎, 编译速度快、安装简单。

#### 目前版本 0.1.0-beta

## 安装

#### 源码安装
go运行环境安装请自行google

```
go get -u github.com/cwen-coder/ljgo
make install
```

#### 直接下载编译好的可执行文件

下载地址 ： [ljgo](https://github.com/cwen-coder/ljgo/releases)

## 快速入门

#### 新建站点

```
ljgo new example.com
#执行完毕后，会在生成example.com文件夹
```

`example.com` 文件夹目录结构

```
- config.yml    // 站点配置文件
- source // 保存文章目录
- - | - - about.md // 关于页面内容
- - | - - article.md // 演示文章内容
- themes  // 保存所有主题
- - | - - default // 站点默认主题

```

`config.yml` 配置格式


```
site:
    title: 网站标题
    introduce: 网站描述
    limit: 每页可显示的文章数目
    theme: 网站主题目录   ＃eg: themes/default
    url: 站点域名
    comment: 评论插件变量(默认为Disqus账户名)
    github: github.com 地址 # 可选
    facebook: facebook 地址  # 可选
    twitter: twitter 地址  # 可选
serve:
    addr: ljgo serve 监听地址 # eg: "localhost:3000"

publish:
    cmd: |
        ljgo publish 命令将会执行的脚本
```

创建文章

在source目录中建立任意.md文件（可置于子文件夹），使用如下格式：


```
title: 文章标题
author: 文章作者
date: 2016-08-02
update: 2016-08-02
tags:
    - 设计
    - 写作

---

文章预览内容
	<!--more-->
文章其它内容
(文章的全部内容＝预览＋其他)

```


#### 生成静态页面

```
ljgo build
# 执行完毕在站点文件下生成public文件夹，包含所有静态文件
```

在站点文件夹中直接执行 `ljgo build` ， 或是在站点文件夹外执行但是得指定站点路径 eg
: `ljgo build example.com`
`ljgo serve` `ljgo publis` 都是同样的使用姿势


#### 本地预览

```
ligo serve
# 打来浏览器, 访问你在站点配置中填入的端口地址
# 默认是 http://localhost:3000
```

当然你也可以直接将 `ljgo serve` 运行在 `vps` 上


#### 部署

你可以使用 [github pages](https://pages.github.com/) 等服务，或者放到你的自己的vps下，因为是纯静态文件,不需要php/mysql/java等环境的支持


```
ljgo publish
# 执行站点配置中填写的发布脚本
```

eg : 使用 `github`服务， 初始化好 `public｀ 文件夹后，我们只需要在 `config.yml` 文件中的填写如下内容：

```
publish:
    cmd: |
        git add -A
        git commit -m "update"
        git push origin
```

这样我们在每次编辑完博客后直接运行 `ljgo publish` 就一切ok


## 关于主题

由于自己比较懒，目前的默认主题是从 [start bootstrapt](https://startbootstrap.com) 中的 [clean-blog](https://startbootstrap.com/template-overviews/clean-blog/) 修改而来
当导入其他主题，需要把主题文件夹复制到 `example.com/themes/` 文件夹下，并修改站点配置 `config.yml` 中主题路径


#### 十分欢迎大家贡献第三方主题 👏

## 正在使用

* [cwen's blog](http://www.cwen.pw)           - me

期待更多的用户

## 反馈贡献

非常欢迎任何人的任何贡献。如有问题可报告至 [https://github.com/cwen-coder/ljgo/issues](https://github.com/cwen-coder/ljgo/issues)。
或是直接发邮件To me [yincwengo@gmail.com](mailto:yincwengo@gmail.com)















