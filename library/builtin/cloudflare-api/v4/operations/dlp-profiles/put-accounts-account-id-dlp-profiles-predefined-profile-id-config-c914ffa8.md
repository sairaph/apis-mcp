---
title: Update predefined profile config
page_id: operation-put-accounts-account-id-dlp-profiles-predefined-profile-id-config-aaf834ea
path: operations/dlp-profiles
description: |-
    This is similar to `update_predefined` but only returns entries that are enabled.
    This is needed for our terraform API
    Updates a DLP predefined profile. Only supports enabling/disabling entries.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /accounts/{account_id}/dlp/profiles/predefined/{profile_id}/config
operation_ids:
    - dlp-profiles-update-predefined-profile-config
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update predefined profile config

`PUT /accounts/{account_id}/dlp/profiles/predefined/{profile_id}/config`

Operation ID: `dlp-profiles-update-predefined-profile-config`

This is similar to `update_predefined` but only returns entries that are enabled.
This is needed for our terraform API
Updates a DLP predefined profile. Only supports enabling/disabling entries.

## Definition

```yaml
{"operationId": "dlp-profiles-update-predefined-profile-config", "summary": "Update predefined profile config", "description": "This is similar to `update_predefined` but only returns entries that are enabled.\nThis is needed for our terraform API\nUpdates a DLP predefined profile. Only supports enabling/disabling entries.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string"}}, {"name": "profile_id", "in": "path", "required": true, "schema": {"type": "string", "format": "uuid"}}], "requestBody": {"description": "The updated parameters for the predefined profile.", "required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dlp_PredefinedProfileConfigUpdate"}}}}, "responses": {"200": {"description": "Update predefined profile response.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/dlp_api-response-single"}, {"properties": {"result": {"$ref": "#/components/schemas/dlp_PredefinedProfileConfig"}}, "type": "object"}]}}}}, "4XX": {"description": "Update predefined profile failure response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dlp_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["DLP Profiles"], "x-api-token-group": ["Zero Trust Write"]}
```
