---
title: List PCAPs Bucket Ownership
page_id: operation-get-accounts-account-id-pcaps-ownership-457654c4
path: operations/magic-pcap-collection
description: List all buckets configured for use with PCAPs API.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/pcaps/ownership
operation_ids:
    - magic-pcap-collection-list-pca-ps-bucket-ownership
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List PCAPs Bucket Ownership

`GET /accounts/{account_id}/pcaps/ownership`

Operation ID: `magic-pcap-collection-list-pca-ps-bucket-ownership`

List all buckets configured for use with PCAPs API.

## Definition

```yaml
{"operationId": "magic-pcap-collection-list-pca-ps-bucket-ownership", "summary": "List PCAPs Bucket Ownership", "description": "List all buckets configured for use with PCAPs API.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/magic-visibility-pcaps_identifier"}}], "responses": {"200": {"description": "List PCAPs Bucket Ownership response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/magic-visibility-pcaps_pcaps_ownership_collection"}}}}, "default": {"description": "List PCAPs Bucket Ownership response failure.", "content": {"application/json": {"schema": {"oneOf": [{"$ref": "#/components/schemas/magic-visibility-pcaps_pcaps_ownership_collection"}, {"$ref": "#/components/schemas/magic-visibility-pcaps_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Magic PCAP collection"], "x-api-token-group": ["Magic Firewall Packet Captures - Write PCAPs API", "Magic Firewall Packet Captures - Read PCAPs API"], "x-cfPlanAvailability": {"business": false, "enterprise": true, "free": false, "pro": false}}
```
