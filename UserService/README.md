# USER SERVICE

## 1. Introduction
- This service is responsible for managing user's information

## 2. Service Architecture <br/>
### 2.1. Configuration
- I use Viper to read configuration file<br/>
### 2.2. Project Structure <br/>
#### 2.2.1. config folder 
- Contains config files <br/>
#### 2.2.2. db folder 
- Interacting with UserDB (MySQL) <br/>
#### 2.2.3. protobuf
- Generating interfaces and implementing handler codes for UserService <br/>
#### 2.2.4. utils
- Utility functions, helper functions <br/>
### 2.3. Project Workflow
- This project will handle RPC requests relating to user's information of another services
- The following picture shows how project handles an RPC call
- ![image desc](../docs/UserServiceProjectStructure.jpg)

## 3. Service API
### 3.1. Register
- See protobuf/Register.go
- Register a new account
- Input: Username, Password
- Output: Error if an error occurs, otherwise return empty message
- Note: <br/>
    > Username and password's length must be in range [8, 50].<br/>
    > Username is unique. <br/>
    > Password is hashed (see utils/Encrypt.go)

### 3.2 Login
- See protobuf/Login.go
- Handle request login account
- Input: Username, Password
- Output: <br/>
    > Error if error occurs <br/>
    > Otherwise return user's information: Username, UserId, Avatar. <br/>


