---
title: Set Pre-Shared Keys (PSK) for IPsec tunnels
page_id: operation-post-accounts-account-id-magic-ipsec-tunnels-psk-527a7321
path: operations/magic-ipsec-tunnels
description: Sets Pre-Shared Keys for multiple IPsec tunnels associated with an account. Use `?validate_only=true` as an optional query parameter to only run validation without persisting changes. After PSKs are applied, they are immediately persisted to Cloudflare's edge and cannot be retrieved later. Store the PSKs in a safe place.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/magic/ipsec_tunnels/psk
operation_ids:
    - magic-ipsec-tunnels-set-pre-shared-keys-for-ipsec-tunnels
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Set Pre-Shared Keys (PSK) for IPsec tunnels

`POST /accounts/{account_id}/magic/ipsec_tunnels/psk`

Operation ID: `magic-ipsec-tunnels-set-pre-shared-keys-for-ipsec-tunnels`

Sets Pre-Shared Keys for multiple IPsec tunnels associated with an account. Use `?validate_only=true` as an optional query parameter to only run validation without persisting changes. After PSKs are applied, they are immediately persisted to Cloudflare's edge and cannot be retrieved later. Store the PSKs in a safe place.

## Definition

```yaml
{"operationId": "magic-ipsec-tunnels-set-pre-shared-keys-for-ipsec-tunnels", "summary": "Set Pre-Shared Keys (PSK) for IPsec tunnels", "description": "Sets Pre-Shared Keys for multiple IPsec tunnels associated with an account. Use `?validate_only=true` as an optional query parameter to only run validation without persisting changes. After PSKs are applied, they are immediately persisted to Cloudflare's edge and cannot be retrieved later. Store the PSKs in a safe place.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/magic_identifier"}}, {"name": "validate_only", "in": "query", "description": "If `true`, only run validation without persisting changes.", "schema": {"type": "boolean"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/magic_ipsec_tunnels_psk_request"}}}}, "responses": {"200": {"description": "Set Pre-Shared Keys (PSK) for IPsec tunnels response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/magic_ipsec_tunnels_psk_response"}}}}, "4XX": {"description": "Set Pre-Shared Keys (PSK) for IPsec tunnels response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/magic_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Magic IPsec tunnels"], "x-api-token-group": ["Magic WAN Write", "Magic Transit Write"], "x-cfPlanAvailability": {"business": false, "enterprise": true, "free": false, "pro": false}}
```
