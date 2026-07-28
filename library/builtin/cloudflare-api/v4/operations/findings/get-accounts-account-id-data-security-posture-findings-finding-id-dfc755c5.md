---
title: Get a finding type
page_id: operation-get-accounts-account-id-data-security-posture-findings-finding-id-f01cd68e
path: operations/findings
description: Gets a security Finding that has been identified as being problematic.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/data-security/posture/findings/{finding_id}
operation_ids:
    - GetFinding
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get a finding type

`GET /accounts/{account_id}/data-security/posture/findings/{finding_id}`

Operation ID: `GetFinding`

Gets a security Finding that has been identified as being problematic.

## Definition

```yaml
{"operationId": "GetFinding", "summary": "Get a finding type", "description": "Gets a security Finding that has been identified as being problematic.", "parameters": [{"$ref": "#/components/parameters/posture-api_AccountTag"}, {"$ref": "#/components/parameters/posture-api_FindingIdByte"}], "responses": {"200": {"description": "OK: Successful HTTP request", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/posture-api_finding-response"}}}}, "400": {"description": "Bad Request: Invalid finding ID", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/posture-api_error-response"}}}}, "401": {"description": "Unauthorized: Authentication required", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/posture-api_error-response"}}}}, "403": {"description": "Forbidden: Insufficient permissions", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/posture-api_error-response"}}}}, "404": {"description": "Not Found: Finding not found", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/posture-api_error-response"}}}}}, "security": [{"api_token": []}], "tags": ["findings"], "x-api-token-group": ["Zero Trust Read", "Zero Trust Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "data-security.posture", "x-fern-sdk-method-name": "get", "x-forge-hidden": true, "x-stability": "beta"}
```
