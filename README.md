# Public API

## Description
This is a simple public API that provides basic information including an email, current datetime in ISO 8601 format (UTC), and a GitHub link. The API is built using Golang and deployed on Railway.

## Features
- Accepts `GET` requests.
- Returns a JSON response with dynamically generated timestamp.
- Provides appropriate HTTP status codes.
- CORS enabled for cross-origin requests.
- Fast response time (<500ms).

## Endpoint
### `GET /api/info`

#### Response Format:
```json
{
  "email": "chiamaevans10@gmail.com",
  "datetime": "2025-01-30T12:00:00Z",
  "github": "https://github.com/evans-sudo/PublicApi"
}
```

#### Example Usage:
```sh
curl -X GET https://publicapi-production-904c.up.railway.app/api/info
```

## Running the Project Locally

### Prerequisites:
- Go (Golang) installed

### Steps:
1. Clone the repository:
   ```sh
   git clone https://github.com/evans-sudo/PublicApi.git
   ```
2. Navigate to the project directory:
   ```sh
   cd PublicApi
   ```
3. Run the application:
   ```sh
   go run main.go
   ```
4. Open your browser or use `curl` to test:
   ```sh
   curl -X GET https://github.com/evans-sudo/PublicApi.git
   ```

## Deployment
This API is deployed on Railway and can be accessed at:
**[https://publicapi-production-904c.up.railway.app/api/info](https://publicapi-production-904c.up.railway.app/api/info)**

## Hiring Links
Looking to hire a backend developer? Check out these links:

- [Hire Golang Developers](https://hng.tech/hire/golang-developers)

## License
This project is open-source and available under the MIT License.

