---
title: Get the Access key configuration
page_id: operation-get-accounts-account-id-access-keys-27042057
path: operations/access-key-configuration
description: Gets the Access key rotation settings for an account.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/access/keys
operation_ids:
    - access-key-configuration-get-the-access-key-configuration
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get the Access key configuration

`GET /accounts/{account_id}/access/keys`

Operation ID: `access-key-configuration-get-the-access-key-configuration`

Gets the Access key rotation settings for an account.

## Definition

```yaml
{"operationId": "access-key-configuration-get-the-access-key-configuration", "summary": "Get the Access key configuration", "description": "Gets the Access key rotation settings for an account.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/access_identifier"}}], "responses": {"200": {"description": "Get the Access key configuration response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_single_response-12"}}}}, "4XX": {"description": "Get the Access key configuration response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Access key configuration"], "x-api-token-group": ["Access: Organizations, Identity Providers, and Groups Write", "Access: Organizations, Identity Providers, and Groups Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.access.keys", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
