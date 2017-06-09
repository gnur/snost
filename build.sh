CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .
docker build -t gnur/snost .
docker run -p 8123:80 gnur/snost
