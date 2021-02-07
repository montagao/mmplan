PROJECT_NAME=mmplan

api:
	swagger generate server -A ${PROJECT_NAME} -f ./api.yml --exclude-main
install:
	go install ./cmd/${PROJECT_NAME}-server/
clean:
	rm -rf cmd/* models/* restapi/*
