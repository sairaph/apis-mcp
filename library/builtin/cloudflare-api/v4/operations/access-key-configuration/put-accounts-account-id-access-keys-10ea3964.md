---
title: Update the Access key configuration
page_id: operation-put-accounts-account-id-access-keys-5adf29c2
path: operations/access-key-configuration
description: Updates the Access key rotation settings for an account.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /accounts/{account_id}/access/keys
operation_ids:
    - access-key-configuration-update-the-access-key-configuration
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update the Access key configuration

`PUT /accounts/{account_id}/access/keys`

Operation ID: `access-key-configuration-update-the-access-key-configuration`

Updates the Access key rotation settings for an account.

## Definition

```yaml
{"operationId": "access-key-configuration-update-the-access-key-configuration", "summary": "Update the Access key configuration", "description": "Updates the Access key rotation settings for an account.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/access_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"properties": {"key_rotation_interval_days": {"$ref": "#/components/schemas/access_key_rotation_interval_days"}}, "required": ["key_rotation_interval_days"]}}}}, "responses": {"200": {"description": "Update the Access key configuration response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_single_response-12"}}}}, "4XX": {"description": "Update the Access key configuration response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Access key configuration"], "x-api-token-group": ["Access: Organizations, Identity Providers, and Groups Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.access.keys", "x-fern-sdk-method-name": "update", "x-forge-hidden": true}
```
