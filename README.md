# Go TODO + OTP

Go Todo application with user registration with JWT and SMS OTP.

## Teach stack

- Go + Gin
- Templ
- HTMX
- MySQL
- more go tools...

## Prerequisites

- Go programming language installed on your system
- MySQL or another compatible database installed and running
- `.env` file containing environment variables (e.g., database connection details, server port)

## Installation

1. Clone this repository to your local machine:

```bash
git clone https://github.com/Kei-K23/go-otp.git
```

2. Navigate to the project directory:

```bash
cd go-otp

```

3. Create a .env file in the project root and add the following environment variables:

```bash
JWT_SECRET_KEY=<YOUR_JWT_SECRET_KEY>
DB_CONNECTION=<YOUR_DB_CONNECTION>
SENDER=<YOUR_TWILIO_SENDER_NUMBER>
TWILIO_AUTH_TOKEN=<YOUR_TWILIO_AUTH_TOKEN>
TWILIO_ACCOUNT_SID=<YOUR_TWILIO_ACCOUNT_SID>
```

## Usage

1. Install dependencies:

```bash
go mod tidy
```

1. Run migration

```bash
make migration
```

2. Push database table

```bash
make migrate-up
```

2. Run server

```bash
make run
```

This will serve the server at `http://localhost:8080`

1. Access the API endpoints using tools like cURL or Postman.

## API endpoints

All endpoints are available under `http://localhost:4000/api/v1`. Make sure prefix with your localhost with `/api/v1`.

Authentication

- `POST /register`: Register page
- `POST /login`: Login page
- `POST /verify`: Verify page

### Below defined endpoints are protected with JWT authentication. Make sure valid JWT token exist in bearer authentication header

Users

- `GET /users`: User page
