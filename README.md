# auth-service
This is a simple restful api that performs basic user registration and login with rate limiting

to run: 
1. spin up a postgres instance

`podman run -p 5432:5432 -e POSTGRES_PASSWORD=yourpassword postgres -d db --it`

2. run `cp .env.example .env` and populate it with appropriate values
3. run `go mod tidy && go run main.go`

4. to run tests, make sure to populate .env.test files properly and run 
`go test ./...`

Please note that this repository is intended for demonstration purposes only and should not be used in a production environment.
