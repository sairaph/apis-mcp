---
title: Add Account Custom Nameserver
page_id: operation-post-accounts-account-id-custom-ns-ad5fbc86
path: operations/account-level-custom-nameservers
description: Adds a custom nameserver to the account for use as a vanity nameserver on zones.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/custom_ns
operation_ids:
    - account-level-custom-nameservers-add-account-custom-nameserver
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Add Account Custom Nameserver

`POST /accounts/{account_id}/custom_ns`

Operation ID: `account-level-custom-nameservers-add-account-custom-nameserver`

Adds a custom nameserver to the account for use as a vanity nameserver on zones.

## Definition

```yaml
{"operationId": "account-level-custom-nameservers-add-account-custom-nameserver", "summary": "Add Account Custom Nameserver", "description": "Adds a custom nameserver to the account for use as a vanity nameserver on zones.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/dns-custom-nameservers_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dns-custom-nameservers_CustomNSInput"}}}}, "responses": {"200": {"description": "Add Account Custom Nameserver response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dns-custom-nameservers_acns_response_single"}}}}, "4XX": {"description": "Add Account Custom Nameserver response failure", "content": {"application/json": {"schema": {"type": "object", "allOf": [{"$ref": "#/components/schemas/dns-custom-nameservers_acns_response_single"}, {"$ref": "#/components/schemas/dns-custom-nameservers_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Account-Level Custom Nameservers"], "x-api-token-group": ["Account Settings Write"], "x-cfPermissionsRequired": {"enum": ["#organization:edit"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": false, "pro": false}}
```
