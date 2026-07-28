---
title: Create predefined profile
page_id: operation-post-accounts-account-id-dlp-profiles-predefined-2dd1bb24
path: operations/dlp-profiles
description: Creates a DLP predefined profile. Only supports enabling/disabling entries.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/dlp/profiles/predefined
operation_ids:
    - dlp-profiles-create-predefined-profile
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create predefined profile

`POST /accounts/{account_id}/dlp/profiles/predefined`

Operation ID: `dlp-profiles-create-predefined-profile`

Creates a DLP predefined profile. Only supports enabling/disabling entries.

## Definition

```yaml
{"operationId": "dlp-profiles-create-predefined-profile", "summary": "Create predefined profile", "description": "Creates a DLP predefined profile. Only supports enabling/disabling entries.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string"}}], "requestBody": {"description": "Predefined profiles can not be created. This endpoint will only update an existing predefined profiles settings.", "required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dlp_NewPredefinedProfile"}}}}, "responses": {"200": {"description": "Create predefined profile response.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/dlp_api-response-single"}, {"properties": {"result": {"$ref": "#/components/schemas/dlp_Profile"}}, "type": "object"}]}}}}, "4XX": {"description": "Create predefined profile failure response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dlp_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["DLP Profiles"], "x-api-token-group": ["Zero Trust Write"]}
```
