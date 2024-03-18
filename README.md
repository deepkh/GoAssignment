## About 
[![workflow](https://github.com/deepkh/go-recommendation-system/actions/workflows/go.yml/badge.svg)](https://github.com/deepkh/go-recommendation-system/actions)


> This project is about a Golang assignment and demonstrates a project on how to use the following technologies as well. I hope this project can be easy to understand, and I would be happy if you guys could obtain something from it as well. 

✅ Golang-1.22\
✅ BeeGo for API server\
✅ Go-Redis for database caching\
✅ Go-Redsync for distrubte locker\
✅ Gorm for mysql database operation\
✅ Grpc for micro services\
✅ JWT for token generation and parsing


## Requirements
✅ Register: GET /v1/reg?email=?&pass=?
  - [x] The account number is email
  - [x] The password must be no less than six characters and no more than 16 characters. It must have one uppercase character, one lowercase character and a special symbol ()[]{}<>+-*/?,.:;"'_\|~`! @#$%^&=

✅ Verify email (choose one of the two, or both): GET /v1/confirm?email=?&token=?
  - [x] Verify the email after clicking the link in the email
  - [x] Send a verification code to the user's email. After the user enters the verification code, the email is verified.

✅ Login: GET /v1/auth?email=?&pass=?
  - [x] Users can log in using email and password. The response will inculded a Authed Token. Use this token to GET /recommendation.

✅ GET /v1/recommendation?token=?
  - [x] Need to log in to use
  - [x] Respond to information on recommended products

## Create database and the necessary tables

We use a database `recommendation-system` with two table of `Users` and `recommendation` in this project.
Hence, we create a Hepler script to create a database and the necessary tables for convenient purposes. 
But it would be okay if you preferred to create the database manually.
We strongly recommend that you read the script before executing it as well, because the environment settings might not be the same.

```cd db/scripts && ./recommendation-system.sh```

## Build

Use the following command to build all the necessary libraries, tools and code. 
I've successfully built, tested, and run it on <ins>Debian Bookworm</ins> with 
<ins>mysql Ver 8.0.36 for Linux on x86_64 (MySQL Community Server, GPL)</ins> and 
<ins>redis-cli 7.2.4</ins>.

```GOOS=linux GOARCH=amd64 GO111MODULE="on" go build -a -o server server_main.go```

The target binary will show on **{workspace}/server** if the build process is successful.

## Run Api-Server and Grpc-Services-Server 

For simple purposes, we use different go channels to start the **Api-server** and **Grpc-Serviecs-Server** in the same binary. 
Hence, we could just start both of them with 

```./server```

## Test

Use **${workspace}/main_test.go** to test APIs.
