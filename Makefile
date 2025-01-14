build:
	docker build --no-cache --progress=plain -t simple-api:1.0 -f ./Dockerfile .
run:
	docker run -d --rm -p 8585:8585 --name simple-api-cont simple-api:1.0