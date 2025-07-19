# Church Attendance Management System API Reference

**Base URL:**  
`http://localhost:8080/api/v1`

---

## Authentication

### Register (Step 1)
- **POST** `/auth/register`
- **Headers:** `Content-Type: application/json`
- **Body:**
  | Field                            | Type   | Required | Description                                 |
  |----------------------------------|--------|----------|---------------------------------------------|
  | email                           | string | Yes      | User's email address (must be unique)        |
  | user_password                   | string | Yes      | Password (min 8 chars,  strong recommended)  |

- **Sample Request**
      `let headersList = {
      "Accept": "*/*",
      "User-Agent": "Thunder Client (https://www.thunderclient.com)",
      "Content-Type": "application/json"
      }

      let bodyContent = JSON.stringify({
        "Email": "seun@gmail.com",
        "Password": "SecurePassword123!"
      });

      let response = await fetch("http://localhost:8080/api/v1/auth/register", { 
        method: "POST",
        body: bodyContent,
        headers: headersList
      });

      let data = await response.text();
      console.log(data);`

- **Sample Response**
      `{
        "success": true,
        "message": "User registered successfully",
        "data": {
          "user_id": "CCIMRB-32527",
          "email": "seun@gmail.com",
          "created_at": "2025-07-12T10:56:56.552188+01:00"
        }
      }`



### Register (Step 2)
- **POST** `/auth/register/complete`
- **Headers:** `Content-Type: application/json`
- **Body:** (fields as required for completion by the DTO, see sample request body below for guidance)
- **Sample Request**
    `let headersList = {
    "Accept": "*/*",
    "User-Agent": "Local API Client (local)",
    "Content-Type": "application/json"
    }

    let bodyContent = JSON.stringify({
      "FirstName": "Seun",
      "LastName": "Yusuf",
      "Email": "yusuf@gmail.com",
      "Password": "Graphene@work3",
      "Bio": "This is a complete registration for Seun test",
      "DateOfBirth": "1996-12-11",
      "Gender": "Male",
      "Member": true,
      "Visitor": false,
      "Usher": false,
      "Admin": false,
      "DateJoinedChurch": "2002-11-10",
      "FamilyHead": false,
      "UserCampus": "Utako",
      "CampusState": "Abuja",
      "CampusCountry": "Nigeria",
      "UserWorkDepartment":"3",
      "Profession": "Painter",
      "UserHouseAddress": "Maximiseing the Maitama penthouse",
      "PhoneNumber": "08155679876",
      "InstagraHhandle": "@testing",
      "FamilyMemberId": {
        "$numberLong": "1"
      },
      "EmergencyContatcName": "Michaela",
      "EmergencyContactPhone": "Suleman",
      "EmergencyContactEmail": "suleman@example.com",
      "EmergencyCOntactRelationship": "Brother"
    });

    let response = await fetch("http://localhost:8080/api/v1/auth/register/complete", { 
      method: "POST",
      body: bodyContent,
      headers: headersList
    });

    let data = await response.text();
    console.log(data);`

- **Sample Response**
    `{
      "success": true,
      "message": "User registered successfully",
      "data": {
        "user_id": "CCIMRB-89489",
        "email": "yusuf@gmail.com",
        "created_at": "2025-07-12T11:12:34.738426+01:00"
      }
    }`

`###I still need to Check the complete register endpoint for some fields that are not saving properly? Also check and standardise the way family member list should be stored for a family head user`



### Login
- **POST** `/auth/login`
- **Headers:** `Content-Type: application/json`
- **Body:**
  | Field         | Type   | Required | Description           |
  |---------------|--------|----------|-----------------------|
  | email         | string | Yes      | User's email address  |
  | user_password | string | Yes      | User's password       |

### Refresh Token
- **POST** `/auth/refresh`
- **Headers:** `Content-Type: application/json`
- **Body:**
  | Field       | Type   | Required | Description         |
  |-------------|--------|----------|---------------------|
  | refresh_token | string | Yes    | JWT refresh token   |

### Logout
- **POST** `/logout`
- **Headers:**  
  - `Content-Type: application/json`  
  - `Authorization: Bearer <JWT_ACCESS_TOKEN>`
  - Just pass the bearer token as authorisation, and the user will be logged out succesfully

---

## Users

### Get All Users
- **GET** `/users`
- **Headers:** `Authorization: Bearer <JWT_ACCESS_TOKEN>`

### Search Users
- **GET** `/users/search?query=<search>`
- **Headers:** `Authorization: Bearer <JWT_ACCESS_TOKEN>`

### Filter Users
- **GET** `/users/filter?role=<role>&member=<bool>`
- **Headers:** `Authorization: Bearer <JWT_ACCESS_TOKEN>`

filter using any field and value

---

## Attendance

### Create Attendance
- **POST** `/attendance`
- **Headers:** `Authorization: Bearer <JWT_ACCESS_TOKEN>`
- **Body:**
  | Field                   | Type   | Required | Description                       |
  |-------------------------|--------|----------|-----------------------------------|
  | user_id                 | string | Yes      | User's unique ID                  |


### QR Check-in
- **POST** `/attendance/qr-checkin`
- **Headers:** `Authorization: Bearer <JWT_ACCESS_TOKEN>`
- **Body:**
  | Field     | Type   | Required | Description        |
  |-----------|--------|----------|--------------------|
  | qr_code   | string | Yes      | QR code string     |

- **Sample Request**
      `let headersList = {
      "Accept": "*/*",
      "User-Agent": "Client name",
      "Authorization": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiQ0NJTVJCLTcwNjk4IiwiZTIwMDZ9.dAkT3Bx1hRcpYbuHUratTpdFgrwjWw_pm2X0UQ2K1Xw",
      "Content-Type": "application/json"
      }

      let bodyContent = JSON.stringify({
        "qr_code_token":"uT7ojQyOaaNu3FmA7Vk79tVacclZ88ZVa7dZXEsu-Ok="
      });

      let response = await fetch("http://localhost:8080/api/v1/attendance/qr-checkin", { 
        method: "POST",
        body: bodyContent,
        headers: headersList
      });

      let data = await response.text();
      console.log(data);`

