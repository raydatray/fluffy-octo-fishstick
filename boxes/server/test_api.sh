#!/bin/bash

# API Test Script for User CRUD Operations
# Make sure the server is running on port 8080 before executing this script

echo "=== Testing User CRUD API ==="
echo

# Test 1: Health Check
echo "1. Testing health endpoint..."
curl -X GET http://localhost:8080/health
echo -e "\n"

# Test 2: Create a new user (Professor)
echo "2. Creating a new professor..."
curl -X POST http://localhost:8080/api/users \
  -H "Content-Type: application/json" \
  -d '{
    "email": "professor@university.edu",
    "password": "securepassword123",
    "role": "PROFESSOR"
  }'
echo -e "\n"

# Test 3: Create a new user (Student) - default role
echo "3. Creating a new student (default role)..."
curl -X POST http://localhost:8080/api/users \
  -H "Content-Type: application/json" \
  -d '{
    "email": "student@university.edu",
    "password": "studentpass456"
  }'
echo -e "\n"

# Test 4: Create a new user (TA)
echo "4. Creating a new TA..."
curl -X POST http://localhost:8080/api/users \
  -H "Content-Type: application/json" \
  -d '{
    "email": "ta@university.edu",
    "password": "tapassword789",
    "role": "TA"
  }'
echo -e "\n"

# Test 5: Get all users
echo "5. Getting all users..."
curl -X GET http://localhost:8080/api/users
echo -e "\n"

# Test 6: Get user by ID
echo "6. Getting user by ID (ID: 1)..."
curl -X GET http://localhost:8080/api/users/1
echo -e "\n"

# Test 7: Update user
echo "7. Updating user (ID: 2) - changing email..."
curl -X PUT http://localhost:8080/api/users/2 \
  -H "Content-Type: application/json" \
  -d '{
    "email": "updated_student@university.edu"
  }'
echo -e "\n"

# Test 8: Update user role
echo "8. Updating user (ID: 2) - changing role to TA..."
curl -X PUT http://localhost:8080/api/users/2 \
  -H "Content-Type: application/json" \
  -d '{
    "role": "TA"
  }'
echo -e "\n"

# Test 9: Try to create user with duplicate email (should fail)
echo "9. Testing duplicate email (should return error)..."
curl -X POST http://localhost:8080/api/users \
  -H "Content-Type: application/json" \
  -d '{
    "email": "professor@university.edu",
    "password": "anotherpassword",
    "role": "STUDENT"
  }'
echo -e "\n"

# Test 10: Try to create user with invalid role (should fail)
echo "10. Testing invalid role (should return error)..."
curl -X POST http://localhost:8080/api/users \
  -H "Content-Type: application/json" \
  -d '{
    "email": "invalid@university.edu",
    "password": "password123",
    "role": "INVALID_ROLE"
  }'
echo -e "\n"

# Test 11: Get non-existent user (should return 404)
echo "11. Testing non-existent user (should return 404)..."
curl -X GET http://localhost:8080/api/users/999
echo -e "\n"

# Test 12: Delete a user
echo "12. Deleting user (ID: 3)..."
curl -X DELETE http://localhost:8080/api/users/3
echo -e "\n"

# Test 13: Verify deletion by getting all users
echo "13. Getting all users after deletion..."
curl -X GET http://localhost:8080/api/users
echo -e "\n"

# Test 14: Try to delete non-existent user (should return 404)
echo "14. Testing deletion of non-existent user (should return 404)..."
curl -X DELETE http://localhost:8080/api/users/999
echo -e "\n"

echo "=== API Testing Complete ==="
