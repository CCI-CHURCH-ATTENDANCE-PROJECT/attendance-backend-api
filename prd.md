# Church Attendance Management System API - Product Requirements Document

## 1. Overview

### 1.1 Purpose
The Church Attendance Management System API is designed to manage church member attendance, roles, family relationships, sermons, 
announcements, and notifications. The system provides comprehensive tracking of church activities and member engagement.

### 1.2 Technology Stack
- **Backend Framework**: Go (Golang) and Echo framework
- **Database**: MongoDB (use mongo-go-driver)
- **Authentication**: JWT (JSON Web Tokens)
- **Time Zone**: UTC+1
- **Data Validation**: Standard Go validation libraries
- **Security**: Industry-standard security practices
- **API Interface**: Swagger documentation

### 1.3 Key Features
- User management with role-based access control
- QR code-based attendance tracking
- Family member management
- Sermon and announcement management
- Comprehensive reporting and analytics
- Secure authentication and authorization

## 2. Data Models

### 2.1 User Model
```json
{
  "id": "uint (auto-generated)",
  "user_id": "string (CCIMRB-prefixed)",
  "fname": "string",
  "lname": "string",
  "email": "string (unique)",
  "user_password": "string (hashed)",
  "bio": "string",
  "date_of_birth": "date",
  "gender": "string",
  "member": "boolean",
  "visitor": "boolean",
  "usher": "boolean",
  "admin": "boolean",
  "user_work_department": "ObjectId (foreign key)",
  "date_joined_church": "date",
  "qr_code_token": "string",
  "family_head": "boolean",
  "user_campus": "string",
  "campus_state": "string",
  "campus_country": "string",
  "profession": "string",
  "user_house_address": "string",
  "phone_number": "string",
  "instagram_handle": "string",
  "family_member_id": "uint",
  "date_joined": "datetime (UTC+1)",
  "date_updated": "datetime (UTC+1)",
  "role": "ObjectId (foreign key)",
  "emergency_contact_name": "string",
  "emergency_contact_phone": "string",
  "emergency_contact_email": "string",
  "emergency_contact_relationship": "string"
}
```

### 2.2 Attendance Model
```json
{
  "id": "uint (auto-generated)",
  "user": "ObjectId (foreign key - User)",
  "date_time_of_attendance": "datetime (UTC+1)",
  "qrcode_based_checkin": "boolean",
  "late": "boolean",
  "manual_checkin": "boolean",
  "sermon_topic": "ObjectId (foreign key - Sermon)",
  "announcements": "ObjectId (foreign key - Announcements)"
}
```

### 2.3 Family Member Model
```json
{
  "id": "uint (auto-generated)",
  "family_head": "ObjectId (foreign key - User)",
  "family_members": "ObjectId (foreign key - User)",
  "date_joined": "datetime (UTC+1)"
}
```

### 2.4 Sermon Model
```json
{
  "id": "uint (auto-generated)",
  "preacher": "string",
  "date_of_meeting": "datetime (UTC+1)",
  "sermon_topic": "string",
  "sermon_summary": "text",
  "entry_made_by": "ObjectId (foreign key - User)"
}
```

### 2.5 Announcements Model
```json
{
  "id": "uint (auto-generated)",
  "title": "string",
  "announcement_content": "text",
  "announcement_date": "datetime (UTC+1)",
  "announcement_due_date": "datetime (UTC+1)",
  "announcement_entry_made_by": "ObjectId (foreign key - User)",
  "date_added": "datetime (UTC+1)",
  "status": "string (choices: 'Pending', 'Done')"
}
```

### 2.6 Departments Model
```json
{
  "id": "uint (auto-generated)",
  "name": "string",
  "total_members": "int",
  "active_rate": "int",
  "monthly_attendance": "int",
  "average_attendance": "int",
  "performance": "string (choices: 'GOOD', 'POOR', 'EXCELLENT', 'NEEDS_ATTENTION', 'UNDECIDED')"
}
```

### 2.7 Role Model
```json
{
  "id": "uint (auto-generated)",
  "role_name": "string",
  "role_description": "string",
  "total_members": "int",
  "permissions": "string (choices: 'USHER', 'ADMIN', 'YOUTH LEADER', 'MEMBER', 'PASTOR')",
  "date_added": "datetime (UTC+1)",
  "date_updated": "datetime (UTC+1)"
}
```