- **Sample Response**
      `{
        "success": true,
        "message": "QR check-in successful",
        "data": {
          "id": "68722c2e565074bb89212dc5",
          "user_id": "CCIMRB-70698",
          "date_time_of_attendance": "2025-07-12T10:34:38.938171+01:00",
          "qrcode_based_checkin": true,
          "late": true,
          "manual_checkin": false,
          "visitor": false,
          "member": false
        }
      }`





### Attendance History
- **GET** `/attendance/history`
- **Headers:** `Authorization: Bearer <JWT_ACCESS_TOKEN>`

### Attendance Analytics
- **GET** `/attendance/analytics?date=<datetime>`
- **Headers:** `Authorization: Bearer <JWT_ACCESS_TOKEN>`

- **Body:**
  | Field | Type  | Required  | Description                                                           |
  |-------|-------|-----------|---------------------------------------------------------------------  |
  | date  | date  | yes       | This is the date range that the analytics data should be spooled for  |


## QR Code

### Generate QR Code
- **POST** `/qr/generate`
- **Headers:** `Authorization: Bearer <JWT_ACCESS_TOKEN>`
- **Body:**
  | Field     | Type   | Required | Description        |
  |-----------|--------|----------|--------------------|
  | user_id   | string | Yes      | User's unique ID   |

- **Sample Request:**
        `let headersList = {
        "Accept": "*/*",
        "User-Agent": "client name",
        "Authorization": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiQ0NJTVJCLTcwNjk4TIzMTIwMDZ9.dAkT3Bx1hRcpYbuHUratTpdFgrwjWw_pm2X0UQ2K1Xw",
        "Content-Type": "application/json"
        }

        let bodyContent = JSON.stringify({
          "user_id":"CCIMRB-70698"
        });

        let response = await fetch("http://localhost:8080/api/v1/qr/generate", { 
          method: "POST",
          body: bodyContent,
          headers: headersList
        });

        let data = await response.text();
        console.log(data);`
- **Sample Response of Success**
        `{
          "success": true,
          "data": {
            "qr_code_token": "uT7ojQyOaaNu3FmA7Vk79tVacclZ88ZVa7dZXEsu-Ok=",
            "qr_code_image": "iVBORw0KGgoAAAANSUhEUgAAAQAAAAEAAQMAAABmvDolAAAABlBMVEX///8AAABVwtN+AAAB+klEQVR42uyYPbLkIAyEmyJwyBE4Cjfzz804CkcgJKDcWxL2G7v2bbLRyGUlUzN8EwgJqSW89tpr/2cLSbZIz+JayJgKEOQ3PgsA4BpSD3KS+8QVIR8HdgBPbg3oiBsrgCgAWQwCMXdg1iiV2TDg5QSJDa4/EdCcjATiPlX8K2m/HRj1IfdQ5qkm3+LWfysg3w0MEzfhyIyJ7L+U9C8HFiBtJHvgOoK1e3H8Gs0HACC5jw+phr5pfB0LLAEL6/GyNJriYgfYYApAD3mWnJSXJQ+qqJstrpYArQ8sGHmW1L+gsudZAOQeVD8cL4sdyV1bkgWANW3qphbzPkk08y6dyRBw1IfUEfchf+CZd0RTwNJHnpFFW5Jy8pflWcB5HTWuE2UWmn1Njp9wmwDIvLQoFdxRVBBmX3FLWgMAVH9GSUb3MwulS7BMAEd9SD3wGOWomfc0wOtXaUnrGIJGTn5UkAVAoqk5Kd00HCoo8SIwTADDzeRrFHmNqcgDu84PFoBD5CQN1pinD9nzKOCc9XyN1AuI2xDeiyXgnLt1FRLGxBruMskCMJaHIobWMT64Dux/bRe/HBj7KLLi3EepuvtouQcBo5jL+ACtiLAK1DEE6ao089azDAA/W3epcrrp9beVmw3gXB4iSmOFBuu2yH0E8Nprr93tTwAAAP//vvDCp6xAOnoAAAAASUVORK5CYII="
          }
        }`
---

## Roles (Admin Only)

### Create Role
- **POST** `/roles`
- **Headers:**  
  - `Authorization: Bearer <JWT_ACCESS_TOKEN>`  
  - Must be admin
- **Body:**  
  | Field       | Type   | Required   | Description                                   |
  |-------------|--------|------------|-----------------------------------------------|
  | name        | string | Yes        | Role name                                     |
  | permissions | list   | Yes        |  List of permissions you want the role to have|

- **Sample Request**
    `let headersList = {
    "Accept": "*/*",
    "User-Agent": "Thunder Client (https://www.thunderclient.com)",
    "Authorization": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiQ0NJaWwiOiJ5dXN1ZkBnbWFpbC5jb20iLCJhZG1pbiI6dHJ1ZSwiaXNzIjoiY2h1cmNoLWF0dGVuZGFuY2UtYXBpIiwiZXhwIjoxNzUyMzE3MTMyLCJpYXQiOjE3NTIzMTYyMzJ9.aqqQYwNa8RQ5VweqEJT3Cg4bqEfym45MQ4iYGWwLpYA",
    "Content-Type": "application/json"
    }

    let bodyContent = JSON.stringify({
      "role_name": "Usher",
      "role_description":"This is the basic role created for an Usher",
      "permissions":"Can Do QR Code Checkin and mark attendance manually"
    });

    let response = await fetch("http://localhost:8080/api/v1/roles", { 
      method: "POST",
      body: bodyContent,
      headers: headersList
    });

    let data = await response.text();
    console.log(data);`


