# Reward System - Complete API Design

## Base URL
```
https://api.4sale-rewards.com/api/v1
```

## Authentication
All protected endpoints require JWT token in the Authorization header:
```
Authorization: Bearer <jwt_token>
```

---

## 1. User Authentication Routes

### 1.1 User Registration
**`POST /users/signup`**

**Request Body:**
```json
{
  "first_name": "John",
  "last_name": "Doe", 
  "username": "johndoe",
  "email": "john@example.com",
  "password": "SecurePass123!"
}
```

**Response:**
- `201 Created`: User successfully created
```json
{
  "message": "User created successfully",
  "user": {
    "id": 123,
    "first_name": "John",
    "last_name": "Doe",
    "username": "johndoe",
    "email": "john@example.com",
    "created_at": "2025-06-26T10:30:00Z"
  }
}
```
- `400 Bad Request`: Invalid request body
- `409 Conflict`: Username or email already exists
- `500 Internal Server Error`

### 1.2 User Login
**`POST /users/login`**

**Request Body:**
```json
{
  "email": "john@example.com",
  "password": "SecurePass123!"
}
```

**Response:**
- `200 OK`: Login successful
```json
{
  "message": "Login successful",
  "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
  "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
  "user": {
    "id": 123,
    "first_name": "John",
    "last_name": "Doe",
    "email": "john@example.com"
  }
}
```
- `401 Unauthorized`: Invalid credentials
- `500 Internal Server Error`

### 1.3 Forgot Password
**`POST /users/forgot-password`**

**Request Body:**
```json
{
  "email": "john@example.com"
}
```

**Response:**
- `200 OK`: Password reset email sent
```json
{
  "message": "Password reset instructions sent to your email"
}
```
- `404 Not Found`: Email not found
- `500 Internal Server Error`

### 1.4 Reset Password
**`POST /users/reset-password/:token`**

**Request Body:**
```json
{
  "password": "NewSecurePass123!",
  "confirm_password": "NewSecurePass123!"
}
```

**Response:**
- `200 OK`: Password reset successful
- `400 Bad Request`: Invalid token or passwords don't match
- `401 Unauthorized`: Token expired
- `500 Internal Server Error`

### 1.5 Change Password
**`PUT /users/change-password`** *(Protected)*

**Request Body:**
```json
{
  "current_password": "OldPass123!",
  "new_password": "NewSecurePass123!",
  "confirm_password": "NewSecurePass123!"
}
```

**Response:**
- `200 OK`: Password changed successfully
- `400 Bad Request`: Invalid request body
- `401 Unauthorized`: Current password incorrect
- `500 Internal Server Error`

---

## 2. User Profile Routes

### 2.1 Get User Profile
**`GET /users/profile`** *(Protected)*

**Response:**
- `200 OK`:
```json
{
  "user": {
    "id": 123,
    "first_name": "John",
    "last_name": "Doe",
    "username": "johndoe",
    "email": "john@example.com",
    "created_at": "2025-06-26T10:30:00Z"
  }
}
```

### 2.2 Update User Profile
**`PUT /users/profile`** *(Protected)*

**Request Body:**
```json
{
  "first_name": "John",
  "last_name": "Smith",
  "username": "johnsmith"
}
```

**Response:**
- `200 OK`: Profile updated successfully
- `400 Bad Request`: Invalid data
- `409 Conflict`: Username already exists

---

## 3. Credit Package Routes

### 3.1 Get All Credit Packages
**`GET /credit-packages`**

**Query Parameters:**
- `page` (optional): Page number (default: 1)
- `limit` (optional): Items per page (default: 20)
- `active` (optional): Filter by active status (true/false)

**Response:**
- `200 OK`:
```json
{
  "packages": [
    {
      "id": 1,
      "name": "Basic Package",
      "price": 100.00,
      "credits": 100,
      "reward_points": 10,
      "is_active": true,
      "created_at": "2025-06-26T10:30:00Z"
    },
    {
      "id": 2,
      "name": "Standard Package",
      "price": 500.00,
      "credits": 550,
      "reward_points": 75,
      "is_active": true,
      "created_at": "2025-06-26T10:30:00Z"
    }
  ],
  "pagination": {
    "current_page": 1,
    "total_pages": 5,
    "total_items": 100,
    "items_per_page": 20
  }
}
```

### 3.2 Get Credit Package by ID
**`GET /credit-packages/:id`**

