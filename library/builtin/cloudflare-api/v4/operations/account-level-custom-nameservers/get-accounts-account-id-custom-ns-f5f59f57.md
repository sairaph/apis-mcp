---
title: List Account Custom Nameservers
page_id: operation-get-accounts-account-id-custom-ns-54cc1737
path: operations/account-level-custom-nameservers
description: List an account's custom nameservers.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/custom_ns
operation_ids:
    - account-level-custom-nameservers-list-account-custom-nameservers
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List Account Custom Nameservers

`GET /accounts/{account_id}/custom_ns`

Operation ID: `account-level-custom-nameservers-list-account-custom-nameservers`

List an account's custom nameservers.

## Definition

```yaml
{"operationId": "account-level-custom-nameservers-list-account-custom-nameservers", "summary": "List Account Custom Nameservers", "description": "List an account's custom nameservers.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/dns-custom-nameservers_identifier"}}], "responses": {"200": {"description": "List Account Custom Nameservers response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dns-custom-nameservers_acns_response_collection"}}}}, "4XX": {"description": "List Account Custom Nameservers response failure", "content": {"application/json": {"schema": {"type": "object", "allOf": [{"$ref": "#/components/schemas/dns-custom-nameservers_acns_response_collection"}, {"$ref": "#/components/schemas/dns-custom-nameservers_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Account-Level Custom Nameservers"], "x-api-token-group": ["Account Settings Write", "Account Settings Read"], "x-cfPermissionsRequired": {"enum": ["#organization:read"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": false, "pro": false}}
```