- **Sample Response**
    `{
      "code": "ROLE_CREATED",
      "message": "Role created successfully",
      "data": {
        "id": "68723b47949fcaa17c3b88e7",
        "role_name": "Usher",
        "role_description": "This is the basic role created for an Usher",
        "permissions": "Can Do QR Code Checkin and mark attendance manually",
        "total_members": 0,
        "date_added": "2025-07-12T11:39:03.159103+01:00",
        "date_updated": "2025-07-12T11:39:03.159103+01:00"
      }
    }`




### Fetch List of all Roles
- **GET** `/roles`
- **Headers:**
  - `Authorization: Bearer <JWT_ACESS_TOKEN>`
  - Must be admin
- **Body:**
- **Sample Request**
    `let headersList = {
    "Accept": "*/*",
    "User-Agent": "Local Client",
    "Authorization": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiQ0NJTVJCLTg5NDg5IiwiZW1haWwiOiJ5dXN1ZkBnbWFpbC5jb20iLCJhZG1pbiI6dHJ1ZSwiaXNzIjoiY2h1cmNoLWF0dGVuZGFuY2UtYXBpIiwiZXhwIjoxNzUyMzE3MTMyLCJpYXQiOjE3NTIzMTYyMzJ9.aqqQYwNa8RQ5VweqEJT3Cg4bqEfym45MQ4iYGWwLpYA"
    }

    let response = await fetch("http://localhost:8080/api/v1/roles", { 
      method: "GET",
      headers: headersList
    });

    let data = await response.text();
    console.log(data);`

- **Sample Response**
      `{
        "code": "ROLES_RETRIEVED",
        "message": "Roles retrieved successfully",
        "data": {
          "data": [
            {
              "id": "68723d38949fcaa17c3b88eb",
              "role_name": "Member",
              "role_description": "This is for the member",
              "permissions": "Can do all things through Christ who gives him or her the strenght",
              "total_members": 0,
              "date_added": "2025-07-12T10:47:20.788Z",
              "date_updated": "2025-07-12T10:47:20.788Z"
            },
            {
              "id": "68723ca2949fcaa17c3b88e9",
              "role_name": "Youth Leader",
              "role_description": "This is the youth role created for an Big bosses",
              "permissions": "Can Do all things through Christ who strengthens him",
              "total_members": 0,
              "date_added": "2025-07-12T10:44:50.649Z",
              "date_updated": "2025-07-12T10:44:50.649Z"
            },
            {
              "id": "68723c93949fcaa17c3b88e8",
              "role_name": "Admin",
              "role_description": "This is the super admin role created for an Big bosses",
              "permissions": "Can Do all things through Christ who strengthens him",
              "total_members": 0,
              "date_added": "2025-07-12T10:44:35.268Z",
              "date_updated": "2025-07-12T10:44:35.268Z"
            },
            {
              "id": "68723b47949fcaa17c3b88e7",
              "role_name": "Usher",
              "role_description": "This is the basic role created for an Usher",
              "permissions": "Can Do QR Code Checkin and mark attendance manually",
              "total_members": 0,
              "date_added": "2025-07-12T10:39:03.159Z",
              "date_updated": "2025-07-12T10:39:03.159Z"
            }
          ],
          "pagination": {
            "page": 1,
            "limit": 10,
            "total": 4,
            "total_pages": 1
          }
        }
      }`


### Update Role
- **PUT** `/roles/:id`
- **Headers:** `Authorization: Bearer <JWT_ACCESS_TOKEN>` (admin)
- **Body:**  
  | Field | Type   | Required | Description      |
  |-------|--------|----------|------------------|
  | name  | string | Yes      | Role name        |

- **Sample Request:**
      `let headersList = {
      "Accept": "*/*",
      "User-Agent": "Local Client",
      "Authorization": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiQ0NJTVJCLTg5MTcxODN9.6grb-nkJqQPz1CU7Ch79T4DLMsJQeqn__AdfqmZcE6c",
      "Content-Type": "application/json"
      }

      let bodyContent = JSON.stringify({
        "role_name": "Member2",
        "role_description":"This is for the second member",
        "permissions":"Can do all things through Christ who gives him or her the strenght"
      });

      let response = await fetch("http://localhost:8080/api/v1/roles/68723d38949fcaa17c3b88eb", { 
        method: "PUT",
        body: bodyContent,
        headers: headersList
      });

      let data = await response.text();
      console.log(data);`



- **Sample Response:**
    `{
      "code": "ROLE_UPDATED",
      "message": "Role updated successfully",
      "data": {
        "id": "68723d38949fcaa17c3b88eb",
        "role_name": "Member2",
        "role_description": "This is for the second member",
        "permissions": "Can do all things through Christ who gives him or her the strenght",
        "total_members": 0,
        "date_added": "2025-07-12T10:47:20.788Z",
        "date_updated": "2025-07-12T11:52:42.976094+01:00"
      }
    }`

### Delete Role
- **DELETE** `/roles/:id`
- **Headers:** `Authorization: Bearer <JWT_ACCESS_TOKEN>` (admin)
- **Sample Request:**
      `let headersList = {
      "Accept": "*/*",
      "User-Agent": "Local Client",
      "Authorization": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiQ0NJTVJCLTg5NDg5IiwiZW1haWwiOiJ5dXN1ZkBnbWFpbC5jb20iLCJhZG1pbiI6dHJ1ZSwiaXNzIjoiY2h1cmNoLWF0dGVuZGFuY2UtYXBpIiwiZXhwIjoxNzUyMzE4ODE1LCJpYXQiOjE3NTIzMTc5MTV9.WKlKhQG5CITr6fRVWPimMR1dh6mAdn545PZ-u7ur9MM"
      }

      let response = await fetch("http://localhost:8080/api/v1/roles/68723efd949fcaa17c3b88ec", { 
        method: "DELETE",
        headers: headersList
      });

      let data = await response.text();
      console.log(data);

- **Sample Response:**
      {
        "code": "ROLE_DELETED",
        "message": "Role deleted successfully"
      }`

