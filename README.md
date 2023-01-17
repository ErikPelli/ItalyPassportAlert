# ItalyPassportAlert
A script to check for available slots to apply for a passport in Italy.
Send one notify and a sound to the user when it's available (and then the faster one wins).

Copy `.env.example` into `.env` and set up it as you want.

## Requisites
[Go 1.19+](https://go.dev/dl/)

## Run
```
git clone https://github.com/ErikPelli/ItalyPassportAlert.git
cd ItalyPassportAlert

# Option 1: Build an executable and use it
go build cmd/checker.go
# Linux: ./checker
# Windows: checker.exe or click on it

# Option 2
go run cmd/checker.go
```