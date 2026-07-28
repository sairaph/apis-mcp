---
title: Add an mTLS certificate
page_id: operation-post-accounts-account-id-access-certificates-cfdb1e51
path: operations/access-mtls-authentication
description: Adds a new mTLS root certificate to Access.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/access/certificates
operation_ids:
    - access-mtls-authentication-add-an-mtls-certificate
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Add an mTLS certificate

`POST /accounts/{account_id}/access/certificates`

Operation ID: `access-mtls-authentication-add-an-mtls-certificate`

Adds a new mTLS root certificate to Access.

## Definition

```yaml
{"operationId": "access-mtls-authentication-add-an-mtls-certificate", "summary": "Add an mTLS certificate", "description": "Adds a new mTLS root certificate to Access.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/access_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"properties": {"associated_hostnames": {"$ref": "#/components/schemas/access_associated_hostnames"}, "certificate": {"description": "The certificate content.", "type": "string", "example": "-----BEGIN CERTIFICATE-----\nMIIGAjCCA+qgAwIBAgIJAI7kymlF7CWT...N4RI7KKB7nikiuUf8vhULKy5IX10\nDrUtmu/B\n-----END CERTIFICATE-----"}, "name": {"$ref": "#/components/schemas/access_name-7"}}, "required": ["name", "certificate"]}}}}, "responses": {"201": {"description": "Add an mTLS certificate response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_single_response-6"}}}}, "4XX": {"description": "Add an mTLS certificate response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Access mTLS authentication"], "x-api-token-group": ["Access: Mutual TLS Certificates Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.access.certificates", "x-fern-sdk-method-name": "create", "x-forge-hidden": true}
```
