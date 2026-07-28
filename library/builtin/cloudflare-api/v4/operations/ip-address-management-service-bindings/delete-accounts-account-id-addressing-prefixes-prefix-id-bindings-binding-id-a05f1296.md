---
title: Delete Service Binding
page_id: operation-delete-accounts-account-id-addressing-prefixes-prefix-id-bindings-bindin-5c8e907d
path: operations/ip-address-management-service-bindings
description: Delete a Service Binding
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/addressing/prefixes/{prefix_id}/bindings/{binding_id}
operation_ids:
    - ip-address-management-service-bindings-delete-service-binding
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete Service Binding

`DELETE /accounts/{account_id}/addressing/prefixes/{prefix_id}/bindings/{binding_id}`

Operation ID: `ip-address-management-service-bindings-delete-service-binding`

Delete a Service Binding

## Definition

```yaml
{"operationId": "ip-address-management-service-bindings-delete-service-binding", "summary": "Delete Service Binding", "description": "Delete a Service Binding", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/addressing_account_identifier"}}, {"name": "prefix_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/addressing_prefix_identifier"}}, {"name": "binding_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/addressing_service_binding_identifier"}}], "responses": {"200": {"description": "Service Binding deleted", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/addressing_api-response-common"}}}}, "4XX": {"description": "Delete Service Binding response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/addressing_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["IP Address Management Service Bindings"], "x-api-token-group": ["IP Prefixes: Write"]}
```
