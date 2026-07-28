---
title: Delete a service token
page_id: operation-delete-accounts-account-id-access-service-tokens-service-token-id-0a8394e6
path: operations/access-service-tokens
description: Deletes a service token.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/access/service_tokens/{service_token_id}
operation_ids:
    - access-service-tokens-delete-a-service-token
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete a service token

`DELETE /accounts/{account_id}/access/service_tokens/{service_token_id}`

Operation ID: `access-service-tokens-delete-a-service-token`

Deletes a service token.

## Definition

```yaml
{"operationId": "access-service-tokens-delete-a-service-token", "summary": "Delete a service token", "description": "Deletes a service token.", "parameters": [{"name": "service_token_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/access_uuid"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/access_identifier"}}], "responses": {"200": {"description": "Delete a service token response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_single_response-2"}}}}, "4XX": {"description": "Delete a service token response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Access service tokens"], "x-api-token-group": ["Access: Service Tokens Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.access.service-tokens", "x-fern-sdk-method-name": "delete", "x-forge-hidden": true}
```
