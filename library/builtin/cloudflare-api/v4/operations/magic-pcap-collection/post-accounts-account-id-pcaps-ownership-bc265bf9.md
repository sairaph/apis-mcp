---
title: Add buckets for full packet captures
page_id: operation-post-accounts-account-id-pcaps-ownership-16d78a90
path: operations/magic-pcap-collection
description: Adds an AWS or GCP bucket to use with full packet captures.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/pcaps/ownership
operation_ids:
    - magic-pcap-collection-add-buckets-for-full-packet-captures
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Add buckets for full packet captures

`POST /accounts/{account_id}/pcaps/ownership`

Operation ID: `magic-pcap-collection-add-buckets-for-full-packet-captures`

Adds an AWS or GCP bucket to use with full packet captures.

## Definition

```yaml
{"operationId": "magic-pcap-collection-add-buckets-for-full-packet-captures", "summary": "Add buckets for full packet captures", "description": "Adds an AWS or GCP bucket to use with full packet captures.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/magic-visibility-pcaps_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/magic-visibility-pcaps_pcaps_ownership_request"}}}}, "responses": {"200": {"description": "Add buckets for full packet captures response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/magic-visibility-pcaps_pcaps_ownership_single_response"}}}}, "default": {"description": "Add buckets for full packet captures response failure.", "content": {"application/json": {"schema": {"oneOf": [{"$ref": "#/components/schemas/magic-visibility-pcaps_pcaps_ownership_single_response"}, {"$ref": "#/components/schemas/magic-visibility-pcaps_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Magic PCAP collection"], "x-api-token-group": ["Magic Firewall Packet Captures - Write PCAPs API"], "x-cfPlanAvailability": {"business": false, "enterprise": true, "free": false, "pro": false}}
```
