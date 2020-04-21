```

```

# Url shortener description 

### Create short Url:

>Method: Post
>
>Path:**/**
>
>Body:

```json5
{
   "url": "http://goole.com",
}
```
> Response body:

```json5
 {
   "shortUrl": "http://domain/jx_ekffjxa"
 }
```
### Get origin url by short:

 >  Method: GET  
>  
> Path: **/{shortUrl}**  
>
>
 
## Docker

1. Сборка: `docker build -t url_shortener .`
2. Запуск: `docker run -p 8000:8000 url_shortener:latest .`
3. доступен по адресу `http://{dockerhost}:8000/`
