---
title: Reset severity for a finding back to the default
page_id: operation-post-accounts-account-id-data-security-posture-findings-finding-id-reset-3fa758b2
path: operations/findings
description: |-
    If a Finding's severity has been changed, reset it back to default value.
    Does nothing if no override exists.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/data-security/posture/findings/{finding_id}/reset_finding_severity
operation_ids:
    - ResetFindingSeverity
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Reset severity for a finding back to the default

`POST /accounts/{account_id}/data-security/posture/findings/{finding_id}/reset_finding_severity`

Operation ID: `ResetFindingSeverity`

If a Finding's severity has been changed, reset it back to default value.
Does nothing if no override exists.

## Definition

```yaml
{"operationId": "ResetFindingSeverity", "summary": "Reset severity for a finding back to the default", "description": "If a Finding's severity has been changed, reset it back to default value.\nDoes nothing if no override exists.", "parameters": [{"$ref": "#/components/parameters/posture-api_AccountTag"}, {"$ref": "#/components/parameters/posture-api_FindingIdByte"}], "responses": {"200": {"description": "OK: Successful HTTP request", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/posture-api_finding-response"}}}}, "400": {"description": "Bad Request: Invalid request parameters"}, "404": {"description": "Not Found: Finding not found"}, "500": {"description": "Internal Server Error: Unexpected failure resetting finding severity", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/posture-api_error-response"}}}}}, "security": [{"api_token": []}], "tags": ["findings"], "x-api-token-group": ["Zero Trust Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "data-security.posture", "x-fern-sdk-method-name": "create", "x-forge-hidden": true, "x-stability": "beta"}
```
