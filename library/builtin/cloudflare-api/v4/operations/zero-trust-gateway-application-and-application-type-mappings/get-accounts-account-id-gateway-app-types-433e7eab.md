---
title: List application and application type mappings
page_id: operation-get-accounts-account-id-gateway-app-types-818dd108
path: operations/zero-trust-gateway-application-and-application-type-mappings
description: List all application and application type mappings.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/gateway/app_types
operation_ids:
    - zero-trust-gateway-application-and-application-type-mappings-list-application-and-application-type-mappings
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List application and application type mappings

`GET /accounts/{account_id}/gateway/app_types`

Operation ID: `zero-trust-gateway-application-and-application-type-mappings-list-application-and-application-type-mappings`

List all application and application type mappings.

## Definition

```yaml
{"operationId": "zero-trust-gateway-application-and-application-type-mappings-list-application-and-application-type-mappings", "summary": "List application and application type mappings", "description": "List all application and application type mappings.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/zero-trust-gateway_identifier-3"}}], "responses": {"200": {"description": "List application and application type mappings response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/zero-trust-gateway_response_collection-7"}}}}, "4XX": {"description": "List application and application type mappings response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/zero-trust-gateway_response_collection-7"}, {"$ref": "#/components/schemas/zero-trust-gateway_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Zero Trust Gateway application and application type mappings"], "x-api-token-group": ["Zero Trust Read", "Zero Trust Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.gateway.app-types", "x-fern-sdk-method-name": "list", "x-forge-hidden": true}
```
