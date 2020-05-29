export $(cat ../configs/.env)
export GIN_MODE=debug
go run ../cmd/tenants/main.go