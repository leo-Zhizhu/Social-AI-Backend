# Postman Testing Guide for SocialAI Service

## 🌐 Base URL
```
https://socialai-dot-socialai-477621.ue.r.appspot.com
```

---

## 1. Sign Up (User Registration)

**Method:** `POST`  
**URL:** `https://socialai-dot-socialai-477621.ue.r.appspot.com/signup`

**Headers:**
```
Content-Type: application/json
```

**Body (raw JSON):**
```json
{
  "username": "john_doe",
  "password": "mypassword123",
  "age": 25,
  "gender": "M"
}
```

**Expected Response:**
- Status: `200 OK` (if successful)
- Body: Empty (or success message)

**Example:**
- Username: `john_doe`
- Password: `mypassword123`
- Age: `25`
- Gender: `M` or `F`

---

## 2. Sign In (Get JWT Token)

**Method:** `POST`  
**URL:** `https://socialai-dot-socialai-477621.ue.r.appspot.com/signin`

**Headers:**
```
Content-Type: application/json
```

**Body (raw JSON):**
```json
{
  "username": "john_doe",
  "password": "mypassword123"
}
```

**Expected Response:**
- Status: `200 OK`
- Body: JWT token string (e.g., `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...`)

**⚠️ Important:** Copy this token! You'll need it for protected endpoints.

---

## 3. Upload Post (Protected - Requires JWT)

**Method:** `POST`  
**URL:** `https://socialai-dot-socialai-477621.ue.r.appspot.com/upload`

**Headers:**
```
Authorization: Bearer YOUR_JWT_TOKEN_HERE
Content-Type: multipart/form-data
```

**Body (form-data):**
- Key: `user` (Text) - Value: `john_doe` (or username from token)
- Key: `message` (Text) - Value: `This is my first post!`
- Key: `media_file` (File) - Select an image or video file

**Supported Media Types:**
- Images: `.jpeg`, `.jpg`, `.gif`, `.png`
- Videos: `.mov`, `.mp4`, `.avi`, `.flv`, `.wmv`

**Expected Response:**
- Status: `200 OK`
- Body: Success message

**Example:**
- user: `john_doe`
- message: `Check out this amazing sunset!`
- media_file: Select a file from your computer

---

## 4. Search Posts by Keywords (Protected - Requires JWT)

**Method:** `GET`  
**URL:** `https://socialai-dot-socialai-477621.ue.r.appspot.com/search?keywords=your+search+terms`

**Headers:**
```
Authorization: Bearer YOUR_JWT_TOKEN_HERE
```

**Query Parameters:**
- `keywords`: Your search terms (e.g., `sunset`, `amazing`, `hello world`)

**Example URLs:**
```
https://socialai-dot-socialai-477621.ue.r.appspot.com/search?keywords=sunset
https://socialai-dot-socialai-477621.ue.r.appspot.com/search?keywords=amazing+post
```

**Expected Response:**
- Status: `200 OK`
- Body: JSON array of posts
```json
[
  {
    "id": "uuid-here",
    "user": "john_doe",
    "message": "Check out this amazing sunset!",
    "url": "https://storage.googleapis.com/...",
    "type": "image"
  }
]
```

**Note:** If `keywords` is empty, it returns all posts.

---

## 5. Search Posts by User (Protected - Requires JWT)

**Method:** `GET`  
**URL:** `https://socialai-dot-socialai-477621.ue.r.appspot.com/search?user=username`

**Headers:**
```
Authorization: Bearer YOUR_JWT_TOKEN_HERE
```

**Query Parameters:**
- `user`: Username to search for (e.g., `john_doe`)

**Example URL:**
```
https://socialai-dot-socialai-477621.ue.r.appspot.com/search?user=john_doe
```

**Expected Response:**
- Status: `200 OK`
- Body: JSON array of posts by that user

---

## 6. Delete Post (Protected - Requires JWT)

**Method:** `DELETE`  
**URL:** `https://socialai-dot-socialai-477621.ue.r.appspot.com/post/{id}`

