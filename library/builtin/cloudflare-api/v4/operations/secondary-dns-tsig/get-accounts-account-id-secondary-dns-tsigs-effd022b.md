---
title: List TSIGs
page_id: operation-get-accounts-account-id-secondary-dns-tsigs-f1cf9442
path: operations/secondary-dns-tsig
description: List TSIGs.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/secondary_dns/tsigs
operation_ids:
    - secondary-dns-(-tsig)-list-tsi-gs
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List TSIGs

`GET /accounts/{account_id}/secondary_dns/tsigs`

Operation ID: `secondary-dns-(-tsig)-list-tsi-gs`

List TSIGs.

## Definition

```yaml
{"operationId": "secondary-dns-(-tsig)-list-tsi-gs", "summary": "List TSIGs", "description": "List TSIGs.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/secondary-dns_account_identifier"}}], "responses": {"200": {"description": "List TSIGs response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/secondary-dns_response_collection"}}}}, "4XX": {"description": "List TSIGs response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/secondary-dns_response_collection"}, {"$ref": "#/components/schemas/secondary-dns_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Secondary DNS (TSIG)"], "x-api-token-group": ["Account Settings Write", "Account Settings Read"], "x-cfPlanAvailability": {"business": false, "enterprise": true, "free": false, "pro": false}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "dns.zone-transfers.tsigs", "x-fern-sdk-method-name": "list"}
```
