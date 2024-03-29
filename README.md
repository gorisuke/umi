
# Umi🏖️

This repository is for a web application called "umi".

Basically, it is built using [Go](https://go.dev/). [gin](https://gin-gonic.com/), a Go web framework, is used.
## Deployment

Lecture on development methods.
1. Install Air
    ```bash
    docker compose up -d
    ```
## API Reference

#### Get all items

```http
  GET /
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `api_key` | `string` | **Required**. Your API key |

#### Get item

```http
  POST /
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `string` | **Required**. Id of item to fetch |




## Acknowledgements

 - [Go Air Hotreload](https://blog.juge6.jp/go/)

## Authors

- [@gorisuke](https://www.github.com/gorisuke)

