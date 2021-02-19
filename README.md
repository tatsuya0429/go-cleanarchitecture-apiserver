# Go Clean Architecture

- echo
- gorm
- mysql
- docker
- swagger

## build

```
docker-compose up -d --build
```

## API

[Postman](https://www.postman.com/) and [Fiddler](https://www.telerik.com/fiddler) is recommended to use.

```bash
curl -X POST "http://localhost:8080/users" -H "Content-Type: application/json" -d '"{"lastname": "Murakami", "firstname": "Kohei"}"'
curl -X GET "http://localhost:8080/users/1"
curl -X GET "http://localhost:8080/users"
curl -X PUT "http://localhost:8080/users/1" -H "Content-Type: application/json" -d '"{"lastname": "David", "firstname": "San"}"'
curl -X DELETE "http://localhost:8080/users/1"
```

## document

http://localhost/docs/

## Thanks

I learned a lot from these articles.

https://nrslib.com/clean-architecture/
https://www.slideshare.net/masuda220/powered-by-spring
https://qiita.com/tono-maron/items/345c433b86f74d314c8d
https://medium.com/eureka-engineering/go-dependency-inversion-principle-8ffaf7854a55
https://qiita.com/hirotakan/items/698c1f5773a3cca6193e
https://qiita.com/Le0tk0k/items/c2945c260f28f7ee2d47
https://qiita.com/so-heee/items/0cca93008eae635c642a
https://qiita.com/ryokky59/items/6c2b35169fb6acafce15
