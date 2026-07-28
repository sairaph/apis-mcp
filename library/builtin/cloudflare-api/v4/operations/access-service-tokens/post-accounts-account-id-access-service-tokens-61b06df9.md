---
title: Create a service token
page_id: operation-post-accounts-account-id-access-service-tokens-8c90ef1a
path: operations/access-service-tokens
description: Generates a new service token. **Note:** This is the only time you can get the Client Secret. If you lose the Client Secret, you will have to rotate the Client Secret or create a new service token.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/access/service_tokens
operation_ids:
    - access-service-tokens-create-a-service-token
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create a service token

`POST /accounts/{account_id}/access/service_tokens`

Operation ID: `access-service-tokens-create-a-service-token`

Generates a new service token. **Note:** This is the only time you can get the Client Secret. If you lose the Client Secret, you will have to rotate the Client Secret or create a new service token.

## Definition

```yaml
{"operationId": "access-service-tokens-create-a-service-token", "summary": "Create a service token", "description": "Generates a new service token. **Note:** This is the only time you can get the Client Secret. If you lose the Client Secret, you will have to rotate the Client Secret or create a new service token.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/access_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"properties": {"client_secret_version": {"$ref": "#/components/schemas/access_client_secret_version"}, "duration": {"$ref": "#/components/schemas/access_duration"}, "name": {"$ref": "#/components/schemas/access_name-2"}, "previous_client_secret_expires_at": {"$ref": "#/components/schemas/access_previous_client_secret_expires_at"}}, "required": ["name"]}}}}, "responses": {"201": {"description": "Create a service token response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_create_response"}}}}, "4XX": {"description": "Create a service token response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Access service tokens"], "x-api-token-group": ["Access: Service Tokens Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.access.service-tokens", "x-fern-sdk-method-name": "create", "x-forge-hidden": true}
```
