---
title: Get Service Binding
page_id: operation-get-accounts-account-id-addressing-prefixes-prefix-id-bindings-binding-i-17131267
path: operations/ip-address-management-service-bindings
description: Fetch a single Service Binding
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/addressing/prefixes/{prefix_id}/bindings/{binding_id}
operation_ids:
    - ip-address-management-service-bindings-get-service-binding
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get Service Binding

`GET /accounts/{account_id}/addressing/prefixes/{prefix_id}/bindings/{binding_id}`

Operation ID: `ip-address-management-service-bindings-get-service-binding`

Fetch a single Service Binding

## Definition

```yaml
{"operationId": "ip-address-management-service-bindings-get-service-binding", "summary": "Get Service Binding", "description": "Fetch a single Service Binding", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/addressing_account_identifier"}}, {"name": "prefix_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/addressing_prefix_identifier"}}, {"name": "binding_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/addressing_service_binding_identifier"}}], "responses": {"200": {"description": "The Service Binding with the requested ID", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/addressing_api-response-common"}, {"properties": {"result": {"$ref": "#/components/schemas/addressing_service_binding"}}}]}}}}, "4XX": {"description": "Get Service Binding response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/addressing_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["IP Address Management Service Bindings"], "x-api-token-group": ["IP Prefixes: Write", "IP Prefixes: Read"]}
```
