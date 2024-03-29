build:
	docker build -f docker/Dockerfile . -t 34234247632/otus-ha-hw:v1.0

push:
	docker push 34234247632/otus-ha-hw:v1.0

docker-start:
	cd docker && docker-compose up -d

docker-stop:
	cd docker && docker-compose down

newman:
	newman run postman/collection.json

