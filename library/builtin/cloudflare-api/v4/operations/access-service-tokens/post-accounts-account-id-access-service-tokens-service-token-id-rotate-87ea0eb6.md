---
title: Rotate a service token
page_id: operation-post-accounts-account-id-access-service-tokens-service-token-id-rotate-cd4cfd77
path: operations/access-service-tokens
description: Generates a new Client Secret for a service token and revokes the old one.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/access/service_tokens/{service_token_id}/rotate
operation_ids:
    - access-service-tokens-rotate-a-service-token
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Rotate a service token

`POST /accounts/{account_id}/access/service_tokens/{service_token_id}/rotate`

Operation ID: `access-service-tokens-rotate-a-service-token`

Generates a new Client Secret for a service token and revokes the old one.

## Definition

```yaml
{"operationId": "access-service-tokens-rotate-a-service-token", "summary": "Rotate a service token", "description": "Generates a new Client Secret for a service token and revokes the old one.", "parameters": [{"name": "service_token_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/access_uuid"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/access_identifier"}}], "requestBody": {"content": {"application/json": {"schema": {"properties": {"previous_client_secret_expires_at": {"description": "The expiration of the previous `client_secret`. If not provided, it defaults to the current timestamp in order to immediately expire the previous secret.", "type": "string", "format": "date-time", "example": "2014-01-01T05:20:00.12345Z", "x-auditable": true}}}}}}, "responses": {"200": {"description": "Rotate a service token response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_create_response"}}}}, "4XX": {"description": "Rotate a service token response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Access service tokens"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.access.service-tokens", "x-fern-sdk-method-name": "rotate", "x-forge-hidden": true}
```
