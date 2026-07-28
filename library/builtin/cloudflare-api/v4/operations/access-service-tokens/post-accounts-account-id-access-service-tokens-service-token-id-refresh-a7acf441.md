---
title: Refresh a service token
page_id: operation-post-accounts-account-id-access-service-tokens-service-token-id-refresh-8e823246
path: operations/access-service-tokens
description: Refreshes the expiration of a service token.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/access/service_tokens/{service_token_id}/refresh
operation_ids:
    - access-service-tokens-refresh-a-service-token
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Refresh a service token

`POST /accounts/{account_id}/access/service_tokens/{service_token_id}/refresh`

Operation ID: `access-service-tokens-refresh-a-service-token`

Refreshes the expiration of a service token.

## Definition

```yaml
{"operationId": "access-service-tokens-refresh-a-service-token", "summary": "Refresh a service token", "description": "Refreshes the expiration of a service token.", "parameters": [{"name": "service_token_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/access_uuid"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/access_identifier"}}], "responses": {"200": {"description": "Refresh a service token response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_single_response-2"}}}}, "4XX": {"description": "Refresh a service token response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Access service tokens"], "x-api-token-group": ["Access: Service Tokens Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.access.service-tokens", "x-fern-sdk-method-name": "refresh", "x-forge-hidden": true}
```
