---
title: Remove the archive marking from a finding instance
page_id: operation-post-accounts-account-id-data-security-posture-findings-finding-id-insta-5d536b42
path: operations/findings
description: Remove the archive marking from one or more finding instances.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/data-security/posture/findings/{finding_id}/instances/unarchive
operation_ids:
    - UnarchiveFindingInstance
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Remove the archive marking from a finding instance

`POST /accounts/{account_id}/data-security/posture/findings/{finding_id}/instances/unarchive`

Operation ID: `UnarchiveFindingInstance`

Remove the archive marking from one or more finding instances.

## Definition

```yaml
{"operationId": "UnarchiveFindingInstance", "summary": "Remove the archive marking from a finding instance", "description": "Remove the archive marking from one or more finding instances.", "parameters": [{"$ref": "#/components/parameters/posture-api_AccountTag"}, {"$ref": "#/components/parameters/posture-api_FindingIdByte"}], "requestBody": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/posture-api_FindingInstanceBulkActionRequest"}}}}, "responses": {"200": {"description": "OK: Successful HTTP request", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/posture-api_finding-instance-response"}}}}, "400": {"description": "Bad Request: Invalid request parameters"}, "404": {"description": "Not Found: Finding not found"}}, "security": [{"api_token": []}], "tags": ["findings"], "x-api-token-group": ["Zero Trust Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "data-security.posture", "x-fern-sdk-method-name": "create", "x-forge-hidden": true, "x-stability": "beta"}
```
