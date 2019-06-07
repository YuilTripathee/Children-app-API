# API documentation for Children app API

## Routes design

Routes (as in `routes.json`):
```json
[
  {
    "method": "GET",
    "path": "/alphabets/local",
    "name": "main.getAllAlpLocal"
  },
  {
    "method": "GET",
    "path": "/alphabets/remote",
    "name": "main.getAllAlpRemote"
  },
  {
    "method": "GET",
    "path": "/alphabet",
    "name": "main.getOneAlp"
  },
  {
    "method": "GET",
    "path": "/img/write/:id",
    "name": "main.getAlpImgWrite"
  },
  {
    "method": "GET",
    "path": "/img/for/:id",
    "name": "main.getAlpImg"
  },
  {
    "method": "GET",
    "path": "/aud/norm/:id",
    "name": "main.getAlpAud"
  },
  {
    "method": "GET",
    "path": "/aud/for/:id",
    "name": "main.getAlpAudFor"
  },
  {
    "method": "GET",
    "path": "/",
    "name": "main.main.func2"
  },
  {
    "method": "GET",
    "path": "/*",
    "name": "github.com/labstack/echo.common.static.func1"
  }
]
```

### Sending entire data (for local + remote resource)

Request:

```bash
REQUEST METHOD  :GET
REQUEST PATH    :/alphabets/remote
```

Response:

```json
{
  "data": [
    {
      "audio_for_ref": "http://127.0.0.1:8000/aud/for/a",
      "audio_norm_ref": "http://127.0.0.1:8000/aud/norm/a",
      "id": 0,
      "img_ref": "http://pngimg.com/uploads/apple/apple_PNG12405.png",
      "l_case": "a",
      "u_case": "A",
      "vowel_conf": true,
      "write_img_ref": "https://upload.wikimedia.org/wikipedia/commons/c/c7/A_cursiva.gif"
    },
    {
      "audio_for_ref": "http://127.0.0.1:8000/aud/for/b",
      "audio_norm_ref": "http://127.0.0.1:8000/aud/norm/b",
      "id": 1,
      "img_ref": "http://pngimg.com/uploads/football/football_PNG52789.png",
      "l_case": "b",
      "u_case": "B",
      "vowel_conf": false,
      "write_img_ref": "https://upload.wikimedia.org/wikipedia/commons/4/48/B_cursiva.gif"
    } // ...
  ],
  "response_code": 1,
  "status_code": 200,
  "status_message": "Everything's working properly"
```

### Sending entire data locally

Request:

```bash
REQUEST METHOD  :GET
REQUEST PATH    :/alphabets/local
```

Response:

```json
{
  "data": [
    {
      "audio_for_ref": "http://127.0.0.1:8000/aud/for/a",
      "audio_norm_ref": "http://127.0.0.1:8000/aud/norm/a",
      "id": 0,
      "img_ref": "http://127.0.0.1:8000/img/for/a",
      "l_case": "a",
      "u_case": "A",
      "vowel_conf": true,
      "write_img_ref": "http://127.0.0.1:8000/img/write/a"
    },
    {
      "audio_for_ref": "http://127.0.0.1:8000/aud/for/b",
      "audio_norm_ref": "http://127.0.0.1:8000/aud/norm/b",
      "id": 1,
      "img_ref": "http://127.0.0.1:8000/img/for/b",
      "l_case": "b",
      "u_case": "B",
      "vowel_conf": false,
      "write_img_ref": "http://127.0.0.1:8000/img/write/b"
    } //...
  ], 
  "response_code": 1,
  "status_code": 200,
  "status_message": "Everything's working properly"
```

### Individual data (local data reference served)

Request:

```bash
REQUEST METHOD  :GET
REQUEST PATH    :/alphabet?id=ID
                :/alphabet?id=ID&source=remote # for accessing remote data source
```

Response:

```json
// local version
{
  "data": {
    "audio_for_ref": "http://127.0.0.1:8000/aud/for/a",
    "audio_norm_ref": "http://127.0.0.1:8000/aud/norm/a",
    "id": 0,
    "img_ref": "http://127.0.0.1:8000/img/for/a",
    "l_case": "a",
    "u_case": "A",
    "vowel_conf": true,
    "write_img_ref": "http://127.0.0.1:8000/img/write/a"
  },
  "response_code": 1,
  "status_code": 200,
  "status_message": "Everything's working properly"
}
```

```json
// remote version
{
  "data": {
    "audio_for_ref": "http://127.0.0.1:8000/aud/for/a",
    "audio_norm_ref": "http://127.0.0.1:8000/aud/norm/a",
    "id": 0,
    "img_ref": "http://pngimg.com/uploads/apple/apple_PNG12405.png",
    "l_case": "a",
    "u_case": "A",
    "vowel_conf": true,
    "write_img_ref": "https://upload.wikimedia.org/wikipedia/commons/c/c7/A_cursiva.gif"
  },
  "response_code": 1,
  "status_code": 200,
  "status_message": "Everything's working properly"
}
```

