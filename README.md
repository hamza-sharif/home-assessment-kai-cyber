# home-assessment-kai-cyber
A Go-based microservice that scans repositories with json files and save data in SQLite using GORM. This app supports native execution as well as Docker-based deployment.

---

## Requirements For Native (Go) Setup
- Go 1.20+
- SQLite3
  
### commands to run
  `go build cmd/main.go` <br>
  `go test ./...` <br>
  `go run cmd/main.go` <br>

## Requirements For Docker Setup
- Docker 
- Docker Compose

### commands to run
  `docker-compose up -d`    #   to run the sevice in background  <br>
  `docker-compose down`   # to stop the running server<br>

---

## Run with Makefile (Native Go)

You can use the `Makefile` to automate tests, prepare builds, and run the app locally Docker.

### Run Everything:

`make prepare`  # to prepare the images for build <br>
`make build`    # to run the build command <br>
`make run`      # to run build, test, and run the server. <br>

## Live Api Testing
Download Postman and import `Home Assessment Kai Cyber.postman_collection.json` and test APIs with live data.

