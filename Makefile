TEST_FILES=$(shell find -name '*_test.go')
build: task
	go version
task:
	go test ./post_service
	go test ./post_app_service
