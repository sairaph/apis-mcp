---
title: Delete buckets for full packet captures
page_id: operation-delete-accounts-account-id-pcaps-ownership-ownership-id-f5206b01
path: operations/magic-pcap-collection
description: Deletes buckets added to the packet captures API.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/pcaps/ownership/{ownership_id}
operation_ids:
    - magic-pcap-collection-delete-buckets-for-full-packet-captures
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete buckets for full packet captures

`DELETE /accounts/{account_id}/pcaps/ownership/{ownership_id}`

Operation ID: `magic-pcap-collection-delete-buckets-for-full-packet-captures`

Deletes buckets added to the packet captures API.

## Definition

```yaml
{"operationId": "magic-pcap-collection-delete-buckets-for-full-packet-captures", "summary": "Delete buckets for full packet captures", "description": "Deletes buckets added to the packet captures API.", "parameters": [{"name": "ownership_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/magic-visibility-pcaps_identifier"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/magic-visibility-pcaps_identifier"}}], "responses": {"204": {"description": "Delete buckets for full packet captures response."}, "default": {"description": "Delete buckets for full packet captures response failure.", "content": {"application/json": {}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Magic PCAP collection"], "x-api-token-group": ["Magic Firewall Packet Captures - Write PCAPs API"], "x-cfPlanAvailability": {"business": false, "enterprise": true, "free": false, "pro": false}}
```
