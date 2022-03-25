# go-webserver

## Dependencies

[Go](https://go.dev/doc/install)

## Build app
Build the executable from the root directory
```
go build .
```

## Run the executable
Double click the `go-webserver` file or open with command line
```
./go-webserver
```

## Endpoints

Get all users
```
[GET] http://localhost:3000/users
```

Add new user
```
[POST] http://localhost:3000/users
[Body] {"FirstName": "Bob", "LastName": "Loblaw"}
```

Get user by id
```
[GET] http://localhost:3000/users/1
```

Update user 
```
[PUT] http://localhost:3000/users/1
[Body] {"ID": 1, "FirstName": "Bobbie", "LastName": "Loblawbie"}
```

Delete user
```
[DELETE] http://localhost:3000/users/1
```