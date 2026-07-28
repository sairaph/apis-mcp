---
title: List certificate logs
page_id: operation-get-radar-ct-logs-1f047da0
path: operations/radar-certificate-transparency
description: Retrieves a list of certificate logs.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /radar/ct/logs
operation_ids:
    - radar-get-certificate-logs
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List certificate logs

`GET /radar/ct/logs`

Operation ID: `radar-get-certificate-logs`

Retrieves a list of certificate logs.

## Definition

```yaml
{"operationId": "radar-get-certificate-logs", "summary": "List certificate logs", "description": "Retrieves a list of certificate logs.", "parameters": [{"name": "limit", "in": "query", "description": "Limits the number of objects returned in the response.", "schema": {"description": "Limits the number of objects returned in the response.", "type": "integer", "example": 5, "default": 5, "exclusiveMinimum": true, "minimum": 0}}, {"name": "offset", "in": "query", "description": "Skips the specified number of objects before fetching the results.", "schema": {"description": "Skips the specified number of objects before fetching the results.", "type": "integer", "minimum": 0}}, {"name": "format", "in": "query", "description": "Format in which results will be returned.", "schema": {"description": "Format in which results will be returned.", "type": "string", "example": "json", "enum": ["JSON", "CSV"]}}], "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"type": "object", "properties": {"result": {"type": "object", "properties": {"certificateLogs": {"type": "array", "items": {"properties": {"api": {"description": "The API standard that the certificate log follows.", "type": "string", "enum": ["RFC6962", "STATIC"]}, "description": {"description": "A brief description of the certificate log.", "type": "string", "example": "Google 'Argon2024' log"}, "endExclusive": {"description": "The end date and time for when the log will stop accepting certificates.", "type": "string", "format": "date-time", "example": "2025-01-01T00:00:00Z"}, "operator": {"description": "The organization responsible for operating the certificate log.", "type": "string", "example": "Google"}, "slug": {"description": "A URL-friendly, kebab-case identifier for the certificate log.", "type": "string", "example": "argon2024"}, "startInclusive": {"description": "The start date and time for when the log starts accepting certificates.", "type": "string", "format": "date-time", "example": "2024-01-01T00:00:00Z"}, "state": {"description": "The current state of the certificate log. More details about log states can be found here: https://googlechrome.github.io/CertificateTransparency/log_states.html", "type": "string", "enum": ["USABLE", "PENDING", "QUALIFIED", "READ_ONLY", "RETIRED", "REJECTED"]}, "stateTimestamp": {"description": "Timestamp of when the log state was last updated.", "type": "string", "format": "date-time", "example": "2025-02-01T08:53:20Z"}, "url": {"description": "The URL for the certificate log.", "type": "string", "example": "https://ct.googleapis.com/logs/us1/argon2024/"}}, "required": ["slug", "description", "operator", "api", "state", "stateTimestamp", "startInclusive", "endExclusive", "url"], "type": "object"}}}, "required": ["certificateLogs"]}, "success": {"type": "boolean", "example": true}}, "required": ["result", "success"]}}}}, "400": {"description": "Bad request.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"message": {"type": "string"}}, "required": ["message"], "type": "object"}}, "result": {"type": "object"}, "success": {"type": "boolean", "example": false}}, "required": ["result", "success", "errors"]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Radar Certificate Transparency"], "x-api-token-group": ["User Details Write", "User Details Read"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "radar.ct.logs", "x-fern-sdk-method-name": "list", "x-forge-hidden": true}
```
