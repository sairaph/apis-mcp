---
title: Download current certificate in PEM format
page_id: operation-get-accounts-account-id-access-saml-certificates-saml-cert-set-id-pem-707dc56c
path: operations/access-saml-encryption-certificates
description: Downloads the current SAML encryption certificate's public key in PEM format for the specified certificate set. This endpoint is useful for providing the certificate to Identity Providers for SAML assertion encryption configuration.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/access/saml_certificates/{saml_cert_set_id}/pem
operation_ids:
    - access-saml-certificates-get-pem
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Download current certificate in PEM format

`GET /accounts/{account_id}/access/saml_certificates/{saml_cert_set_id}/pem`

Operation ID: `access-saml-certificates-get-pem`

Downloads the current SAML encryption certificate's public key in PEM format for the specified certificate set. This endpoint is useful for providing the certificate to Identity Providers for SAML assertion encryption configuration.

## Definition

```yaml
{"operationId": "access-saml-certificates-get-pem", "summary": "Download current certificate in PEM format", "description": "Downloads the current SAML encryption certificate's public key in PEM format for the specified certificate set. This endpoint is useful for providing the certificate to Identity Providers for SAML assertion encryption configuration.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/access_identifier"}}, {"name": "saml_cert_set_id", "in": "path", "description": "UID of the SAML certificate set.", "required": true, "schema": {"$ref": "#/components/schemas/access_uuid"}}], "responses": {"200": {"description": "PEM certificate file", "headers": {"Content-Disposition": {"schema": {"type": "string", "example": "attachment; filename=\"saml_certificate.pem\""}}}, "content": {"application/x-pem-file": {"schema": {"type": "string", "example": "-----BEGIN CERTIFICATE-----\nMIIGAjCCA+qgAwIBAgIJAI7kymlF7CWT...\n...certificate content...\n-----END CERTIFICATE-----\n"}}}}, "404": {"description": "SAML certificate set does not exist", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_api-response-common-failure"}}}}, "4XX": {"description": "Get PEM certificate response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Access SAML encryption certificates"]}
```
