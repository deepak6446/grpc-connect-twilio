# OTP Generation Services

## Overview
This repository contains two services: 
1. Authentication Service: Manages user accounts and profiles. 
2. OTP Generation Service: Dedicated to generating OTPs.

## Installation
1. For each service, install dependencies by moving into respective folders and running 
`go mod tidy`. 

## Starting the Services
### Authentication Service
To start the authentication service, run the following command: 
```
cd auth-service
go run ./cmd/server/main.go 
```

### OTP Generation Service
To start the OTP generation service, run the following command: 
```
cd otp-service
update .env with your credentials for twilio
go run ./cmd/server/main.go 
```

## Install grpcurl
To install grpcurl follow command in ``` install.sh``` in each folder.

### Auth (Authentication Service)
To signup with a phone number: 
```
./grpcurl -protoset <(buf build -o -) -plaintext -d '{"first_name": "deepak", "phone_number": "8888760147", "email": "deepak.r.poojari@gmail.com", "last_name": "poojari"}' localhost:8080 auth.v1.AuthService/SignupWithPhoneNumber 
```

To Verify Account: 
```
./grpcurl -protoset <(buf build -o -) -plaintext  -d '{"phone_number": "8888760147", "otp": "6022"}' localhost:8080 auth.v1.AuthService/LoginWithPhoneNumber
```

To Login: 
```
./grpcurl -protoset <(buf build -o -) -plaintext  -d '{"phone_number": "8888760147", "otp": "6022"}' localhost:8080 auth.v1.AuthService/ValidatePhoneNumberLogin
```

To Get Profile: 
```
./grpcurl -protoset <(buf build -o -) -plaintext  -d '{"phone_number": "8888760147"}' localhost:8080 auth.v1.AuthService/GetProfile
```

### OTP Generation (OTP Generation Service)
To generate an OTP using gRPC: 
```
./grpcurl -protoset <(buf build -o -) -plaintext \ -d '{"phone_number": "8888760147"}' \ localhost:8081 otp.v1.OtpService/VerifyPhoneNumber
```

## generate code from proto file
```
buf generate
```

## Additional Resources
For more information on getting started with gRPC, refer to the [ConnectRPC documentation](https://connectrpc.com/docs/go/getting-started/).

