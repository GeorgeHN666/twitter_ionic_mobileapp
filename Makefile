SHELL=cmd

## build: builds all binaries
build: clean build_api
	@echo All binaries built!

## clean: cleans all binaries and runs go clean
clean:
	@echo Cleaning...
	@echo y | DEL /S dist
	@go clean
	@echo Cleaned and deleted binaries

## build_back: builds the back end
build_api:
	@echo Building backend...
	@go build -o dist/go-server.exe ./src/api/main.go
	@echo Backend built!

## start: starts front and back end
start: start_api


## start_back: starts the back end -dsn=${DSN}
start_api: build_api
	@echo Starting the back end...
	start /B .\dist\go-server.exe 
	@echo Backend running in the background!

## stop: stops the front and back end
stop: stop_api
	@echo All applications stopped


## stop_back: stops the back end
stop_api:
	@echo Stopping the backend...
	@taskkill /IM go-server.exe /F
	@echo Stopped backend