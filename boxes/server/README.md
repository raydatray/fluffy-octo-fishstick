# University Management System API

A REST API built with Go's standard `net/http` library and Ent ORM for managing university courses, users, and discussion boards.

## Features

- User management (Professors, TAs, Students)
- Course and section management
- Discussion boards and posts
- SQLite database with automatic schema migration
- Standard HTTP status codes and JSON responses

## Getting Started

### Prerequisites

- Go 1.24.6 or higher
- SQLite3

### Installation

1. Clone the repository
2. Navigate to the server directory:
   ```bash
   cd boxes/server
   ```
3. Install dependencies:
   ```bash
   go mod tidy
   ```
4. Build the application:
   ```bash
   go build
   ```
5. Run the server:
   ```bash
   ./server
   ```

The server will start on port 8080 by default. You can override this by setting the `PORT` environment variable.

## API Endpoints

### Health Check

#### GET /health
Returns the health status of the API.

**Response:**
```json
{
  "status": "healthy"
}
```

### User Management

#### POST /api/users
Create a new user.

**Request Body:**
```json
{
  "email": "user@example.com",
  "password": "securepassword",
  "role": "STUDENT" // Optional: PROFESSOR, TA, or STUDENT (default)
}
```

**Response (201 Created):**
```json
{
  "id": 1,
  "email": "user@example.com",
  "role": "STUDENT"
}
```

**Error Responses:**
- `400 Bad Request`: Invalid JSON, missing required fields, or invalid role
- `409 Conflict`: Email already exists

#### GET /api/users
Get all users.

**Response (200 OK):**
```json
[
  {
    "id": 1,
    "email": "user@example.com",
    "role": "STUDENT"
  },
  {
    "id": 2,
    "email": "professor@example.com",
    "role": "PROFESSOR"
  }
]
```

#### GET /api/users/{id}
Get a specific user by ID.

**Response (200 OK):**
```json
{
  "id": 1,
  "email": "user@example.com",
  "role": "STUDENT"
}
```

**Error Responses:**
- `400 Bad Request`: Invalid user ID
- `404 Not Found`: User not found

#### PUT /api/users/{id}
Update a user by ID.

**Request Body (all fields optional):**
```json
{
  "email": "newemail@example.com",
  "password": "newpassword",
  "role": "TA"
}
```

**Response (200 OK):**
```json
{
  "id": 1,
  "email": "newemail@example.com",
  "role": "TA"
}
```

**Error Responses:**
- `400 Bad Request`: Invalid JSON, invalid user ID, or invalid role
- `404 Not Found`: User not found
- `409 Conflict`: Email already exists

#### DELETE /api/users/{id}
Delete a user by ID.

**Response (204 No Content):** Empty response body

**Error Responses:**
- `400 Bad Request`: Invalid user ID
- `404 Not Found`: User not found

## User Roles

The system supports three user roles:

- **PROFESSOR**: Can teach course sections
- **TA**: Can assist in course sections
- **STUDENT**: Can enroll in course sections (default role)

## Database Schema

The system uses the following entities:

- **User**: Email, password, role
- **Course**: Code (unique)
- **CourseSection**: Number, associated with a course
- **DiscussionBoard**: Associated with a course
- **Post**: Title, content, author, discussion board
- **Reply**: Content, author, post

### Relationships

- Course ↔ DiscussionBoard (one-to-one)
- Course → CourseSection (one-to-many)
- User → Post (one-to-many, as author)
- User → Reply (one-to-many, as author)
- Post → Reply (one-to-many)
- User ↔ CourseSection (many-to-many, as professors, TAs, or students)

## Testing

A test script is provided to test all user endpoints:

```bash
./test_api.sh
```

Make sure the server is running before executing the test script.

## Development

### Adding New Endpoints

1. Create new handler functions in the `handlers` package
2. Register the routes in `main.go`
3. Update this README with the new endpoints

### Database Changes

If you modify the Ent schema files in `ent/schema/`, regenerate the code:

```bash
go generate ./ent
```

## Error Handling

The API returns appropriate HTTP status codes:

- `200 OK`: Successful GET/PUT requests
- `201 Created`: Successful POST requests
- `204 No Content`: Successful DELETE requests
- `400 Bad Request`: Invalid request data
- `404 Not Found`: Resource not found
- `405 Method Not Allowed`: Unsupported HTTP method
- `409 Conflict`: Resource conflict (e.g., duplicate email)
- `500 Internal Server Error`: Server-side errors

All error responses include a descriptive error message in the response body.

## Security Considerations

⚠️ **This is a development version. For production use, consider implementing:**

- Password hashing (currently passwords are stored in plain text)
- Authentication and authorization middleware
- Input validation and sanitization
- Rate limiting
- HTTPS/TLS encryption
- Database connection pooling
- Logging and monitoring