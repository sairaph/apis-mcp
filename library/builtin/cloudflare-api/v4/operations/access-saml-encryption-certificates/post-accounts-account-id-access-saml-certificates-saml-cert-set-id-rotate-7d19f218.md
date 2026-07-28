---
title: Rotate SAML certificate
page_id: operation-post-accounts-account-id-access-saml-certificates-saml-cert-set-id-rotat-27da489b
path: operations/access-saml-encryption-certificates
description: "Rotates the SAML encryption certificates within the specified certificate set. This generates a new\ncertificate and moves the current certificate to the previous slot. If a previous certificate exists,\nit will be deactivated and removed.\n\nThis endpoint ensures zero-downtime rotation by maintaining both current and previous certificates \nduring the transition period, allowing IdPs time to update their configurations. Automated rotation\nhappens 30 days before a current certificate's expiration."
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/access/saml_certificates/{saml_cert_set_id}/rotate
operation_ids:
    - access-saml-certificates-rotate-certificate
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Rotate SAML certificate

`POST /accounts/{account_id}/access/saml_certificates/{saml_cert_set_id}/rotate`

Operation ID: `access-saml-certificates-rotate-certificate`

Rotates the SAML encryption certificates within the specified certificate set. This generates a new
certificate and moves the current certificate to the previous slot. If a previous certificate exists,
it will be deactivated and removed.

This endpoint ensures zero-downtime rotation by maintaining both current and previous certificates
during the transition period, allowing IdPs time to update their configurations. Automated rotation
happens 30 days before a current certificate's expiration.

## Definition

```yaml
{"operationId": "access-saml-certificates-rotate-certificate", "summary": "Rotate SAML certificate", "description": "Rotates the SAML encryption certificates within the specified certificate set. This generates a new\ncertificate and moves the current certificate to the previous slot. If a previous certificate exists,\nit will be deactivated and removed.\n\nThis endpoint ensures zero-downtime rotation by maintaining both current and previous certificates \nduring the transition period, allowing IdPs time to update their configurations. Automated rotation\nhappens 30 days before a current certificate's expiration.\n", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/access_identifier"}}, {"name": "saml_cert_set_id", "in": "path", "description": "UID of the SAML certificate set to rotate.", "required": true, "schema": {"$ref": "#/components/schemas/access_uuid"}}], "responses": {"200": {"description": "Rotate SAML certificate response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_certificate_set_response"}}}}, "404": {"description": "SAML certificate set does not exist", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_api-response-common-failure"}}}}, "4XX": {"description": "Rotate SAML certificate response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Access SAML encryption certificates"]}
```
