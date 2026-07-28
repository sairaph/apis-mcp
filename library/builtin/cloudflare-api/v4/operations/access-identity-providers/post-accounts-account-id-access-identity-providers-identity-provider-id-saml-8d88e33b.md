---
title: Create SAML encryption certificate for Identity Provider
page_id: operation-post-accounts-account-id-access-identity-providers-identity-provider-id-64d9bf00
path: operations/access-identity-providers
description: "Creates a new SAML encryption certificate set and assigns it to the specified \nSAML Identity Provider. This endpoint is idempotent - if the IdP already has \na certificate set assigned, the existing certificate set is returned with a 200 status.\n\n**Workflow for enabling SAML encryption:**\n1. Call this endpoint to create and assign a certificate set to the IdP\n2. Update the IdP configuration (PUT `/identity_providers/{id}`) with:\n   - `config.enable_encryption: true`\n   - `saml_certificate_set_id: <uid from step 1>`\n3. Configure the certificate's public key in your external SAML Identity Provider"
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/access/identity_providers/{identity_provider_id}/saml_certificate
operation_ids:
    - access-identity-providers-create-saml-certificate-for-identity-provider
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create SAML encryption certificate for Identity Provider

`POST /accounts/{account_id}/access/identity_providers/{identity_provider_id}/saml_certificate`

Operation ID: `access-identity-providers-create-saml-certificate-for-identity-provider`

Creates a new SAML encryption certificate set and assigns it to the specified
SAML Identity Provider. This endpoint is idempotent - if the IdP already has
a certificate set assigned, the existing certificate set is returned with a 200 status.

**Workflow for enabling SAML encryption:**
1. Call this endpoint to create and assign a certificate set to the IdP
2. Update the IdP configuration (PUT `/identity_providers/{id}`) with:
   - `config.enable_encryption: true`
   - `saml_certificate_set_id: <uid from step 1>`
3. Configure the certificate's public key in your external SAML Identity Provider

## Definition

```yaml
{"operationId": "access-identity-providers-create-saml-certificate-for-identity-provider", "summary": "Create SAML encryption certificate for Identity Provider", "description": "Creates a new SAML encryption certificate set and assigns it to the specified \nSAML Identity Provider. This endpoint is idempotent - if the IdP already has \na certificate set assigned, the existing certificate set is returned with a 200 status.\n\n**Workflow for enabling SAML encryption:**\n1. Call this endpoint to create and assign a certificate set to the IdP\n2. Update the IdP configuration (PUT `/identity_providers/{id}`) with:\n   - `config.enable_encryption: true`\n   - `saml_certificate_set_id: <uid from step 1>`\n3. Configure the certificate's public key in your external SAML Identity Provider\n", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/access_identifier"}}, {"name": "identity_provider_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/access_uuid"}}], "responses": {"200": {"description": "IdP already has a certificate set assigned (idempotent)", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_saml_certificate_set_response"}}}}, "201": {"description": "SAML certificate set created and assigned to IdP", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_saml_certificate_set_response"}}}}, "403": {"description": "SAML encryption not enabled for organization", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_api-response-common-failure"}}}}, "404": {"description": "Identity provider not found", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_api-response-common-failure"}}}}, "4XX": {"description": "Create SAML certificate for IdP response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Access identity providers"]}
```