---

## Sermons

### Create Sermon
- **POST** `/sermons`
- **Headers:** `Authorization: Bearer <JWT_ACCESS_TOKEN>`
- **Body:**  
  | Field         | Type   | Required | Description                |
  |---------------|--------|----------|----------------------------|
  | title         | string | Yes      | Sermon title               |
  | date_of_meeting | string | Yes    | Date of sermon (YYYY-MM-DD)|
  | entry_made_by | string | Yes      | User ObjectId              |
  | ...           | ...    | ...      | Other sermon fields        |

- **Sample Request:**
      `let headersList = {
      "Accept": "*/*",
      "User-Agent": "Thunder Client (https://www.thunderclient.com)",
      "Authorization": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiQ0NJTVJCLTg5NDg5IiwiZW1haWwiOiJ5dXN1ZkBnbWFpbC5jb20iLCJhZG1pbiI6dHJ1ZSwiaXNzIjoiY2h1cmNoLWF0dGVuZGFuY2UtYXBpIiwiZXhwIjoxNzUyNTgyNzc3LCJpYXQiOjE3NTI1ODE4Nzd9.1rFHpRzmT5YAoyvbh22rbJWxcdZpjakOwffFlVUMxWw",
      "Content-Type": "application/json"
      }

      let bodyContent = JSON.stringify({
        "title": "Orthodoxy",
        "speaker": "Pastor Lazarus",
        "date": "1996-11-11",
        "video_url": "https://www.youtubbe.com/music/audioe4",
        "audio_url": "https://www.youtube.com/videso",
        "notes": "This is the summary of the teaching on Orhtodoxy",
        "scripture": "Matthew 4:3",
        "series": "Landmarks",
        "tags": []
      });

      let response = await fetch("http://localhost:8080/api/v1/sermons", { 
        method: "POST",
        body: bodyContent,
        headers: headersList
      });

      let data = await response.text();
      console.log(data);`

- **Sample Response:**
      `{
        "code": "SERMON_CREATED",
        "message": "Sermon created successfully",
        "data": {
          "id": "68764730b58062ebfdfc59d6",
          "title": "Orthodoxy",
          "speaker": "Pastor Lazarus",
          "date": "1996-11-11",
          "video_url": "https://www.youtubbe.com/music/audioe4",
          "audio_url": "https://www.youtube.com/videso",
          "notes": "This is the summary of the teaching on Orhtodoxy",
          "scripture": "Matthew 4:3",
          "series": "Landmarks",
          "tags": null,
          "date_added": "2025-07-15T13:18:56.72782+01:00",
          "date_updated": "2025-07-15T13:18:56.72782+01:00"
        }
      }`



### Fetch list of sermons
- **GET** `/sermons`
- **Headers:** `Authorization: Bearer <JWT_ACESS_TOKEN>`
- **Sample Request:**
        `let headersList = {
        "Accept": "*/*",
        "User-Agent": "Local Client",
        "Authorization": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiQ0NJTVJC1NzgxOTJ9.Jj41tDRp-R8U-UiXt_psJsVIhqm0d5AsnwxluWVb9ZY"
        }

        let response = await fetch("http://localhost:8080/api/v1/sermons", { 
          method: "GET",
          headers: headersList
        });

        let data = await response.text();
        console.log(data);`


- **Sample Response:**
      `{
        "code": "SERMONS_RETRIEVED",
        "message": "Sermons retrieved successfully",
        "data": {
          "data": [
            {
              "id": "68764730b58062ebfdfc59d6",
              "title": "Orthodoxy",
              "speaker": "Pastor Lazarus",
              "date": "1996-11-11",
              "video_url": "https://www.youtubbe.com/music/audioe4",
              "audio_url": "https://www.youtube.com/videso",
              "notes": "This is the summary of the teaching on Orhtodoxy",
              "scripture": "Matthew 4:3",
              "series": "Landmarks",
              "tags": null,
              "date_added": "2025-07-15T13:30:02.891345+01:00",
              "date_updated": "2025-07-15T13:30:02.891345+01:00"
            },
            {
              "id": "687647c2b58062ebfdfc59d7",
              "title": "Charisma",
              "speaker": "Pastor Iren",
              "date": "1946-10-21",
              "video_url": "https://www.youtubbe.com/music/audioe4",
              "audio_url": "https://www.youtube.com/videso",
              "notes": "This is the note of the teaching on by pastor iren today",
              "scripture": "Matthew 4:3",
              "series": "Landmarks",
              "tags": null,
              "date_added": "2025-07-15T13:30:02.891345+01:00",
              "date_updated": "2025-07-15T13:30:02.891346+01:00"
            },
            {
              "id": "687647e2b58062ebfdfc59d8",
              "title": "Charisma3",
              "speaker": "Pastor Iren",
              "date": "1946-10-21",
              "video_url": "https://www.youtubbe.com/music/audioe4",
              "audio_url": "https://www.youtube.com/videso",
              "notes": "This is the note of the teaching on by pastor iren today",
              "scripture": "Matthew 4:3",
              "series": "Landmarks",
              "tags": null,
              "date_added": "2025-07-15T13:30:02.891346+01:00",
              "date_updated": "2025-07-15T13:30:02.891346+01:00"
            }
          ],
          "pagination": {
            "page": 1,
            "limit": 10,
            "total": 3,
            "total_pages": 1
          }
        }
      }`


