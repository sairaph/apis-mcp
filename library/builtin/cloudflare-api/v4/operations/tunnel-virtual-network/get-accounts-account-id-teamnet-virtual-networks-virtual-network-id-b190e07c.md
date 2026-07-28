---
title: Get a virtual network
page_id: operation-get-accounts-account-id-teamnet-virtual-networks-virtual-network-id-c42c0728
path: operations/tunnel-virtual-network
description: Get a virtual network.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/teamnet/virtual_networks/{virtual_network_id}
operation_ids:
    - tunnel-virtual-network-get
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get a virtual network

`GET /accounts/{account_id}/teamnet/virtual_networks/{virtual_network_id}`

Operation ID: `tunnel-virtual-network-get`

Get a virtual network.

## Definition

```yaml
{"operationId": "tunnel-virtual-network-get", "summary": "Get a virtual network", "description": "Get a virtual network.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/tunnel_account_id"}}, {"name": "virtual_network_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/tunnel_virtual_network_id"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"type": "object", "properties": {"comment": {"$ref": "#/components/schemas/tunnel_virtual_network_comment"}, "is_default_network": {"$ref": "#/components/schemas/tunnel_is_default_network"}, "name": {"$ref": "#/components/schemas/tunnel_virtual_network_name"}}}}}}, "responses": {"200": {"description": "A virtual network response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/tunnel_vnet_response_single"}}}}, "4XX": {"description": "A virtual network response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/tunnel_vnet_response_single"}, {"$ref": "#/components/schemas/tunnel_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": []}, {"api_token": []}, {"user_service_key": []}], "tags": ["Tunnel Virtual Network"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.networks.virtual-networks", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
