# Rikka API 简介

[English version][version-en]

注意：Rikka 0.1.0 版本之后才有 API 功能。

另外，Rikka 默认不开启 CORS 支持。如果需要在其他前端页面使用上传 API，请在启动时添加 `-corsAllowOrigin '*'` 参数（0.8.0 版本后可用）。

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

## 错误格式

如果 API 请求出错，返回格式如下：

```json
{
    "Error": "error message"
}
```

[version-en]: https://github.com/7sDream/rikka/blob/master/api/README.md
