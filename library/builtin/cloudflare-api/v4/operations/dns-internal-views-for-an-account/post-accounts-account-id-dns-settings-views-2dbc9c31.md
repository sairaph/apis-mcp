---
title: Create Internal DNS View
page_id: operation-post-accounts-account-id-dns-settings-views-22104d08
path: operations/dns-internal-views-for-an-account
description: Create Internal DNS View for an account
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/dns_settings/views
operation_ids:
    - dns-views-for-an-account-create-internal-dns-views
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create Internal DNS View

`POST /accounts/{account_id}/dns_settings/views`

Operation ID: `dns-views-for-an-account-create-internal-dns-views`

Create Internal DNS View for an account

## Definition

```yaml
{"operationId": "dns-views-for-an-account-create-internal-dns-views", "summary": "Create Internal DNS View", "description": "Create Internal DNS View for an account", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/dns-settings_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dns-settings_dns-view-post"}}}}, "responses": {"200": {"description": "Create Internal DNS View response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dns-settings_dns_view_response_single"}}}}, "4XX": {"description": "Create Internal DNS View response failure", "content": {"application/json": {"schema": {"type": "object", "allOf": [{"$ref": "#/components/schemas/dns-settings_dns_view_response_single"}, {"$ref": "#/components/schemas/dns-settings_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["DNS Internal Views for an Account"], "x-api-token-group": ["DNS View Write"], "x-cfPermissionsRequired": {"enum": ["#dns.view:create"]}, "x-cfPlanAvailability": {"business": false, "enterprise": true, "free": false, "pro": false}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "dns.settings.account.views", "x-fern-sdk-method-name": "create"}
```
