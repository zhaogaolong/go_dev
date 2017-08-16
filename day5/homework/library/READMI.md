## 设计思路

### 图书设计
- 同一种类型的本书设计一个结构体
- 每本书有采用切片管理
- 书的切片里内容包括每一本书的数量map
- 书的map数据保存为user：userinfo指针, bookID:bookID指针

### 图书
- 书名
- 数量
- 作者
- 出版日期
- id

### 用户设计
- 每个用户都是一个独立的结构体
- 每个用户借的书都采用切片的方式存储


### 用户
- 姓名
- 年级
- 身份证
- 性别
- 借书的列表
- id


## 使用说明

> 文件介绍

- book.db 图书数据库文件，man.exe在那运行就把它放在同一级目录下
- student.db 学生用户文件。默认没有，需要注册后就会自动生成
- library.log 所有图书系统的日志记录，包括用户登陆、借书、还书、退出等

### 生成二进制文件

    $ go build go_dev\day5\homework\library\main
    
    
#### running     
    $ ./main.exe
>注册

    -----------Welcome Library--------------
    T(top10), B(books), R(register), L(login), E(exit)
    r
    -----Welcome Register User -----
    Please input your name
    zhaogaolong  // 输入用户名
    Input your password
    123  //输入密码
    Input your ID
    17982301783487983942 // 输入用户id
    Please wait a moment, creating user
    User register success, please login


> 登录

    -----------Welcome Library--------------
    T(top10), B(books), R(register), L(login), E(exit)
    l // l 登录
    -----Welcome Login ----------
    Please input your name.
    zhaogaolong //用户名
    Please input your password.
    123 //密码
    ---User zhaogaolong logined

> 借书流程


    --------Welcome zhaogaolong ---------
    T(top10), B(borrow 借书), R(return,还书), L(books list,书列表) C(Book Car), O(logout),
    t //前十最热门书籍
    ID       Name    BorrowCount
    1       book6   3
    2       book4   2
    3       book2   1
    4       book0   0
    5       book3   0
    6       book5   0
    7       book1   0
    8       book7   0
    9       book8   0
    10      book9   0
    --------Welcome zhaogaolong ---------
    T(top10), B(borrow 借书), R(return,还书), L(books list,书列表) C(Book Car), O(logout),
    b //借书
    Please input book name
    book2 //书名
    Please input number:
    4 // 借几本
    Borrow sucess
    --------Welcome zhaogaolong ---------
    T(top10), B(borrow 借书), R(return,还书), L(books list,书列表) C(Book Car), O(logout),
    c // 查看自己的借书车
    ----zhaogaolong Borrow Car---
    Book     Count
    book2    4
    --------Welcome zhaogaolong ---------
    T(top10), B(borrow 借书), R(return,还书), L(books list,书列表) C(Book Car), O(logout),
    t // 再查看top10 发现book2已经升到第一了
    ID       Name    BorrowCount
    1       book2   5
    2       book6   3
    3       book4   2
    4       book0   0
    5       book3   0
    6       book5   0
    7       book1   0
    8       book7   0
    9       book8   0
    10      book9   0
    --------Welcome zhaogaolong ---------

> 退出

    T(top10), B(borrow 借书), R(return,还书), L(books list,书列表) C(Book Car), O(logout),
    o // 用户退出
    -----------Welcome Library--------------
    T(top10), B(books), R(register), L(login), E(exit)
    e // 退出图书馆系统
    Exit Library 


> 还书流程

    PS H:\360_update\oldboy_go> .\main.exe
    -----------Welcome Library--------------
    T(top10), B(books), R(register), L(login), E(exit)
    l
    -----Welcome Login ----------
    Please input your name.
    zhaogaolong
    Please input your password.
    123
    ---User zhaogaolong logined
    --------Welcome zhaogaolong ---------
    T(top10), B(borrow 借书), R(return,还书), L(books list,书列表) C(Book Car), O(logout),
    c
    ----zhaogaolong Borrow Car---
    Book     Count
    book2    4
    --------Welcome zhaogaolong ---------
    T(top10), B(borrow 借书), R(return,还书), L(books list,书列表) C(Book Car), O(logout),
    r
    Please input book name
    book2
    Please input return books count:
    4

    Borrow sucess
    --------Welcome zhaogaolong ---------
    T(top10), B(borrow 借书), R(return,还书), L(books list,书列表) C(Book Car), O(logout),
    c
    ----zhaogaolong Borrow Car---
    Book     Count
    --------Welcome zhaogaolong ---------
    T(top10), B(borrow 借书), R(return,还书), L(books list,书列表) C(Book Car), O(logout),
    o
    -----------Welcome Library--------------
    T(top10), B(books), R(register), L(login), E(exit)
    q
    Exit Library
        


### log

使用notepad++ 或者linux编辑器查看，因为换行符是\n换行，如果使用windows的记事本的话就会不懂到换行