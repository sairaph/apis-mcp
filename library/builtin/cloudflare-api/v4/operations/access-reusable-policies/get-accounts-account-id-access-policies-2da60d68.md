---
title: List Access reusable policies
page_id: operation-get-accounts-account-id-access-policies-9b8a33e9
path: operations/access-reusable-policies
description: Lists Access reusable policies.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/access/policies
operation_ids:
    - access-policies-list-access-reusable-policies
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List Access reusable policies

`GET /accounts/{account_id}/access/policies`

Operation ID: `access-policies-list-access-reusable-policies`

Lists Access reusable policies.

## Definition

```yaml
{"operationId": "access-policies-list-access-reusable-policies", "summary": "List Access reusable policies", "description": "Lists Access reusable policies.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/access_identifier"}}, {"$ref": "#/components/parameters/access_page"}, {"name": "per_page", "in": "query", "schema": {"description": "Number of results per page.", "type": "integer", "default": 100, "maximum": 1000}}], "responses": {"200": {"description": "List Access reusable policies response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_response_collection-9"}}}}, "4XX": {"description": "List Access reusable policies response failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Access reusable policies"], "x-api-token-group": ["Access: Apps and Policies Write", "Access: Apps and Policies Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.access.policies", "x-fern-sdk-method-name": "list", "x-forge-hidden": true}
```