**Headers:**
```
Authorization: Bearer YOUR_JWT_TOKEN_HERE
```

**URL Parameters:**
- `{id}`: The post ID to delete (e.g., `35cf8466-dc7d-4991-8960-bd0fe15ee43a`)

**Example URL:**
```
https://socialai-dot-socialai-477621.ue.r.appspot.com/post/35cf8466-dc7d-4991-8960-bd0fe15ee43a
```

**Expected Response:**
- Status: `200 OK` (if successful)
- Body: Empty or success message

**Important Notes:**
- ⚠️ You can only delete posts that belong to your authenticated user
- The post ID must match a post created by the user in the JWT token
- This is a DELETE request, not GET!

---

## 7. Get Post by ID (Protected - Requires JWT)

**Method:** `GET`  
**URL:** `https://socialai-dot-socialai-477621.ue.r.appspot.com/search?user=username`

**Headers:**
```
Authorization: Bearer YOUR_JWT_TOKEN_HERE
```

**Query Parameters:**
- `user`: Username to search for (e.g., `john_doe`)

**Example URL:**
```
https://socialai-dot-socialai-477621.ue.r.appspot.com/search?user=john_doe
```

**Expected Response:**
- Status: `200 OK`
- Body: JSON array of posts by that user

---

## 📋 Step-by-Step Testing Workflow

### Step 1: Create a User
1. Use **Sign Up** endpoint
2. Create a test user (e.g., `testuser`, `password123`)

### Step 2: Get Authentication Token
1. Use **Sign In** endpoint with the same credentials
2. Copy the JWT token from the response

### Step 3: Test Protected Endpoints
1. Add the token to the `Authorization` header as `Bearer <token>`
2. Test **Upload Post** with a media file
3. Test **Search** endpoints

---

## 🔧 Postman Setup Tips

### Setting Up Authorization (Recommended)

1. Create a **Postman Environment**:
   - Variable: `base_url` = `https://socialai-dot-socialai-477621.ue.r.appspot.com`
   - Variable: `jwt_token` = (leave empty, will be set after signin)

2. Create a **Collection** with all endpoints

3. Use **Collection Variables**:
   - `{{base_url}}` for the base URL
   - `{{jwt_token}}` for the token

4. **Auto-save Token** (Optional):
   - In Sign In request, add a **Test** script:
   ```javascript
   if (pm.response.code === 200) {
       pm.environment.set("jwt_token", pm.response.text());
   }
   ```

5. **Use Token in Requests**:
   - In protected endpoints, set Authorization header to:
   ```
   Bearer {{jwt_token}}
   ```

---

## 🧪 Quick Test Examples

### Test 1: Complete Flow
```
1. POST /signup → Create user "testuser"
2. POST /signin → Get token
3. POST /upload → Upload a post with image
4. GET /search?keywords=test → Search for posts
5. GET /search?user=testuser → Get user's posts
```

### Test 2: Error Cases
```
1. POST /signin with wrong password → Should return 401
2. GET /search without token → Should return 401
3. POST /upload without token → Should return 401
```

---

## ⚠️ Common Issues

1. **401 Unauthorized**: 
   - Check if JWT token is valid
   - Make sure token is in format: `Bearer <token>`
   - Token might have expired (24 hours)

2. **400 Bad Request**:
   - Check JSON format in request body
   - Verify all required fields are present

3. **500 Internal Server Error**:
   - Check application logs: `gcloud app logs read -s socialai`
   - Verify Elasticsearch is accessible

---

## 📊 Response Status Codes

- `200 OK` - Request successful
- `400 Bad Request` - Invalid request data
- `401 Unauthorized` - Missing or invalid JWT token
- `405 Method Not Allowed` - Wrong HTTP method
- `500 Internal Server Error` - Server error

---

## 🔗 Quick Links

- **Service URL:** https://socialai-dot-socialai-477621.ue.r.appspot.com
- **View Logs:** `gcloud app logs read -s socialai`
- **Service Status:** `gcloud app services describe socialai`

