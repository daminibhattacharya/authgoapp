# Go Backend API

## Overview

This backend API provides user registration functionality with Auth0 authentication integration and sqlite data storage. It's built using Go and the Gin web framework.

## Features

- User registration with first name, last name, email and password
- Input validation and error handling
- Integration with Auth0 for secure authentication
- Data persistence in sqlite
- Prevention of duplicate user registrations

## Prerequisites

- Go
- sqlite
- Auth0 account and API credentials

## Installation

1. Clone the repository.
2. Copy `.env.default` to a new `.env` file and add values for Auth0 URL, etc. (or use the `.env` file provided in the email).
3. Run the command `go run main.go` in the project directory.
4. Open the `http://localhost:8080` file in your web browser.
5. Register yourself as a user.