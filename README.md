# API


## API 명세

**Health Check**
- URL: /health
- Method: GET
- Response 성공시: 
```json
OK
```

**Get Users**

- URL: /user
- Method: GET
- Response 성공시:

```json
[
  {
    "id": 1,
    "username": "user1",
    "email": "user1@example.com"
  },
  {
    "id": 2,
    "username": "user2",
    "email": "user2@example.com"
  }
]
```

**Get User by ID**

- URL: /user/{id}
- Method: GET
- Path Parameters:
  - id: 사용자 ID (integer)
- Response 성공시:
```json
{
  "id": 1,
  "username": "user1",
  "email": "user1@example.com"
}
```

**Create User**
- URL: /user
- Method: POST
- Request 예시: 
```json
{
  "id": 1,
  "username": "user1",
  "email": "user1@example.com"
}
```
- Response 성공시:
```
201 Created
```

**Update User**
- URL: /user/{id}
- Method: PUT
- Path Parameters:
    - id: 사용자 ID (integer)
- Request 에시:
```json
{
  "id": 1,
  "username": "user1change",
  "email": "user1@example.com"
}
```
- Response 성공시:
```
200 OK
```

**Delete User**
- URL: /user/{id}
- Method: DELETE
- Path Parameters:
- id: 사용자 ID (integer)
- Response 성공시:
```
204 (No Content)
```