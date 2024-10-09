run:
	go run ./cmd/tg-bot/main.go --config="./config/prod.yaml"

docker-build:
	docker build -t tg-bot:latest .