---
title: Get raw response
page_id: operation-get-accounts-account-id-urlscanner-v2-responses-response-id-c37316c1
path: operations/url-scanner
description: Returns the raw response of the network request. Find the `response_id` in the `data.requests.response.hash`.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/urlscanner/v2/responses/{response_id}
operation_ids:
    - urlscanner-get-response-v2
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get raw response

`GET /accounts/{account_id}/urlscanner/v2/responses/{response_id}`

Operation ID: `urlscanner-get-response-v2`

Returns the raw response of the network request. Find the `response_id` in the `data.requests.response.hash`.

## Definition

```yaml
{"operationId": "urlscanner-get-response-v2", "summary": "Get raw response", "description": "Returns the raw response of the network request. Find the `response_id` in the `data.requests.response.hash`.", "parameters": [{"name": "response_id", "in": "path", "description": "Response hash.", "required": true, "schema": {"description": "Response hash.", "type": "string"}}, {"name": "account_id", "in": "path", "description": "Account ID.", "required": true, "schema": {"description": "Account ID.", "type": "string"}}], "responses": {"200": {"description": "Get the raw response by its hash.", "content": {"text/plain": {"schema": {"description": "Web resource or image.", "type": "string"}}}}, "400": {"description": "Invalid input.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"detail": {"type": "string"}, "status": {"description": "Status code.", "type": "integer", "example": 400}, "title": {"type": "string", "example": "Invalid url"}}, "required": ["title", "detail", "status"], "type": "object"}}, "message": {"type": "string"}, "status": {"description": "Status code.", "type": "integer", "example": 400}}, "required": ["message", "status", "errors"]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["URL Scanner"], "x-api-token-group": ["URL Scanner Write", "URL Scanner Read"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "url-scanner.responses", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
