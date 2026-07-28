---
title: List Peers
page_id: operation-get-accounts-account-id-secondary-dns-peers-1518b783
path: operations/secondary-dns-peer
description: List Peers.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/secondary_dns/peers
operation_ids:
    - secondary-dns-(-peer)-list-peers
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List Peers

`GET /accounts/{account_id}/secondary_dns/peers`

Operation ID: `secondary-dns-(-peer)-list-peers`

List Peers.

## Definition

```yaml
{"operationId": "secondary-dns-(-peer)-list-peers", "summary": "List Peers", "description": "List Peers.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/secondary-dns_account_identifier"}}], "responses": {"200": {"description": "List Peers response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/secondary-dns_response_collection-2"}}}}, "4XX": {"description": "List Peers response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/secondary-dns_response_collection-2"}, {"$ref": "#/components/schemas/secondary-dns_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Secondary DNS (Peer)"], "x-api-token-group": ["Account Settings Write", "Account Settings Read"], "x-cfPlanAvailability": {"business": false, "enterprise": true, "free": false, "pro": false}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "dns.zone-transfers.peers", "x-fern-sdk-method-name": "list"}
```
