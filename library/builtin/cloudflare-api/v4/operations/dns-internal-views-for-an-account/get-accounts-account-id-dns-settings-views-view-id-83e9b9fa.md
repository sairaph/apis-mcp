---
title: DNS Internal View Details
page_id: operation-get-accounts-account-id-dns-settings-views-view-id-7716ddc8
path: operations/dns-internal-views-for-an-account
description: Get DNS Internal View
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/dns_settings/views/{view_id}
operation_ids:
    - dns-views-for-an-account-get-internal-dns-view
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# DNS Internal View Details

`GET /accounts/{account_id}/dns_settings/views/{view_id}`

Operation ID: `dns-views-for-an-account-get-internal-dns-view`

Get DNS Internal View

## Definition

```yaml
{"operationId": "dns-views-for-an-account-get-internal-dns-view", "summary": "DNS Internal View Details", "description": "Get DNS Internal View", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/dns-settings_identifier"}}, {"name": "view_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/dns-settings_identifier"}}], "responses": {"200": {"description": "Get DNS Internal View response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dns-settings_dns_view_response_single"}}}}, "4XX": {"description": "List Internal DNS Views response failure", "content": {"application/json": {"schema": {"type": "object", "allOf": [{"$ref": "#/components/schemas/dns-settings_dns_view_response_single"}, {"$ref": "#/components/schemas/dns-settings_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["DNS Internal Views for an Account"], "x-api-token-group": ["DNS View Write", "DNS View Read"], "x-cfPermissionsRequired": {"enum": ["#dns.view:read"]}, "x-cfPlanAvailability": {"business": false, "enterprise": true, "free": false, "pro": false}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "dns.settings.account.views", "x-fern-sdk-method-name": "get"}
```
