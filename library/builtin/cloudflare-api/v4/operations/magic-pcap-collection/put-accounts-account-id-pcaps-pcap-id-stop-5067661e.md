---
title: Stop full PCAP
page_id: operation-put-accounts-account-id-pcaps-pcap-id-stop-9c20a0ea
path: operations/magic-pcap-collection
description: Stop full PCAP.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /accounts/{account_id}/pcaps/{pcap_id}/stop
operation_ids:
    - magic-pcap-collection-stop-full-pcap
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Stop full PCAP

`PUT /accounts/{account_id}/pcaps/{pcap_id}/stop`

Operation ID: `magic-pcap-collection-stop-full-pcap`

Stop full PCAP.

## Definition

```yaml
{"operationId": "magic-pcap-collection-stop-full-pcap", "summary": "Stop full PCAP", "description": "Stop full PCAP.", "parameters": [{"name": "pcap_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/magic-visibility-pcaps_identifier"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/magic-visibility-pcaps_identifier"}}], "responses": {"204": {"description": "Stop full PCAP response."}, "default": {"description": "Stop full PCAP response failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/magic-visibility-pcaps_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Magic PCAP collection"], "x-api-token-group": ["Magic Firewall Packet Captures - Write PCAPs API"], "x-cfPlanAvailability": {"business": false, "enterprise": true, "free": false, "pro": false}}
```
