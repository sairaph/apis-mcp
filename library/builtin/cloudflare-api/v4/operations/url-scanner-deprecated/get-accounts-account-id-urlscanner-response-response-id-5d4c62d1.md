---
title: Get raw response
page_id: operation-get-accounts-account-id-urlscanner-response-response-id-4f4f93b7
path: operations/url-scanner-deprecated
description: Returns the plain response of the network request.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/urlscanner/response/{response_id}
operation_ids:
    - urlscanner-get-response-text
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get raw response

`GET /accounts/{account_id}/urlscanner/response/{response_id}`

Operation ID: `urlscanner-get-response-text`

Returns the plain response of the network request.

## Definition

```yaml
{"operationId": "urlscanner-get-response-text", "summary": "Get raw response", "description": "Returns the plain response of the network request.", "parameters": [{"name": "response_id", "in": "path", "description": "Response hash.", "required": true, "schema": {"description": "Response hash.", "type": "string"}}, {"name": "account_id", "in": "path", "description": "Account ID.", "required": true, "schema": {"description": "Account ID.", "type": "string"}}], "responses": {"200": {"description": "When `har.log.entries[].response._cf.contentAvailable` is `true`, use `response._cf.hash` value to fetch the raw response.", "content": {"text/plain": {"schema": {"description": "Web resource text/image.", "type": "string"}}}}, "400": {"description": "Invalid params.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"message": {"type": "string", "example": "Scan ID is not a valid UUID."}}, "required": ["message"], "type": "object"}}, "messages": {"type": "array", "items": {"properties": {"message": {"type": "string"}}, "required": ["message"], "type": "object"}}, "success": {"description": "Whether request was successful or not", "type": "boolean"}}, "required": ["messages", "errors", "success"]}}}}, "404": {"description": "Scan not found.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"message": {"type": "string", "example": "Scan not found."}}, "required": ["message"], "type": "object"}}, "messages": {"type": "array", "items": {"properties": {"message": {"type": "string"}}, "required": ["message"], "type": "object"}}, "success": {"description": "Whether request was successful or not", "type": "boolean"}}, "required": ["messages", "errors", "success"]}}}}}, "deprecated": true, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["URL Scanner (Deprecated)"], "x-api-token-group": ["URL Scanner Write", "URL Scanner Read"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "deprecated", "x-fern-sdk-group-name": "url-scanner.response", "x-fern-sdk-method-name": "get", "x-forge-hidden": true, "x-forge-sunset": {"date": "2030-01-01T00:00:00Z"}, "x-stainless-deprecation-message": "Use [V2](https://developers.cloudflare.com/api/resources/url_scanner/subresources/responses/methods/get/) instead."}
```
