# Go Clean Architecture

- echo
- gorm
- mysql
- docker
- swagger

<img width="500" alt="ulwlu/Go-CleanArchitecture-APIServer" src="https://user-images.githubusercontent.com/41639488/108515829-e0009800-7308-11eb-93aa-2823cefb83c3.png">


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

<img width="300" alt="ulwlu/Go-CleanArchitecture-APIServer" src="https://user-images.githubusercontent.com/41639488/108515020-f35f3380-7307-11eb-91b2-9378d35d9324.png">

## document

http://localhost/docs/

<img width="300" alt="ulwlu/Go-CleanArchitecture-APIServer" src="https://user-images.githubusercontent.com/41639488/108515026-f4906080-7307-11eb-8496-796cfb60393e.png">
<img width="300" alt="ulwlu/Go-CleanArchitecture-APIServer" src="https://user-images.githubusercontent.com/41639488/108515028-f528f700-7307-11eb-8bab-4773211bd161.png">

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
