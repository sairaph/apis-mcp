---
title: Remove ignore marker from a finding
page_id: operation-post-accounts-account-id-data-security-posture-findings-unignore-85e74b08
path: operations/findings
description: Ability to un-ignore a Finding if it's previously been ignored. Does nothing if the Finding is not ignored.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/data-security/posture/findings/unignore
operation_ids:
    - UnIgnoreFinding
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Remove ignore marker from a finding

`POST /accounts/{account_id}/data-security/posture/findings/unignore`

Operation ID: `UnIgnoreFinding`

Ability to un-ignore a Finding if it's previously been ignored. Does nothing if the Finding is not ignored.

## Definition

```yaml
{"operationId": "UnIgnoreFinding", "summary": "Remove ignore marker from a finding", "description": "Ability to un-ignore a Finding if it's previously been ignored. Does nothing if the Finding is not ignored.", "parameters": [{"$ref": "#/components/parameters/posture-api_AccountTag"}], "requestBody": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/posture-api_FindingBulkActionRequest"}}}}, "responses": {"200": {"description": "OK: Successful HTTP request", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/posture-api_finding-response"}}}}, "400": {"description": "Bad Request: Invalid request parameters"}}, "security": [{"api_token": []}], "tags": ["findings"], "x-api-token-group": ["Zero Trust Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "data-security.posture", "x-fern-sdk-method-name": "create", "x-forge-hidden": true, "x-stability": "beta"}
```
