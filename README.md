# go-webserver

## Run the server
Build the executable from the root directory
```
go build .
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