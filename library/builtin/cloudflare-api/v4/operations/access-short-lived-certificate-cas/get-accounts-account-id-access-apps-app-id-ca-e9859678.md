---
title: Get a short-lived certificate CA
page_id: operation-get-accounts-account-id-access-apps-app-id-ca-4ba58d13
path: operations/access-short-lived-certificate-cas
description: Fetches a short-lived certificate CA and its public key.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/access/apps/{app_id}/ca
operation_ids:
    - access-short-lived-certificate-c-as-get-a-short-lived-certificate-ca
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get a short-lived certificate CA

`GET /accounts/{account_id}/access/apps/{app_id}/ca`

Operation ID: `access-short-lived-certificate-c-as-get-a-short-lived-certificate-ca`

Fetches a short-lived certificate CA and its public key.

## Definition

```yaml
{"operationId": "access-short-lived-certificate-c-as-get-a-short-lived-certificate-ca", "summary": "Get a short-lived certificate CA", "description": "Fetches a short-lived certificate CA and its public key.", "parameters": [{"name": "app_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/access_uuid"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/access_identifier"}}], "responses": {"200": {"description": "Get a short-lived certificate CA response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_single_response-5"}}}}, "4XX": {"description": "Get a short-lived certificate CA response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Access short-lived certificate CAs"], "x-api-token-group": ["Access: Apps and Policies Write", "Access: Apps and Policies Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.access.applications.cas", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
