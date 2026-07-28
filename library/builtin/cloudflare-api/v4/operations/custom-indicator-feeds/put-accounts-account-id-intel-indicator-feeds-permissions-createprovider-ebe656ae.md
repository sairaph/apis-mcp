---
title: Create indicator feed provider
page_id: operation-put-accounts-account-id-intel-indicator-feeds-permissions-createprovider-3de214eb
path: operations/custom-indicator-feeds
description: Creates a new indicator feed provider for an account. Only available to Intel accounts.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /accounts/{account_id}/intel/indicator-feeds/permissions/createProvider
operation_ids:
    - custom-indicator-feeds-create-provider
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create indicator feed provider

`PUT /accounts/{account_id}/intel/indicator-feeds/permissions/createProvider`

Operation ID: `custom-indicator-feeds-create-provider`

Creates a new indicator feed provider for an account. Only available to Intel accounts.

## Definition

```yaml
{"operationId": "custom-indicator-feeds-create-provider", "summary": "Create indicator feed provider", "description": "Creates a new indicator feed provider for an account. Only available to Intel accounts.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/custom-indicator-feeds_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/custom-indicator-feeds_create_provider_request"}}}}, "responses": {"200": {"description": "Create provider response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/custom-indicator-feeds_create_provider_response"}}}}, "4XX": {"description": "Create provider response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/custom-indicator-feeds_create_provider_response"}, {"$ref": "#/components/schemas/custom-indicator-feeds_api_response_common_failure"}]}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["Custom Indicator Feeds"]}
```
