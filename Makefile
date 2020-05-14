build:
	go build -a -o gin-user-center .

docker:
	docker build -t gin-user-center:latest .