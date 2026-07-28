---
title: Update Prefix Dynamic Advertisement Status
page_id: operation-patch-accounts-account-id-addressing-prefixes-prefix-id-bgp-status-2bcebbd8
path: operations/ip-address-management-dynamic-advertisement
description: "Advertise or withdraw the BGP route for a prefix.\n\n**Deprecated:** Prefer the BGP Prefixes endpoints, which additionally allow for advertising and withdrawing \nsubnets of an IP prefix."
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PATCH
api_endpoints:
    - /accounts/{account_id}/addressing/prefixes/{prefix_id}/bgp/status
operation_ids:
    - ip-address-management-dynamic-advertisement-update-prefix-dynamic-advertisement-status
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update Prefix Dynamic Advertisement Status

`PATCH /accounts/{account_id}/addressing/prefixes/{prefix_id}/bgp/status`

Operation ID: `ip-address-management-dynamic-advertisement-update-prefix-dynamic-advertisement-status`

Advertise or withdraw the BGP route for a prefix.

**Deprecated:** Prefer the BGP Prefixes endpoints, which additionally allow for advertising and withdrawing
subnets of an IP prefix.

## Definition

```yaml
{"operationId": "ip-address-management-dynamic-advertisement-update-prefix-dynamic-advertisement-status", "summary": "Update Prefix Dynamic Advertisement Status", "description": "Advertise or withdraw the BGP route for a prefix.\n\n**Deprecated:** Prefer the BGP Prefixes endpoints, which additionally allow for advertising and withdrawing \nsubnets of an IP prefix.\n", "parameters": [{"name": "prefix_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/addressing_prefix_identifier"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/addressing_account_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"properties": {"advertised": {"$ref": "#/components/schemas/addressing_components-schemas-advertised"}}, "required": ["advertised"]}}}}, "responses": {"200": {"description": "Update Prefix Dynamic Advertisement Status response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/addressing_advertised_response"}}}}, "4XX": {"description": "Update Prefix Dynamic Advertisement Status response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/addressing_advertised_response"}, {"$ref": "#/components/schemas/addressing_api-response-common-failure"}]}}}}}, "deprecated": true, "security": [{"api_email": [], "api_key": []}], "tags": ["IP Address Management Dynamic Advertisement"], "x-api-token-group": ["Magic Transit Write", "IP Prefixes: Write", "IP Prefixes: BGP On Demand Write"]}
```
