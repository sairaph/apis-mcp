---
title: List Access applications
page_id: operation-get-accounts-account-id-access-apps-1a1f5dbc
path: operations/access-applications
description: Lists all Access applications in an account.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/access/apps
operation_ids:
    - access-applications-list-access-applications
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List Access applications

`GET /accounts/{account_id}/access/apps`

Operation ID: `access-applications-list-access-applications`

Lists all Access applications in an account.

## Definition

```yaml
{"operationId": "access-applications-list-access-applications", "summary": "List Access applications", "description": "Lists all Access applications in an account.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/access_identifier"}}, {"name": "name", "in": "query", "schema": {"description": "The name of the app.", "type": "string"}}, {"name": "domain", "in": "query", "schema": {"description": "The domain of the app.", "type": "string"}}, {"name": "aud", "in": "query", "schema": {"description": "The aud of the app.", "type": "string"}}, {"name": "target_attributes", "in": "query", "schema": {"description": "Target Criteria attributes in key=value format.", "type": "string"}}, {"name": "exact", "in": "query", "schema": {"description": "True for only exact string matches against passed name/domain query parameters.", "type": "boolean"}}, {"name": "search", "in": "query", "schema": {"description": "Search for apps by other listed query parameters.", "type": "string"}}, {"$ref": "#/components/parameters/access_page"}, {"$ref": "#/components/parameters/access_per_page"}], "responses": {"200": {"description": "List Access applications response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_response_collection-7"}}}}, "4XX": {"description": "List Access applications response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Access applications"], "x-api-token-group": ["Access: Apps and Policies Revoke", "Access: Apps and Policies Write", "Access: Apps and Policies Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.access.applications", "x-fern-sdk-method-name": "list", "x-forge-hidden": true, "x-stability": "beta"}
```
