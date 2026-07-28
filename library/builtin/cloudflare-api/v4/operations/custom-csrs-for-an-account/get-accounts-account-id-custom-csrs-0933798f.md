---
title: List Custom CSRs
page_id: operation-get-accounts-account-id-custom-csrs-9389d020
path: operations/custom-csrs-for-an-account
description: List all custom Certificate Signing Requests (CSRs) for an account.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/custom_csrs
operation_ids:
    - custom-csrs-for-an-account-list-custom-csrs
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List Custom CSRs

`GET /accounts/{account_id}/custom_csrs`

Operation ID: `custom-csrs-for-an-account-list-custom-csrs`

List all custom Certificate Signing Requests (CSRs) for an account.

## Definition

```yaml
{"operationId": "custom-csrs-for-an-account-list-custom-csrs", "summary": "List Custom CSRs", "description": "List all custom Certificate Signing Requests (CSRs) for an account.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_identifier"}}, {"name": "page", "in": "query", "schema": {"description": "Page number of paginated results.", "type": "number", "default": 1, "minimum": 1}}, {"name": "per_page", "in": "query", "schema": {"description": "Number of custom CSRs per page.", "type": "number", "default": 50, "maximum": 50, "minimum": 5}}, {"name": "order", "in": "query", "schema": {"description": "The field to sort the returned custom CSRs by.", "type": "string", "default": "created_at", "enum": ["name", "account_tag", "created_at"]}}, {"name": "direction", "in": "query", "schema": {"description": "The direction to sort by.", "type": "string", "default": "asc", "enum": ["asc", "desc"]}}], "responses": {"200": {"description": "List Custom CSRs response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_custom_csr_response_collection-2"}}}}, "4XX": {"description": "List Custom CSRs response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/tls-certificates-and-hostnames_custom_csr_response_collection-2"}, {"$ref": "#/components/schemas/tls-certificates-and-hostnames_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Custom CSRs for an Account"], "x-api-token-group": ["Account: SSL and Certificates Read", "Account: SSL and Certificates Write"], "x-cfPermissionsRequired": {"enum": ["#ssl:read"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": false, "pro": false}}
```
