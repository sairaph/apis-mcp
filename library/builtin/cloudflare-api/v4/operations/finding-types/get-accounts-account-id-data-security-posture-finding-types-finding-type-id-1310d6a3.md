---
title: Get finding by ID
page_id: operation-get-accounts-account-id-data-security-posture-finding-types-finding-type-6b96fe2e
path: operations/finding-types
description: Retrieve a specific finding type by its unique identifier.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/data-security/posture/finding_types/{finding_type_id}
operation_ids:
    - GetFindingType
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get finding by ID

`GET /accounts/{account_id}/data-security/posture/finding_types/{finding_type_id}`

Operation ID: `GetFindingType`

Retrieve a specific finding type by its unique identifier.

## Definition

```yaml
{"operationId": "GetFindingType", "summary": "Get finding by ID", "description": "Retrieve a specific finding type by its unique identifier.", "parameters": [{"$ref": "#/components/parameters/posture-api_AccountTag"}, {"$ref": "#/components/parameters/posture-api_FindingTypeId"}], "responses": {"200": {"description": "OK: Successful HTTP request", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/posture-api_finding-type-response"}}}}, "400": {"description": "Bad Request: Invalid request parameters"}, "404": {"description": "Not Found: Finding type not found"}}, "security": [{"api_token": []}], "tags": ["finding_types"], "x-api-token-group": ["Zero Trust Read", "Zero Trust Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "data-security.posture", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
