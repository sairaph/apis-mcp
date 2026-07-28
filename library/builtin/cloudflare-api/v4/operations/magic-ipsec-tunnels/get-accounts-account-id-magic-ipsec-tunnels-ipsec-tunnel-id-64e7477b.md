---
title: List IPsec tunnel details
page_id: operation-get-accounts-account-id-magic-ipsec-tunnels-ipsec-tunnel-id-fdcec31a
path: operations/magic-ipsec-tunnels
description: Lists details for a specific IPsec tunnel.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/magic/ipsec_tunnels/{ipsec_tunnel_id}
operation_ids:
    - magic-ipsec-tunnels-list-ipsec-tunnel-details
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List IPsec tunnel details

`GET /accounts/{account_id}/magic/ipsec_tunnels/{ipsec_tunnel_id}`

Operation ID: `magic-ipsec-tunnels-list-ipsec-tunnel-details`

Lists details for a specific IPsec tunnel.

## Definition

```yaml
{"operationId": "magic-ipsec-tunnels-list-ipsec-tunnel-details", "summary": "List IPsec tunnel details", "description": "Lists details for a specific IPsec tunnel.", "parameters": [{"name": "ipsec_tunnel_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/magic_identifier"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/magic_identifier"}}, {"name": "x-magic-new-hc-target", "in": "header", "description": "If true, the health check target in the response body will be presented using the new object format. Defaults to false.", "schema": {"type": "boolean"}}], "responses": {"200": {"description": "List IPsec tunnel details response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/magic_schemas-tunnel_single_response"}}}}, "4XX": {"description": "List IPsec tunnel details response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/magic_schemas-tunnel_single_response"}, {"$ref": "#/components/schemas/magic_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Magic IPsec tunnels"], "x-api-token-group": ["Magic WAN Write", "Magic WAN Read", "Magic Transit Read", "Magic Transit Write"], "x-cfPlanAvailability": {"business": false, "enterprise": true, "free": false, "pro": false}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "magic-transit.ipsec-tunnels", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
