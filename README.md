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

## Result
In the province of Padova there was only one place where it was possible to make reservations, in the city center.
There wasn't even a single time slot available, but every morning at 8:00 a.m. you can reserve a slot for 4 months
in the future, and by 8:03 a.m. they are already gone.

I tried my own script by running it at 9:15 p.m. on 17/01/2023, and at 9:33 p.m. I was able to get a place for 2 days later!
You have to be quick to get the seats of those who cancel (with the help of this bot), in fact I had already opened
the SPID session in the browser, I had tried previously with a date that would be in 1 week but literally the reservations
go so fast that after 30 seconds it was gone.