### 2.8 Local Church Model
```json
{
  "id": "uint (auto-generated)",
  "church_name": "string",
  "church_phone": "string",
  "church_email": "string",
  "church_address": "string",
  "state_county": "string",
  "country": "strig",
  "church_sunday_meeting_time": "int",
  "midweek_meeting_day": "string",
  "midweek_meeting_time": "int"
}
```

### 2.9 Notifications Model
```json
{
  "id": "uint (auto-generated)",
  "title": "string",
  "user_created": "ObjectId (foreign key - USer)",
  "body": "string",
  "email": "boolean",
  "sms": "boolean",
  "event": "boolean",
  "newsletter": "boolean",
  "date_created": "datetime (UTC+1)"
}
```

### 2.10 User Notifications Model
```json
{
  "id": "uint (auto-generated)",
  "user": "ObjectId (foreign key - User)",
  "title": "string",
  "body": "string",
  "email": "boolean",
  "sms": "boolean",
  "event": "boolean",
  "newsletter": "boolean",
  "date_delivered": "datetime (UTC+1)"
}
```

## 3. API Endpoints

### 3.1 Authentication Endpoints

#### 3.1.1 Basic User Registration
**POST** `/api/v1/auth/register`

**Request Body:**
```json
{
  "email": "user@example.com",
  "password": "securePassword123"
}
```

**Response:**
```json
{
  "success": true,
  "message": "User registered successfully",
  "data": {
    "user_id": "CCIMRB-12345",
    "email": "user@example.com",
    "created_at": "2025-07-07T10:00:00+01:00"
  }
}
```

#### 3.1.2 User Login
**POST** `/api/v1/auth/login`

**Request Body:**
```json
{
  "email": "user@example.com",
  "password": "securePassword123"
}
```

**Response:**
```json
{
  "success": true,
  "message": "Login successful",
  "data": {
    "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
    "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
    "user": {
      "user_id": "CCIMRB-12345",
      "fname": "John",
      "lname": "Doe",
      "email": "user@example.com"
    }
  }
}
```

#### 3.1.3 Complete User Registration
**POST** `/api/v1/auth/register/complete`

**Request Body:**
```json
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
  "emergency_contact_name": "Sola Kosoko",
  "emergency_contact_phone": "+2348987654321",
  "emergency_contact_email": "jane@example.com",
  "emergency_contact_relationship": "Spouse"
}
```

#### 3.1.4 Token Refresh
**POST** `/api/v1/auth/refresh`

**Request Body:**
```json
{
  "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
}
```

**Response:**
```json
{
  "success": true,
  "data": {
    "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
    "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
  }
}
```

### 3.2 User Management Endpoints

#### 3.2.1 Search Users
**GET** `/api/v1/users/search`

**Query Parameters:**
- `q` (required): Search query (name, email, or user_id)
- `page` (optional): Page number (default: 1)
- `limit` (optional): Items per page (default: 10)

**Response:**
```json
{
  "success": true,
  "data": {
    "users": [
      {
        "user_id": "CCIMRB-12345",
        "fname": "John",
        "lname": "Doe",
        "email": "john@example.com"
      }
    ],
    "pagination": {
      "page": 1,
      "limit": 10,
      "total": 1,
      "total_pages": 1
    }
  }
}
```

#### 3.2.2 Get All Users (Paginated)
**GET** `/api/v1/users`

**Query Parameters:**
- `page` (optional): Page number (default: 1)
- `limit` (optional): Items per page (default: 10)

#### 3.2.3 Filter Users
**GET** `/api/v1/users/filter`

**Query Parameters:**
- `field` (required): Field to filter by
- `value` (required): Value to filter
- `page` (optional): Page number
- `limit` (optional): Items per page

**Example:** `/api/v1/users/filter?field=member&value=true&page=1&limit=10`

### 3.3 Attendance Endpoints

#### 3.3.1 Create Attendance Record
**POST** `/api/v1/attendance`

**Headers:**
- `Authorization: Bearer <token>`

**Request Body:**
```json
{
  "user_id": "CCI-12345"
}
```

**Response:**
```json
{
  "success": true,
  "message": "Attendance recorded successfully",
  "data": {
    "id": "attendance_id",
    "user_id": "CCI-12345",
    "date_time_of_attendance": "2025-07-07T10:00:00+01:00",
    "qrcode_based_checkin": false,
    "late": false,
    "manual_checkin": true
  }
}
```

#### 3.3.2 QR Code Check-in
**POST** `/api/v1/attendance/qr-checkin`

**Request Body:**
```json
{
  "qr_code_token": "encoded_user_id_token"
}
```

#### 3.3.3 Attendance History
**GET** `/api/v1/attendance/history`

