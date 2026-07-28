---
title: Validate a Web Analytics site hostname
page_id: operation-get-accounts-account-id-rum-site-info-validate-hostname-52ebe056
path: operations/web-analytics
description: Validates that the provided hostname is well-formed, does not contain wildcards, and has a valid TLD. Returns an empty result on success.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/rum/site_info/validate/{hostname}
operation_ids:
    - web-analytics-validate-site-hostname
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Validate a Web Analytics site hostname

`GET /accounts/{account_id}/rum/site_info/validate/{hostname}`

Operation ID: `web-analytics-validate-site-hostname`

Validates that the provided hostname is well-formed, does not contain wildcards, and has a valid TLD. Returns an empty result on success.

## Definition

```yaml
{"operationId": "web-analytics-validate-site-hostname", "summary": "Validate a Web Analytics site hostname", "description": "Validates that the provided hostname is well-formed, does not contain wildcards, and has a valid TLD. Returns an empty result on success.\n", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/rum_identifier"}}, {"name": "hostname", "in": "path", "description": "The hostname to validate (e.g. example.com). The `pattern` below validates the hostname's structure (label syntax) only. In addition, the hostname must end in a valid public suffix (TLD). For the list of valid suffixes and how it is used, see the Public Suffix List: https://wiki.mozilla.org/Public_Suffix_List/Use_Cases\n", "required": true, "schema": {"type": "string", "example": "example.com", "maxLength": 253, "pattern": "^(([a-zA-Z0-9]|[a-zA-Z0-9][a-zA-Z0-9\\-]*[a-zA-Z0-9]?)\\.){1,}([A-Za-z0-9]|[A-Za-z0-9][A-Za-z0-9\\-]*[A-Za-z0-9]?){2,}$"}}], "responses": {"200": {"description": "Hostname is valid.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/rum_empty-response"}}}}, "4XX": {"description": "Failure response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/rum_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Web Analytics"], "x-api-token-group": ["Account Settings Read"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}}
```
