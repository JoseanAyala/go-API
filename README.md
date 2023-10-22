# Go API with Mux, Auth0, and SQL

## Description

This is a simple Go API that demonstrates how to build a web application using the Mux router for routing, Auth0 for authentication, and SQL for database operations. It provides a starting point for building secure and scalable APIs in Go.

## Prerequisites

Before you begin, make sure you have the following dependencies installed:

- Go: https://golang.org/dl/
- PostgreSQL or MySQL: Install and configure your SQL database.
- Auth0 Account: Sign up for an Auth0 account at https://auth0.com/.

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/joseanayala/go-api.git
   ```

2. Change to the project directory:

   ```bash
   cd go-api
   ```

3. Install Go dependencies:

   ```bash
   go mod download
   ```

4. Create a `.env` file in the project root and configure the following environment variables:

   ```plaintext
   DSN=postgres://username:password@localhost/database_name
   AUTH0_DOMAIN=your-auth0-domain
   AUTH0_AUDIENCE=your-auth0-audience
   ```

5. Start the API:

   ```bash
   go run main.go
   ```

Your API should now be running at `http://localhost:8080`.
