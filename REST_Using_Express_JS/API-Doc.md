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
#### 3.5.1 *Login ID is not the requester's ID*
**Request**:
```text
http://127.0.0.1:1234/api/users/info?id=1
```

**Reponse**:
```json
{
    "id": 1,
    "name": "Name",
    "email": "example@email.com",
    "address": "Address"
}
```

#### 3.5.2 *Login ID is the requester's ID*
**Request**: 
```text
http://127.0.0.1:1234/api/users/info?id=1
```

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

#### 3.5.3 *Get information using email*
**Request**:
```text
http://127.0.0.1:1234/api/users/info?email=example@email.com
```

**Reponse**:
```json
{
    "id": 2,
    "name": "User Name",
    "email": "example@email.com",
    "address": "Adderss"
}
```
## 4. User Update
> Update information.

### 4.1 URL
`http://${host}:1234/api/users/update`

### 4.2 Method
- [ ] POST
- [ ] GET
- [ ] DELETE
- [x] PUT

### 4.3 Request Parameters
Using: `Body`
|Params|Type|Required|Description|Example|
|:--:|:--:|:--:|:--:|:--:|
|name|String||User New Name|New Name|
|email|String||User New Email|new_example@mail.com|
|passwordOld|String|Required if the new password has a value|Old Pawword|Password123546|
|passwordNew|String||New Password|123456Password|
|address|String||User New Address|New Address|

### 4.4 Result Description
|Params||Type|Description|Example|
|:--:|:--:|:--:|:--:|:--:|
|id||Number|User Id|1|
|name||String|User New Name|New Name|
|email||String|User New Email|new_example@email.com|
|password||String|Encrypted password|fvmnisufnij|
|address||String|New User Address|New Address|

### 4.5 Request and Response
#### 4.5.1
**Request**:
```json
{
    "name": "New Name",
    "email": "new_example@email.com",
    "passwordOld": "",
    "passwordNew": "",
    "address": "New Address"
}
```
**Reponse**:
```json
{
    "id": 2,
    "name": "New Name",
    "email": "new_example@email.com",
    "password": "$2a$10$at",
    "address": "New Address"
}
```

## 5. User Delete
> Update information.

### 5.1 URL
`http://${host}:1234/api/users/delete`

### 5.2 Method
- [ ] POST
- [ ] GET
- [x] DELETE
- [ ] PUT

### 5.3 Request Parameters
Using: `Body`
|Params|Type|Required|Description|Example|
|:--:|:--:|:--:|:--:|:--:|
|name|String|:heavy_check_mark:|User Name|Name|
|email|String|:heavy_check_mark:|User Email|example@mail.com|
|password|String|:heavy_check_mark:|Password|Password123546|

### 5.4 Result Description
|Params||Type|Description|Example|
|:--:|:--:|:--:|:--:|:--:|
||||||

### 5.5 Request and Response
#### 5.5.1
**Request**:
```json
{
    "name": "New Name",
    "email": "new_example@email.com",
    "password": "Password123456"
}
```
**Reponse**:
```text
Delete Successful.
```

## 6. Create Company
> Create a company.

### 6.1 URL
`http://${host}:1234/api/company/create`

### 6.2 Method
- [x] POST
- [ ] GET
- [ ] DELETE
- [ ] PUT

### 6.3 Request Parameters
Using: `Body`
|Params|Type|Required|Description|Example|
|:--:|:--:|:--:|:--:|:--:|
|name|String|:heavy_check_mark:|Company Name|Company Name|
|address|String|:heavy_check_mark:|Company Address|Company Address|

*Users must be logged in to create a company.*

### 6.4 Result Description
|Params||Type|Description|Example|
|:--:|:--:|:--:|:--:|:--:|
|id||Number|Company Id in database|1|
|name||String|Company Name|Company Name|
|address||String|Company Address|Company Address|

### 6.5 Request and Response
#### 6.5.1
**Request**:
```json
{
    "name": "Company Name",
    "address": "Company Address"
}
```
**Reponse**:
```json
{
    "id": 1,
    "name": "Company Name",
    "address": "Company Address"
}
```

## 7. Get Company Information
> Get Company Information.

### 7.1 URL
`http://${host}:1234/api/company/info`

### 7.2 Method
- [ ] POST
- [x] GET
- [ ] DELETE
- [ ] PUT

### 7.3 Request Parameters
Using: `Params`
|Params|Type|Required|Description|Example|
|:--:|:--:|:--:|:--:|:--:|
|id|Number||ID of Compnay|1|
|name|String||Company Name|Company Name|

*Select at least one.*

### 7.4 Result Description
|Params||Type|Description|Example|
|:--:|:--:|:--:|:--:|:--:|
|id||Number|ID of Company|1|
|name||String|Company Name|Company Name|
|address||String|Company Address|Address|
|users||List|Include users in the company|null|
||userId|Number|User Id|1|
||companyId|String|Company ID|1|
||role|Number|User Permissions|0|

### 7.5 Request and Response
#### 7.5.1 *Get information by company ID*
**Request**:
```text
http://127.0.0.1:1234/api/company/info?id=1
```
**Reponse**:
```json
{
    "id": 1,
    "name": "Company Name",
    "address": "Company Address"
}
```

#### 7.5.2 *Get information by company name*
**Request**:
```text
http://127.0.0.1:1234/api/company/info?name=CompanyName
```
**Reponse**:
```json
{
    "id": 1,
    "name": "Company Name",
    "address": "Company Address"
}
```

#### 7.5.3 *Get information when the user searches for a company*
**Request**:
```text
http://127.0.0.1:1234/api/company/info?name=CompanyName&id=1
```
**Reponse**:
```json
{
    "id": 1,
    "name": "Company Name",
    "address": "Company Address",
    "users": [
        {
            "userId": 1,
            "companyId": 1,
            "role": 0
        }
    ]
}
```

## 8. Update Company Information
> Update Company Information.

### 8.1 URL
`http://${host}:1234/api/company/update`

### 8.2 Method
- [ ] POST
- [ ] GET
- [ ] DELETE
- [x] PUT

### 8.3 Request Parameters
Using: `Params`
|Params|Type|Required|Description|Example|
|:--:|:--:|:--:|:--:|:--:|
|id|Number|:heavy_check_mark:|Company ID|1|

Using: `body`
|Params|Type|Required|Description|Example|
|:--:|:--:|:--:|:--:|:--:|
|name|String||New Company Name|New Company Name|
|address|String||New Company address|Address-new|

*Users must be logged in to update a company.*

### 8.4 Result Description
|Params||Type|Description|Example|
|:--:|:--:|:--:|:--:|:--:|
|id||Number|Company Id in database|1|
|name||String|Company Name|New Company Name|
|address||String|Company Address|New Company Address|

### 8.5 Request and Response
#### 8.5.1 *Two parameters*
**Request**:

*`Params`*
```text
http://127.0.0.1:1234/api/company/update?id=1
```
*`body`*
```json
{
    "name": "Company Example",
    "address": "New-Address"
}
```
**Reponse**:
```json
{
    "id": 1,
    "name": "Company Example",
    "address": "New-Address"
}
```

#### 8.5.2 *Company Name*
**Request**:

*`Params`*
```text
http://127.0.0.1:1234/api/company/update?id=1
```
*`body`*
```json
{
    "name": "Company Example",
    "address": ""
}
```
**Reponse**:
```json
{
    "id": 1,
    "name": "Company Example",
    "address": "Old-Address"
}
```

#### 8.5.3 *Company Address*
**Request**:

*`Params`*
```text
http://127.0.0.1:1234/api/company/update?id=1
```
*`body`*
```json
{
    "name": "Company Name",
    "address": "New-Address"
}
```
**Reponse**:
```json
{
    "id": 1,
    "name": "Company Name",
    "address": "New-Address"
}
```

## 9. Delete Company
> Delete Company.

### 9.1 URL
`http://${host}:1234/api/company/delete`

### 9.2 Method
- [ ] POST
- [ ] GET
- [x] DELETE
- [ ] PUT

### 9.3 Request Parameters
Using: `Params`
|Params|Type|Required|Description|Example|
|:--:|:--:|:--:|:--:|:--:|
|id|Number|:heavy_check_mark:|Company ID|1|

Using: `body`
|Params|Type|Required|Description|Example|
|:--:|:--:|:--:|:--:|:--:|
|companyName|String|:heavy_check_mark:|Manual input to prevent misoperation|Company Name|
|adminEmail|String|:heavy_check_mark:|Manual input to prevent misoperation|example@email.com|
|adminPassword|String|:heavy_check_mark:|Manual input to prevent misoperation|admin123456|

### 9.4 Result Description
|Params||Type|Description|Example|
|:--:|:--:|:--:|:--:|:--:|
||||||

### 9.5 Request and Response
#### 9.5.1
**Request**:

*`Params`*
```text
http://127.0.0.1:1234/api/company/delete?id=1
```
*`body`*
```json
{
    "companyName": "Company",
    "adminEmail": "example@email.com",
    "adminPassword": "admin123456"
}
```
**Reponse**:
```text
Delete Successful.
```
## 10. Assign User
> Assign user to a company.

### 10.1 URL
`http://${host}:1234/api/company/manage/assign`

### 10.2 Method
- [x] POST
- [ ] GET
- [ ] DELETE
- [ ] PUT

### 10.3 Request Parameters
Using: `body`
|Params|Type|Required|Description|Example|
|:--:|:--:|:--:|:--:|:--:|
|userId|Number|:heavy_check_mark:|The user ID you want to assign|2|
|role|Number|:heavy_check_mark:|Permission, 0 is admin, 1 is user|0|

### 10.4 Result Description
|Params||Type|Description|Example|
|:--:|:--:|:--:|:--:|:--:|
|userId||Number|Assigned User ID|2|
|companyId||Number|The company id to which the user has been assigned|1|
|role||Number|Permission|0|

### 10.5 Request and Response
#### 10.5.1
**Request**:

*`body`*
```json
{
    "userId": 2,
    "role": 1
}
```
**Reponse**:
```json
{
    "userId": 2,
    "companyId": 1,
    "role": 1
}
```
## 11. Remove User
> Remove user from a company.

### 11.1 URL
`http://${host}:1234/api/company/manage/remove`

### 11.2 Method
- [ ] POST
- [ ] GET
- [x] DELETE
- [ ] PUT

### 11.3 Request Parameters
Using: `body`
|Params|Type|Required|Description|Example|
|:--:|:--:|:--:|:--:|:--:|
|userId|Number|:heavy_check_mark:|The user ID you want to remove|2|

### 11.4 Result Description
|Params||Type|Description|Example|
|:--:|:--:|:--:|:--:|:--:|
||||||

### 11.5 Request and Response
#### 11.5.1
**Request**:

*`body`*
```json
{
    "userId": 2
}
```
**Reponse**:
```text
Remove Successful.
```
## 12. User Role
> Modify the user role.

### 12.1 URL
`http://${host}:1234/api/company/manage/modify-role`

### 12.2 Method
- [ ] POST
- [ ] GET
- [ ] DELETE
- [x] PUT

### 12.3 Request Parameters
Using: `body`
|Params|Type|Required|Description|Example|
|:--:|:--:|:--:|:--:|:--:|
|userId|Number|:heavy_check_mark:|The user ID you want to modify|2|
|roleNew|Number|:heavy_check_mark:|New role|0|

### 12.4 Result Description
|Params||Type|Description|Example|
|:--:|:--:|:--:|:--:|:--:|
|userId||Number|User ID|2|
|companyId||Number|The company id|1|
|role||Number|Permission|0|

### 12.5 Request and Response
#### 12.5.1
**Request**:

*`body`*
```json
{
    "userId": 2,
    "roleNew": 0
}
```
**Reponse**:
```json
{
    "userId": 2,
    "companyId": 1,
    "role": 0
}
```
---
<div align="right">

###### *Last Modified by [SeeChen](https://github.com/SeeChen/) @ 25-OCT-2024 19:48 UTC +08:00*
</div>