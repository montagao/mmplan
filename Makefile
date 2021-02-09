PROJECT_NAME=mmplan

run:
	go run ./cmd/${PROJECT_NAME}-server/main.go
api:
	swagger generate server -A ${PROJECT_NAME} -f ./api.yml --exclude-main
install:
	go install ./cmd/${PROJECT_NAME}-server/main.go 
clean:
	rm -rf cmd/* models/* restapi/*
