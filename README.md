
gox --osarch "linux/amd64 darwin/386" ./...

docker build -t multichat .

docker run -p 8000:5000 multichat

docker exec -i -t multichat ash
docker exec -i -t 01c4b21de79a ash
