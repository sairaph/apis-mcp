---
title: Get screenshot
page_id: operation-get-accounts-account-id-urlscanner-scan-scan-id-screenshot-85cea870
path: operations/url-scanner-deprecated
description: Get scan's screenshot by resolution (desktop/mobile/tablet).
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/urlscanner/scan/{scan_id}/screenshot
operation_ids:
    - urlscanner-get-scan-screenshot
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get screenshot

`GET /accounts/{account_id}/urlscanner/scan/{scan_id}/screenshot`

Operation ID: `urlscanner-get-scan-screenshot`

Get scan's screenshot by resolution (desktop/mobile/tablet).

## Definition

```yaml
{"operationId": "urlscanner-get-scan-screenshot", "summary": "Get screenshot", "description": "Get scan's screenshot by resolution (desktop/mobile/tablet).", "parameters": [{"name": "scan_id", "in": "path", "description": "Scan UUID.", "required": true, "schema": {"description": "Scan UUID.", "type": "string", "format": "uuid"}}, {"name": "account_id", "in": "path", "description": "Account ID.", "required": true, "schema": {"description": "Account ID.", "type": "string"}}, {"name": "resolution", "in": "query", "description": "Target device type.", "schema": {"description": "Target device type.", "type": "string", "default": "desktop", "enum": ["desktop", "mobile", "tablet"]}}], "responses": {"200": {"description": "Returns the scan's requested screenshot.", "content": {"image/png": {"schema": {"description": "PNG Image.", "type": "string"}}}}, "202": {"description": "Scan is in progress. Check current status in `result.scan.task.status`. Possible statuses: `Queued`,`InProgress`,`InPostProcessing`,`Finished`.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"message": {"type": "string"}}, "required": ["message"], "type": "object"}}, "messages": {"type": "array", "items": {"properties": {"message": {"type": "string", "example": "In Progress"}}, "required": ["message"], "type": "object"}}, "result": {"type": "object", "properties": {"scan": {"type": "object", "properties": {"task": {"type": "object", "properties": {"effectiveUrl": {"type": "string", "example": "http://example.com/"}, "errors": {"type": "array", "items": {"properties": {"message": {"type": "string"}}, "required": ["message"], "type": "object"}}, "location": {"type": "string", "example": "PT"}, "region": {"type": "string", "example": "enam"}, "status": {"type": "string", "example": "InProgress"}, "success": {"type": "boolean", "example": true}, "time": {"type": "string", "example": "2023-05-03T17:05:04.843Z"}, "url": {"type": "string", "example": "http://example.com"}, "uuid": {"type": "string", "example": "2ee568d0-bf70-4827-b922-b7088c0f056f"}, "visibility": {"type": "string", "example": "Public"}}, "required": ["uuid", "url", "status", "success", "errors", "time", "visibility", "location", "region", "effectiveUrl"]}}, "required": ["task"]}}, "required": ["scan"]}, "success": {"description": "Whether request was successful or not", "type": "boolean"}}, "required": ["messages", "errors", "success", "result"]}}}}, "400": {"description": "Invalid params.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"message": {"type": "string", "example": "Scan ID is not a valid UUID."}}, "required": ["message"], "type": "object"}}, "messages": {"type": "array", "items": {"properties": {"message": {"type": "string"}}, "required": ["message"], "type": "object"}}, "success": {"description": "Whether request was successful or not", "type": "boolean"}}, "required": ["messages", "errors", "success"]}}}}, "404": {"description": "Scan not found.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"message": {"type": "string", "example": "Scan not found."}}, "required": ["message"], "type": "object"}}, "messages": {"type": "array", "items": {"properties": {"message": {"type": "string"}}, "required": ["message"], "type": "object"}}, "success": {"description": "Whether request was successful or not", "type": "boolean"}}, "required": ["messages", "errors", "success"]}}}}}, "deprecated": true, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["URL Scanner (Deprecated)"], "x-api-token-group": ["URL Scanner Write", "URL Scanner Read"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "deprecated", "x-fern-sdk-group-name": "url-scanner.scan.screenshot", "x-fern-sdk-method-name": "get", "x-forge-hidden": true, "x-forge-sunset": {"date": "2030-01-01T00:00:00Z"}, "x-stainless-deprecation-message": "Use [V2](https://developers.cloudflare.com/api/resources/url_scanner/subresources/scans/methods/screenshot/) instead."}
```
