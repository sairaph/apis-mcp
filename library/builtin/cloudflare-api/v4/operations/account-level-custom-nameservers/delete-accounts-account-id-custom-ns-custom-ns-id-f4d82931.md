---
title: Delete Account Custom Nameserver
page_id: operation-delete-accounts-account-id-custom-ns-custom-ns-id-30b3ca5d
path: operations/account-level-custom-nameservers
description: Removes a custom nameserver from the account.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/custom_ns/{custom_ns_id}
operation_ids:
    - account-level-custom-nameservers-delete-account-custom-nameserver
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete Account Custom Nameserver

`DELETE /accounts/{account_id}/custom_ns/{custom_ns_id}`

Operation ID: `account-level-custom-nameservers-delete-account-custom-nameserver`

Removes a custom nameserver from the account.

## Definition

```yaml
{"operationId": "account-level-custom-nameservers-delete-account-custom-nameserver", "summary": "Delete Account Custom Nameserver", "description": "Removes a custom nameserver from the account.", "parameters": [{"name": "custom_ns_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/dns-custom-nameservers_ns_name"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/dns-custom-nameservers_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {}}}, "responses": {"200": {"description": "Delete Account Custom Nameserver response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dns-custom-nameservers_empty_response"}}}}, "4XX": {"description": "Delete Account Custom Nameserver response failure", "content": {"application/json": {"schema": {"type": "object", "allOf": [{"$ref": "#/components/schemas/dns-custom-nameservers_empty_response"}, {"$ref": "#/components/schemas/dns-custom-nameservers_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Account-Level Custom Nameservers"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": false, "pro": false}}
```
