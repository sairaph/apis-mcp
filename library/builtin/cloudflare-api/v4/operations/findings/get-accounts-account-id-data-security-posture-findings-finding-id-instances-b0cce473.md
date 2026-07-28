---
title: Get a finding instance using an instance ID
page_id: operation-get-accounts-account-id-data-security-posture-findings-finding-id-instan-b03020bc
path: operations/findings
description: Gets a security Finding instance by id.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/data-security/posture/findings/{finding_id}/instances/{instance_id}
operation_ids:
    - GetFindingInstance
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get a finding instance using an instance ID

`GET /accounts/{account_id}/data-security/posture/findings/{finding_id}/instances/{instance_id}`

Operation ID: `GetFindingInstance`

Gets a security Finding instance by id.

## Definition

```yaml
{"operationId": "GetFindingInstance", "summary": "Get a finding instance using an instance ID", "description": "Gets a security Finding instance by id.", "parameters": [{"$ref": "#/components/parameters/posture-api_AccountTag"}, {"$ref": "#/components/parameters/posture-api_FindingIdByte"}, {"$ref": "#/components/parameters/posture-api_FindingInstanceId"}], "responses": {"200": {"description": "OK: Successful HTTP request", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/posture-api_finding-instance-response"}}}}, "400": {"description": "Bad Request: Invalid request parameters"}, "404": {"description": "Not Found: Finding instance not found"}}, "security": [{"api_token": []}], "tags": ["findings"], "x-api-token-group": ["Zero Trust Read", "Zero Trust Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "data-security.posture", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
