# API Documentation

## 1. User Register
> Registers a new user by providing user details.

### 1.1 URL
`http://${host}:1234/api/users/create`

### 1.2 Method
- [x] POST
- [ ] GET
- [ ] DELETE
- [ ] PUT
- [ ] GET

### 1.3 Request Parameters
Using: `Body`
|Params|Type|Required|Description|Example|
|:--:|:--:|:--:|:--:|:--:|
|name|String|:heavy_check_mark:|Name of User|SeeChen|
|email|String|:heavy_check_mark:|User Email|example@mail.com|
|password|String|:heavy_check_mark:|User Password|1234567|
|address|String|:heavy_check_mark:|User Address|Address|

### 1.4 Result Description
|Params||Type|Description|Example|
|:--:|:--:|:--:|:--:|:--:|
|id||Number|Automatically assigned id|1|
|name||String|User Name|SeeChen|
|email||String|User Email|example@email.com|
|address||String|User Address|Address|
|userCompany||Object|The company to which the user belongs and their permissions.|null|

### 1.4 Request and Response
#### 1.4.1
**Request**:
```json
{
    "name": "User Name",
    "email": "example@email.com",
    "password": "Password123456",
    "address": "Adderss"
}
```
**Reponse**:
```json
{
    "id": 1,
    "name": "User Name",
    "email": "example@email.com",
    "address": "Adderss",
    "userCompany": null
}
```

## 2. User Login
> User login interface after registration.

### 1.1 URL
`http://${host}:1234/api/users/login`

### 1.2 Method
- [x] POST
- [ ] GET
- [ ] DELETE
- [ ] PUT
- [ ] GET

### 1.3 Request Parameters
Using: `Body`
|Params|Type|Required|Description|Example|
|:--:|:--:|:--:|:--:|:--:|
|email|String|:heavy_check_mark:|User Email|example@mail.com|
|password|String|:heavy_check_mark:|User Password|1234567|

### 1.4 Result Description
|Params||Type|Description|Example|
|:--:|:--:|:--:|:--:|:--:|
|message||String|Login Result|Login Successful|
|data||Object|||
||id|Number|User ID|2|
||name|String|User Name|SEECHEN|
||email|String|User Email|example@email.com|
||address|String|User Address|Address|

### 1.4 Request and Response
#### 1.4.1
**Request**:
```json
{
    "email": "example@email.com",
    "password": "Password123456"
}
```
**Reponse**:
```json
{
    "message": "Login Successful",
    "data": {
        "id": 2,
        "name": "User Name",
        "email": "example@email.com",
        "address": "Adderss"
    }
}
```
