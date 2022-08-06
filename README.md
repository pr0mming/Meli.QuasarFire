## Meli.QuasarFire :fire:
Small application for decode the secret message and get the location received from the other enemy ship!

Using Go and a lot of math :)

## Technologies :rocket:

 - Go v1.17
 - Gin Gonic v1.8.1
 - Wire v0.5.0

## Startup :vertical_traffic_light:

For run the project in localhost you can follow this steps:

- Use `make generate` from the project root for generate the dependencies with Wire
- Use `make run` for start the project without update the Wire files
- Use `make all` (recommended) is the plus between `make generate` and `make run`
- For default, Gin will open the service on `http://localhost:8080`
