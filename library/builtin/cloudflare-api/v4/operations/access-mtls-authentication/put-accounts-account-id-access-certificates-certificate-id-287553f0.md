---
title: Update an mTLS certificate
page_id: operation-put-accounts-account-id-access-certificates-certificate-id-2544af3c
path: operations/access-mtls-authentication
description: Updates a configured mTLS certificate.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /accounts/{account_id}/access/certificates/{certificate_id}
operation_ids:
    - access-mtls-authentication-update-an-mtls-certificate
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update an mTLS certificate

`PUT /accounts/{account_id}/access/certificates/{certificate_id}`

Operation ID: `access-mtls-authentication-update-an-mtls-certificate`

Updates a configured mTLS certificate.

## Definition

```yaml
{"operationId": "access-mtls-authentication-update-an-mtls-certificate", "summary": "Update an mTLS certificate", "description": "Updates a configured mTLS certificate.", "parameters": [{"name": "certificate_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/access_uuid"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/access_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"properties": {"associated_hostnames": {"$ref": "#/components/schemas/access_associated_hostnames"}, "name": {"$ref": "#/components/schemas/access_name-7"}}, "required": ["associated_hostnames"]}}}}, "responses": {"200": {"description": "Update an mTLS certificate response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_single_response-6"}}}}, "4XX": {"description": "Update an mTLS certificate response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Access mTLS authentication"], "x-api-token-group": ["Access: Mutual TLS Certificates Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.access.certificates", "x-fern-sdk-method-name": "update", "x-forge-hidden": true}
```
