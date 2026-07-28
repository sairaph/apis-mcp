---
title: Generate Pre-Shared Key (PSK) for IPsec tunnels
page_id: operation-post-accounts-account-id-magic-ipsec-tunnels-ipsec-tunnel-id-psk-generat-b898113b
path: operations/magic-ipsec-tunnels
description: Generates a Pre-Shared Key for a specific IPsec tunnel used in the IKE session. Use `?validate_only=true` as an optional query parameter to only run validation without persisting changes. After a PSK is generated, the PSK is immediately persisted to Cloudflare's edge and cannot be retrieved later. Store the PSK in a safe place.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/magic/ipsec_tunnels/{ipsec_tunnel_id}/psk_generate
operation_ids:
    - magic-ipsec-tunnels-generate-pre-shared-key-(-psk)-for-ipsec-tunnels
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Generate Pre-Shared Key (PSK) for IPsec tunnels

`POST /accounts/{account_id}/magic/ipsec_tunnels/{ipsec_tunnel_id}/psk_generate`

Operation ID: `magic-ipsec-tunnels-generate-pre-shared-key-(-psk)-for-ipsec-tunnels`

Generates a Pre-Shared Key for a specific IPsec tunnel used in the IKE session. Use `?validate_only=true` as an optional query parameter to only run validation without persisting changes. After a PSK is generated, the PSK is immediately persisted to Cloudflare's edge and cannot be retrieved later. Store the PSK in a safe place.

## Definition

```yaml
{"operationId": "magic-ipsec-tunnels-generate-pre-shared-key-(-psk)-for-ipsec-tunnels", "summary": "Generate Pre-Shared Key (PSK) for IPsec tunnels", "description": "Generates a Pre-Shared Key for a specific IPsec tunnel used in the IKE session. Use `?validate_only=true` as an optional query parameter to only run validation without persisting changes. After a PSK is generated, the PSK is immediately persisted to Cloudflare's edge and cannot be retrieved later. Store the PSK in a safe place.", "parameters": [{"name": "ipsec_tunnel_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/magic_identifier"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/magic_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {}}}, "responses": {"200": {"description": "Generate Pre-Shared Key (PSK) for IPsec tunnels response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/magic_psk_generation_response"}}}}, "4XX": {"description": "Generate Pre-Shared Key (PSK) for IPsec tunnels response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/magic_psk_generation_response"}, {"$ref": "#/components/schemas/magic_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Magic IPsec tunnels"], "x-api-token-group": ["Magic WAN Write", "Magic Transit Write"], "x-cfPlanAvailability": {"business": false, "enterprise": true, "free": false, "pro": false}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "magic-transit.ipsec-tunnels", "x-fern-sdk-method-name": "psk-generate", "x-forge-hidden": true}
```
