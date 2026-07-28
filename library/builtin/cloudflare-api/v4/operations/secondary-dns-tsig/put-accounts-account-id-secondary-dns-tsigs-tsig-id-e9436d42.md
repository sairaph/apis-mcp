---
title: Update TSIG
page_id: operation-put-accounts-account-id-secondary-dns-tsigs-tsig-id-17ee0f9f
path: operations/secondary-dns-tsig
description: Modify TSIG.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /accounts/{account_id}/secondary_dns/tsigs/{tsig_id}
operation_ids:
    - secondary-dns-(-tsig)-update-tsig
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update TSIG

`PUT /accounts/{account_id}/secondary_dns/tsigs/{tsig_id}`

Operation ID: `secondary-dns-(-tsig)-update-tsig`

Modify TSIG.

## Definition

```yaml
{"operationId": "secondary-dns-(-tsig)-update-tsig", "summary": "Update TSIG", "description": "Modify TSIG.", "parameters": [{"name": "tsig_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/secondary-dns_identifier-2"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/secondary-dns_account_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/secondary-dns_tsig"}}}}, "responses": {"200": {"description": "Update TSIG response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/secondary-dns_single_response"}}}}, "4XX": {"description": "Update TSIG response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/secondary-dns_single_response"}, {"$ref": "#/components/schemas/secondary-dns_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Secondary DNS (TSIG)"], "x-api-token-group": ["Account Settings Write"], "x-cfPlanAvailability": {"business": false, "enterprise": true, "free": false, "pro": false}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "dns.zone-transfers.tsigs", "x-fern-sdk-method-name": "update"}
```
