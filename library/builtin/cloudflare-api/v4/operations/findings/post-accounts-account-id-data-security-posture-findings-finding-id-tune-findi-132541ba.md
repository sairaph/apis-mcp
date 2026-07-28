---
title: Update the severity for a finding
page_id: operation-post-accounts-account-id-data-security-posture-findings-finding-id-tune-93c44315
path: operations/findings
description: |-
    Update the severity of a Finding.
    This will update the `severity_override` field on the Finding payload with the new severity value.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/data-security/posture/findings/{finding_id}/tune_finding_severity
operation_ids:
    - ChangeFindingSeverity
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update the severity for a finding

`POST /accounts/{account_id}/data-security/posture/findings/{finding_id}/tune_finding_severity`

Operation ID: `ChangeFindingSeverity`

Update the severity of a Finding.
This will update the `severity_override` field on the Finding payload with the new severity value.

## Definition

```yaml
{"operationId": "ChangeFindingSeverity", "summary": "Update the severity for a finding", "description": "Update the severity of a Finding.\nThis will update the `severity_override` field on the Finding payload with the new severity value.", "parameters": [{"$ref": "#/components/parameters/posture-api_AccountTag"}, {"$ref": "#/components/parameters/posture-api_FindingIdByte"}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/posture-api_TuneFindingSeverityRequest"}}}}, "responses": {"200": {"description": "OK: Successful HTTP request", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/posture-api_finding-response"}}}}, "400": {"description": "Bad Request: Invalid request parameters"}, "404": {"description": "Not Found: Finding not found"}, "500": {"description": "Internal Server Error: Unexpected failure updating finding severity", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/posture-api_error-response"}}}}}, "security": [{"api_token": []}], "tags": ["findings"], "x-api-token-group": ["Zero Trust Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "data-security.posture", "x-fern-sdk-method-name": "create", "x-forge-hidden": true, "x-stability": "beta"}
```
