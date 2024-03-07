# One-Two-One-Two

This is a simple mock API built for development. It will mirror the JSON data sent in the request body. You can also send the status code you'd like to receive back as a URL query param so you can test successful and failed responses of your app.

For learning purposes, I have also written an expansion of this which connects to MongoDB, using hexagonal architecture. [Here is a link](https://github.com/josenymad/hexagonal-architecture-api)

## Installation

- Fork and clone this repo
- Navigate to the repo folder on your machine

## Usage

- Run the mock API with `go run .`
- By default, it runs on `http://localhost:8000`, you can change this in the main file should you need to
- The API endpoint would then be:

`http://localhost:8000/test?status=<status code you'd like to receive back>`
