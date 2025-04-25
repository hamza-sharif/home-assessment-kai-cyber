# home-assessment-kai-cyber
A Go-based microservice that scans repositories with json files and save data in SQLite using GORM. This app supports native execution as well as Docker-based deployment.

---

## 📦 Requirements

### For Native (Go) Setup
- Go 1.20+
- SQLite3
- go build cmd/main.go
- go test ./...
- go run cmd/main.go

### For Docker Setup
- Docker
- Docker Compose
- docker-compose up -d  # to run the sevice in background
- docker-compose down   # to stop the running server

---

## 🚀 Run with Makefile (Native Go)

You can use the `Makefile` to automate tests, prepare builds, and run the app locally Docker.

### Run Everything:

- make prepare  # to prepare the images for build
- make build    # to run the build command
- make run      # to run build, test and run the server.
-  