**Response:**
- `200 OK`: Package details
- `404 Not Found`: Package not found

### 3.3 Create Credit Package
**`POST /credit-packages`** *(Admin Only)*

**Request Body:**
```json
{
  "name": "Premium Package",
  "price": 1000.00,
  "credits": 1200,
  "reward_points": 200,
  "is_active": true
}
```

**Response:**
- `201 Created`: Package created successfully
- `400 Bad Request`: Invalid data
- `403 Forbidden`: Admin access required

### 3.4 Update Credit Package
**`PUT /credit-packages/:id`** *(Admin Only)*

**Request Body:**
```json
{
  "name": "Premium Package Updated",
  "price": 1000.00,
  "credits": 1200,
  "reward_points": 200,
  "is_active": true
}
```

**Response:**
- `200 OK`: Package updated successfully
- `400 Bad Request`: Invalid data
- `403 Forbidden`: Admin access required
- `404 Not Found`: Package not found

### 3.5 Delete Credit Package
**`DELETE /credit-packages/:id`** *(Admin Only)*

**Response:**
- `204 No Content`: Package deleted successfully
- `403 Forbidden`: Admin access required
- `404 Not Found`: Package not found

---

## 4. Purchase Routes

### 4.1 Create Purchase
**`POST /purchases`** *(Protected)*

**Request Body:**
```json
{
  "credit_package_id": 2,
  "payment_method": "credit_card",
  "payment_details": {
    "card_number": "1234567890123456",
    "expiry_date": "12/25",
    "cvv": "123"
  }
}
```

**Response:**
- `201 Created`: Purchase created successfully
```json
{
  "purchase": {
    "id": 456,
    "user_id": 123,
    "credit_package_id": 2,
    "amount": 500.00,
    "credits_earned": 550,
    "reward_points_earned": 75,
    "status": "completed",
    "created_at": "2025-06-26T10:30:00Z"
  }
}
```
- `400 Bad Request`: Invalid data
- `402 Payment Required`: Payment failed
- `404 Not Found`: Package not found

### 4.2 Get User Purchases
**`GET /purchases`** *(Protected)*

**Query Parameters:**
- `page` (optional): Page number (default: 1)
- `limit` (optional): Items per page (default: 20)
- `status` (optional): Filter by status

**Response:**
- `200 OK`:
```json
{
  "purchases": [
    {
      "id": 456,
      "credit_package": {
        "id": 2,
        "name": "Standard Package",
        "price": 500.00
      },
      "amount": 500.00,
      "credits_earned": 550,
      "reward_points_earned": 75,
      "status": "completed",
      "created_at": "2025-06-26T10:30:00Z"
    }
  ],
  "pagination": {
    "current_page": 1,
    "total_pages": 3,
    "total_items": 24,
    "items_per_page": 20
  }
}
```

### 4.3 Get Purchase by ID
**`GET /purchases/:id`** *(Protected)*

**Response:**
- `200 OK`: Purchase details
- `403 Forbidden`: Not owner of purchase
- `404 Not Found`: Purchase not found

---

## 5. Wallet Routes

### 5.1 Get User Wallet Info
**`GET /wallets`** *(Protected)*

**Response:**
- `200 OK`:
```json
{
  "wallet": {
    "user_id": 123,
    "points_balance": 1230,
    "credits_balance": 2450,
    "updated_at": "2025-06-26T10:30:00Z"
  }
}
```

---

## 6. Product Routes

### 6.1 Get All Products
**`GET /products`**

**Query Parameters:**
- `page` (optional): Page number (default: 1)
- `limit` (optional): Items per page (default: 20)
- `category_id` (optional): Filter by category
- `is_offer` (optional): Filter by offer status (true/false)
- `min_points` (optional): Minimum reward points
- `max_points` (optional): Maximum reward points
- `sort_by` (optional): Sort field (name, reward_points, stock_quantity)
- `sort_order` (optional): Sort order (asc, desc)

**Response:**
- `200 OK`:
```json
{
  "products": [
    {
      "id": 1,
      "name": "iPhone 15",
      "description": "Latest iPhone model",
      "category": {
        "id": 1,
        "name": "Electronics",
        "parent_category_id": null
      },
      "reward_points": 500,
      "stock_quantity": 50,
      "is_offer": true,
      "created_at": "2025-06-26T10:30:00Z"
    }
  ],
  "pagination": {
    "current_page": 1,
    "total_pages": 10,
    "total_items": 200,
    "items_per_page": 20
  }
}
```

