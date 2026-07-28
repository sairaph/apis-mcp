---
title: Validate buckets for full packet captures
page_id: operation-post-accounts-account-id-pcaps-ownership-validate-85c2fefe
path: operations/magic-pcap-collection
description: Validates buckets added to the packet captures API.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/pcaps/ownership/validate
operation_ids:
    - magic-pcap-collection-validate-buckets-for-full-packet-captures
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Validate buckets for full packet captures

`POST /accounts/{account_id}/pcaps/ownership/validate`

Operation ID: `magic-pcap-collection-validate-buckets-for-full-packet-captures`

Validates buckets added to the packet captures API.

## Definition

```yaml
{"operationId": "magic-pcap-collection-validate-buckets-for-full-packet-captures", "summary": "Validate buckets for full packet captures", "description": "Validates buckets added to the packet captures API.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/magic-visibility-pcaps_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/magic-visibility-pcaps_pcaps_ownership_validate_request"}}}}, "responses": {"200": {"description": "Validate buckets for full packet captures response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/magic-visibility-pcaps_pcaps_ownership_single_response"}}}}, "default": {"description": "Validate buckets for full packet captures response failure.", "content": {"application/json": {"schema": {"oneOf": [{"$ref": "#/components/schemas/magic-visibility-pcaps_pcaps_ownership_single_response"}, {"$ref": "#/components/schemas/magic-visibility-pcaps_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Magic PCAP collection"], "x-api-token-group": ["Magic Firewall Packet Captures - Write PCAPs API"], "x-cfPlanAvailability": {"business": false, "enterprise": true, "free": false, "pro": false}}
```
