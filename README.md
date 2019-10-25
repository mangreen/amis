# Mini Project

## Description
[Software Engineer Project](./SoftwareEngineerProject.pdf)

## Environment
- Go 1.12

Please operating under amis folder.

### golang module
start GO111MODULE mode
```
export GO111MODULE=on
```

### run
start web server
```
go run cmd/main.go
```

### Docker
If you had install Docker...

build image
```
docker build -t <amis-img> .
```

run container
```
docker run -it --rm <amis-img>
```

run web server
```
go run cmd/main.go
```

### Docker-Compose
If you had install Docker-Compose...

start docker-compose
```
docker-compose up -d --build
```

watch log of container
```
docker logs -f <amis_amis_1>
```

exec container
```
docker exec -it <amis_amis_1> bash
```

close docker-compose
```
docker-compose down
```

## API doc

http://localhost:1323

## GET /api/currencies

Display pricemovement for cryptocurrencies.

### Request

#### Parameters

| Name  | Value  | Description |
| ----- | :----- | :--------- |
| coin | string | Coin Name, only accept "btc" or "eth" |
| start | string | Start Time, only accept after "20190101"|

### Response

#### Headers

| Name  | Value  | Description |
| ----- | :----- | :--------- |
| Content-Type | application/json |  |

#### Response fields

| Name  | Value  | Description |
| ----- | :----- | :--------- |
| sources | []string | Data Sources |
| twd | int | Average TWD Price |
| usd | float | Average USD Price |
| time | string | Data Time |

#### Response example

<details>
<summary>GET http://localhost:1323/api/currencies?coin=btc&start=20190801</summary>

```javascript
{
    sources: [
        "MAX SDK",
        "coingecko.com"
    ],
    twd: 310904,
    usd: 10012.11,
    time: "20190801"
}
```
</details>
