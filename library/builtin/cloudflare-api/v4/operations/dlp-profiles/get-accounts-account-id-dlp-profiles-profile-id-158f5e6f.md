---
title: Get DLP Profile
page_id: operation-get-accounts-account-id-dlp-profiles-profile-id-eaf11585
path: operations/dlp-profiles
description: Fetches a DLP profile by ID.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/dlp/profiles/{profile_id}
operation_ids:
    - dlp-profiles-get-dlp-profile
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get DLP Profile

`GET /accounts/{account_id}/dlp/profiles/{profile_id}`

Operation ID: `dlp-profiles-get-dlp-profile`

Fetches a DLP profile by ID.

## Definition

```yaml
{"operationId": "dlp-profiles-get-dlp-profile", "summary": "Get DLP Profile", "description": "Fetches a DLP profile by ID.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string"}}, {"name": "profile_id", "in": "path", "required": true, "schema": {"type": "string", "format": "uuid"}}], "responses": {"200": {"description": "Get profile response.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/dlp_api-response-single"}, {"properties": {"result": {"$ref": "#/components/schemas/dlp_Profile"}}, "type": "object"}]}}}}, "4XX": {"description": "Get profile failure response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dlp_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["DLP Profiles"], "x-api-token-group": ["Zero Trust Read", "Zero Trust Write"]}
```
