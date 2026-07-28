---
title: Delete a virtual network
page_id: operation-delete-accounts-account-id-teamnet-virtual-networks-virtual-network-id-0eec0793
path: operations/tunnel-virtual-network
description: Deletes an existing virtual network.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/teamnet/virtual_networks/{virtual_network_id}
operation_ids:
    - tunnel-virtual-network-delete
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete a virtual network

`DELETE /accounts/{account_id}/teamnet/virtual_networks/{virtual_network_id}`

Operation ID: `tunnel-virtual-network-delete`

Deletes an existing virtual network.

## Definition

```yaml
{"operationId": "tunnel-virtual-network-delete", "summary": "Delete a virtual network", "description": "Deletes an existing virtual network.", "parameters": [{"name": "virtual_network_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/tunnel_virtual_network_id"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/tunnel_account_id"}}], "requestBody": {"required": true, "content": {"application/json": {}}}, "responses": {"200": {"description": "Delete a virtual network response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/tunnel_vnet_response_single"}}}}, "4XX": {"description": "Delete a virtual network response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/tunnel_vnet_response_single"}, {"$ref": "#/components/schemas/tunnel_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": []}, {"api_token": []}, {"user_service_key": []}], "tags": ["Tunnel Virtual Network"], "x-api-token-group": ["Cloudflare One Networks Write", "Cloudflare Tunnel Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.networks.virtual-networks", "x-fern-sdk-method-name": "delete", "x-forge-hidden": true}
```
