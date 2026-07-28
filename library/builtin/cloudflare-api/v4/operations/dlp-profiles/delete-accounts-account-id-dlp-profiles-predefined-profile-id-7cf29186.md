---
title: Delete predefined profile
page_id: operation-delete-accounts-account-id-dlp-profiles-predefined-profile-id-b50efc64
path: operations/dlp-profiles
description: This is a no-op as predefined profiles can't be deleted but is needed for our generated terraform API.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/dlp/profiles/predefined/{profile_id}
operation_ids:
    - dlp-profiles-delete-predefined-profile
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete predefined profile

`DELETE /accounts/{account_id}/dlp/profiles/predefined/{profile_id}`

Operation ID: `dlp-profiles-delete-predefined-profile`

This is a no-op as predefined profiles can't be deleted but is needed for our generated terraform API.

## Definition

```yaml
{"operationId": "dlp-profiles-delete-predefined-profile", "summary": "Delete predefined profile", "description": "This is a no-op as predefined profiles can't be deleted but is needed for our generated terraform API.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string"}}, {"name": "profile_id", "in": "path", "required": true, "schema": {"type": "string", "format": "uuid"}}], "responses": {"200": {"description": "Delete predefined profile response.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/dlp_api-response-single"}, {"properties": {"result": {"$ref": "#/components/schemas/dlp_Empty"}}, "type": "object"}]}}}}, "4XX": {"description": "Delete predefined profile failure response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dlp_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["DLP Profiles"], "x-api-token-group": ["Zero Trust Write"]}
```
