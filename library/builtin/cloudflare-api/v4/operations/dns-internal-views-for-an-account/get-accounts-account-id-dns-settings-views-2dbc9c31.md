---
title: List Internal DNS Views
page_id: operation-get-accounts-account-id-dns-settings-views-7573ad1e
path: operations/dns-internal-views-for-an-account
description: List DNS Internal Views for an Account
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/dns_settings/views
operation_ids:
    - dns-views-for-an-account-list-internal-dns-views
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List Internal DNS Views

`GET /accounts/{account_id}/dns_settings/views`

Operation ID: `dns-views-for-an-account-list-internal-dns-views`

List DNS Internal Views for an Account

## Definition

```yaml
{"operationId": "dns-views-for-an-account-list-internal-dns-views", "summary": "List Internal DNS Views", "description": "List DNS Internal Views for an Account", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/dns-settings_identifier"}}, {"name": "name", "in": "query", "schema": {"description": "Exact value of the DNS view name. This is a convenience alias for `name.exact`.\n", "type": "string", "example": "my view"}}, {"name": "name.exact", "in": "query", "schema": {"description": "Exact value of the DNS view name.\n", "type": "string", "example": "my view"}}, {"name": "name.contains", "in": "query", "schema": {"description": "Substring of the DNS view name.\n", "type": "string", "example": "view"}}, {"name": "name.startswith", "in": "query", "schema": {"description": "Prefix of the DNS view name.\n", "type": "string", "example": "my"}}, {"name": "name.endswith", "in": "query", "schema": {"description": "Suffix of the DNS view name.\n", "type": "string", "example": "ew"}}, {"name": "zone_id", "in": "query", "schema": {"description": "A zone ID that exists in the zones list for the view.\n", "type": "string", "example": "ae29bea30e2e427ba9cd8d78b628177b"}}, {"name": "zone_name", "in": "query", "schema": {"description": "A zone name that exists in the zones list for the view.\n", "type": "string", "example": "www.example.com"}}, {"name": "match", "in": "query", "schema": {"$ref": "#/components/schemas/dns-settings_match"}}, {"name": "page", "in": "query", "schema": {"$ref": "#/components/schemas/dns-settings_page"}}, {"name": "per_page", "in": "query", "schema": {"$ref": "#/components/schemas/dns-settings_per_page"}}, {"name": "order", "in": "query", "schema": {"$ref": "#/components/schemas/dns-settings_order"}}, {"name": "direction", "in": "query", "schema": {"$ref": "#/components/schemas/dns-settings_direction"}}], "responses": {"200": {"description": "List Internal DNS Views response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dns-settings_dns_view_response_collection"}}}}, "4XX": {"description": "List Internal DNS Views response failure", "content": {"application/json": {"schema": {"type": "object", "allOf": [{"$ref": "#/components/schemas/dns-settings_dns_view_response_collection"}, {"$ref": "#/components/schemas/dns-settings_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["DNS Internal Views for an Account"], "x-api-token-group": ["DNS View Write", "DNS View Read"], "x-cfPermissionsRequired": {"enum": ["#dns.view:list"]}, "x-cfPlanAvailability": {"business": false, "enterprise": true, "free": false, "pro": false}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "dns.settings.account.views", "x-fern-sdk-method-name": "list"}
```