### 6.2 Get Product by ID
**`GET /products/:id`**

**Response:**
- `200 OK`: Product details
- `404 Not Found`: Product not found

### 6.3 Search Products
**`GET /products/search`**

**Query Parameters:**
- `query` (required): Search query
- `page` (optional): Page number (default: 1)
- `limit` (optional): Items per page (default: 20)
- `category_id` (optional): Filter by category
- `is_offer` (optional): Filter by offer status
- `min_points` (optional): Minimum reward points
- `max_points` (optional): Maximum reward points

**Response:**
- `200 OK`:
```json
{
  "products": [
    {
      "id": 1,
      "name": "iPhone 15",
      "description": "Latest iPhone model with advanced features",
      "category": {
        "id": 1,
        "name": "Electronics"
      },
      "reward_points": 500,
      "stock_quantity": 50,
      "is_offer": true,
      "relevance_score": 0.95
    }
  ],
  "search_meta": {
    "query": "iPhone",
    "total_results": 15,
    "search_time_ms": 23
  },
  "pagination": {
    "current_page": 1,
    "total_pages": 2,
    "total_items": 15,
    "items_per_page": 20
  }
}
```

### 6.4 Create Product
**`POST /products`** *(Admin Only)*

**Request Body:**
```json
{
  "name": "iPhone 15 Pro",
  "description": "Premium iPhone model",
  "category_id": 1,
  "reward_points": 500,
  "stock_quantity": 25,
  "is_offer": true
}
```

**Response:**
- `201 Created`: Product created successfully
- `400 Bad Request`: Invalid data
- `403 Forbidden`: Admin access required

### 6.5 Update Product
**`PUT /products/:id`** *(Admin Only)*

**Request Body:**
```json
{
  "name": "iPhone 15 Pro Updated",
  "description": "Premium iPhone model with new features",
  "category_id": 1,
  "reward_points": 600,
  "stock_quantity": 30,
  "is_offer": true
}
```

**Response:**
- `200 OK`: Product updated successfully
- `400 Bad Request`: Invalid data
- `403 Forbidden`: Admin access required
- `404 Not Found`: Product not found

### 6.6 Delete Product
**`DELETE /products/:id`** *(Admin Only)*

**Response:**
- `204 No Content`: Product deleted successfully
- `403 Forbidden`: Admin access required
- `404 Not Found`: Product not found

---

## 7. Category Routes

### 7.1 Get All Categories
**`GET /categories`**

**Query Parameters:**
- `parent_id` (optional): Filter by parent category
- `include_children` (optional): Include child categories (true/false)

**Response:**
- `200 OK`:
```json
{
  "categories": [
    {
      "id": 1,
      "name": "Electronics",
      "description": "Electronic devices and accessories",
      "parent_category_id": null,
      "children": [
        {
          "id": 2,
          "name": "Smartphones",
          "description": "Mobile phones and accessories",
          "parent_category_id": 1
        }
      ]
    }
  ]
}
```

### 7.2 Create Category
**`POST /categories`** *(Admin Only)*

**Request Body:**
```json
{
  "name": "Gaming",
  "description": "Gaming consoles and accessories",
  "parent_category_id": 1
}
```

**Response:**
- `201 Created`: Category created successfully
- `400 Bad Request`: Invalid data
- `403 Forbidden`: Admin access required

### 7.3 Update Category
**`PUT /categories/:id`** *(Admin Only)*

**Request Body:**
```json
{
  "name": "Gaming Updated",
  "description": "Gaming consoles, accessories, and games",
  "parent_category_id": 1
}
```

**Response:**
- `200 OK`: Category updated successfully
- `400 Bad Request`: Invalid data
- `403 Forbidden`: Admin access required
- `404 Not Found`: Category not found

### 7.4 Delete Category
**`DELETE /categories/:id`** *(Admin Only)*

**Response:**
- `204 No Content`: Category deleted successfully
- `403 Forbidden`: Admin access required
- `404 Not Found`: Category not found

---

## 8. Redemption Routes

### 8.1 Create Redemption
**`POST /redemptions`** *(Protected)*

**Request Body:**
```json
{
  "product_id": 1,
  "quantity": 1,
  "shipping_address": {
    "street": "123 Main St",
    "city": "Cairo",
    "country": "Egypt",
    "postal_code": "12345"
  }
}
```

