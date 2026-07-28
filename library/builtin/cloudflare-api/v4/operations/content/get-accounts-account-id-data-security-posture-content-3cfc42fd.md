---
title: List DLP content findings
page_id: operation-get-accounts-account-id-data-security-posture-content-141934c0
path: operations/content
description: List DLP content findings
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/data-security/posture/content
operation_ids:
    - ListContentAssets
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List DLP content findings

`GET /accounts/{account_id}/data-security/posture/content`

Operation ID: `ListContentAssets`

List DLP content findings

## Definition

```yaml
{"operationId": "ListContentAssets", "summary": "List DLP content findings", "description": "List DLP content findings", "parameters": [{"$ref": "#/components/parameters/posture-api_AccountTag"}, {"$ref": "#/components/parameters/posture-api_Direction"}, {"name": "dlp_profile_id", "in": "query", "description": "Filter by an DLP profile ID", "schema": {"type": "string", "format": "uuid"}}, {"$ref": "#/components/parameters/posture-api_IntegrationId"}, {"$ref": "#/components/parameters/posture-api_MaxAfflictionDate"}, {"$ref": "#/components/parameters/posture-api_MinAfflictionDate"}, {"name": "order", "in": "query", "description": "Which field to use when ordering content assets.", "schema": {"type": "string", "enum": ["asset_name", "dlp_profile_count", "integration_name", "latest_affliction_date"]}}, {"$ref": "#/components/parameters/posture-api_Page"}, {"$ref": "#/components/parameters/posture-api_PerPage"}, {"$ref": "#/components/parameters/posture-api_Search"}, {"$ref": "#/components/parameters/posture-api_Vendor"}], "responses": {"200": {"description": "OK: Successful HTTP request", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/posture-api_PaginatedContentAssetList"}}}}, "400": {"description": "Bad Request: Invalid request parameters"}, "429": {"description": "Too Many Requests: Request was rate limited. May include a\n`Retry-After` header.", "headers": {"Retry-After": {"description": "Optional. Indicates how many seconds to wait before retrying.", "schema": {"type": "string"}}}, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/posture-api_api-response-common"}}}}}, "security": [{"api_token": []}], "tags": ["content"], "x-api-token-group": ["Zero Trust Read", "Zero Trust Write"]}
```
