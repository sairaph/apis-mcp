---
title: Get a single export job
page_id: operation-get-accounts-account-id-data-security-posture-exports-id-cdd7fc6b
path: operations/exports
description: Retrieves a single export job by its unique identifier
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/data-security/posture/exports/{id}
operation_ids:
    - GetExportJob
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get a single export job

`GET /accounts/{account_id}/data-security/posture/exports/{id}`

Operation ID: `GetExportJob`

Retrieves a single export job by its unique identifier

## Definition

```yaml
{"operationId": "GetExportJob", "summary": "Get a single export job", "description": "Retrieves a single export job by its unique identifier", "parameters": [{"$ref": "#/components/parameters/posture-api_AccountTag"}, {"name": "id", "in": "path", "description": "Unique identifier for the export job", "required": true, "schema": {"type": "string", "pattern": "^[\\da-f]{8}-([\\da-f]{4}-){3}[\\da-f]{12}$"}}], "responses": {"200": {"description": "OK: Successful HTTP request", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/posture-api_export-job-response"}}}}, "400": {"description": "Bad Request: Invalid request parameters"}, "404": {"description": "Not Found: Export job not found"}}, "security": [{"api_token": []}], "tags": ["exports"], "x-api-token-group": ["Zero Trust Read", "Zero Trust Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "data-security.posture", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
