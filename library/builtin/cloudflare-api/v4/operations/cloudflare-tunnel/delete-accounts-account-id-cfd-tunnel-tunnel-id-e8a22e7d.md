---
title: Delete a Cloudflare Tunnel
page_id: operation-delete-accounts-account-id-cfd-tunnel-tunnel-id-e2c9b5f2
path: operations/cloudflare-tunnel
description: Deletes a Cloudflare Tunnel from an account.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/cfd_tunnel/{tunnel_id}
operation_ids:
    - cloudflare-tunnel-delete-a-cloudflare-tunnel
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete a Cloudflare Tunnel

`DELETE /accounts/{account_id}/cfd_tunnel/{tunnel_id}`

Operation ID: `cloudflare-tunnel-delete-a-cloudflare-tunnel`

Deletes a Cloudflare Tunnel from an account.

## Definition

```yaml
{"operationId": "cloudflare-tunnel-delete-a-cloudflare-tunnel", "summary": "Delete a Cloudflare Tunnel", "description": "Deletes a Cloudflare Tunnel from an account.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/tunnel_account_id"}}, {"name": "tunnel_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/tunnel_tunnel_id"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"type": "object"}}}}, "responses": {"200": {"description": "Delete a Cloudflare Tunnel response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/tunnel_cfd-tunnel-response-single"}}}}, "4XX": {"description": "Delete a Cloudflare Tunnel response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/tunnel_cfd-tunnel-response-single"}, {"$ref": "#/components/schemas/tunnel_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": []}, {"api_token": []}, {"user_service_key": []}], "tags": ["Cloudflare Tunnel"], "x-api-token-group": ["Cloudflare One Connectors Write", "Cloudflare One Connector: cloudflared Write", "Cloudflare Tunnel Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.tunnels.cloudflared", "x-fern-sdk-method-name": "delete", "x-forge-hidden": true}
```
