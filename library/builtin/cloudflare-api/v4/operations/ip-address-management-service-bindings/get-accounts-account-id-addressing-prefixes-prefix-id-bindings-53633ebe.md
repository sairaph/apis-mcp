---
title: List Service Bindings
page_id: operation-get-accounts-account-id-addressing-prefixes-prefix-id-bindings-e4611e81
path: operations/ip-address-management-service-bindings
description: |-
    List the Cloudflare services this prefix is currently bound to. Traffic sent to an address within an IP prefix will be routed to the Cloudflare service of the most-specific Service Binding matching the address.
    **Example:** binding `192.0.2.0/24` to Cloudflare Magic Transit and `192.0.2.1/32` to the Cloudflare CDN would route traffic for `192.0.2.1` to the CDN, and traffic for all other IPs in the prefix to Cloudflare Magic Transit.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/addressing/prefixes/{prefix_id}/bindings
operation_ids:
    - ip-address-management-service-bindings-list-service-bindings
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List Service Bindings

`GET /accounts/{account_id}/addressing/prefixes/{prefix_id}/bindings`

Operation ID: `ip-address-management-service-bindings-list-service-bindings`

List the Cloudflare services this prefix is currently bound to. Traffic sent to an address within an IP prefix will be routed to the Cloudflare service of the most-specific Service Binding matching the address.
**Example:** binding `192.0.2.0/24` to Cloudflare Magic Transit and `192.0.2.1/32` to the Cloudflare CDN would route traffic for `192.0.2.1` to the CDN, and traffic for all other IPs in the prefix to Cloudflare Magic Transit.

## Definition

```yaml
{"operationId": "ip-address-management-service-bindings-list-service-bindings", "summary": "List Service Bindings", "description": "List the Cloudflare services this prefix is currently bound to. Traffic sent to an address within an IP prefix will be routed to the Cloudflare service of the most-specific Service Binding matching the address.\n**Example:** binding `192.0.2.0/24` to Cloudflare Magic Transit and `192.0.2.1/32` to the Cloudflare CDN would route traffic for `192.0.2.1` to the CDN, and traffic for all other IPs in the prefix to Cloudflare Magic Transit.\n", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/addressing_account_identifier"}}, {"name": "prefix_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/addressing_prefix_identifier"}}], "responses": {"200": {"description": "Service Bindings attached to the Prefix", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/addressing_api-response-common"}, {"properties": {"result": {"type": "array", "items": {"$ref": "#/components/schemas/addressing_service_binding"}}}}]}}}}, "4XX": {"description": "List Service Bindings response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/addressing_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["IP Address Management Service Bindings"], "x-api-token-group": ["IP Prefixes: Write", "IP Prefixes: Read"]}
```
