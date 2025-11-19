# RedRockBlog

> #### ⚡红岩网校后端考核⚡

![logo](img/logo.png "logo")

## 1. 接口列表

### 1.1 👩UserManager(用户管理器) - 4个接口

#### 1.1.1 Login(OAuth2单点登录)

![logo](https://badgen.net/badge/method/GET/green?icon=bitcoin-lightning "logo")
```
GET /v1/blog/user/login
```

#### 1.1.2 User_info(获取用户信息) 

![logo](https://badgen.net/badge/method/GET/green?icon=bitcoin-lightning "logo")

``` 
GET /v1/blog/user/info
```
- 绑定数据：

| 位置     | 项      | 值                      | 是否必须 | 数据类型   | 备注    | 
|--------|--------|------------------------|------|--------|-------| 
| Header | Cookie | sso_jwt:<your_sso_jwt> | 是    | String | JWT令牌 |

- 返回数据：

| 位置   | 项         | 数据类型   | 备注     | 
|------|-----------|--------|--------|
| Body | CreatedAt | String | 用户创建时间 |
| Body | UpdatedAt | String | 用户更新时间 |
| Body | DeletedAt | String | 用户删除时间 |
| Body | Name      | String | 用户名    |
| Body | ID        | String | 用户ID   |



- 返回值JSON：
```json
{
    "CreatedAt": "2025-11-04T20:56:50.909+08:00",
    "UpdatedAt": "2025-11-04T20:56:50.909+08:00",
    "DeletedAt": null,
    "Name": "TuF3i",
    "ID": "722519f7c64dd624a173eb0549f1cfec"
}
```

#### 1.1.3 Delete_user(删除用户)

![logo](https://badgen.net/badge/method/DELETE/red?icon=bitcoin-lightning "logo")

``` 
DELETE /v1/blog/user/delete
```
- 绑定数据

| 位置     | 项      | 值                      | 是否必须 | 数据类型   | 备注    |
|--------|--------|------------------------|------|--------|-------|
| Header | Cookie | sso_jwt:<your_sso_jwt> | 是    | String | JWT令牌 |
- 返回数据

| 位置   | 项       | 数据类型   | 备注   | 
|------|---------|--------|------|
| Body | Message | String | 返回信息 |

- 返回值JSON：
```json
{
    "Message": "string"
}
```

#### 1.1.4 Logout(注销)

![logo](https://badgen.net/badge/method/GET/green?icon=bitcoin-lightning "logo")

```
GET /v1/blog/user/logout
```
- 绑定数据

| 位置     | 项      | 值                      | 是否必须 | 数据类型   | 备注    |
|--------|--------|------------------------|------|--------|-------|
| Header | Cookie | sso_jwt:<your_sso_jwt> | 是    | String | JWT令牌 |

### 1.2 🌐CommentManager(评论管理器) - 4个接口

#### 1.2.1 Get_comment(获取评论)

![logo](https://badgen.net/badge/method/GET/green?icon=bitcoin-lightning "logo")

``` 
GET /v1/blog/comment/get/{article_id}
```

- 绑定数据

| 位置     | 项      | 值                      | 是否必须 | 数据类型 | 备注   |
|--------|--------|------------------------|------|------|------|
| Path | article_id | article_id | 是    | int  | 文章ID |
- 返回数据

| 位置   | 项             | 数据类型   | 备注           | 
|------|---------------|--------|--------------|
| Body | comments      | list   | 评论           |
| Body | ID            | int    | 评论ID         |
| Body | CreatedAt     | String | 创建时间         |
| Body | UpdatedAt     | String | 更新时间         |
| Body | DeletedAt     | String | 删除时间         |
| Body | Level         | int    | 评论层级         |
| Body | ArticleID     | int    | 评论绑定的文章ID    |
| Body | Content       | String | 评论内容(Base64) |
| Body | IP            | String | 评论者的IP       |
| Body | Location      | String | 评论者的地理位置     |
| Body | Author        | String | 评论者的用户名      |
| Body | AuthorID      | String | 评论者的ID       |
| Body | ParentID      | String | 父评论的ID       |
| Body | childComments | list   | 子评论列表        |


- 返回值JSON：
```json
{
  "comments": [
    {
      "ID": 1,
      "CreatedAt": "2025-11-04T20:58:38.915+08:00",
      "UpdatedAt": "2025-11-04T20:58:38.915+08:00",
      "DeletedAt": null,
      "Level": 0,
      "ArticleID": 33,
      "Content": "4pyFSGVsbG8gV29ybGTinIU=",
      "IP": "127.0.0.1",
      "Location": "",
      "Author": "TuF3i",
      "AuthorID": "722519f7c64dd624a173eb0549f1cfec",
      "ParentID": 0,
      "childComments": [
        {
          "ID": 2,
          "CreatedAt": "2025-11-04T20:59:46.972+08:00",
          "UpdatedAt": "2025-11-04T20:59:46.972+08:00",
          "DeletedAt": null,
          "Level": 1,
          "ArticleID": 33,
          "Content": "4pyFSGVsbG8gV29ybGTinIU=",
          "IP": "127.0.0.1",
          "Location": "",
          "Author": "TuF3i",
          "AuthorID": "",
          "ParentID": 1,
          "childComments": null
        }
      ]
    }
  ]
}
```

#### 1.2.2 Delete_comment(删除评论)

![logo](https://badgen.net/badge/method/DELETE/red?icon=bitcoin-lightning "logo")

``` 
DELETE /v1/blog/comment/delete/{comment_id}
```

- 绑定数据

| 位置     | 项          | 值                      | 是否必须 | 数据类型   | 备注    |
|--------|------------|------------------------|------|--------|-------|
| Path   | comment_id | comment_id             | 是    | String | 评论ID  |
| Header | Cookie     | sso_jwt:<your_sso_jwt> | 是    | String | JWT令牌 |

- 返回数据

| 位置   | 项       | 数据类型   | 备注   | 
|------|---------|--------|------|
| Body | Message | String | 返回信息 |

```json
{
    "Message": "string"
}
```

#### 1.2.3 Add_father_comment(添加父评论)

![logo](https://badgen.net/badge/method/POST/yellow?icon=bitcoin-lightning "logo")

``` 
POST /v1/blog/comment/add
```
- 绑定数据

| 位置     | 项         | 值                      | 是否必须 | 数据类型   | 备注    |
|--------|-----------|------------------------|------|--------|-------|
| Body   | articleID | articleID              | 是    | int    | 评论ID  |
| Body   | content   | content                | 是    | String | 评论内容  |
| Header | Cookie    | sso_jwt:<your_sso_jwt> | 是    | String | JWT令牌 |

```json
{
	"articleID": 33,
	"content": "4pyFSGVsbG8gV29ybGTinIU="
}
```

- 返回数据

| 位置   | 项       | 数据类型   | 备注   | 
|------|---------|--------|------|
| Body | Message | String | 返回信息 |

```json
{
    "Message": "string"
}
```

#### 1.2.4 Add_sub_comment(添加子评论)

![logo](https://badgen.net/badge/method/POST/yellow?icon=bitcoin-lightning "logo")

``` 
POST /v1/blog/comment/add-sub
```
- 绑定数据

| 位置     | 项         | 值                     | 是否必须 | 数据类型   | 备注    |
|--------|-----------|-----------------------|------|--------|-------|
| Body   | ParentID | ParentID              | 是    | int    | 父评论ID |
| Body   | articleID | articleID             | 是    | int    | 评论ID  |
| Body   | content   | content               | 是    | String | 评论内容  |
| Header | Cookie    | sso_jwt:<your_sso_jwt> | 是    | String | JWT令牌 |

```json
{
  "articleID": 33,
  "ParentID": 1,
  "content": "4pyFSGVsbG8gV29ybGTinIU="
}
```

- 返回数据

| 位置   | 项       | 数据类型   | 备注   | 
|------|---------|--------|------|
| Body | Message | String | 返回信息 |

```json
{
  "Message": "string"
}
```

### 1.3 📄ArticleManager(文章管理器) - 9个接口

#### 1.3.1 Get_my_work_list(获取自己的作品)

![logo](https://badgen.net/badge/method/GET/green?icon=bitcoin-lightning "logo")

```
GET /v1/blog/article/my-work-list
```
- 绑定数据

| 位置     | 项      | 值                      | 是否必须 | 数据类型   | 备注    | 
|--------|--------|------------------------|------|--------|-------| 
| Header | Cookie | sso_jwt:<your_sso_jwt> | 是    | String | JWT令牌 |

- 返回数据

| 位置   | 项             | 数据类型   | 备注                | 
|------|---------------|--------|-------------------|
| Body | topWork      | list   | 置顶作品              |
| Body | drafts      | list   | 草稿                |
| Body | languageID      | int    | 语言ID(0: 中文，1: 英文) |
| Body | articleID      | int    | 文章ID              |
| Body | updatedAt      | String | 更新时间              |
| Body | title      | String | 标题                |
| Body | extTitle      | String | 别名                |
| Body | introduction      | String | 作品简介              |
| Body | ifDraft      | bool   | 是否为草稿             |
| Body | ifPrivate      | bool   | 是否私有              |
| Body | ifTop      | bool   | 是否置顶              |



- 返回值JSON
```json
{
    "topWork": [
        {
            "languageID": 0,
            "articleID": 1,
            "updatedAt": "2025-11-04T21:28:50.157+08:00",
            "title": "Test_Article_pub_top",
            "extTitle": "ext-name-pub",
            "introduction": "This is a test 1",
            "ifDraft": false,
            "ifPrivate": false,
            "ifTop": true
        }
    ],
    "drafts": [
        {
            "languageID": 0,
            "articleID": 5,
            "updatedAt": "2025-11-04T21:34:41.496+08:00",
            "title": "Test_Draft",
            "extTitle": "ext-name-draft-4",
            "introduction": "This is a test 4",
            "ifDraft": true,
            "ifPrivate": false,
            "ifTop": false
        }
    ],
    "normalWork": [
        {
            "languageID": 0,
            "articleID": 7,
            "updatedAt": "2025-11-04T23:53:39.258+08:00",
            "title": "Test_Private",
            "extTitle": "ext-name-draft-Private",
            "introduction": "Private Test",
            "ifDraft": false,
            "ifPrivate": true,
            "ifTop": false
        },
        {
            "languageID": 1,
            "articleID": 6,
            "updatedAt": "2025-11-04T23:42:12.177+08:00",
            "title": "Test_Article_pub",
            "extTitle": "ext-name-5",
            "introduction": "This is a test 2",
            "ifDraft": false,
            "ifPrivate": false,
            "ifTop": false
        },
        {
            "languageID": 0,
            "articleID": 3,
            "updatedAt": "2025-11-04T21:32:22.926+08:00",
            "title": "Test_Article_pub",
            "extTitle": "ext-name-pub-2",
            "introduction": "This is a test 2",
            "ifDraft": false,
            "ifPrivate": false,
            "ifTop": false
        }
    ]
}
```

#### 1.3.2 Get_articles_list(获取所有文章，包括私有和公有，不包括草稿)

![logo](https://badgen.net/badge/method/GET/green?icon=bitcoin-lightning "logo")

``` 
GET /v1/blog/article/articles
```

#### 1.3.3 Search_by_ext_name(绝对索引)

![logo](https://badgen.net/badge/method/GET/green?icon=bitcoin-lightning "logo")

``` 
GET /v1/blog/article/search-ext-name/{extName}
```

- 绑定数据

| 位置   | 项      | 值                      | 是否必须 | 数据类型   | 备注   | 
|------|--------|------------------------|------|--------|------| 
| Path | extName | extName | 是    | String | 文章别名 |

- 返回数据

| 位置   | 项             | 数据类型   | 备注                | 
|------|---------------|--------|-------------------|
| Body | SoloWork      | list   | 作品                |
| Body | languageID      | int    | 语言ID(0: 中文，1: 英文) |
| Body | articleID      | int    | 作品ID              |
| Body | updatedAt      | String | 更新时间              |
| Body | title      | String | 文章标题              |
| Body | introduction      | String | 文章简介              |
| Body | extTitle      | String | 文章别名              |
| Body | ifPrivate      | bool   | 是否为私有             |
| Body | ifTop      | bool   | 是否置顶              |

- 返回值JSON
```json
{
    "SoloWork": [
        {
            "languageID": 0,
            "articleID": 1,
            "updatedAt": "0001-01-01T00:00:00Z",
            "title": "Test_Article_pub_top",
            "extTitle": "ext-name-pub",
            "introduction": "This is a test 1",
            "ifPrivate": false,
            "ifTop": true
        }
    ]
}
```

#### 1.3.4 Search_by_name(按名称模糊搜索)

![logo](https://badgen.net/badge/method/GET/green?icon=bitcoin-lightning "logo")

``` 
GET /v1/blog/article/search-mohu-name/{name}
```

- 绑定数据

| 位置   | 项      | 值                      | 是否必须 | 数据类型   | 备注 | 
|------|--------|------------------------|------|--------|----| 
| Path | name | name | 是    | String | 名称 |

- 返回数据

| 位置   | 项             | 数据类型   | 备注                | 
|------|---------------|--------|-------------------|
| Body | topWork      | list   | 置顶作品              |
| Body | normalWork      | list   | 普通作品              |
| Body | languageID      | int    | 语言ID(0: 中文，1: 英文) |
| Body | articleID      | int    | 作品ID              |
| Body | updatedAt      | String | 更新时间              |
| Body | title      | String | 文章标题              |
| Body | introduction      | String | 文章简介              |
| Body | extTitle      | String | 文章别名              |
| Body | ifPrivate      | bool   | 是否为私有             |
| Body | ifTop      | bool   | 是否置顶              |

- 返回值JSON
```json
{
    "topWork": [
        {
            "languageID": 0,
            "articleID": 1,
            "updatedAt": "2025-11-04T21:28:50.157+08:00",
            "title": "Test_Article_pub_top",
            "extTitle": "ext-name-pub",
            "introduction": "This is a test 1",
            "ifPrivate": false,
            "ifTop": true
        }
    ],
    "normalWork": [
        {
            "languageID": 1,
            "articleID": 6,
            "updatedAt": "2025-11-04T23:42:12.177+08:00",
            "title": "Test_Article_pub",
            "extTitle": "ext-name-5",
            "introduction": "This is a test 2",
            "ifPrivate": false,
            "ifTop": false
        }
    ]
}
```

#### 1.3.5 Get_work_content(获取作品的信息，包括Article,Draft)

![logo](https://badgen.net/badge/method/GET/green?icon=bitcoin-lightning "logo")

``` 
GET /v1/blog/article/get-work-content/{article_id}
```

- 绑定数据

| 位置   | 项          | 值                      | 是否必须 | 数据类型   | 备注 | 
|------|------------|------------------------|------|--------|----| 
| Path | article_id | article_id | 是    | String | 文章ID |
| Header | Cookie     | Cookie:<sso_jwt:your_sso_jwt> | 是    | String | JWT令牌 |

- 返回数据

| 位置   | 项             | 数据类型 | 备注     |
  |------|---------------|-----|--------|
  | Body | CreatedAt     | string | 创建时间   |
  | Body | UpdatedAt     | string | 更新时间   |
  | Body | DeletedAt     | bool | 删除时间   |
  | Body | ArticleID     | number | 文章ID   |
  | Body | AuthorID      | string | 作者ID   |
  | Body | LanguageID    | number | 语言ID(0: 中文，1: 英文) |
  | Body | Title         | string | 文章标题   |
  | Body | ExtTitle      | string | 扩展标题   |
  | Body | Introduction  | string | 文章简介   |
  | Body | Content       | string | 文章内容   |
  | Body | IfDraft       | bool| 是否为草稿  |
  | Body | IfPrivate     | bool| 是否为私有  |
  | Body | IfTop         | bool| 是否置顶   |

- 返回值JSON
```json
{
    "CreatedAt": "2025-11-04T23:26:49.486+08:00",
    "UpdatedAt": "2025-11-04T23:42:12.177+08:00",
    "DeletedAt": null,
    "ArticleID": 6,
    "AuthorID": "722519f7c64dd624a173eb0549f1cfec",
    "LanguageID": 1,
    "Title": "Test_Article_pub",
    "ExtTitle": "ext-name-5",
    "Introduction": "This is a test 2",
    "Content": "1111",
    "IfDraft": false,
    "IfPrivate": false,
    "IfTop": false
}
```

#### 1.3.6 Get_article_content(获取文章内容)

![logo](https://badgen.net/badge/method/GET/green?icon=bitcoin-lightning "logo")

``` 
GET /v1/blog/article/get-article-content/{article_id}
```

- 绑定数据

| 位置   | 项          | 值                      | 是否必须 | 数据类型   | 备注 |
  |------|------------|------------------------|------|--------|----|
  | Path | article_id | article_id | 是    | String | 文章ID |

- 返回数据

| 位置   | 项             | 数据类型 | 备注     |
  |------|---------------|-----|--------|
  | Body | CreatedAt     | string | 创建时间   |
  | Body | UpdatedAt     | string | 更新时间   |
  | Body | DeletedAt     | bool | 删除时间   |
  | Body | ArticleID     | number | 文章ID   |
  | Body | AuthorID      | string | 作者ID   |
  | Body | LanguageID    | number | 语言ID(0: 中文，1: 英文) |
  | Body | Title         | string | 文章标题   |
  | Body | ExtTitle      | string | 扩展标题   |
  | Body | Introduction  | string | 文章简介   |
  | Body | Content       | string | 文章内容   |
  | Body | IfDraft       | bool| 是否为草稿  |
  | Body | IfPrivate     | bool| 是否为私有  |
  | Body | IfTop         | bool| 是否置顶   |

- 返回值JSON
```json
{
    "CreatedAt": "2025-11-04T23:53:39.258+08:00",
    "UpdatedAt": "2025-11-04T23:53:39.258+08:00",
    "DeletedAt": null,
    "ArticleID": 7,
    "AuthorID": "722519f7c64dd624a173eb0549f1cfec",
    "LanguageID": 0,
    "Title": "Test_Private",
    "ExtTitle": "ext-name-draft-Private",
    "Introduction": "Private Test",
    "Content": "4pyFSGVsbG8gV29ybGTinIU=",
    "IfDraft": false,
    "IfPrivate": true,
    "IfTop": false
}
```

#### 1.3.7 Add_article(添加作品)

![logo](https://badgen.net/badge/method/POST/yellow?icon=bitcoin-lightning "logo")

``` 
POST /v1/blog/article/add
```

- 绑定数据

| 位置     | 项            | 值                      | 是否必须 | 数据类型   | 备注                |
|--------|--------------|------------------------|------|--------|-------------------|
| Header | Cookie       | <sso_jwt:your_sso_jwt> | 是    | String | JWT令牌             |
| Body   | languageID   | languageID             | 是    | int    | 语言ID(0: 中文，1: 英文) |
| Body   | title        | title                  | 是    | String | 文章标题              |
| Body   | extTitle     | extTitle               | 是    | String | 文章别名              |
| Body   | introduction | introduction           | 是    | String | 文章简介              |
| Body   | content      | content                | 是    | String | 正文(Base64)        |
| Body   | ifDraft      | ifDraft                | 是    | bool   | 是否为草稿             |
| Body   | ifPrivate    | ifPrivate              | 是    | bool   | 是否私有              |
| Body   | ifTop        | ifTop                  | 是    | bool   | 是否置顶              |              

- 绑定数据JSON

```json
{
	"languageID": 0,
	"title": "Test_Private",
	"extTitle": "ext-name-draft-Private",
	"introduction": "Private Test",
	"content": "4pyFSGVsbG8gV29ybGTinIU=",
	"ifDraft": false,
	"ifPrivate": true,
	"ifTop": false
}
```

- 返回数据

| 位置   | 项       | 数据类型   | 备注   | 
|------|---------|--------|------|
| Body | Message | String | 返回信息 |

- 返回值JSON：
```json
{
    "Message": "string"
}
```

#### 1.3.8 Update_work(更新作品)

![logo](https://badgen.net/badge/method/PUT/blue?icon=bitcoin-lightning "logo")

``` 
PUT /v1/blog/article/update
```

- 绑定数据

| 位置     | 项            | 值                      | 是否必须 | 数据类型   | 备注                |
|--------|--------------|------------------------|------|--------|-------------------|
| Header | Cookie       | <sso_jwt:your_sso_jwt> | 是    | String | JWT令牌             |
| Body   | articleID    | articleID              | 是    | int    | 文章ID              |
| Body   | languageID   | languageID             | 是    | int    | 语言ID(0: 中文，1: 英文) |
| Body   | title        | title                  | 是    | String | 文章标题              |
| Body   | extTitle     | extTitle               | 是    | String | 文章别名              |
| Body   | introduction | introduction           | 是    | String | 文章简介              |
| Body   | content      | content                | 是    | String | 正文(Base64)        |
| Body   | ifDraft      | ifDraft                | 是    | bool   | 是否为草稿             |
| Body   | ifPrivate    | ifPrivate              | 是    | bool   | 是否私有              |
| Body   | ifTop        | ifTop                  | 是    | bool   | 是否置顶              |              

- 绑定数据JSON

```json
{
  "articleID": 6,
  "languageID": 1,
  "title": "Test_Article_pub",
  "extTitle": "ext-name-5",
  "introduction": "This is a test 2",
  "content": "1111",
  "ifDraft": false,
  "ifPrivate": false,
  "ifTop": false
}
```

- 返回数据

| 位置   | 项       | 数据类型   | 备注   | 
|------|---------|--------|------|
| Body | Message | String | 返回信息 |

- 返回值JSON：
```json
{
    "Message": "string"
}
```

#### 1.3.9 Delete_my_work(删除作品)

![logo](https://badgen.net/badge/method/DELETE/red?icon=bitcoin-lightning "logo")

``` 
DELETE /v1/blog/article/delete/{article_id}
```

- 绑定数据

| 位置     | 项          | 值                      | 是否必须 | 数据类型   | 备注                |
|--------|------------|------------------------|------|--------|-------------------|
| Header | Cookie     | <sso_jwt:your_sso_jwt> | 是    | String | JWT令牌             |
| Path   | article_id | article_id             | 是    | int    | 文章ID              |

- 返回数据

| 位置   | 项       | 数据类型   | 备注   | 
|------|---------|--------|------|
| Body | Message | String | 返回信息 |

- 返回值JSON：
```json
{
    "Message": "string"
}
```

## 2.状态码及错误消息

### 2.1 状态码

| Code | 状态          |
|------|-------------|
| 200  | 请求成功        |
| 301  | Redirect重定向 |
| 404  | 资源未找到       |
| 500  | 服务器错误       |

### 2.2 消息示例

#### 2.2.1 `POST`-`200`-`Success`

| Method | Code | Status |
|--------|------|--------|
| POST   | 200  | 成功     |

- 响应体JSON:
```json
{
    "Message": "string"
}
```

#### 2.2.1 `POST`-`500/403`-`failed`

| Method | Code    | Status |
|--------|---------|--------|
| POST   | 500/404 | 失败     |

- 响应体JSON:
```json
{
    "Error": "string"
}
```

#### 2.2.1 `GET`-`301`-`Redirect`

| Method | Code | Status |
|--------|------|--------|
| GET    | 301  | 重定向    |

> [!NOTE]
> 
> 重定向只会在调用`GET /v1/blog/user/login`的`OAuth2`单点登录跳转Github时被触发

## 3.配置文件

- 配置说明

| 位置   | 项                | 数据类型   | 备注                          |
|------|------------------|--------|-----------------------------|
| Body | OidcProvider     | string | OIDC提供者地址                |
| Body | ClientID         | string | 客户端ID             |
| Body | ClientSecret     | string | 客户端密钥            |
| Body | Domain           | string | 域名                         |
| Body | apiListeningPort | string | API监听端口                  |
| Body | MySQLAddr        | string | MySQL地址            |
| Body | MySQLPort        | string | MySQL端口                    |
| Body | MySQLUser        | string | MySQL用户名                  |
| Body | MySQLPassword    | string | MySQL密码            |
| Body | MySQLDBName      | string | MySQL数据库名                |
| Body | i18nProfilePath  | list   | 国际化配置文件路径列表        |
| Body | MaxCommentLevel  | int    | 最大评论层级                  |

- 配置文件JSON

```json
{
  "OidcProvider" : "https://token.actions.githubusercontent.com",
  "ClientID" : "",
  "ClientSecret": "",
  "Domain" : "127.0.0.1",
  "apiListeningPort" : "8080",

  "MySQLAddr" : "",
  "MySQLPort": "3306",
  "MySQLUser" : "root",
  "MySQLPassword" : "",
  "MySQLDBName" : "blog",

  "i18nProfilePath" : [
    "data/i18n/active.en.json",
    "data/i18n/active.zh.json"
  ],

  "MaxCommentLevel": 3

}
```

## 4.i18n国际化

- 通过修改请求头`Header`中的`Accept-Language`字段即可实现语言切换
  - 目前支持的语言有中文`zh-CN`和英语`en-US`
  - 默认语言为中文
- 返回值示例：
```json
{
  "Message": "Operation Success"
}
```

```json
{
  "Error": "你未登录"
}
```

## 5.Banner

``` 
____                 __      ____                        __
/\  _`\              /\ \    /\  _`\                     /\ \
\ \ \L\ \      __    \_\ \   \ \ \L\ \     ___     ___   \ \ \/'\
 \ \ ,  /    /'__`\  /'_` \   \ \ ,  /    / __`\  /'___\  \ \ , <
  \ \ \\ \  /\  __/ /\ \L\ \   \ \ \\ \  /\ \L\ \/\ \__/   \ \ \\`\
   \ \_\ \_\\ \____\\ \___,_\   \ \_\ \_\\ \____/\ \____\   \ \_\ \_\
    \/_/\/ / \/____/ \/__,_ /    \/_/\/ / \/___/  \/____/    \/_/\/_/

=======================================================================
ProjectName: RedRockBlog
Author: TuF3i
GitHub: https://github.com/TuF3i
=======================================================================
```