## Accessing media resources via API

### Accessing image resources (remotely)

Request:

```bash
REQUEST TYPE    :GET
REQUEST URL     : "the URL given in alphabet data"

```

Response:

```
Image file of the specific request.
```

### Accessing image resources (locally)

Request:

```bash
REQUEST METHOD  :GET
REQUEST PATH    :/img/for/:id # for images referring to letter usage
REQUEST PATH    :img/write/:id # for images to help the alphabet get written (in GIF form)
```

Response:

```
Image file for the specified request.
```

### Accessing audio resources (local option only)

Request:

```bash
REQUEST TYPE    :GET
REQUEST URL     :/aud/norm/:id # for pronounciation
REQUEST URL     :/aud/for/:id # for word usage
```

Response:

```
Audio (mp3) file for the specified request.
```

## Miscellanous routes

### API Info

The basic information regarding the API being used.

Request:

```bash
REQUEST METHOD  :GET
REQUEST PATH    :/
```

Response:

```json
{
  "bytes_out": "0",
  "developer": "Yuil Tripathee",
  "ip": "127.0.0.1",
  "method": "GET",
  "name": "Children_app_api",
  "protocol": "HTTP/1.1",
  "status_code": "200",
  "time": "2019/06/06 - 09:07:32",
  "url": "/",
  "version": "v1.0"
}
```



### Landing page

The contents for landing page is placed in `static\` directory.

Request:

```bash
REQUEST METHOD  :GET
REQUEST PATH    :/*
```

Response: Based on the resource you're accessing (png, html, css, js, ...)

## Validation and errors

Format for the error message:

```json
[
    {
        "status_code" : 500,
        "response_code" : 0,
        "status_message" : "Invalid Input"
    },
    {
        "status_code" : 404,
        "response_code" : -1,
        "status_message" : "Must provide an argument"
    }
]
```

The error handling is done within the route handler function.

## Running servers

Firing up server in  development environment:

```c
./<server_binary> <arg1:DEV> >> <log_file>.log // after compilation of binary
// no LOG needed for debuggnig purpose but can be done over the need.
```

Firing up server in production environment:

```c
./<server_binary> <arg1:PROD> >> <log_file>.log // after compilation of binary
// LOG seems to be optional here also.
```

> **Note**:  
> Double debugging can be done by supplying custom port by setting up a command line flag with ease at any instant.

## About the log format

### LOG for DEVELOPMENT

Log output for debugging includes only the essential a backend developer may need at development phase of software development.

```cs
app_name-environment timestamp | status_code | protocol | client_ip | method path
```

Example:

```cs
[APP-DEBUG] 2019/04/25 - 00:06:12 | 300 | HTTP/1.1 |  127.0.0.1 | DELETE  /v1/app/ // method and status are colored by their nature
```

### LOG for PRODUCTION

Log output for production includes all the data essential for application analytics, tailored specifically to be processable is easy to use formats like CSV (and JSON as well). The logger for production is worth storing in a LOG for further appliments.

```cs
app_name-environment | host | time_custom | status | latency_human | remote_ip | bytes_in bytes_in | bytes_out bytes_out | method | uri
```

Example:

```cs
APP-PROD | 127.0.0.1:8000 | 2019/04/25 00:06:12 | 300 | 215.117µs | 127.0.0.1 | 0 bytes_in | 41 bytes_out | GET | /v1/app/ // clean - no color
```

## Encoding log to analyzable CSV format

Algorithm to do so:

1. GET the CSV file with custom delimeter
2. Strip all the JSON based log printouts onto a separate file.
3. Relace all commas in a file (use URI decoding or Unicode decoding in advance).
4. Remove all the unwanted strings, texts, spaces from the file.
5. Replace all pipes '|' with comma, since the pipes are used to as a delimeter.

> **Note:**  
> Python is the best suited scripting language for this purpose. Although, GO or C++ can be applied on binary level.

## About saving binary log on linux machine

If you run binary on linux machine you could use shell script.

Overwrite into a file

```shell
./binaryapp > binaryapp.log
```

append into a file

```shell
./binaryapp >> binaryapp.log
```

overwrite stderr into a file

```shell
./binaryapp &> binaryapp.error.log
```

append stderr into a file

```shell
./binaryapp &>> binalyapp.error.log
```

It can be more dynamic using shell script file.  
Reference : [Solution from StackOverflow](https://stackoverflow.com/questions/19965795/go-golang-write-log-to-file)

# Containerization with Docker

Building the image for the project:

```bash
docker build --tag children_app_api .
# where,
# children_app_api is tagname for your image
```

Running the app within the container:

```bash
docker run -p 4000:8000 children_app_api
# where,
# 4000 is a port of system
# 8000 is a port of virtual system
# children_app_api is the image
```

## References

1. https://hackernoon.com/golang-docker-microservices-for-enterprise-model-5c79addfa811