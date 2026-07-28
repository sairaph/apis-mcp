---
title: TSIG Details
page_id: operation-get-accounts-account-id-secondary-dns-tsigs-tsig-id-437aa0c5
path: operations/secondary-dns-tsig
description: Get TSIG.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/secondary_dns/tsigs/{tsig_id}
operation_ids:
    - secondary-dns-(-tsig)-tsig-details
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# TSIG Details

`GET /accounts/{account_id}/secondary_dns/tsigs/{tsig_id}`

Operation ID: `secondary-dns-(-tsig)-tsig-details`

Get TSIG.

## Definition

```yaml
{"operationId": "secondary-dns-(-tsig)-tsig-details", "summary": "TSIG Details", "description": "Get TSIG.", "parameters": [{"name": "tsig_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/secondary-dns_identifier-2"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/secondary-dns_account_identifier"}}], "responses": {"200": {"description": "TSIG Details response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/secondary-dns_single_response"}}}}, "4XX": {"description": "TSIG Details response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/secondary-dns_single_response"}, {"$ref": "#/components/schemas/secondary-dns_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Secondary DNS (TSIG)"], "x-api-token-group": ["Account Settings Write", "Account Settings Read"], "x-cfPlanAvailability": {"business": false, "enterprise": true, "free": false, "pro": false}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "dns.zone-transfers.tsigs", "x-fern-sdk-method-name": "get"}
```
