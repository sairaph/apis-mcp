---
title: Get DNS Record Usage for Account
page_id: operation-get-accounts-account-id-dns-records-usage-8e9a5797
path: operations/dns-records-for-an-account
description: Get the current DNS record usage and quota for an account. May include internal DNS usage and quota.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/dns_records/usage
operation_ids:
    - dns-records-for-an-account-get-usage
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get DNS Record Usage for Account

`GET /accounts/{account_id}/dns_records/usage`

Operation ID: `dns-records-for-an-account-get-usage`

Get the current DNS record usage and quota for an account. May include internal DNS usage and quota.

## Definition

```yaml
{"operationId": "dns-records-for-an-account-get-usage", "summary": "Get DNS Record Usage for Account", "description": "Get the current DNS record usage and quota for an account. May include internal DNS usage and quota.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/dns-records_identifier"}}], "responses": {"200": {"description": "Get DNS Record Usage response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dns-records_dns_response_account_usage"}}}}, "4XX": {"description": "Get DNS Record Usage response failure.", "content": {"application/json": {"schema": {"type": "object", "allOf": [{"$ref": "#/components/schemas/dns-records_dns_response_account_usage"}, {"$ref": "#/components/schemas/dns-records_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["DNS Records for an Account"], "x-api-token-group": ["Account DNS Settings Write", "Account DNS Settings Read"], "x-cfPermissionsRequired": {"enum": ["#dns_records:read"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "dns.usage.account", "x-fern-sdk-method-name": "get"}
```
