1. run command
 `migrate -path ./migrations -database 'postgres://user:userpass@0.0.0.0:8081/userdb?sslmode=disable' up` 

2. How to come migrations back 
 `migrate -path ./migrations -database 'postgres://user:userpass@0.0.0.0:8081/userdb?sslmode=disable' down` 

3. How to fix database 
`migrate -path user/migrations -database 'postgres://user:userpass@0.0.0.0:8081/userdb?sslmode=disable' force 1 `

4. add migr files
`migrate create -ext sql -dir migrations -seq add_permissions `
