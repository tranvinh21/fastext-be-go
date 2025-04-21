run: 
	cd cmd && go run main.go
build:
	go build -o app cmd/main.go

clean:
	rm -f app
air:
	air
