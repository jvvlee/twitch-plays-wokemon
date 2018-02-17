export $(cat .env | grep -v ^# | xargs)
go run twitch-reader.go