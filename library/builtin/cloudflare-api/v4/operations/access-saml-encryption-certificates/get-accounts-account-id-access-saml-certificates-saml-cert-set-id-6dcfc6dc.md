---
title: Get SAML certificate set
page_id: operation-get-accounts-account-id-access-saml-certificates-saml-cert-set-id-3900f8f6
path: operations/access-saml-encryption-certificates
description: Retrieves a specific SAML encryption certificate set by its UID, including both current and previous certificates if available.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/access/saml_certificates/{saml_cert_set_id}
operation_ids:
    - access-saml-certificates-get-certificate-set
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get SAML certificate set

`GET /accounts/{account_id}/access/saml_certificates/{saml_cert_set_id}`

Operation ID: `access-saml-certificates-get-certificate-set`

Retrieves a specific SAML encryption certificate set by its UID, including both current and previous certificates if available.

## Definition

```yaml
{"operationId": "access-saml-certificates-get-certificate-set", "summary": "Get SAML certificate set", "description": "Retrieves a specific SAML encryption certificate set by its UID, including both current and previous certificates if available.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/access_identifier"}}, {"name": "saml_cert_set_id", "in": "path", "description": "UID of the SAML certificate set.", "required": true, "schema": {"$ref": "#/components/schemas/access_uuid"}}], "responses": {"200": {"description": "Get SAML certificate set response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_certificate_set_response"}}}}, "404": {"description": "SAML certificate set does not exist", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_api-response-common-failure"}}}}, "4XX": {"description": "Get SAML certificate set response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Access SAML encryption certificates"]}
```
