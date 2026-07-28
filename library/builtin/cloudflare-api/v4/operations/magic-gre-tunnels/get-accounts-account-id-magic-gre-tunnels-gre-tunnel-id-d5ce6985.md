---
title: List GRE Tunnel Details
page_id: operation-get-accounts-account-id-magic-gre-tunnels-gre-tunnel-id-512d3f14
path: operations/magic-gre-tunnels
description: Lists informtion for a specific GRE tunnel.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/magic/gre_tunnels/{gre_tunnel_id}
operation_ids:
    - magic-gre-tunnels-list-gre-tunnel-details
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List GRE Tunnel Details

`GET /accounts/{account_id}/magic/gre_tunnels/{gre_tunnel_id}`

Operation ID: `magic-gre-tunnels-list-gre-tunnel-details`

Lists informtion for a specific GRE tunnel.

## Definition

```yaml
{"operationId": "magic-gre-tunnels-list-gre-tunnel-details", "summary": "List GRE Tunnel Details", "description": "Lists informtion for a specific GRE tunnel.", "parameters": [{"name": "gre_tunnel_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/magic_identifier"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/magic_identifier"}}, {"name": "x-magic-new-hc-target", "in": "header", "description": "If true, the health check target in the response body will be presented using the new object format. Defaults to false.", "schema": {"type": "boolean"}}], "responses": {"200": {"description": "List GRE Tunnel Details response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/magic_tunnel_single_response"}}}}, "4XX": {"description": "List GRE Tunnel Details response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/magic_tunnel_single_response"}, {"$ref": "#/components/schemas/magic_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Magic GRE tunnels"], "x-api-token-group": ["Magic WAN Write", "Magic WAN Read", "Magic Transit Read", "Magic Transit Write"], "x-cfPlanAvailability": {"business": false, "enterprise": true, "free": false, "pro": false}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "magic-transit.gre-tunnels", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
