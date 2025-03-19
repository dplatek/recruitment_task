
# Recruitment task for GoSolve
This project is a simple web service build in Go. It loads data from input file and provides an API endpoint to find index of closest matching value.

## Features
- API for http with `GET` method.
	- There is GET method implemented for endpoint named `endpoint`.
- Functionality for searching `index` for `given` value.
	- The code is assuming that numbers in `input.txt` file are sorted.
	- In case of sorted numbers, best algorithm to use is binary search that would lead to time complexity of O(logn).
- Logging with 3 levels (Info, Debug, Error).
- Add possibility to use configuration file where you can specify service port and log level.
	- There is `config.json` file in main folder.
- Unit tests for created components.
- File `README.md` to describing service.
- Automated running tests with `make` file.
	- There are only test and clean methods.
- Project has correct code structure.
- Solution is uploaded into `GitHub` account.

## Configuration
The service reads configuration from `config.json`:
```
{
  "port": "8080",
  "log_level": "Info"
}
```

## Running the Service
Start the service with:
```
go run main.go
```
It runs on `localhost:8080` by default.

## API Usage
#### Example Request
```
GET /endpoint/21
```

#### Example Response (Exact Match)
```
{
  "error": "",
  "index": 1
}
```

#### Example Response (Closest Match)
```
{
  "error": "Value 21 not found, but closest match 20 found at index 2",
  "index": 2
}
```

#### Example Response (Not Found)
```
{
  "error": "Value 200 not found",
  "index": -1
}
```

## Running Tests
Run tests using:
```
make test
```
or manually:
```
go test ./...
```