---
title: Update IPsec Tunnel
page_id: operation-put-accounts-account-id-magic-ipsec-tunnels-ipsec-tunnel-id-89c930cb
path: operations/magic-ipsec-tunnels
description: Updates a specific IPsec tunnel associated with an account. Use `?validate_only=true` as an optional query parameter to only run validation without persisting changes.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /accounts/{account_id}/magic/ipsec_tunnels/{ipsec_tunnel_id}
operation_ids:
    - magic-ipsec-tunnels-update-ipsec-tunnel
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update IPsec Tunnel

`PUT /accounts/{account_id}/magic/ipsec_tunnels/{ipsec_tunnel_id}`

Operation ID: `magic-ipsec-tunnels-update-ipsec-tunnel`

Updates a specific IPsec tunnel associated with an account. Use `?validate_only=true` as an optional query parameter to only run validation without persisting changes.

## Definition

```yaml
{"operationId": "magic-ipsec-tunnels-update-ipsec-tunnel", "summary": "Update IPsec Tunnel", "description": "Updates a specific IPsec tunnel associated with an account. Use `?validate_only=true` as an optional query parameter to only run validation without persisting changes.", "parameters": [{"name": "ipsec_tunnel_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/magic_identifier"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/magic_identifier"}}, {"name": "x-magic-new-hc-target", "in": "header", "description": "If true, the health check target in the request and response bodies will be presented using the new object format. Defaults to false.", "schema": {"type": "boolean"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/magic_ipsec_tunnel_add_single_request"}}}}, "responses": {"200": {"description": "Update IPsec Tunnel response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/magic_schemas-tunnel_modified_response"}}}}, "4XX": {"description": "Update IPsec Tunnel response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/magic_schemas-tunnel_modified_response"}, {"$ref": "#/components/schemas/magic_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Magic IPsec tunnels"], "x-api-token-group": ["Magic WAN Write", "Magic Transit Write"], "x-cfPlanAvailability": {"business": false, "enterprise": true, "free": false, "pro": false}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "magic-transit.ipsec-tunnels", "x-fern-sdk-method-name": "update", "x-forge-hidden": true}
```
