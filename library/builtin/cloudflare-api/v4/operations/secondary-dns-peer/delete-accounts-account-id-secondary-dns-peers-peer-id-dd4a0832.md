---
title: Delete Peer
page_id: operation-delete-accounts-account-id-secondary-dns-peers-peer-id-d0bfc217
path: operations/secondary-dns-peer
description: Delete Peer.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/secondary_dns/peers/{peer_id}
operation_ids:
    - secondary-dns-(-peer)-delete-peer
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete Peer

`DELETE /accounts/{account_id}/secondary_dns/peers/{peer_id}`

Operation ID: `secondary-dns-(-peer)-delete-peer`

Delete Peer.

## Definition

```yaml
{"operationId": "secondary-dns-(-peer)-delete-peer", "summary": "Delete Peer", "description": "Delete Peer.", "parameters": [{"name": "peer_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/secondary-dns_identifier-3"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/secondary-dns_account_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {}}}, "responses": {"200": {"description": "Delete Peer response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/secondary-dns_id_response-3"}}}}, "4XX": {"description": "Delete Peer response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/secondary-dns_id_response-3"}, {"$ref": "#/components/schemas/secondary-dns_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Secondary DNS (Peer)"], "x-api-token-group": ["Account Settings Write"], "x-cfPlanAvailability": {"business": false, "enterprise": true, "free": false, "pro": false}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "dns.zone-transfers.peers", "x-fern-sdk-method-name": "delete"}
```
