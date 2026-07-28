---
title: Update a virtual network
page_id: operation-patch-accounts-account-id-teamnet-virtual-networks-virtual-network-id-23fb87dd
path: operations/tunnel-virtual-network
description: Updates an existing virtual network.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PATCH
api_endpoints:
    - /accounts/{account_id}/teamnet/virtual_networks/{virtual_network_id}
operation_ids:
    - tunnel-virtual-network-update
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update a virtual network

`PATCH /accounts/{account_id}/teamnet/virtual_networks/{virtual_network_id}`

Operation ID: `tunnel-virtual-network-update`

Updates an existing virtual network.

## Definition

```yaml
{"operationId": "tunnel-virtual-network-update", "summary": "Update a virtual network", "description": "Updates an existing virtual network.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/tunnel_account_id"}}, {"name": "virtual_network_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/tunnel_virtual_network_id"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"type": "object", "properties": {"comment": {"$ref": "#/components/schemas/tunnel_virtual_network_comment"}, "is_default_network": {"$ref": "#/components/schemas/tunnel_is_default_network_optional"}, "name": {"$ref": "#/components/schemas/tunnel_virtual_network_name"}}}}}}, "responses": {"200": {"description": "Update a virtual network response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/tunnel_vnet_response_single"}}}}, "4XX": {"description": "Update a virtual network response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/tunnel_vnet_response_single"}, {"$ref": "#/components/schemas/tunnel_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": []}, {"api_token": []}, {"user_service_key": []}], "tags": ["Tunnel Virtual Network"], "x-api-token-group": ["Cloudflare One Networks Write", "Cloudflare Tunnel Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.networks.virtual-networks", "x-fern-sdk-method-name": "edit", "x-forge-hidden": true}
```
