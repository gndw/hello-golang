build-app:
	@echo "Building App..."
	@go build main.go

run-app-win: build-app
	@echo "Running App..."
	@main.exe

run-app: build-app
	@echo "Running App..."
	@./main