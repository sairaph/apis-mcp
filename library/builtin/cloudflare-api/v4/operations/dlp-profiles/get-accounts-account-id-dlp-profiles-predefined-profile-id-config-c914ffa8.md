---
title: Get predefined profile config
page_id: operation-get-accounts-account-id-dlp-profiles-predefined-profile-id-config-f643e1e5
path: operations/dlp-profiles
description: |-
    This is similar to `get_predefined` but only returns entries that are enabled.
    This is needed for our terraform API
    Fetches a predefined DLP profile by id.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/dlp/profiles/predefined/{profile_id}/config
operation_ids:
    - dlp-profiles-get-predefined-profile-config
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get predefined profile config

`GET /accounts/{account_id}/dlp/profiles/predefined/{profile_id}/config`

Operation ID: `dlp-profiles-get-predefined-profile-config`

This is similar to `get_predefined` but only returns entries that are enabled.
This is needed for our terraform API
Fetches a predefined DLP profile by id.

## Definition

```yaml
{"operationId": "dlp-profiles-get-predefined-profile-config", "summary": "Get predefined profile config", "description": "This is similar to `get_predefined` but only returns entries that are enabled.\nThis is needed for our terraform API\nFetches a predefined DLP profile by id.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string"}}, {"name": "profile_id", "in": "path", "required": true, "schema": {"type": "string", "format": "uuid"}}], "responses": {"200": {"description": "Predefined profile response.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/dlp_api-response-single"}, {"properties": {"result": {"$ref": "#/components/schemas/dlp_PredefinedProfileConfig"}}, "type": "object"}]}}}}, "4XX": {"description": "Predefined profile failure response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dlp_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["DLP Profiles"], "x-api-token-group": ["Zero Trust Read", "Zero Trust Write"]}
```
