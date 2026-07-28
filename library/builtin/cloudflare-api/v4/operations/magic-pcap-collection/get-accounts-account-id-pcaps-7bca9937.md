---
title: List packet capture requests
page_id: operation-get-accounts-account-id-pcaps-7f4ced41
path: operations/magic-pcap-collection
description: Lists all packet capture requests for an account.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/pcaps
operation_ids:
    - magic-pcap-collection-list-packet-capture-requests
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List packet capture requests

`GET /accounts/{account_id}/pcaps`

Operation ID: `magic-pcap-collection-list-packet-capture-requests`

Lists all packet capture requests for an account.

## Definition

```yaml
{"operationId": "magic-pcap-collection-list-packet-capture-requests", "summary": "List packet capture requests", "description": "Lists all packet capture requests for an account.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/magic-visibility-pcaps_identifier"}}], "responses": {"200": {"description": "List packet capture requests response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/magic-visibility-pcaps_pcaps_collection_response"}}}}, "default": {"description": "List packet capture requests response failure.", "content": {"application/json": {"schema": {"oneOf": [{"$ref": "#/components/schemas/magic-visibility-pcaps_pcaps_collection_response"}, {"$ref": "#/components/schemas/magic-visibility-pcaps_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Magic PCAP collection"], "x-api-token-group": ["Magic Firewall Packet Captures - Write PCAPs API", "Magic Firewall Packet Captures - Read PCAPs API"], "x-cfPlanAvailability": {"business": false, "enterprise": true, "free": false, "pro": false}}
```
