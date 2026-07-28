---
title: Get URL scan's DOM
page_id: operation-get-accounts-account-id-urlscanner-v2-dom-scan-id-cd597ce8
path: operations/url-scanner
description: Returns a plain text response, with the scan's DOM content as rendered by Chrome.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/urlscanner/v2/dom/{scan_id}
operation_ids:
    - urlscanner-get-scan-dom-v2
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get URL scan's DOM

`GET /accounts/{account_id}/urlscanner/v2/dom/{scan_id}`

Operation ID: `urlscanner-get-scan-dom-v2`

Returns a plain text response, with the scan's DOM content as rendered by Chrome.

## Definition

```yaml
{"operationId": "urlscanner-get-scan-dom-v2", "summary": "Get URL scan's DOM", "description": "Returns a plain text response, with the scan's DOM content as rendered by Chrome.", "parameters": [{"name": "scan_id", "in": "path", "description": "Scan UUID.", "required": true, "schema": {"description": "Scan UUID.", "type": "string", "format": "uuid"}}, {"name": "account_id", "in": "path", "description": "Account ID.", "required": true, "schema": {"description": "Account ID.", "type": "string"}}], "responses": {"200": {"description": "Returns a plain text response, with the scan's DOM content as rendered by Chrome.", "content": {"text/plain": {"schema": {"description": "HTML of webpage.", "type": "string"}}}}, "400": {"description": "Invalid input.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"detail": {"type": "string"}, "status": {"description": "Status code.", "type": "integer", "example": 400}, "title": {"type": "string", "example": "Invalid url"}}, "required": ["title", "detail", "status"], "type": "object"}}, "message": {"type": "string"}, "status": {"description": "Status code.", "type": "integer", "example": 400}}, "required": ["message", "status", "errors"]}}}}, "404": {"description": "Scan not found or in progress.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"detail": {"type": "string", "example": "Queued"}, "status": {"description": "Status code.", "type": "integer", "example": 404}, "title": {"type": "string", "example": "Scan is not finished yet."}}, "required": ["title", "detail", "status"], "type": "object"}}, "message": {"description": "Scan not found or in progress.", "type": "string", "example": "Scan is not finished yet."}, "status": {"description": "Status code.", "type": "integer", "example": 404}, "task": {"type": "object", "properties": {"status": {"type": "string", "example": "Queued"}, "time": {"type": "string"}, "url": {"type": "string"}, "uuid": {"type": "string"}, "visibility": {"type": "string", "example": "public"}}, "required": ["uuid", "url", "time", "status", "visibility"]}}, "required": ["message", "status", "errors", "task"]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["URL Scanner"], "x-api-token-group": ["URL Scanner Write", "URL Scanner Read"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "url-scanner.scans", "x-fern-sdk-method-name": "dom", "x-forge-hidden": true}
```
