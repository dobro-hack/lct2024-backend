build:
	GOOS=linux GOARCH=amd64 go build -o ./dist/api ./cmd/api

test:
	go test -cover -race -v -timeout 30s ./internal/app/...
	go fmt ./internal/app/...		
	go vet -composites=false ./internal/app/...

run: 
	go fmt ./internal/app/...
	go fmt ./cmd/...
	go vet -composites=false ./internal/app/...
	go vet -composites=false ./cmd/...
	go test -cover -race -v -timeout 30s ./internal/app/...
	go run ./cmd/api

migrate-create:
	migrate create -ext sql -dir migrations/ -seq $(name)

swagger:
	go get github.com/swaggo/swag && ~/bin/swag init -g cmd/api/main.go --exclude dist/

scp:
	/usr/bin/scp ./dist/api 91:/home/bitrix/www/hack2024/

kill:
	ssh -f -n 91 'screen -S lct2024 -X quit && pkill api && rm -f ./api'

exec:
	$(eval REVISION = $(shell git log -1 --pretty=format:"%H"))
	ssh -f 91 'export VERSION="${git_tag}" && export REVISION="$(REVISION)" && export APP_TIER="prod" && cd /home/bitrix/www/hack2024 && screen -dmS lct2024 ./api /dev/null 2>&1 &'