### Update Sermon
- **PUT** `/sermons/:id`
- **Headers:** `Authorization: Bearer <JWT_ACCESS_TOKEN>`
- **Body:** 
- **Sample Request:**
    `let headersList = {
    "Accept": "*/*",
    "User-Agent": "Local Client",
    "Authorization": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiQ0NJTVJCLTg5NDgDV9.Rt3kV6Phl2fGsfm-IYRTWJXQL4aYH0ytnJwYrbEGAFI",
    "Content-Type": "application/json"
    }

    let bodyContent = JSON.stringify({
      "title": "CharismaUpdate",
      "speaker": "Pastor Iren New",
      "date": "1946-14-2100",
      "video_url": "https://www.youtubbe.com/music/audioe4",
      "audio_url": "https://www.youtube.com/videso",
      "notes": "This is the note of the teaching on by pastor iren today",
      "scripture": "Matthew 4:3updated",
      "series": "Landmarks",
      "tags": ["tag"]
    });

    let response = await fetch("http://localhost:8080/api/v1/sermons/687647e2b58062ebfdfc59d8", { 
      method: "PUT",
      body: bodyContent,
      headers: headersList
    });

    let data = await response.text();
    console.log(data);`

- **Sample Response:**
    `{
      "code": "SERMON_UPDATED",
      "message": "Sermon updated successfully",
      "data": {
        "id": "687647e2b58062ebfdfc59d8",
        "title": "CharismaUpdate",
        "speaker": "Pastor Iren New",
        "date": "1946-14-2100",
        "video_url": "https://www.youtubbe.com/music/audioe4",
        "audio_url": "https://www.youtube.com/videso",
        "notes": "This is the note of the teaching on by pastor iren today",
        "scripture": "Matthew 4:3",
        "series": "Landmarks",
        "tags": null,
        "date_added": "2025-07-15T13:36:15.31688+01:00",
        "date_updated": "2025-07-15T13:36:15.31688+01:00"
      }
    }`


### Delete Sermon
- **DELETE** `/sermons/:id`
- **Headers:** `Authorization: Bearer <JWT_ACCESS_TOKEN>`
- **Sample Request:**
      `let headersList = {
      "Accept": "*/*",
      "User-Agent": "Local Client",
      "Authorization": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiQ0NJTVJCLTg5N5MDV9.Rt3kV6Phl2fGsfm-IYRTWJXQL4aYH0ytnJwYrbEGAFI"
      }

      let response = await fetch("http://localhost:8080/api/v1/sermons/687647e2b58062ebfdfc59d8", { 
        method: "DELETE",
        headers: headersList
      });

      let data = await response.text();
      console.log(data);`
- **Sample Response:**
    `{
      "code": "SERMON_DELETED",
      "message": "Sermon deleted successfully"
    }`
---

## Announcements

### Create Announcement
- **POST** `/announcements`
- **Headers:** `Authorization: Bearer <JWT_ACCESS_TOKEN>`
- **Body:**  
  | Field         | Type   | Required | Description                |
  |---------------|--------|----------|----------------------------|
  | title         | string | Yes      | Announcement title         |
  | announcement_date | string | Yes  | Date (YYYY-MM-DD)          |
  | status        | string | Yes      | 'Pending' or 'Done'        |
  | announcement_entry_made_by | string | Yes | User ObjectId      |

- **Sample Request:**
    `let headersList = {
    "Accept": "*/*",
    "User-Agent": "Local Client",
    "Authorization": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiQ0NJ3NTI4NTM4OTh9.aK59e2tpNwv88r6-eG1qccRhMtP8Cn6JXw4fTFisFd8",
    "Content-Type": "application/json"
    }

    let bodyContent = JSON.stringify({
      "title": "Announcement three",
      "content": "the date of the next prayer stretch is very important for all to attend",
      "announcement_due_date": "2023-02-23",
      "start_date": "2026-01-02",
      "end_date": "2026-02-02",
      "type": "event",
      "priority": "high",
      "target_users": ["gaspers"],
      "image_url": "htpps://www.image.url/ihijidhubus"
    });

    let response = await fetch("http://localhost:8080/api/v1/announcements", { 
      method: "POST",
      body: bodyContent,
      headers: headersList
    });

    let data = await response.text();
    console.log(data);`

- **Sample Response:**
    `{
      "code": "ANNOUNCEMENT_CREATED",
      "message": "Announcement created successfully",
      "data": {
        "id": "687a701f6e2ce0eefa473a9e",
        "title": "Announcement three",
        "content": "the date of the next prayer stretch is very important for all to attend",
        "type": "event",
        "announcement_due_date": "2023-02-23T00:00:00Z",
        "start_date": "2026-01-02T00:00:00Z",
        "end_date": "2026-02-02T00:00:00Z",
        "priority": "high",
        "target_users": [
          "gaspers"
        ],
        "image_url": "htpps://www.image.url/ihijidhubus",
        "status": "Pending",
        "date_added": "2025-07-18T17:02:39.892673+01:00",
        "date_updated": "2025-07-18T17:02:39.892673+01:00",
        "entry_made_by": "000000000000000000000000"
      }
    }`

### Fetch all the Announcements
- **GET** `/announcements/`
- **HEaders:** `Authorization: Bearer <JWT_BEARER_TOKEN>`
- **Sample Request:**
   `let headersList = {
    "Accept": "*/*",
    "User-Agent": "Local Client",
    "Authorization": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiQ0NJTVJCLTg54NTY0MDV9.wwWPaZ1gsd8zNHPL3XszXMe72NXny56QOSxDC6_Bbuw"
    }

    let response = await fetch("http://localhost:8080/api/v1/announcements", { 
      method: "GET",
      headers: headersList
    });

    let data = await response.text();
    console.log(data);`

