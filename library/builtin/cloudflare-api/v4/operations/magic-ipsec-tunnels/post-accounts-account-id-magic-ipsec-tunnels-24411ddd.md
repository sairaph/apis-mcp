---
title: Create an IPsec tunnel
page_id: operation-post-accounts-account-id-magic-ipsec-tunnels-046ca227
path: operations/magic-ipsec-tunnels
description: Creates a new IPsec tunnel associated with an account. Use `?validate_only=true` as an optional query parameter to only run validation without persisting changes.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/magic/ipsec_tunnels
operation_ids:
    - magic-ipsec-tunnels-create-ipsec-tunnels
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create an IPsec tunnel

`POST /accounts/{account_id}/magic/ipsec_tunnels`

Operation ID: `magic-ipsec-tunnels-create-ipsec-tunnels`

Creates a new IPsec tunnel associated with an account. Use `?validate_only=true` as an optional query parameter to only run validation without persisting changes.

## Definition

```yaml
{"operationId": "magic-ipsec-tunnels-create-ipsec-tunnels", "summary": "Create an IPsec tunnel", "description": "Creates a new IPsec tunnel associated with an account. Use `?validate_only=true` as an optional query parameter to only run validation without persisting changes.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/magic_identifier"}}, {"name": "x-magic-new-hc-target", "in": "header", "description": "If true, the health check target in the request and response bodies will be presented using the new object format. Defaults to false.", "schema": {"type": "boolean"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/magic_ipsec_tunnel_add_request"}}}}, "responses": {"200": {"description": "Create IPsec tunnels response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/magic_schemas-create_ipsec_tunnel_response"}}}}, "4XX": {"description": "Create IPsec tunnels response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/magic_schemas-create_ipsec_tunnel_response"}, {"$ref": "#/components/schemas/magic_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Magic IPsec tunnels"], "x-api-token-group": ["Magic WAN Write", "Magic Transit Write"], "x-cfPlanAvailability": {"business": false, "enterprise": true, "free": false, "pro": false}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "magic-transit.ipsec-tunnels", "x-fern-sdk-method-name": "create", "x-forge-hidden": true}
```
