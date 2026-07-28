---
title: Delete GRE Tunnel
page_id: operation-delete-accounts-account-id-magic-gre-tunnels-gre-tunnel-id-b8df2d55
path: operations/magic-gre-tunnels
description: Disables and removes a specific static GRE tunnel. Use `?validate_only=true` as an optional query parameter to only run validation without persisting changes.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/magic/gre_tunnels/{gre_tunnel_id}
operation_ids:
    - magic-gre-tunnels-delete-gre-tunnel
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete GRE Tunnel

`DELETE /accounts/{account_id}/magic/gre_tunnels/{gre_tunnel_id}`

Operation ID: `magic-gre-tunnels-delete-gre-tunnel`

Disables and removes a specific static GRE tunnel. Use `?validate_only=true` as an optional query parameter to only run validation without persisting changes.

## Definition

```yaml
{"operationId": "magic-gre-tunnels-delete-gre-tunnel", "summary": "Delete GRE Tunnel", "description": "Disables and removes a specific static GRE tunnel. Use `?validate_only=true` as an optional query parameter to only run validation without persisting changes.", "parameters": [{"name": "gre_tunnel_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/magic_identifier"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/magic_identifier"}}, {"name": "x-magic-new-hc-target", "in": "header", "description": "If true, the health check target in the response body will be presented using the new object format. Defaults to false.", "schema": {"type": "boolean"}}], "requestBody": {"required": true, "content": {"application/json": {}}}, "responses": {"200": {"description": "Delete GRE Tunnel response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/magic_tunnel_deleted_response"}}}}, "4XX": {"description": "Delete GRE Tunnel response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/magic_tunnel_deleted_response"}, {"$ref": "#/components/schemas/magic_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Magic GRE tunnels"], "x-api-token-group": ["Magic WAN Write", "Magic Transit Write"], "x-cfPlanAvailability": {"business": false, "enterprise": true, "free": false, "pro": false}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "magic-transit.gre-tunnels", "x-fern-sdk-method-name": "delete", "x-forge-hidden": true}
```
