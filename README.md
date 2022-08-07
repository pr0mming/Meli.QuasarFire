## Meli.QuasarFire :fire:
Small application for decode the secret message and get the location received from the other enemy starship!

Using Go and a lot of math :)

## Technologies :rocket:

 - Go v1.17
 - Gin Gonic v1.8.1
 - Wire v0.5.0

## Startup :vertical_traffic_light:

For run the project in localhost you can follow this steps:

- Use `make run` for start the project without update the Wire files
- For default, Gin will open the service on `http://localhost:8080` and the Swagger on `http://localhost:8080/swagger/index.html`

Otherwise if you has made changes in the project you can use this commands:

- Use `make generate` from the project root for generate the dependencies with Wire
- Use `make swagger` from the project root for generate the Swagger Documentation
- Use `make all` (recommended) is the plus between `make generate` `make swagger` and `make run`
- Use `make test` for test the `GetLocation` and `GetMessage` methods

## References 

- Implementation with Wire fro manage Dependency Injection taken from [here](https://clavinjune.dev/en/blogs/golang-dependency-injection-using-wire/)
- The math solution for discover the enemy starship from [here](https://stackoverflow.com/questions/24970605/finding-third-points-of-triangle-using-other-two-points-with-known-distances/24980145?noredirect=1#comment129325370_24980145)
- If you want to test easly the location, you can use [this tool](https://www.mathsisfun.com/data/cartesian-coordinates-interactive.html)
- Swagger for Gin Gonic from [here](https://santoshk.dev/posts/2022/how-to-integrate-swagger-ui-in-go-backend-gin-edition/)