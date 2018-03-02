export $(cat .env | grep -v ^# | xargs)
go build
./pwitch_plays_tokemon