**Query Parameters:**
- `start_date` (optional): Start date (YYYY-MM-DD)
- `end_date` (optional): End date (YYYY-MM-DD)
- `page` (optional): Page number
- `limit` (optional): Items per page

**Response:**
```json
{
  "success": true,
  "data": {
    "history": [
      {
        "date": "2025-07-07",
        "members": 45,
        "visitors": 12,
        "total_attendance": 57
      }
    ],
    "pagination": {
      "page": 1,
      "limit": 10,
      "total": 1,
      "total_pages": 1
    }
  }
}
```

#### 3.3.4 Attendance Analytics
**GET** `/api/v1/attendance/analytics`

**Headers:**
- `Authorization: Bearer <token>`

**Query Parameters:**
- `date` (required): Date for analytics (YYYY-MM-DD)

**Response:**
```json
{
  "success": true,
  "data": {
    "total_active_users_all_time": 150,
    "total_attendance_for_date": 57,
    "members_for_month": 45,
    "visitors_count": 12
  }
}
```

### 3.4 QR Code Endpoints

#### 3.4.1 Generate QR Code
**POST** `/api/v1/qr/generate`

**Headers:**
- `Authorization: Bearer <token>`

**Request Body:**
```json
{
  "user_id": "CCI-12345"
}
```

**Response:**
```json
{
  "success": true,
  "data": {
    "qr_code_token": "encoded_token",
    "qr_code_image": "base64_encoded_image"
  }
}
```

### 3.5 Role Management Endpoints

#### 3.5.1 Create Role
**POST** `/api/v1/roles`

**Headers:**
- `Authorization: Bearer <token>`

**Request Body:**
```json
{
  "role_name": "Youth Leader",
  "role_description": "Leads youth activities",
  "permissions": "view_users,create_events,manage_youth"
}
```

#### 3.5.2 Get Roles and Permissions
**GET** `/api/v1/roles`

**Headers:**
- `Authorization: Bearer <token>`

**Response:**
```json
{
  "success": true,
  "data": {
    "roles": [
      {
        "id": "role_id",
        "role_name": "Youth Leader",
        "role_description": "Leads youth activities",
        "permissions": "view_users,create_events,manage_youth",
        "total_members": 5
      }
    ],
    "available_permissions": [
      "view_users",
      "create_events",
      "manage_youth",
      "admin_access"
    ]
  }
}
```

### 3.6 Family Management Endpoints

#### 3.6.1 Create Family Member
**POST** `/api/v1/family-members`

**Headers:**
- `Authorization: Bearer <token>`

**Request Body:**
```json
{
  "family_name": "Doe Family",
  "family_info": "Nuclear family of 4",
  "family_email": "family@example.com",
  "family_relationship": "Father",
  "family_phone_number": "+2348123456789"
}
```

#### 3.6.2 Get Family Members
**GET** `/api/v1/family-members`

**Headers:**
- `Authorization: Bearer <token>`

**Query Parameters:**
- `family_head_id` (optional): User ID of family head

**Response:**
```json
{
  "success": true,
  "data": {
    "family_members": [
      {
        "id": "member_id",
        "family_head": "CCI-12345",
        "family_members": "CCI-12346",
        "date_joined": "2025-07-07T10:00:00+01:00"
      }
    ]
  }
}
```

### 3.7 Sermon Management Endpoints

#### 3.7.1 Create Sermon
**POST** `/api/v1/sermons`

**Headers:**
- `Authorization: Bearer <token>`

**Request Body:**
```json
{
  "preacher": "Pastor John Smith",
  "sermon_summary": "A message about faith and perseverance in difficult times."
}
```

#### 3.7.2 Get Sermons
**GET** `/api/v1/sermons`

**Query Parameters:**
- `page` (optional): Page number
- `limit` (optional): Items per page
- `start_date` (optional): Filter by date range
- `end_date` (optional): Filter by date range

**Response:**
```json
{
  "success": true,
  "data": {
    "sermons": [
      {
        "id": "sermon_id",
        "preacher": "Pastor John Smith",
        "sermon_topic": "Faith and Perseverance",
        "sermon_summary": "A message about faith and perseverance...",
        "date_of_meeting": "2025-07-07T10:00:00+01:00",
        "entry_made_by": "CCI-12345"
      }
    ],
    "pagination": {
      "page": 1,
      "limit": 10,
      "total": 1,
      "total_pages": 1
    }
  }
}
```

### 3.8 Announcements Endpoints

