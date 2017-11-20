### 一、概述
设计一个 web 小应用，展示静态文件服务、js 请求支持、模板输出、表单处理、Filter 中间件设计等方面的能力。（不需要数据库支持）

### 二、任务要求
编程 web 应用程序 cloudgo-io。 请在项目 README.MD 给出完成任务的证据！

**基本要求**

>支持静态文件服务
>
>支持简单 js 访问
>
>提交表单，并输出一个表格
>
>对 `/unknown` 给出开发中的提示，返回码 `5xx`

---
**由于这次是在上次的基础上继续实验（我没有使用上次的martini框架，改成了老师提供的框架），所以其他过程不再赘述，只展示测试过程。**

## **图片水印是我自己的CSDN博客，我是在博客编写的README.md文件。**

---

### 三、测试

#### 1.静态文件服务,建立/api/test，并在服务器上创建目录，以存放静态内容，输出ID，age，content。下面是/api/test实现后的输出情况。
![这里写图片描述](http://img.blog.csdn.net/20171118211043005?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvcXFfMzM0NTQxMTI=/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast)

![这里写图片描述](http://img.blog.csdn.net/20171118211054496?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvcXFfMzM0NTQxMTI=/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast)

![](http://img.blog.csdn.net/20171118211103952?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvcXFfMzM0NTQxMTI=/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast)

####２.下面依次展示了服务器添加文件结构以及文件夹下添加css、js文件后的情况。

![这里写图片描述](http://img.blog.csdn.net/20171118210347558?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvcXFfMzM0NTQxMTI=/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast)

![这里写图片描述](http://img.blog.csdn.net/20171118210340438?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvcXFfMzM0NTQxMTI=/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast)

![这里写图片描述](http://img.blog.csdn.net/20171118210328201?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvcXFfMzM0NTQxMTI=/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast)

#### 3.添加完成后，再增加index.html文件，访问端口并查看监听情况。
![这里写图片描述](http://img.blog.csdn.net/20171118210856644?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvcXFfMzM0NTQxMTI=/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast)

![这里写图片描述](http://img.blog.csdn.net/20171118211229460?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvcXFfMzM0NTQxMTI=/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast)

#### 4.使用模板template输出。
![这里写图片描述](http://img.blog.csdn.net/20171118211347767?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvcXFfMzM0NTQxMTI=/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast)

#### 5.输出一个登录账户密码的表单表格。
![这里写图片描述](http://img.blog.csdn.net/20171118211436069?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvcXFfMzM0NTQxMTI=/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast)

#### 6./api/unknown正在开发中的501未实现信息提示。
![这里写图片描述](http://img.blog.csdn.net/20171118211526245?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvcXFfMzM0NTQxMTI=/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast)