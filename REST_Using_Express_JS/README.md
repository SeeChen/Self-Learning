<div align=center>

# REST Using Express JS

[Summary](#1-summary)</br>
[Project Setup](#2-project-setup)</br>
[Project Architecture](#3-project-architecture)</br>
[API Test](#4-api-test)</br>
[API Documentation](#5-api-documentation)</br>

</div>

## 1. Summary
This task involves building a server with REST API using Express.js and TypeScript. The task also requires implementing basic database CRUD operations, as outlined in the provided [task file](./Coding_Question_Set_1.txt).

## 2. Project Setup
### 2.1 Dependencies
The dependencies for this project are as follows:
```
rest_using_express_js@1.0.0 \SeeChen\Self-Learning\REST_Using_Express_JS
+-- @prisma/client@5.21.0
+-- @types/bcryptjs@2.4.6
+-- @types/express@5.0.0
+-- @types/jsonwebtoken@9.0.7
+-- @types/node@22.7.6
+-- bcryptjs@2.4.3
+-- cookie-parser@1.4.7
+-- express@4.21.1
+-- jsonwebtoken@9.0.2
+-- prisma@5.21.0
+-- ts-node@10.9.2
+-- tsx@4.19.1
`-- typescript@5.6.3
```

Additionally, you can view the [package.json](./package.json) or [package-lock.json](./package-lock.json) for more details.

### 2.2 Database
I am use SQLite for development. The schema is defined in [schema.prisma](./prisma/schema.prisma). 

The database structure is as follows:
```prisma
model User {
  id          Int    @id @default(autoincrement())
  name        String
  email       String @unique
  password    String
  address     String
  userCompany UserCompany?
}

model Company {
  id      Int     @id @default(autoincrement())
  name    String  @unique 
  address String
  users   UserCompany[]
}

model UserCompany {
  userId    Int   @unique
  companyId Int
  role      Int // 0: Admin, 1: User

  user      User    @relation(fields: [userId], references: [id], onDelete: Cascade)
  company   Company @relation(fields: [companyId], references: [id], onDelete: Cascade)

  @@id([userId, companyId])
}
```

## 3. Project Architecture
This task utilizes a simple layered architecture, as shown below:
```
src
├───auth
├───controllers
├───dao
├───middleware
├───routes
└───Services
```
Currently, the architecture consists of six layers: auth, controllers, dao, middleware, routes, and Services.

### **Auth**
`ROLE`:
> The auth layer abstracts all authentication-related operations, encapsulating them into reusable functions. By consolidating these operations, it simplifies the handling of authentication and authorization across the application.

`FUNCTION`:
> The auth layer provides methods like verifyToken, verifyPassword, and genToken, which are commonly used in various parts of the app. This layer ensures consistency and security in authentication workflows by centralizing similar code into one file.

### **Controller**
`ROLE`:
> The controller layer acts as an intermediary between the routes and the business logic of the application. It processes incoming requests, interacts with the model (or database), and prepares the response.

`FUNCTION`:
> Controllers handle the core functionality of the application. They receive data from the routes, perform operations (such as fetching or saving data), and return the appropriate response to the client.

### **Routes**
`ROLE`:
> The routes layer defines the endpoints of the application and maps incoming requests to the appropriate controller functions.

`FUNCTION`:
> Routes specify how the application responds to client requests. They determine what happens when a particular URL is accessed and specify which controller should handle the request.

### **Middleware**
`ROLE`:
> Middleware functions are functions that have access to the request object, response object, and the next middleware function in the application’s request-response cycle.

`Function`:
> Middleware can modify the request or response, end the request-response cycle, or call the next middleware in the stack. They are commonly used for tasks such as authentication, logging, and error handling.

### **DAO (Data Access Object)**
`ROLE`:
> The DAO layer provides an abstraction over direct database operations, making it easier to manage data access and manipulation consistently across the application.

`Function`:
> DAOs handle all database queries and operations (such as creating, reading, updating, and deleting records) by providing a standardized interface. This keeps data logic separated from business logic and supports more maintainable and scalable database interactions.

### **Services**
`ROLE`:
> The services layer encapsulates the core business logic and workflows of the application. It serves as an intermediary between controllers and DAOs or external data sources, handling the essential rules and processes.

`Function`:
> Services define how the application processes data and implement reusable functions that encapsulate specific workflows and rules. This helps maintain consistency in business logic across different parts of the app and supports a clear separation of concerns.

## 4. API Test
You can run the server with the command `npx tsx server.ts`.

To test my API and understand its functionality, you can use Postman. Download the collection file [REST.postman_collection.json](./REST.postman_collection.json) or [click here](https://yiyingpiaopiao.postman.co/workspace/user-api~6adfe4bb-2bc2-4b1d-8ba6-66928967946e/collection/25397697-e8cf12f6-9fa3-44d8-a088-68589638c2f2?action=share&creator=25397697) to access the Postman website.

## 5. API Documentation
[API Documetation](./API-Doc.md)

---
<div align="right">

###### *Last Modified by [SeeChen](https://github.com/SeeChen/) @ 25-OCT-2024 19:01 UTC +08:00*
</div>