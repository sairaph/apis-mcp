---
title: Rotate Access keys
page_id: operation-post-accounts-account-id-access-keys-rotate-1e7ecd80
path: operations/access-key-configuration
description: Perfoms a key rotation for an account.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/access/keys/rotate
operation_ids:
    - access-key-configuration-rotate-access-keys
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Rotate Access keys

`POST /accounts/{account_id}/access/keys/rotate`

Operation ID: `access-key-configuration-rotate-access-keys`

Perfoms a key rotation for an account.

## Definition

```yaml
{"operationId": "access-key-configuration-rotate-access-keys", "summary": "Rotate Access keys", "description": "Perfoms a key rotation for an account.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/access_identifier"}}], "responses": {"200": {"description": "Rotate Access keys response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_single_response-12"}}}}, "4XX": {"description": "Rotate Access keys response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Access key configuration"], "x-api-token-group": ["Access: Organizations, Identity Providers, and Groups Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.access.keys", "x-fern-sdk-method-name": "rotate", "x-forge-hidden": true}
```
