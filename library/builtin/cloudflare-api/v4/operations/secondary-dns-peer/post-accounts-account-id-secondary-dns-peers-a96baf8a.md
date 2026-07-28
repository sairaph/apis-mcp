---
title: Create Peer
page_id: operation-post-accounts-account-id-secondary-dns-peers-e09db4ec
path: operations/secondary-dns-peer
description: Create Peer.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/secondary_dns/peers
operation_ids:
    - secondary-dns-(-peer)-create-peer
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create Peer

`POST /accounts/{account_id}/secondary_dns/peers`

Operation ID: `secondary-dns-(-peer)-create-peer`

Create Peer.

## Definition

```yaml
{"operationId": "secondary-dns-(-peer)-create-peer", "summary": "Create Peer", "description": "Create Peer.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/secondary-dns_account_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"type": "object", "properties": {"name": {"$ref": "#/components/schemas/secondary-dns_name-3"}}, "required": ["name"]}}}}, "responses": {"200": {"description": "Create Peer response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/secondary-dns_single_response-2"}}}}, "4XX": {"description": "Create Peer response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/secondary-dns_single_response-2"}, {"$ref": "#/components/schemas/secondary-dns_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Secondary DNS (Peer)"], "x-api-token-group": ["Account Settings Write"], "x-cfPlanAvailability": {"business": false, "enterprise": true, "free": false, "pro": false}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "dns.zone-transfers.peers", "x-fern-sdk-method-name": "create"}
```