#### 3.8.1 Create Announcement
**POST** `/api/v1/announcements`

**Headers:**
- `Authorization: Bearer <token>`

**Request Body:**
```json
{
  "title": "Youth Fellowship Meeting",
  "announcement_content": "Join us for our monthly youth fellowship meeting this Saturday at 4 PM."
}
```

#### 3.8.2 Get Announcements
**GET** `/api/v1/announcements`

**Query Parameters:**
- `page` (optional): Page number
- `limit` (optional): Items per page
- `status` (optional): Filter by status (Pending, Done)

### 3.9 Church Management Endpoints

#### 3.9.1 Create Church Data
**POST** `/api/v1/church`

**Headers:**
- `Authorization: Bearer <token>`

**Request Body:**
```json
{
  "church_name": "Christ Chapel International",
  "church_phone": "+2348123456789",
  "church_email": "info@cci.org",
  "church_address": "123 Church Street, Lagos, Nigeria",
  "church_sunday_meeting_time": 9,
  "midweek_meeting_day": "Wednesday",
  "midweek_meeting_time": 18
}
```

## 4. Security Requirements

### 4.1 Authentication & Authorization
- JWT-based authentication with access and refresh tokens
- Role-based access control (RBAC)
- Password hashing using bcrypt
- Token expiration: Access tokens (15 minutes), Refresh tokens (7 days)

### 4.2 Data Validation
- Input validation for all endpoints
- Email format validation
- Password strength requirements (minimum 8 characters, mixed case, numbers)
- Phone number format validation
- Date format validation

### 4.3 Security Measures
- Rate limiting on all endpoints
- HTTPS only communication
- CORS configuration
- Input sanitization
- SQL injection prevention
- XSS protection headers

### 4.4 Data Privacy
- User data encryption at rest
- Secure password storage
- PII data handling compliance
- User consent management

## 5. Error Handling

### 5.1 Standard Error Response Format
```json
{
  "success": false,
  "error": {
    "code": "VALIDATION_ERROR",
    "message": "Invalid input data",
    "details": [
      {
        "field": "email",
        "message": "Email format is invalid"
      }
    ]
  }
}
```

### 5.2 HTTP Status Codes
- `200`: Success
- `201`: Created
- `400`: Bad Request
- `401`: Unauthorized
- `403`: Forbidden
- `404`: Not Found
- `422`: Validation Error
- `429`: Rate Limit Exceeded
- `500`: Internal Server Error

## 6. Database Configuration

### 6.1 MongoDB Setup
- Database name: `church_attendance_db`
- Connection with authentication
- Indexes on frequently queried fields
- Data backup and recovery procedures

### 6.2 Collections
- `users`
- `attendance`
- `family_members`
- `sermons`
- `announcements`
- `departments`
- `roles`
- `church_info`
- `notifications`

## 7. Performance Requirements

### 7.1 Response Times
- Authentication endpoints: < 500ms
- Data retrieval endpoints: < 1000ms
- Complex queries: < 2000ms

### 7.2 Scalability
- Support for 1000+ concurrent users
- Horizontal scaling capability
- Database optimization for large datasets

## 8. Testing Requirements

### 8.1 Unit Tests
- Test coverage > 80%
- All endpoint functionality testing
- Database operation testing

### 8.2 Integration Tests
- End-to-end API testing
- Authentication flow testing
- Database integration testing

### 8.3 Security Tests
- Authentication bypass testing
- Authorization testing
- Input validation testing

## 9. Deployment & Monitoring

### 9.1 Environment Configuration
- Development, staging, and production environments
- Environment-specific configuration
- Secrets management

### 9.2 Monitoring
- API response time monitoring
- Error rate monitoring
- Database performance monitoring
- Security event logging

## 10. Documentation

### 10.1 API Documentation
- OpenAPI/Swagger specification
- Interactive API documentation
- Code examples for all endpoints

### 10.2 Developer Documentation
- Setup and installation guide
- Database schema documentation
- Authentication guide
- Best practices guide

## 11. Future Enhancements

### 11.1 Phase 2 Features
- Mobile app integration
- Push notifications
- Advanced analytics dashboard
- Bulk operations support

### 11.2 Phase 3 Features
- Multi-church support
- Event management
- Financial tracking
- Social features

## 12. Compliance & Standards

### 12.1 Data Protection
- GDPR compliance considerations
- Data retention policies
- Right to be forgotten implementation

### 12.2 API Standards
- RESTful API design principles
- JSON:API specification adherence
- Consistent naming conventions
- Proper HTTP method usage
