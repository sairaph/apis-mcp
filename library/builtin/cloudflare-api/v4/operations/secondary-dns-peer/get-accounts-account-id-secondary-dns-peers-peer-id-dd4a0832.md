---
title: Peer Details
page_id: operation-get-accounts-account-id-secondary-dns-peers-peer-id-4b8a87d0
path: operations/secondary-dns-peer
description: Get Peer.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/secondary_dns/peers/{peer_id}
operation_ids:
    - secondary-dns-(-peer)-peer-details
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Peer Details

`GET /accounts/{account_id}/secondary_dns/peers/{peer_id}`

Operation ID: `secondary-dns-(-peer)-peer-details`

Get Peer.

## Definition

```yaml
{"operationId": "secondary-dns-(-peer)-peer-details", "summary": "Peer Details", "description": "Get Peer.", "parameters": [{"name": "peer_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/secondary-dns_identifier-3"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/secondary-dns_account_identifier"}}], "responses": {"200": {"description": "Peer Details response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/secondary-dns_single_response-2"}}}}, "4XX": {"description": "Peer Details response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/secondary-dns_single_response-2"}, {"$ref": "#/components/schemas/secondary-dns_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Secondary DNS (Peer)"], "x-api-token-group": ["Account Settings Write", "Account Settings Read"], "x-cfPlanAvailability": {"business": false, "enterprise": true, "free": false, "pro": false}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "dns.zone-transfers.peers", "x-fern-sdk-method-name": "get"}
```
