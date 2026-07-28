---
title: Add a new email domain
page_id: operation-post-accounts-account-id-email-security-settings-domains-4c47e19d
path: operations/email-security-settings
description: Protects a new email domain by adding it to Email Security. Accepts a flat configuration object covering all delivery modes. Returns the newly created domain configuration.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/email-security/settings/domains
operation_ids:
    - email_security_create_domains
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Add a new email domain

`POST /accounts/{account_id}/email-security/settings/domains`

Operation ID: `email_security_create_domains`

Protects a new email domain by adding it to Email Security. Accepts a flat configuration object covering all delivery modes. Returns the newly created domain configuration.

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/email-security_account_id"}]
```

## Definition

```yaml
{"operationId": "email_security_create_domains", "summary": "Add a new email domain", "description": "Protects a new email domain by adding it to Email Security. Accepts a flat configuration object covering all delivery modes. Returns the newly created domain configuration.", "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/email-security_CreateDomain"}}}}, "responses": {"201": {"description": "Created domain.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/email-security_api-response-single"}, {"properties": {"result": {"$ref": "#/components/schemas/email-security_Domain"}}, "type": "object"}]}}}}, "4XX": {"$ref": "#/components/responses/email-security_ClientError"}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Email Security Settings"], "x-api-token-group": ["Cloud Email Security: Write"], "x-stability": "beta"}
```
