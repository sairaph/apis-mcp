---
title: Create custom profile
page_id: operation-post-accounts-account-id-dlp-profiles-custom-581386f0
path: operations/dlp-profiles
description: Creates a DLP custom profile.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/dlp/profiles/custom
operation_ids:
    - dlp-profiles-create-custom-profiles
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create custom profile

`POST /accounts/{account_id}/dlp/profiles/custom`

Operation ID: `dlp-profiles-create-custom-profiles`

Creates a DLP custom profile.

## Definition

```yaml
{"operationId": "dlp-profiles-create-custom-profiles", "summary": "Create custom profile", "description": "Creates a DLP custom profile.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string"}}], "requestBody": {"description": "A new profile to create.", "required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dlp_NewCustomProfile"}}}}, "responses": {"200": {"description": "New custom profile response.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/dlp_api-response-single"}, {"properties": {"result": {"$ref": "#/components/schemas/dlp_Profile"}}, "type": "object"}]}}}}, "4XX": {"description": "New custom profile failure response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dlp_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["DLP Profiles"], "x-api-token-group": ["Zero Trust Write"]}
```