**Response:**
- `201 Created`: Redemption created successfully
```json
{
  "redemption": {
    "id": 101,
    "user_id": 123,
    "product": {
      "id": 1,
      "name": "iPhone 15",
      "reward_points": 500
    },
    "quantity": 1,
    "points_used": 500,
    "status": "pending",
    "shipping_address": {
      "street": "123 Main St",
      "city": "Cairo",
      "country": "Egypt",
      "postal_code": "12345"
    },
    "created_at": "2025-06-26T10:30:00Z",
    "estimated_delivery": "2025-07-03T00:00:00Z"
  }
}
```
- `400 Bad Request`: Invalid data or insufficient points
- `404 Not Found`: Product not found
- `409 Conflict`: Insufficient stock

### 8.2 Get User Redemptions
**`GET /redemptions`** *(Protected)*

**Query Parameters:**
- `page` (optional): Page number (default: 1)
- `limit` (optional): Items per page (default: 20)
- `status` (optional): Filter by status

**Response:**
- `200 OK`:
```json
{
  "redemptions": [
    {
      "id": 101,
      "product": {
        "id": 1,
        "name": "iPhone 15"
      },
      "quantity": 1,
      "points_used": 500,
      "status": "delivered",
      "created_at": "2025-06-26T10:30:00Z",
      "delivered_at": "2025-07-01T14:30:00Z"
    }
  ],
  "pagination": {
    "current_page": 1,
    "total_pages": 1,
    "total_items": 8,
    "items_per_page": 20
  }
}
```

### 8.3 Get Redemption by ID
**`GET /redemptions/:id`** *(Protected)*

**Response:**
- `200 OK`: Redemption details
- `403 Forbidden`: Not owner of redemption
- `404 Not Found`: Redemption not found

### 8.4 Cancel Redemption
**`DELETE /redemptions/:id`** *(Protected)*

**Response:**
- `200 OK`: Redemption cancelled (if status allows)
- `400 Bad Request`: Cannot cancel redemption
- `403 Forbidden`: Not owner of redemption
- `404 Not Found`: Redemption not found

---

## 9. Admin Routes

### 9.1 Get Admin Dashboard Stats
**`GET /admin/dashboard`** *(Admin Only)*

**Response:**
- `200 OK`:
```json
{
  "stats": {
    "total_users": 1234,
    "total_orders": 5678,
    "credits_issued": "2.5M",
    "points_earned": "890K",
    "recent_transactions": [
      {
        "id": 1,
        "user": "john@example.com",
        "type": "purchase",
        "amount": 500.00,
        "date": "2025-06-26T10:30:00Z",
        "status": "completed"
      }
    ]
  }
}
```

### 9.2 Get All Users
**`GET /admin/users`** *(Admin Only)*

**Query Parameters:**
- `page` (optional): Page number (default: 1)
- `limit` (optional): Items per page (default: 20)
- `search` (optional): Search by name or email
- `sort_by` (optional): Sort field
- `sort_order` (optional): Sort order (asc, desc)

**Response:**
- `200 OK`: List of users with pagination

### 9.3 Get All Purchases
**`GET /admin/purchases`** *(Admin Only)*

**Query Parameters:**
- `page` (optional): Page number (default: 1)
- `limit` (optional): Items per page (default: 20)
- `status` (optional): Filter by status
- `date_from` (optional): Filter from date
- `date_to` (optional): Filter to date

**Response:**
- `200 OK`: List of purchases with pagination

### 9.4 Get All Redemptions
**`GET /admin/redemptions`** *(Admin Only)*

**Query Parameters:**
- `page` (optional): Page number (default: 1)
- `limit` (optional): Items per page (default: 20)
- `status` (optional): Filter by status
- `date_from` (optional): Filter from date
- `date_to` (optional): Filter to date

**Response:**
- `200 OK`: List of redemptions with pagination

### 9.5 Update Redemption Status
**`PUT /admin/redemptions/:id/status`** *(Admin Only)*

**Request Body:**
```json
{
  "status": "delivered",
  "notes": "Package delivered successfully"
}
```

**Response:**
- `200 OK`: Status updated successfully
- `400 Bad Request`: Invalid status
- `404 Not Found`: Redemption not found

### 9.6 Manage User Credits
**`POST /admin/users/:id/credits`** *(Admin Only)*

**Request Body:**
```json
{
  "action": "add", // or "subtract"
  "amount": 100,
  "reason": "Bonus credits for loyal customer"
}
```

