This is a project build in golang using gin server.

Libraries used
1. gin : Web framework for go
2. gorm : This is an ORM(Object Relational Mapper) for Golang. 
3. postgres driver for gorm : GORM dialect for postgres is installed to enable connections to postgres database.
4. crypto : Provides supplementary Go cryptography libraries
5. jwt/v4 : A Go implementation for JSON Web Tokens
6. GoDotEnv : This will help in managing environment variables

Run the following commands to get started
1. cp .env .env.local   #Would create a local env file where you can update your config

Files/Folders
1. .env : store all the environment variables required
2. models : store two models - a. User b. Entry
3. 