- **Sample Response:**
    `{
      "code": "ANNOUNCEMENTS_RETRIEVED",
      "message": "Announcements retrieved successfully",
      "data": {
        "data": [
          {
            "id": "687a5e978aca1efa2ca227dd",
            "title": "Announcement one",
            "content": "the date of the next prayer stretch is very important for all to attend",
            "type": "event",
            "announcement_due_date": "2023-01-23T00:00:00Z",
            "start_date": "0001-01-01T00:00:00Z",
            "end_date": "2020-02-02T14:00:00Z",
            "priority": "high",
            "target_users": [
              "gaspers"
            ],
            "image_url": "htpps://www.image.url/ihijidhubus",
            "status": "Pending",
            "date_added": "2025-07-18T14:47:51.816Z",
            "date_updated": "2025-07-18T14:47:51.816Z",
            "entry_made_by": "000000000000000000000000"
          },
          . . . 
          {
            "id": "687a701f6e2ce0eefa473a9e",
            "title": "Announcement updated!",
            "content": "the date of the updated next prayer stretch is very important for all to attend",
            "type": "event",
            "announcement_due_date": "2023-02-23T00:00:00Z",
            "start_date": "2026-01-02T00:00:00Z",
            "end_date": "2026-02-02T00:00:00Z",
            "priority": "high",
            "target_users": [
              "gaspers"
            ],
            "image_url": "htpps://www.image.url/ihijidhubus",
            "status": "Pending",
            "date_added": "2025-07-18T16:02:39.892Z",
            "date_updated": "2025-07-18T16:02:39.892Z",
            "entry_made_by": "000000000000000000000000"
          }
        ],
        "pagination": {
          "page": 1,
          "limit": 10,
          "total": 5,
          "total_pages": 1
        }
      }
    }`


### Get Announcement by ID
- **GET** `/announcements/:id`
- **Headers:** `Authorization: Bearer <JWT_ACCESS_TOKEN>`
- **sample Request:**
    `let headersList = {
    "Accept": "*/*",
    "User-Agent": "Local Client",
    "Authorization": "Bearer eyJhbGciOY0MDV9.wwWPaZ1gsd8zNHPL3XszXMe72NXny56QOSxDC6_Bbuw"
    }

    let response = await fetch("http://localhost:8080/api/v1/announcements/687a701f6e2ce0eefa473a9e", { 
      method: "GET",
      headers: headersList
    });

    let data = await response.text();
    console.log(data);`

- **sample Response:**
    `{
      "code": "ANNOUNCEMENT_RETRIEVED",
      "message": "Announcement retrieved successfully",
      "data": {
        "id": "687a701f6e2ce0eefa473a9e",
        "title": "Announcement updated!",
        "content": "the date of the updated next prayer stretch is very important for all to attend",
        "type": "event",
        "announcement_due_date": "2023-02-23T00:00:00Z",
        "start_date": "2026-01-02T00:00:00Z",
        "end_date": "2026-02-02T00:00:00Z",
        "priority": "high",
        "target_users": [
          "gaspers"
        ],
        "image_url": "htpps://www.image.url/ihijidhubus",
        "status": "Pending",
        "date_added": "2025-07-18T16:02:39.892Z",
        "date_updated": "2025-07-18T16:02:39.892Z",
        "entry_made_by": "000000000000000000000000"
      }
    }`



### Update Announcement
- **PUT** `/announcements/:id`
- **Headers:** `Authorization: Bearer <JWT_ACCESS_TOKEN>`
- **Body:** (same as create)
- **sample Request:**
    `let headersList = {
    "Accept": "*/*",
    "User-Agent": "Local Client",
    "Authorization": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiQ0NJTE3NTI4NTQ4MzV9.OBx3aUZp7dv0zD7wtIcs5CH_2XggcHnKqU1mym3jrEY",
    "Content-Type": "application/json"
    }

    let bodyContent = JSON.stringify({
      "title": "Announcement updated!",
      "content": "the date of the updated next prayer stretch is very important for all to attend",
      "announcement_due_date": "2024-02-23",
      "start_date": "2026-01-02",
      "end_date": "2026-02-02",
      "type": "event",
      "priority": "medium",
      "target_users": ["jaffar"],
      "image_url": "https://www.image.urls/ihijidhubus"
    });

    let response = await fetch("http://localhost:8080/api/v1/announcements/687a701f6e2ce0eefa473a9e", { 
      method: "PUT",
      body: bodyContent,
      headers: headersList
    });

    let data = await response.text();
    console.log(data);`


- **Sample Response:**
    `{
      "code": "ANNOUNCEMENT_UPDATED",
      "message": "Announcement updated successfully",
      "data": {
        "id": "687a701f6e2ce0eefa473a9e",
        "title": "Announcement updated!",
        "content": "the date of the updated next prayer stretch is very important for all to attend",
        "type": "event",
        "announcement_due_date": "2023-02-23T00:00:00Z",
        "start_date": "2026-01-02T00:00:00Z",
        "end_date": "2026-02-02T00:00:00Z",
        "priority": "high",
        "target_users": [
          "gaspers"
        ],
        "image_url": "htpps://www.image.url/ihijidhubus",
        "status": "Pending",
        "date_added": "2025-07-18T16:02:39.892Z",
        "date_updated": "2025-07-18T16:02:39.892Z",
        "entry_made_by": "000000000000000000000000"
      }
    }`

### Delete Announcement
- **DELETE** `/announcements/:id`
- **Headers:** `Authorization: Bearer <JWT_ACCESS_TOKEN>`
- **Sample Request:**
    `let headersList = {
    "Accept": "*/*",
    "User-Agent": "Thunder Client (https://www.thunderclient.com)",
    "Authorization": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiQ0NJTVJCLTg5NY0MDV9.wwWPaZ1gsd8zNHPL3XszXMe72NXny56QOSxDC6_Bbuw"
    }

    let response = await fetch("http://localhost:8080/api/v1/announcements/687a701f6e2ce0eefa473a9e", { 
      method: "DELETE",
      headers: headersList
    });

    let data = await response.text();
    console.log(data);`

- **Sample Response:**
    `{
      "code": "ANNOUNCEMENT_DELETED",
      "message": "Announcement deleted successfully"
    }`
---

