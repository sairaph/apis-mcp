---
title: List remediation types for a finding type
page_id: operation-get-accounts-account-id-data-security-posture-finding-types-finding-type-bd5aaf25
path: operations/finding-types
description: |-
    List all remediation types for a given finding type.
    This endpoint supports both cursor and offset pagination.
    Note that `cursor` and `page` are mutually exclusive.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/data-security/posture/finding_types/{finding_type_id}/remediation_types
operation_ids:
    - GetRemediationTypesForFindingType
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List remediation types for a finding type

`GET /accounts/{account_id}/data-security/posture/finding_types/{finding_type_id}/remediation_types`

Operation ID: `GetRemediationTypesForFindingType`

List all remediation types for a given finding type.
This endpoint supports both cursor and offset pagination.
Note that `cursor` and `page` are mutually exclusive.

## Definition

```yaml
{"operationId": "GetRemediationTypesForFindingType", "summary": "List remediation types for a finding type", "description": "List all remediation types for a given finding type.\nThis endpoint supports both cursor and offset pagination.\nNote that `cursor` and `page` are mutually exclusive.", "parameters": [{"$ref": "#/components/parameters/posture-api_AccountTag"}, {"$ref": "#/components/parameters/posture-api_FindingTypeIdParam"}, {"$ref": "#/components/parameters/posture-api_IntegrationId"}, {"name": "cursor", "in": "query", "description": "A cursor for pagination.", "schema": {"type": "string"}}, {"$ref": "#/components/parameters/posture-api_Page"}, {"$ref": "#/components/parameters/posture-api_PerPage"}], "responses": {"200": {"description": "OK: Successful HTTP request", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/posture-api_PaginatedRemediationTypeList"}}}}, "400": {"description": "Bad Request: Invalid request parameters"}, "404": {"description": "Not Found: Finding type not found"}}, "security": [{"api_token": []}], "tags": ["finding_types"], "x-api-token-group": ["Zero Trust Read", "Zero Trust Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "data-security.posture", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
