# Phone Validator
a single page application in Go (Frameworks allowed) that uses the provided database (SQLite 3) to list and categorize country phone numbers.

# API
``GET /phones?pageNumber=1&pageSize=10`` List all phones with (optional) paging enabled.

``GET /phones/:countryName?pageNumber=1&pageSize=10`` List all phones filtered by country name with (optional) paging enabled.

# Make jobs
- build: runs build for Linux OS.
- build-osx: runs build for macOSX.
- run: for local running.
- test: runs unit tests with coverage.
- lint: runs lint checking.
- docker-build: build base docker image
- docker-run: runs base docker image.
# Run instructions:
- Docker build `$ make docker-build`
- Docker run `$ make docker-run`