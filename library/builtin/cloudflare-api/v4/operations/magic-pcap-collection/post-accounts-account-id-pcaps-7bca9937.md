---
title: Create PCAP request
page_id: operation-post-accounts-account-id-pcaps-4b891bca
path: operations/magic-pcap-collection
description: Create new PCAP request for account.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/pcaps
operation_ids:
    - magic-pcap-collection-create-pcap-request
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create PCAP request

`POST /accounts/{account_id}/pcaps`

Operation ID: `magic-pcap-collection-create-pcap-request`

Create new PCAP request for account.

## Definition

```yaml
{"operationId": "magic-pcap-collection-create-pcap-request", "summary": "Create PCAP request", "description": "Create new PCAP request for account.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/magic-visibility-pcaps_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/magic-visibility-pcaps_pcaps_request_pcap"}}}}, "responses": {"200": {"description": "Create PCAP request response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/magic-visibility-pcaps_pcaps_single_response"}}}}, "default": {"description": "Create PCAP request response failure.", "content": {"application/json": {"schema": {"oneOf": [{"$ref": "#/components/schemas/magic-visibility-pcaps_pcaps_single_response"}, {"$ref": "#/components/schemas/magic-visibility-pcaps_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Magic PCAP collection"], "x-api-token-group": ["Magic Firewall Packet Captures - Write PCAPs API"], "x-cfPlanAvailability": {"business": false, "enterprise": true, "free": false, "pro": false}}
```
