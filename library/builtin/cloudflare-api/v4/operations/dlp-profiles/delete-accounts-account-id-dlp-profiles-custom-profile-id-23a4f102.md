---
title: Delete custom profile
page_id: operation-delete-accounts-account-id-dlp-profiles-custom-profile-id-caeb9eba
path: operations/dlp-profiles
description: Deletes a DLP custom profile.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/dlp/profiles/custom/{profile_id}
operation_ids:
    - dlp-profiles-delete-custom-profile
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete custom profile

`DELETE /accounts/{account_id}/dlp/profiles/custom/{profile_id}`

Operation ID: `dlp-profiles-delete-custom-profile`

Deletes a DLP custom profile.

## Definition

```yaml
{"operationId": "dlp-profiles-delete-custom-profile", "summary": "Delete custom profile", "description": "Deletes a DLP custom profile.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string"}}, {"name": "profile_id", "in": "path", "required": true, "schema": {"type": "string", "format": "uuid"}}], "responses": {"200": {"description": "Delete custom profile response.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/dlp_api-response-single"}, {"properties": {"result": {"$ref": "#/components/schemas/dlp_Empty"}}, "type": "object"}]}}}}, "4XX": {"description": "Delete custom profile failure response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dlp_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["DLP Profiles"], "x-api-token-group": ["Zero Trust Write"]}
```
