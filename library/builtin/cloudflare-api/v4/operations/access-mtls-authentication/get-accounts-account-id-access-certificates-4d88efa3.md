---
title: List mTLS certificates
page_id: operation-get-accounts-account-id-access-certificates-3a2a9373
path: operations/access-mtls-authentication
description: Lists all mTLS root certificates.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/access/certificates
operation_ids:
    - access-mtls-authentication-list-mtls-certificates
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List mTLS certificates

`GET /accounts/{account_id}/access/certificates`

Operation ID: `access-mtls-authentication-list-mtls-certificates`

Lists all mTLS root certificates.

## Definition

```yaml
{"operationId": "access-mtls-authentication-list-mtls-certificates", "summary": "List mTLS certificates", "description": "Lists all mTLS root certificates.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/access_identifier"}}, {"$ref": "#/components/parameters/access_page"}, {"name": "per_page", "in": "query", "schema": {"description": "Number of results per page.", "type": "integer", "default": 50, "maximum": 1000}}], "responses": {"200": {"description": "List mTLS certificates response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_response_collection-5"}}}}, "4XX": {"description": "List mTLS certificates response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Access mTLS authentication"], "x-api-token-group": ["Access: Mutual TLS Certificates Write", "Access: Mutual TLS Certificates Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.access.certificates", "x-fern-sdk-method-name": "list", "x-forge-hidden": true}
```
