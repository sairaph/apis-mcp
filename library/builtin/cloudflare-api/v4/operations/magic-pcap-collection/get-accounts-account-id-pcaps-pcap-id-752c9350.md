---
title: Get PCAP request
page_id: operation-get-accounts-account-id-pcaps-pcap-id-ec3d113e
path: operations/magic-pcap-collection
description: Get information for a PCAP request by id.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/pcaps/{pcap_id}
operation_ids:
    - magic-pcap-collection-get-pcap-request
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get PCAP request

`GET /accounts/{account_id}/pcaps/{pcap_id}`

Operation ID: `magic-pcap-collection-get-pcap-request`

Get information for a PCAP request by id.

## Definition

```yaml
{"operationId": "magic-pcap-collection-get-pcap-request", "summary": "Get PCAP request", "description": "Get information for a PCAP request by id.", "parameters": [{"name": "pcap_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/magic-visibility-pcaps_identifier"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/magic-visibility-pcaps_identifier"}}], "responses": {"200": {"description": "Get PCAP request response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/magic-visibility-pcaps_pcaps_single_response"}}}}, "default": {"description": "Get PCAP request response failure.", "content": {"application/json": {"schema": {"oneOf": [{"$ref": "#/components/schemas/magic-visibility-pcaps_pcaps_single_response"}, {"$ref": "#/components/schemas/magic-visibility-pcaps_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Magic PCAP collection"], "x-api-token-group": ["Magic Firewall Packet Captures - Write PCAPs API", "Magic Firewall Packet Captures - Read PCAPs API"], "x-cfPlanAvailability": {"business": false, "enterprise": true, "free": false, "pro": false}}
```
