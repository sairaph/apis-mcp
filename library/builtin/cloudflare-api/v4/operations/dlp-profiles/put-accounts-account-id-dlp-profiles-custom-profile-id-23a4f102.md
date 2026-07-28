---
title: Update custom profile
page_id: operation-put-accounts-account-id-dlp-profiles-custom-profile-id-707cf036
path: operations/dlp-profiles
description: Updates a DLP custom profile.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /accounts/{account_id}/dlp/profiles/custom/{profile_id}
operation_ids:
    - dlp-profiles-update-custom-profile
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update custom profile

`PUT /accounts/{account_id}/dlp/profiles/custom/{profile_id}`

Operation ID: `dlp-profiles-update-custom-profile`

Updates a DLP custom profile.

## Definition

```yaml
{"operationId": "dlp-profiles-update-custom-profile", "summary": "Update custom profile", "description": "Updates a DLP custom profile.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string"}}, {"name": "profile_id", "in": "path", "required": true, "schema": {"type": "string", "format": "uuid"}}], "requestBody": {"description": "The updated parameters for the profile.", "required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dlp_CustomProfileUpdate"}}}}, "responses": {"200": {"description": "Update custom profile response.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/dlp_api-response-single"}, {"properties": {"result": {"$ref": "#/components/schemas/dlp_Profile"}}, "type": "object"}]}}}}, "4XX": {"description": "Update custom profile failure response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dlp_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["DLP Profiles"], "x-api-token-group": ["Zero Trust Write"]}
```
