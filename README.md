# Hwsi

**Hwsi** is a file indexer for HTTP web servers with focus on your files.

After specifying the working directory in `config.yaml`, you can browse all the folders and files in the directory through the web. You can also upload files to this folder via the web

## Build and Run

```bash
git clone https://github.com/AkvicorEdwards/hwsi.git
cd hwsi
go mod download
go build hwsi.go
./hwsi

# ListenAndServe: localhost
```

## Docker

```bash
git clone https://github.com/AkvicorEdwards/hwsi.git
cd hwsi
mkdir work config theme
docker-compose up -d

# ListenAndServe: localhost:8080
```
