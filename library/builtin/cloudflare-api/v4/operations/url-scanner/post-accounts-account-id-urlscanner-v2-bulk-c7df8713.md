---
title: Bulk create URL Scans
page_id: operation-post-accounts-account-id-urlscanner-v2-bulk-282ad5e7
path: operations/url-scanner
description: Submit URLs to scan. Check limits at https://developers.cloudflare.com/security-center/investigate/scan-limits/ and take into account scans submitted in bulk have lower priority and may take longer to finish.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/urlscanner/v2/bulk
operation_ids:
    - urlscanner-create-scan-bulk-v2
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Bulk create URL Scans

`POST /accounts/{account_id}/urlscanner/v2/bulk`

Operation ID: `urlscanner-create-scan-bulk-v2`

Submit URLs to scan. Check limits at https://developers.cloudflare.com/security-center/investigate/scan-limits/ and take into account scans submitted in bulk have lower priority and may take longer to finish.

## Definition

```yaml
{"operationId": "urlscanner-create-scan-bulk-v2", "summary": "Bulk create URL Scans", "description": "Submit URLs to scan. Check limits at https://developers.cloudflare.com/security-center/investigate/scan-limits/ and take into account scans submitted in bulk have lower priority and may take longer to finish.", "parameters": [{"name": "account_id", "in": "path", "description": "Account ID.", "required": true, "schema": {"description": "Account ID.", "type": "string"}}], "requestBody": {"content": {"application/json": {"schema": {"description": "List of urls to scan (up to a 100).", "type": "array", "items": {"properties": {"agentReadiness": {"description": "Enable agent readiness checks.", "type": "boolean"}, "customHeaders": {"description": "Set custom headers.", "type": "object", "additionalProperties": {"type": "string"}}, "customagent": {"type": "string", "maxLength": 4096}, "referer": {"type": "string", "maxLength": 4096}, "screenshotsResolutions": {"description": "Take multiple screenshots targeting different device types.", "type": "array", "items": {"default": "desktop", "description": "Device resolutions.", "enum": ["desktop", "mobile", "tablet"], "type": "string"}, "default": ["desktop"]}, "url": {"type": "string", "example": "https://www.example.com"}, "visibility": {"description": "The option `Public` means it will be included in listings like recent scans and search results. `Unlisted` means it will not be included in the aforementioned listings, users will need to have the scan's ID to access it. A a scan will be automatically marked as unlisted if it fails, if it contains potential PII or other sensitive material.", "type": "string", "default": "Public", "enum": ["Public", "Unlisted"]}}, "required": ["url"], "type": "object"}}}}}, "responses": {"200": {"description": "Scan bulk request accepted successfully.", "content": {"application/json": {"schema": {"type": "array", "items": {"properties": {"api": {"description": "URL to api report.", "type": "string"}, "options": {"type": "object", "properties": {"useragent": {"type": "string"}}}, "result": {"description": "URL to report.", "type": "string"}, "url": {"description": "Submitted URL", "type": "string"}, "uuid": {"description": "Scan ID.", "type": "string", "format": "uuid"}, "visibility": {"description": "Submitted visibility status.", "type": "string", "example": "public", "enum": ["public", "unlisted"]}}, "required": ["uuid", "result", "api", "visibility", "url"], "type": "object"}}}}}, "400": {"description": "Invalid input.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"detail": {"type": "string"}, "status": {"description": "Status code.", "type": "integer", "example": 400}, "title": {"type": "string", "example": "Invalid url"}}, "required": ["title", "detail", "status"], "type": "object"}}, "message": {"type": "string"}, "status": {"description": "Status code.", "type": "integer", "example": 400}}, "required": ["message", "status", "errors"]}}}}, "429": {"description": "Scan request denied: rate limited.", "content": {"application/json": {"schema": {"type": "object", "properties": {"description": {"type": "string"}, "errors": {"type": "array", "items": {"properties": {"detail": {"type": "string", "example": "DNS Error - Could not resolve domain."}, "status": {"type": "number", "example": 400}, "title": {"type": "string", "example": "DNS Error - Could not resolve domain."}}, "required": ["title", "detail", "status"], "type": "object"}}, "message": {"type": "string"}, "status": {"type": "number", "example": 429}}, "required": ["message", "status", "errors"]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["URL Scanner"], "x-api-token-group": ["URL Scanner Write", "URL Scanner Read"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "url-scanner.scans", "x-fern-sdk-method-name": "bulk-create", "x-forge-hidden": true}
```
