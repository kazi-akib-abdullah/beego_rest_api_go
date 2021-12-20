# beego_rest_api_go

### Prerequisites
- Install go
- Install beego
- Set GOPATH and GOBIN
### Start With the Project
- Go to the project directory
- create database
- go mod tidy
- bee run [Would run on localhost:/8080 ]

### Database Create
- Default connection settings to database server 
```python
			host     = "localhost"
			port     = 5432
			user     = "postgres"
			password = "newPassword"
			dbname   = "user_db"
```
- Table will automatically Create if there is no table with name "post_data"

### Features
- Only Valid Email, Phone and Date will save the data to database
- Password should be greater than 6 digits.
- password will stored as a hash value
