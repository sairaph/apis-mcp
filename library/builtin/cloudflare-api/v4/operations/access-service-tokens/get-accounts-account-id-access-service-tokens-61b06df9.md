---
title: List service tokens
page_id: operation-get-accounts-account-id-access-service-tokens-e90f4b0b
path: operations/access-service-tokens
description: Lists all service tokens.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/access/service_tokens
operation_ids:
    - access-service-tokens-list-service-tokens
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List service tokens

`GET /accounts/{account_id}/access/service_tokens`

Operation ID: `access-service-tokens-list-service-tokens`

Lists all service tokens.

## Definition

```yaml
{"operationId": "access-service-tokens-list-service-tokens", "summary": "List service tokens", "description": "Lists all service tokens.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/access_identifier"}}, {"name": "name", "in": "query", "schema": {"description": "The name of the service token.", "type": "string"}}, {"name": "search", "in": "query", "schema": {"description": "Search for service tokens by other listed query parameters.", "type": "string"}}, {"$ref": "#/components/parameters/access_page"}, {"$ref": "#/components/parameters/access_per_page"}], "responses": {"200": {"description": "List service tokens response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_response_collection-3"}}}}, "4XX": {"description": "List service tokens response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Access service tokens"], "x-api-token-group": ["Access: Service Tokens Write", "Access: Service Tokens Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.access.service-tokens", "x-fern-sdk-method-name": "list", "x-forge-hidden": true}
```
