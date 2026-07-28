---
title: List short-lived certificate CAs
page_id: operation-get-accounts-account-id-access-apps-ca-bfbf66c6
path: operations/access-short-lived-certificate-cas
description: Lists short-lived certificate CAs and their public keys.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/access/apps/ca
operation_ids:
    - access-short-lived-certificate-c-as-list-short-lived-certificate-c-as
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List short-lived certificate CAs

`GET /accounts/{account_id}/access/apps/ca`

Operation ID: `access-short-lived-certificate-c-as-list-short-lived-certificate-c-as`

Lists short-lived certificate CAs and their public keys.

## Definition

```yaml
{"operationId": "access-short-lived-certificate-c-as-list-short-lived-certificate-c-as", "summary": "List short-lived certificate CAs", "description": "Lists short-lived certificate CAs and their public keys.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/access_identifier"}}, {"$ref": "#/components/parameters/access_page"}, {"name": "per_page", "in": "query", "schema": {"description": "Number of results per page.", "type": "integer", "default": 100, "maximum": 1000}}], "responses": {"200": {"description": "List short-lived certificate CAs response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_response_collection-4"}}}}, "4XX": {"description": "List short-lived certificate CAs response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Access short-lived certificate CAs"], "x-api-token-group": ["Access: Apps and Policies Write", "Access: Apps and Policies Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.access.applications.cas", "x-fern-sdk-method-name": "list", "x-forge-hidden": true}
```
