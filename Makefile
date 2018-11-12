dev:
	docker-compose up
prd:
	docker build -t go-prd-app -f ./build/prd/Dockerfile .
run:
	docker run -p 8080:8080 -it go-prd-app
