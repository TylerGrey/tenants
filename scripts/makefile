schema:
	go-bindata -ignore=\.go -pkg=schema -o=../internal/graphql/bindata.go ../api/schema/...

gqlgen:
	cd ../ && gqlgen generate