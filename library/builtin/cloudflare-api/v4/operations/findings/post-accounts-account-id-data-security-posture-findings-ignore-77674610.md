---
title: Mark a finding as ignored
page_id: operation-post-accounts-account-id-data-security-posture-findings-ignore-9fea742f
path: operations/findings
description: Given a list of findings, mark as ignored. Does nothing if Finding is already ignored.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/data-security/posture/findings/ignore
operation_ids:
    - IgnoreFinding
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Mark a finding as ignored

`POST /accounts/{account_id}/data-security/posture/findings/ignore`

Operation ID: `IgnoreFinding`

Given a list of findings, mark as ignored. Does nothing if Finding is already ignored.

## Definition

```yaml
{"operationId": "IgnoreFinding", "summary": "Mark a finding as ignored", "description": "Given a list of findings, mark as ignored. Does nothing if Finding is already ignored.", "parameters": [{"$ref": "#/components/parameters/posture-api_AccountTag"}], "requestBody": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/posture-api_FindingBulkActionRequest"}}}}, "responses": {"200": {"description": "OK: Successful HTTP request", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/posture-api_finding-response"}}}}, "400": {"description": "Bad Request: Invalid request parameters"}}, "security": [{"api_token": []}], "tags": ["findings"], "x-api-token-group": ["Zero Trust Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "data-security.posture", "x-fern-sdk-method-name": "create", "x-forge-hidden": true, "x-stability": "beta"}
```
