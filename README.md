# URL Shortener

Simple URL shortening service written in Go using Fiber and MySQL.

## Project Overview

- Minimal web service to shorten URLs and redirect via a short code.
- Uses `github.com/gofiber/fiber/v3` for HTTP handling and `github.com/matoous/go-nanoid/v2` to generate short codes.
- Stores mappings in a MySQL database.

## Prerequisites

- Go 1.20+ installed
- MySQL server accessible (create a database for the app)

## Setup

1. Clone the repository and open the project directory.

2. Install dependencies and tidy modules:

```bash
cd urlShortener
go mod tidy
```

3. Database:

- Create a database named `urlshortener` (or change DSN in `main.go`).

Example SQL to create the table:

```sql
CREATE TABLE urls (
  id INT AUTO_INCREMENT PRIMARY KEY,
  original_url TEXT NOT NULL,
  short_code VARCHAR(255) NOT NULL UNIQUE,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

4. Configure DB connection:

- The current `main.go` uses a hard-coded DSN:

```
root:password@tcp(localhost:3306)/urlshortener
```

- Update that DSN as needed for your environment (or replace with an environment variable).

## Run

Start the server:

```bash
go run main.go
```

Server listens on port `3000` by default.

## API Endpoints

- `GET /` — health check

Response example:

```json
{"server":"urlShortener","status":"200ok","health":"Healthy"}
```

- `POST /shorten` — create a shortened URL

Request body (JSON):

```json
{ "url": "https://example.com" }
```

Success response (JSON):

```json
{ "url": "http://127.0.0.1:3000/<short_code>" }
```

Possible status codes:
- `400` — invalid request
- `500` — server error / DB error

- `GET /:shorten` — redirect to the original URL using the short code

Behavior:
- If `short_code` exists, the server responds with an HTTP redirect to the stored `original_url`.
- If not found, returns `404` with JSON error message.

## Database Schema

See the SQL snippet in the Setup section. The table stores `original_url`, `short_code`, and `created_at`.

## Notes & Improvements

- Move DB credentials into environment variables.
- Add input validation and URL normalization.
- Track click analytics (click count, last accessed).
- Add tests and CI.

## License

MIT
