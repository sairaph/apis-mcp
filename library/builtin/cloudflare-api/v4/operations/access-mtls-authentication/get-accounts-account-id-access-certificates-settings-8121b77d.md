---
title: List all mTLS hostname settings
page_id: operation-get-accounts-account-id-access-certificates-settings-3e960f21
path: operations/access-mtls-authentication
description: List all mTLS hostname settings for this account.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/access/certificates/settings
operation_ids:
    - access-mtls-authentication-list-mtls-certificates-hostname-settings
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List all mTLS hostname settings

`GET /accounts/{account_id}/access/certificates/settings`

Operation ID: `access-mtls-authentication-list-mtls-certificates-hostname-settings`

List all mTLS hostname settings for this account.

## Definition

```yaml
{"operationId": "access-mtls-authentication-list-mtls-certificates-hostname-settings", "summary": "List all mTLS hostname settings", "description": "List all mTLS hostname settings for this account.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/access_identifier"}}], "responses": {"200": {"description": "List mTLS hostname settings response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_response_collection_hostnames"}}}}, "4XX": {"description": "List mTLS hostname settings response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Access mTLS authentication"], "x-api-token-group": ["Access: Mutual TLS Certificates Write", "Access: Mutual TLS Certificates Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.access.certificates.settings", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
