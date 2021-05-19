# Rikka API Introduction

[中文版][version-zh]

Note: API was added into Rikka after version 0.1.0

BTW，Rikka disable CORS by default, if you need use api in other domain, please add `-corsAllowOrigin '*'` argument on start (available after version 0.8.0).

## Upload

### Path

`POST /api/upload`

### Params

`Content-type` muse be `multipart/form-data`.

- `uploadFile"`：the image file
- `password"`：Rikka password
- `from"`：request source, value must be `api`

### Return Value

A JSON contains task ID if you request successfully, then you can get task state and final image url use the ID.

```json
{
    "TaskID": "2016-09-02-908463234"
}
```

A error JSON when request failed, error message format can be found in the end of article.

## Get Task State

### Path

`GET /api/state/<TaskID>`

### Params

None

### Return Value

A state JSON like bellow if query successfully.

```json
{
    "TaskID": "2016-09-02-908463234",
    "State": "state name",
    "StateCode": 1,
    "Description": "state Description"
}
```

A error JSON when query failed, error message format can be found in the end of article.

## Get Image URL

### Path

`GET /api/url/<TaskID>`

### Params

None

### Return Value

A JSON contains image URL like bellow if query successfully.

```json
{
    "URL": "http://127.0.0.1/files/2016-09-02-908463234"
}
```

A error JSON when query failed, error message format can be found in the end of article.

## Error JSON format

If API request error, return error JSON like bellow：

```json
{
    "Error": "error message"
}
```

[version-zh]: https://github.com/7sDream/rikka/blob/master/api/README.zh.md
