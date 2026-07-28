---
title: List virtual networks
page_id: operation-get-accounts-account-id-teamnet-virtual-networks-ed4359a9
path: operations/tunnel-virtual-network
description: Lists and filters virtual networks in an account.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/teamnet/virtual_networks
operation_ids:
    - tunnel-virtual-network-list-virtual-networks
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List virtual networks

`GET /accounts/{account_id}/teamnet/virtual_networks`

Operation ID: `tunnel-virtual-network-list-virtual-networks`

Lists and filters virtual networks in an account.

## Definition

```yaml
{"operationId": "tunnel-virtual-network-list-virtual-networks", "summary": "List virtual networks", "description": "Lists and filters virtual networks in an account.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/tunnel_account_id"}}, {"name": "id", "in": "query", "schema": {"$ref": "#/components/schemas/tunnel_virtual_network_id"}}, {"name": "name", "in": "query", "schema": {"$ref": "#/components/schemas/tunnel_virtual_network_name"}}, {"name": "is_default", "in": "query", "schema": {"description": "If `true`, only include the default virtual network. If `false`, exclude the default virtual network. If empty, all virtual networks will be included.", "type": "boolean"}, "deprecated": true, "x-stainless-deprecation-message": "Use the is_default_network property instead.", "x-stainless-ignore": true}, {"name": "is_default_network", "in": "query", "schema": {"description": "If `true`, only include the default virtual network. If `false`, exclude the default virtual network. If empty, all virtual networks will be included.", "type": "boolean"}}, {"name": "is_deleted", "in": "query", "schema": {"description": "If `true`, only include deleted virtual networks. If `false`, exclude deleted virtual networks. If empty, all virtual networks will be included.", "type": "boolean"}}], "responses": {"200": {"description": "List virtual networks response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/tunnel_vnet_response_collection"}}}}, "4XX": {"description": "List virtual networks response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/tunnel_vnet_response_collection"}, {"$ref": "#/components/schemas/tunnel_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": []}, {"api_token": []}, {"user_service_key": []}], "tags": ["Tunnel Virtual Network"], "x-api-token-group": ["Cloudflare One Networks Write", "Cloudflare One Networks Read", "Cloudflare Tunnel Write", "Cloudflare Tunnel Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.networks.virtual-networks", "x-fern-sdk-method-name": "list", "x-forge-hidden": true}
```
