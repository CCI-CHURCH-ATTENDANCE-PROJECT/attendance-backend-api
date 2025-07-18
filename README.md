# Church Attendance Management System API

A comprehensive REST API built with Go (Golang) and Echo framework for managing church attendance, members, roles, family relationships, sermons, and announcements.

## Features

- 🔐 **Authentication & Authorization**: JWT-based auth with role-based access control
- 👥 **User Management**: Complete user profiles with search and filtering
- 📅 **Attendance Tracking**: Manual and QR code-based check-in system
- 📱 **QR Code Generation**: Dynamic QR codes for quick attendance
- 👨‍👩‍👧‍👦 **Family Management**: Track family relationships and members
- 🎤 **Sermon Management**: Record and manage church sermons
- 📢 **Announcements**: Create and manage church announcements
- 📊 **Analytics & Reporting**: Comprehensive attendance analytics
- 🔒 **Security**: Industry-standard security practices with rate limiting
- 🌍 **Timezone Support**: UTC+1 (Lagos timezone) support

## Technology Stack

- **Backend**: Go 1.21+ with Echo v4 framework
- **Database**: MongoDB with mongo-go-driver
- **Authentication**: JWT (JSON Web Tokens)
- **Validation**: go-playground/validator
- **QR Codes**: skip2/go-qrcode
- **Security**: bcrypt password hashing, CORS, security headers

## Quick Start

### Prerequisites

- Go 1.21 or higher
- MongoDB 4.4 or higher
- Git

### Installation

1. **Clone the repository**
   ```bash
   git clone <repository-url>
   cd church-attendance-api
   ```

2. **Install dependencies**
   ```bash
   go mod tidy
   ```

3. **Set up environment variables**
   ```bash
   cp .env.example .env
   # Edit .env with your configuration
   ```

4. **Start MongoDB**
   ```bash
   # If using Docker
   docker run -d -p 27017:27017 --name mongodb mongo:latest
   
   # Or start your local MongoDB instance
   mongod
   ```

5. **Run the application**
   ```bash
   go run main.go
   ```

The API will be available at `http://localhost:8080`

### Health Check

```bash
curl http://localhost:8080/health
```

## API Documentation

### Base URL
```
http://localhost:8080/api/v1
```

### Authentication Endpoints

#### Basic User Registration
```http
POST /api/v1/auth/register
Content-Type: application/json

{
  "email": "user@example.com",
  "password": "securePassword123"
}
```

#### Complete User Registration
```http
POST /api/v1/auth/register/complete
Content-Type: application/json

{
  "email": "user@example.com",
  "password": "securePassword123",
  "fname": "John",
  "lname": "Doe",
  "bio": "Church member",
  "date_of_birth": "1990-05-15",
  "gender": "Male",
  "member": true,
  "visitor": false,
  "phone_number": "+2348123456789",
  "profession": "Software Engineer",
  "user_house_address": "123 Church Street, Lagos",
  "campus_state": "Lagos",
  "campus_country": "Nigeria",
  "emergency_contact_name": "Jane Doe",
  "emergency_contact_phone": "+2348987654321",
  "emergency_contact_email": "jane@example.com",
  "emergency_contact_relationship": "Spouse"
}
```

#### User Login
```http
POST /api/v1/auth/login
Content-Type: application/json

{
  "email": "user@example.com",
  "password": "securePassword123"
}
```

#### Token Refresh
```http
POST /api/v1/auth/refresh
Content-Type: application/json

{
  "refresh_token": "your-refresh-token"
}
```

#### Logout
```http
POST /api/v1/auth/logout
Authorization: Bearer <access-token>
```

### User Management Endpoints

#### Search Users
```http
GET /api/v1/users/search?q=john&page=1&limit=10
Authorization: Bearer <access-token>
```

#### Get All Users (Paginated)
```http
GET /api/v1/users?page=1&limit=10
Authorization: Bearer <access-token>
```

#### Filter Users
```http
GET /api/v1/users/filter?field=member&value=true&page=1&limit=10
Authorization: Bearer <access-token>
```

### Attendance Endpoints

#### Create Attendance Record
```http
POST /api/v1/attendance
Authorization: Bearer <access-token>
Content-Type: application/json

{
  "user_id": "CCIMRB-12345"
}
```

