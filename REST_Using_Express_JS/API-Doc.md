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

### 1.5 Request and Response
#### 1.5.1
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

### 2.1 URL
`http://${host}:1234/api/users/login`

### 2.2 Method
- [x] POST
- [ ] GET
- [ ] DELETE
- [ ] PUT
- [ ] GET

### 2.3 Request Parameters
Using: `Body`
|Params|Type|Required|Description|Example|
|:--:|:--:|:--:|:--:|:--:|
|email|String|:heavy_check_mark:|User Email|example@mail.com|
|password|String|:heavy_check_mark:|User Password|1234567|

### 2.4 Result Description
|Params||Type|Description|Example|
|:--:|:--:|:--:|:--:|:--:|
|message||String|Login Result|Login Successful|
|data||Object|||
||id|Number|User ID|2|
||name|String|User Name|SEECHEN|
||email|String|User Email|example@email.com|
||address|String|User Address|Address|

### 2.5 Request and Response
#### 2.5.1
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

## 3. Get User Info
> View user information.

### 3.1 URL
`http://${host}:1234/api/users/info`

### 3.2 Method
- [ ] POST
- [x] GET
- [ ] DELETE
- [ ] PUT
- [ ] GET

### 3.3 Request Parameters
Using: `Params`
|Params|Type|Required|Description|Example|
|:--:|:--:|:--:|:--:|:--:|
|id|Number||User ID|1|
|email|String||User Email|example@mail.com|

*Select at least one.*

### 3.4 Result Description
|Params||Type|Description|Example|
|:--:|:--:|:--:|:--:|:--:|
|id||Number|User ID|1|
|name||String|User Name|SEECHEN|
|email||String|User Email|example@email.com|
|address||String|User Address|Address|
|password||String|Encrypted password|$2a$10$at.iG0NNt3WBXF9pgt|
|userCompany||Object|The company to which the user belongs and their permissions.|null|

### 3.5 Request and Response
#### 3.5.1
**Request**:
```text
http://127.0.0.1:1234/api/users/info?id=1
```
*Login ID is not the requester's ID*

**Reponse**:
```json
{
    "id": 1,
    "name": "Name",
    "email": "example@email.com",
    "address": "Address"
}
```

#### 3.5.2
**Request**:
```text
http://127.0.0.1:1234/api/users/info?id=1
```
*Login ID is the requester's ID*

**Reponse**:
```json
{
    "id": 2,
    "name": "User Name",
    "email": "example@email.com",
    "address": "Adderss",
    "password": "$2a$10$at.iG0NNt3",
    "userCompany": null
}
```

#### 3.5.3
**Request**:
```text
http://127.0.0.1:1234/api/users/info?email=example@email.com
```
*Get information using email*

**Reponse**:
```json
{
    "id": 2,
    "name": "User Name",
    "email": "example@email.com",
    "address": "Adderss"
}
```
