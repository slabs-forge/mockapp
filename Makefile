TAG=sebetln/mockapp:latest

build:
	docker build -t $(TAG) .

push:
	docker push $(TAG)
