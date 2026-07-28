---
title: Create Temporary Access Credentials
page_id: operation-post-accounts-account-id-r2-temp-access-credentials-1d278786
path: operations/r2-bucket
description: Creates temporary access credentials on a bucket that can be optionally scoped to prefixes or objects.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/r2/temp-access-credentials
operation_ids:
    - r2-create-temp-access-credentials
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create Temporary Access Credentials

`POST /accounts/{account_id}/r2/temp-access-credentials`

Operation ID: `r2-create-temp-access-credentials`

Creates temporary access credentials on a bucket that can be optionally scoped to prefixes or objects.

## Definition

```yaml
{"operationId": "r2-create-temp-access-credentials", "summary": "Create Temporary Access Credentials", "description": "Creates temporary access credentials on a bucket that can be optionally scoped to prefixes or objects.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/r2_account_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/r2_temp_access_creds_request"}}}}, "responses": {"200": {"description": "Create temporary access credentials response.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/r2_v4_response"}, {"properties": {"result": {"$ref": "#/components/schemas/r2_temp_access_creds_response"}}, "type": "object"}]}}}}, "4XX": {"description": "Create temporary access credentials response failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/r2_v4_response_failure"}}}}}, "security": [{"api_token": []}], "tags": ["R2 Bucket"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "r2.temporary-credentials", "x-fern-sdk-method-name": "create", "x-forge-hidden": true}
```
