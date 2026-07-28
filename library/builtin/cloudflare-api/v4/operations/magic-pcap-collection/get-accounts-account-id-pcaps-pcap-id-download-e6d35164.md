---
title: Download Simple PCAP
page_id: operation-get-accounts-account-id-pcaps-pcap-id-download-e404422e
path: operations/magic-pcap-collection
description: Download PCAP information into a file. Response is a binary PCAP file.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/pcaps/{pcap_id}/download
operation_ids:
    - magic-pcap-collection-download-simple-pcap
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Download Simple PCAP

`GET /accounts/{account_id}/pcaps/{pcap_id}/download`

Operation ID: `magic-pcap-collection-download-simple-pcap`

Download PCAP information into a file. Response is a binary PCAP file.

## Definition

```yaml
{"operationId": "magic-pcap-collection-download-simple-pcap", "summary": "Download Simple PCAP", "description": "Download PCAP information into a file. Response is a binary PCAP file.", "parameters": [{"name": "pcap_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/magic-visibility-pcaps_identifier"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/magic-visibility-pcaps_identifier"}}], "responses": {"200": {"description": "Download Simple PCAP response.", "content": {"application/vnd.tcpdump.pcap": {}}}, "default": {"description": "Download Simple PCAP response failure.", "content": {"application/json": {}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Magic PCAP collection"], "x-api-token-group": ["Magic Firewall Packet Captures - Write PCAPs API", "Magic Firewall Packet Captures - Read PCAPs API"], "x-cfPlanAvailability": {"business": false, "enterprise": true, "free": false, "pro": false}}
```
