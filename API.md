# Rikka API 简介

## 上传

### 地址

`POST /api/upload`

### 参数

`Content-type` 为 `multipart/form-data`。

- `uploadFile"`：要上传的文件
- `password"`：Rikka 的密码
- `from"`：必须项，指示请求来源

### 返回值

上传成功则返回任务 ID，可通过任务 ID 获取任务状态和图片地址

```json
{
    "TaskID": "2016-09-02-908463234"
}
```

失败返回错误，格式见后。

## 获取任务状态

### 地址

`GET /api/state/<TaskID>`

### 参数

无

### 返回值

查询成功则返回任务状态。

```json
{
    "TaskID": "2016-09-02-908463234",
    "State": "state name",
    "StateCode": 1,
    "Description": "state Description"
}
```

失败返回错误，格式见后。

## 获取图片地址

### 地址

`GET /api/url/<TaskID>`

### 参数

无

### 返回值

查询成功则返回图片原始 URL。

```json
{
    "URL": "http://127.0.0.1/files/2016-09-02-908463234"
}
```

失败返回错误，格式见后。

## 返回错误

如果 API 请求出错，返回格式如下：

```json
{
    "Error": "error message"
}
```
