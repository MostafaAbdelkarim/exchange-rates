# Project Title

`Exchange Rates API`

![Version](https://img.shields.io/badge/version-v1.0.0-blue)
![Contributions Welcome](https://img.shields.io/badge/contributions-Welcome-orange.svg)
[![License](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)

## Basic Overview

Backend API based on Golang - Fiber managing Cryptocurrency and fiat Exchange Rates.

## Features

- **All Endpoints** for exchange rates API.
- **GET Endpoint** for Ethereum Balance API.
- **Global Exception Handler** with scalability for custom-made exceptions.
- **Unified Success/Failure response** following best practice.
- **Validation** for Custom DTOs.

## Dependencies

You can find all dependencies in `go.mod` file

## How to use

- Clone the repo.
  > Note: Configure datasource properties to your MongoDB URL by modifying .env file.
- Run using VSCode/Goland by executing `go mod tidy` -> `go build main.go` -> `go run main.go`.
- You can start experimenting using these links for both APIs:

```text
Rates: http://localhost:3000/api/v1/rates
Balance: http://localhost:3000/api/v1/balance
```

### Enhancement Features

- Add **Security** for Authentication/Authorization
- Add **Soft Delete** for each Model
- **Write logging** in certain file for each day
- **Containerize** the app using **Docker**

### Bug Reports & Feature Requests

> Please use the [issue tracker](https://github.com/MostafaAbdelkarim/exchange-rates/issues) to report any bugs or file feature requests.