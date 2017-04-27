
## signup + json
curl -X POST "http://127.0.0.1:8080/signup" -H  "accept: application/json" -H  "content-type: application/json" -d "{  \"username\": \"demo\",  \"password\": \"demo123\",  \"name\": \"demo\",  \"email\": \"demo@admin.com\"}"

## login + formData
curl -X POST "http://127.0.0.1:8080/login" -H  "accept: application/json" -H  "content-type: application/x-www-form-urlencoded" -d "email=demo%40admin.com&password=demo123"
eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhbGxvd2VkSVAiOiIxMjcuMC4wLjEiLCJhbGxvd2VkVXNlckFnZW50IjoiTW96aWxsYS81LjAgKFgxMTsgVWJ1bnR1OyBMaW51eCB4ODZfNjQ7IHJ2OjUzLjApIEdlY2tvLzIwMTAwMTAxIEZpcmVmb3gvNTMuMCIsImV4cCI6MTQ5MjkzNjI2OSwidXNlcklEIjo5fQ.fdsTRY90ORi44EnIjzhiwe7j3U65AqWK7Ov_C4BvAd8",
    "id": 9,
## token
curl -X POST "http://127.0.0.1:8080/token" -H  "accept: application/json" -H  "content-type: application/json" -d "{  \"refresh_token\": \"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhbGxvd2VkSVAiOiIxMjcuMC4wLjEiLCJhbGxvd2VkVXNlckFnZW50IjoiTW96aWxsYS81LjAgKFgxMTsgVWJ1bnR1OyBMaW51eCB4ODZfNjQ7IHJ2OjUzLjApIEdlY2tvLzIwMTAwMTAxIEZpcmVmb3gvNTMuMCIsImV4cCI6MTQ5MjkzNjI2OSwidXNlcklEIjo5fQ.fdsTRY90ORi44EnIjzhiwe7j3U65AqWK7Ov_C4BvAd8\"}"

Response body
{
"auth_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhbGxvd2VkSVAiOiIxMjcuMC4wLjEiLCJhbGxvd2VkVXNlckFnZW50IjoiTW96aWxsYS81LjAgKFgxMTsgVWJ1bnR1OyBMaW51eCB4ODZfNjQ7IHJ2OjUzLjApIEdlY2tvLzIwMTAwMTAxIEZpcmVmb3gvNTMuMCIsImV4cCI6MTQ5Mjg1MDg2MywidXNlcklEIjo5fQ.UES-PhPY5cPbaDRMP6t5Z67M-JCQhYbAvEDkLgSpeRE"
}
