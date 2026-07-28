---
title: Update Internal DNS View
page_id: operation-patch-accounts-account-id-dns-settings-views-view-id-a31a53bb
path: operations/dns-internal-views-for-an-account
description: Update an existing Internal DNS View
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PATCH
api_endpoints:
    - /accounts/{account_id}/dns_settings/views/{view_id}
operation_ids:
    - dns-views-for-an-account-update-internal-dns-view
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update Internal DNS View

`PATCH /accounts/{account_id}/dns_settings/views/{view_id}`

Operation ID: `dns-views-for-an-account-update-internal-dns-view`

Update an existing Internal DNS View

## Definition

```yaml
{"operationId": "dns-views-for-an-account-update-internal-dns-view", "summary": "Update Internal DNS View", "description": "Update an existing Internal DNS View", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/dns-settings_identifier"}}, {"name": "view_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/dns-settings_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dns-settings_dns-view-patch"}}}}, "responses": {"200": {"description": "Update Internal DNS View response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dns-settings_dns_view_response_single"}}}}, "4XX": {"description": "Update Internal DNS View response failure", "content": {"application/json": {"schema": {"type": "object", "allOf": [{"$ref": "#/components/schemas/dns-settings_dns_view_response_single"}, {"$ref": "#/components/schemas/dns-settings_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["DNS Internal Views for an Account"], "x-api-token-group": ["DNS View Write"], "x-cfPermissionsRequired": {"enum": ["#dns.view:update"]}, "x-cfPlanAvailability": {"business": false, "enterprise": true, "free": false, "pro": false}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "dns.settings.account.views", "x-fern-sdk-method-name": "edit"}
```