#### QR Code Check-in
```http
POST /api/v1/attendance/qr-checkin
Content-Type: application/json

{
  "qr_code_token": "encoded_user_id_token"
}
```

#### Get Attendance History
```http
GET /api/v1/attendance/history?start_date=2025-01-01&end_date=2025-01-31&page=1&limit=10
Authorization: Bearer <access-token>
```

#### Get Attendance Analytics
```http
GET /api/v1/attendance/analytics?date=2025-01-15
Authorization: Bearer <access-token>
```

### QR Code Endpoints

#### Generate QR Code
```http
POST /api/v1/qr/generate
Authorization: Bearer <access-token>
Content-Type: application/json

{
  "user_id": "CCIMRB-12345"
}
```

## Environment Variables

| Variable | Description | Default |
|----------|-------------|---------|
| `DB_HOST` | MongoDB host | `localhost` |
| `DB_PORT` | MongoDB port | `27017` |
| `DB_NAME` | Database name | `church_attendance_db` |
| `DB_USER` | MongoDB username | `` |
| `DB_PASSWORD` | MongoDB password | `` |
| `JWT_SECRET` | JWT signing secret | `the-super-secret-jwt-key-to-be--changed-in-production` |
| `JWT_ACCESS_EXPIRY` | Access token expiry | `15m` |
| `JWT_REFRESH_EXPIRY` | Refresh token expiry | `168h` |
| `PORT` | Server port | `8080` |
| `ENV` | Environment mode | `development` |
| `CORS_ORIGINS` | CORS allowed origins | `http://localhost:3000,http://localhost:8080` |
| `TIMEZONE` | Application timezone | `Africa/Lagos` |

## Database Schema

### Collections

- `users` - User profiles and authentication data
- `attendance` - Attendance records
- `refresh_tokens` - JWT refresh tokens
- `family_members` - Family relationship data
- `sermons` - Sermon information
- `announcements` - Church announcements
- `roles` - User roles and permissions
- `church_info` - Local church information
- `notifications` - System notifications

## Security Features

- 🔐 **JWT Authentication**: Secure access and refresh tokens
- 🛡️ **Password Hashing**: bcrypt for secure password storage
- 🚦 **Rate Limiting**: Protection against abuse
- 🔒 **CORS**: Configurable cross-origin resource sharing
- 🛡️ **Security Headers**: XSS, CSRF, and other security headers
- ⏱️ **Request Timeout**: Prevent hanging requests
- 🔍 **Input Validation**: Comprehensive request validation

## Error Handling

The API uses consistent error response format:

```json
{
  "success": false,
  "error": {
    "code": "ERROR_CODE",
    "message": "Error description",
    "details": [
      {
        "field": "field_name",
        "message": "Field-specific error message"
      }
    ]
  }
}
```

### HTTP Status Codes

- `200` - Success
- `201` - Created
- `400` - Bad Request
- `401` - Unauthorized
- `403` - Forbidden
- `404` - Not Found
- `422` - Validation Error
- `429` - Rate Limit Exceeded
- `500` - Internal Server Error

## Development

### Project Structure

```
church-attendance-api/
├── main.go                     # Application entry point
├── internal/
│   ├── config/                 # Configuration management
│   ├── database/               # Database connection and setup
│   ├── dto/                    # Data Transfer Objects
│   ├── handler/                # HTTP handlers (controllers)
│   ├── middleware/             # Custom middleware
│   ├── models/                 # Database models
│   ├── repository/             # Data access layer
│   ├── service/                # Business logic
│   └── utils/                  # Utility functions
├── .env                        # Environment variables
├── .env.example               # Environment template
├── go.mod                     # Go modules
├── go.sum                     # Go modules checksum
└── README.md                  # This file
```

### Running Tests

```bash
# Run all tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Run tests with verbose output
go test -v ./...
```

### Building for Production

```bash
# Build binary
go build -o church-attendance-api main.go

# Run binary
./church-attendance-api
```

### Docker Support

```dockerfile
# Dockerfile
FROM golang:1.21-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o main .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/

COPY --from=builder /app/main .
COPY --from=builder /app/.env .

CMD ["./main"]
```

## Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Support

For support, email [support@example.com] or join our Slack channel.

## Acknowledgments

- Echo framework for the excellent HTTP router
- MongoDB for the robust database solution
- The Go community for excellent libraries and tools