**Response:**
- `200 OK`: Credits updated successfully
- `400 Bad Request`: Invalid data
- `404 Not Found`: User not found

### 9.7 Moderate Users
**`PUT /admin/users/:id/status`** *(Admin Only)*

**Request Body:**
```json
{
  "status": "suspended", // or "active", "banned"
  "reason": "Violation of terms of service"
}
```

**Response:**
- `200 OK`: User status updated
- `400 Bad Request`: Invalid status
- `404 Not Found`: User not found

---

## 10. AI Recommendation Routes

### 10.1 Get Product Recommendations
**`POST /ai/recommendations`** *(Protected)*

**Request Body:**
```json
{
  "user_preferences": {
    "categories": [1, 2, 3],
    "price_range": {
      "min_points": 100,
      "max_points": 1000
    },
    "exclude_categories": [4, 5]
  },
  "limit": 5,
  "context": "homepage" // or "search", "category", "product_page"
}
```

**Response:**
- `200 OK`:
```json
{
  "recommendations": [
    {
      "product": {
        "id": 1,
        "name": "iPhone 15",
        "description": "Latest iPhone model",
        "category": {
          "id": 1,
          "name": "Electronics"
        },
        "reward_points": 500,
        "stock_quantity": 50,
        "is_offer": true
      },
      "confidence_score": 0.95,
      "reason": "Based on your previous electronics purchases"
    }
  ],
  "recommendation_meta": {
    "algorithm_version": "v2.1",
    "generated_at": "2025-06-26T10:30:00Z",
    "user_segment": "high_value_electronics"
  }
}
```

### 10.2 Get Smart Recommendations
**`GET /ai/smart-recommendations`** *(Protected)*

**Query Parameters:**
- `context` (optional): Context for recommendations
- `limit` (optional): Number of recommendations (default: 5)

**Response:**
- `200 OK`: Personalized recommendations based on user behavior

---

## 11. Utility Routes

### 11.1 Health Check
**`GET /health`**

**Response:**
- `200 OK`:
```json
{
  "status": "healthy",
  "timestamp": "2025-06-26T10:30:00Z",
  "version": "1.0.0",
  "services": {
    "database": "healthy",
    "redis": "healthy",
    "email": "healthy"
  }
}
```

### 11.2 API Version
**`GET /version`**

**Response:**
- `200 OK`:
```json
{
  "version": "1.0.0",
  "build": "12345",
  "commit": "abc123",
  "environment": "production"
}
```

---

## Error Response Format

All error responses follow this format:

```json
{
  "error": {
    "code": "VALIDATION_ERROR",
    "message": "Validation failed",
    "details": [
      {
        "field": "email",
        "message": "Email is required"
      }
    ],
    "timestamp": "2025-06-26T10:30:00Z",
    "request_id": "req_123456789"
  }
}
```

## Common HTTP Status Codes

- `200 OK`: Request successful
- `201 Created`: Resource created successfully
- `204 No Content`: Request successful, no content to return
- `400 Bad Request`: Invalid request data
- `401 Unauthorized`: Authentication required or failed
- `403 Forbidden`: Insufficient permissions
- `404 Not Found`: Resource not found
- `409 Conflict`: Resource conflict (e.g., duplicate data)
- `422 Unprocessable Entity`: Validation errors
- `429 Too Many Requests`: Rate limit exceeded
- `500 Internal Server Error`: Server error

## Rate Limiting

- Public endpoints: 100 requests per minute per IP
- Authenticated endpoints: 1000 requests per minute per user
- Admin endpoints: 500 requests per minute per admin
- Search endpoints: 60 requests per minute per user

## Pagination

All list endpoints support pagination with these query parameters:
- `page`: Page number (default: 1)
- `limit`: Items per page (default: 20, max: 100)

Pagination response format:
```json
{
  "pagination": {
    "current_page": 1,
    "total_pages": 10,
    "total_items": 200,
    "items_per_page": 20,
    "has_next": true,
    "has_previous": false
  }
}
```

## Filtering and Sorting

Most list endpoints support filtering and sorting:
- `sort_by`: Field to sort by
- `sort_order`: `asc` or `desc`
- `filter[field]`: Filter by field value
- `search`: General search query

## Webhooks (Optional)

For real-time notifications:
- `POST /webhooks/purchase-completed`
- `POST /webhooks/redemption-status-changed`
- `POST /webhooks/user-registered`