## Family Members
---- Family member endpoints are yet to be tested!!!----
### Create Family Member
- **POST** `/family-members`
- **Headers:** `Authorization: Bearer <JWT_ACCESS_TOKEN>`
- **Body:**  
  | Field         | Type   | Required | Description                |
  |---------------|--------|----------|----------------------------|
  | user_id       | string | Yes      | User's unique ID           |
  | name          | string | Yes      | Family member's name       |
  | relationship  | string | Yes      | Relationship to user       |
  | date_joined   | string | No       | Date joined (YYYY-MM-DD)   |

### Update Family Member
- **PUT** `/family-members/:id`
- **Headers:** `Authorization: Bearer <JWT_ACCESS_TOKEN>`
- **Body:** (same as create)

### Delete Family Member
- **DELETE** `/family-members/:id`
- **Headers:** `Authorization: Bearer <JWT_ACCESS_TOKEN>`

---

## Local Churches (Admin Only)

### Create Church
- **POST** `/churches`
- **Headers:** `Authorization: Bearer <JWT_ACCESS_TOKEN>` (admin)
- **Body:**  
  | Field         | Type   | Required | Description                |
  |---------------|--------|----------|----------------------------|
  | name          | string | Yes      | Church name                |
  | address       | string | Yes      | Church address             |
  | ...           | ...    | ...      | Other church fields        |

- **Sample Request:**
    `let headersList = {
    "Accept": "*/*",
    "User-Agent": "Local Client",
    "Authorization": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiQ0NJTVJCLTg5NDg5IiwiZW19.b2HRcQgcKWYPVpeWpqNsW7tZV12OwXc8xMgckNIHwYk",
    "Content-Type": "application/json"
    }

    let bodyContent = JSON.stringify({
      "church_name": "Utako CCI",
      "church_phone": "8155657687",
      "church_email": "utako@joincci.org",
      "church_address": "31, eden parks and garden, uatako",
      "state_county": "Federal capital territory",
      "country": "Nigeria",
      "sunday_meeting_time": 9,
      "midweek_meeting_day": "Wednesday",
      "midweek_meeting_time": 17,
      "website": "http://www.joincci.org",
      "social_media": "@cci_utako",
      "pastor_name": "Pastor Mike Micheals",
      "Pastor_phone": "0989876554",
      "pastor_email": "mike@joincci.org",
      "founded_year": 2016,
      "description": "This is the hq of the Northern church"
    });

    let response = await fetch("http://localhost:8080/api/v1/churches", { 
      method: "POST",
      body: bodyContent,
      headers: headersList
    });

    let data = await response.text();
    console.log(data);`


- **Sample Response:**
    `{
      "code": "CHURCH_CREATED",
      "message": "Church created successfully",
      "data": {
        "id": "687a896aa4380825e6c2e7b0",
        "church_name": "Utako CCI",
        "church_phone": "8155657687",
        "church_email": "utako@joincci.org",
        "church_address": "31, eden parks and garden, uatako",
        "state_county": "Federal capital territory",
        "country": "Nigeria",
        "sunday_meeting_time": 9,
        "midweek_meeting_day": "Wednesday",
        "midweek_meeting_time": 17,
        "website": "http://www.joincci.org",
        "social_media": "@cci_utako",
        "pastor_name": "Utako CCI",
        "pastor_phone": "0989876554",
        "pastor_email": "utako@joincci.org",
        "founded_year": 2016,
        "description": "This is the hq of the Northern church",
        "date_added": "2025-07-18T18:50:34.669977+01:00",
        "date_updated": "2025-07-18T18:50:34.669977+01:00"
      }
    }`

### Update Church
- **PUT** `/churches/:id`
- **Headers:** `Authorization: Bearer <JWT_ACCESS_TOKEN>` (admin)
- **Body:** (same as create)

- **Sample Request:**
`let headersList = {
 "Accept": "*/*",
 "User-Agent": "Local Client ",
 "Authorization": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiQ0NJTVJOjE3NTI4NjEzOTd9.VojBn-hyTxUpkj7hHgCXjEcFvNIEPJtLqcoaJG52ztI",
 "Content-Type": "application/json"
}

let bodyContent = JSON.stringify({
  "church_name": "Updated CCI",
  "church_phone": "8155657687",
  "church_email": "mrrb@joincci.org",
  "church_address": "31, Updated Wulvan event centre",
  "state_county": "Federal capital territory",
  "country": "Nigeria",
  "sunday_meeting_time": 9,
  "midweek_meeting_day": "Wednesday",
  "midweek_meeting_time": 17,
  "website": "http://www.joincci.org",
  "social_media": "@cci_marraba",
  "pastor_name": "Pastor Yemi Arowolo",
  "Pastor_phone": "0989876554",
  "pastor_email": "updatedyemi@joincci.org",
  "founded_year": 2021,
  "description": "This is the hq of the Northern mararaba church"
});

let response = await fetch("http://localhost:8080/api/v1/churches/687a8affa4380825e6c2e7b5", { 
  method: "PUT",
  body: bodyContent,
  headers: headersList
});

let data = await response.text();
console.log(data);`

- **Sample Response:**
    `{
      "code": "CHURCH_UPDATED",
      "message": "Church updated successfully",
      "data": {
        "id": "687a8affa4380825e6c2e7b5",
        "church_name": "Updated CCI",
        "church_phone": "8155657687",
        "church_email": "mrrb@joincci.org",
        "church_address": "31, Updated Wulvan event centre",
        "state_county": "Federal capital territory",
        "country": "Nigeria",
        "sunday_meeting_time": 9,
        "midweek_meeting_day": "Wednesday",
        "midweek_meeting_time": 17,
        "website": "http://www.joincci.org",
        "social_media": "@cci_marraba",
        "pastor_name": "Updated CCI",
        "pastor_phone": "0989876554",
        "pastor_email": "updatedyemi@joincci.org",
        "founded_year": 2021,
        "description": "This is the hq of the Northern mararaba church",
        "date_added": "2025-07-18T17:57:19.839Z",
        "date_updated": "2025-07-18T19:05:46.022937+01:00"
      }
    }`


