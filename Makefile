PROJECT=goblaztodos

build:
	rm -rf dist/
	dotnet publish blazfront/blazfront.csproj -c release -o dist

go-build:build
	go build -o ./app ./goback

run:
	./app

docker-build: build
	docker build --rm -t dev-${PROJECT} .

docker-run:
	docker run --rm -p 3000:3000 --name dev-${PROJECT} dev-${PROJECT}