### Get Church by ID
- **GET** `/churches/:id`
- **Headers:** `Authorization: Bearer <JWT_ACCESS_TOKEN>` (admin)
- **Body:** None

- **Sample Request:**
    `let headersList = {
    "Accept": "*/*",
    "User-Agent": "Local Client",
    "Authorization": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiQ0NJTVJCLTI4NjEzOTd9.VojBn-hyTxUpkj7hHgCXjEcFvNIEPJtLqcoaJG52ztI"
    }

    let response = await fetch("http://localhost:8080/api/v1/churches/687a8affa4380825e6c2e7b5", { 
      method: "GET",
      headers: headersList
    });

    let data = await response.text();
    console.log(data);`

- **Sample Response:**
    `{
      "code": "CHURCH_RETRIEVED",
      "message": "Church retrieved successfully",
      "data": {
        "id": "687a8affa4380825e6c2e7b5",
        "church_name": "Test CCI",
        "church_phone": "8155657687",
        "church_email": "mrrb@joincci.org",
        "church_address": "31, Wulvan event centre",
        "state_county": "Federal capital territory",
        "country": "Nigeria",
        "sunday_meeting_time": 9,
        "midweek_meeting_day": "Wednesday",
        "midweek_meeting_time": 17,
        "website": "http://www.joincci.org",
        "social_media": "@cci_marraba",
        "pastor_name": "Test CCI",
        "pastor_phone": "0989876554",
        "pastor_email": "mrrb@joincci.org",
        "founded_year": 2021,
        "description": "This is the hq of the Northern mararaba church",
        "date_added": "2025-07-18T17:57:19.839Z",
        "date_updated": "2025-07-18T17:57:19.839Z"
      }
    }`


### Fetch the list of all the Church
- **GET** `/churches/`
- **Headers:** `Authorization: Bearer <JWT_ACCESS_TOKEN>` (must be admin)

- **Sample Request:**
    `let headersList = {
    "Accept": "*/*",
    "User-Agent": "Local Client",
    "Authorization": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiQ0NJTVJCLTgNTI4NjEzOTd9.VojBn-hyTxUpkj7hHgCXjEcFvNIEPJtLqcoaJG52ztI"
    }

    let response = await fetch("http://localhost:8080/api/v1/churches", { 
      method: "GET",
      headers: headersList
    });

    let data = await response.text();
    console.log(data);`


- **Sample Response:**
`{
  "code": "CHURCHES_RETRIEVED",
  "message": "Churches retrieved successfully",
  "data": {
    "data": [
      {
        "id": "687a896aa4380825e6c2e7b0",
        "church_name": "Utako CCI",
        "church_phone": "8155657687",
        "church_email": "utako@joincci.org",
        "church_address": "31, eden parks and garden, uatako",
        "state_county": "Federal capital territory",
        "country": "Nigeria",
        "sunday_meeting_time": 9,
        "midweek_meeting_day": "Wednesday",
        "midweek_meeting_time": 17,
        "website": "http://www.joincci.org",
        "social_media": "@cci_utako",
        "pastor_name": "Utako CCI",
        "pastor_phone": "0989876554",
        "pastor_email": "utako@joincci.org",
        "founded_year": 2016,
        "description": "This is the hq of the Northern church",
        "date_added": "2025-07-18T17:50:34.669Z",
        "date_updated": "2025-07-18T17:50:34.669Z"
      },
       . . . .,
      {
        "id": "687a8affa4380825e6c2e7b5",
        "church_name": "Test CCI",
        "church_phone": "8155657687",
        "church_email": "mrrb@joincci.org",
        "church_address": "31, Wulvan event centre",
        "state_county": "Federal capital territory",
        "country": "Nigeria",
        "sunday_meeting_time": 9,
        "midweek_meeting_day": "Wednesday",
        "midweek_meeting_time": 17,
        "website": "http://www.joincci.org",
        "social_media": "@cci_marraba",
        "pastor_name": "Test CCI",
        "pastor_phone": "0989876554",
        "pastor_email": "mrrb@joincci.org",
        "founded_year": 2021,
        "description": "This is the hq of the Northern mararaba church",
        "date_added": "2025-07-18T17:57:19.839Z",
        "date_updated": "2025-07-18T17:57:19.839Z"
      }
    ],
    "pagination": {
      "page": 1,
      "limit": 10,
      "total": 5,
      "total_pages": 1
    }
  }
}`

### Delete Church
- **DELETE** `/churches/:id`
- **Headers:** `Authorization: Bearer <JWT_ACCESS_TOKEN>` (must be admin)

- **Sample Request:**
   `let headersList = {
    "Accept": "*/*",
    "User-Agent": "Local Client ",
    "Authorization": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiQ0NJTVJCOjE3NTI4NjEzOTd9.VojBn-hyTxUpkj7hHgCXjEcFvNIEPJtLqcoaJG52ztI"
    }

    let response = await fetch("http://localhost:8080/api/v1/churches/687a8affa4380825e6c2e7b5", { 
      method: "DELETE",
      headers: headersList
    });

    let data = await response.text();
    console.log(data);`


- **Sample REsponse:**
    `{
      "code": "CHURCH_DELETED",
      "message": "Church deleted successfully"
    }`

---------------------------------------------------------------------------

## General Notes

- **All endpoints (except `/auth/*`) require the `Authorization: Bearer <JWT_ACCESS_TOKEN>` header.**
- **Fields marked as 'Yes' in the 'Required' column must be provided in the request.**
- **Date fields should be in `YYYY-MM-DD` format unless otherwise specified.**
- **For endpoints requiring admin privileges, the JWT token must belong to an admin user.**
- **All responses are in JSON format.**

---

###For more details on each field, refer to the DTO definitions in the codebase or contact the backend Engineer (Seun Adeniyi).