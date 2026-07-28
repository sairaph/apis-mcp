---
title: Update a sending domain restriction
page_id: operation-patch-accounts-account-id-email-security-settings-sending-domain-restric-4a86f7bc
path: operations/email-security-settings
description: Updates an existing sending domain restriction. Only provided fields will be modified. Changes affect which domains require TLS and which subdomains are excluded.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PATCH
api_endpoints:
    - /accounts/{account_id}/email-security/settings/sending_domain_restrictions/{sending_domain_restriction_id}
operation_ids:
    - email_security_update_sending_domain_restriction
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update a sending domain restriction

`PATCH /accounts/{account_id}/email-security/settings/sending_domain_restrictions/{sending_domain_restriction_id}`

Operation ID: `email_security_update_sending_domain_restriction`

Updates an existing sending domain restriction. Only provided fields will be modified. Changes affect which domains require TLS and which subdomains are excluded.

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/email-security_account_id"}, {"name": "sending_domain_restriction_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/email-security_SendingDomainRestrictionId"}}]
```

## Definition

```yaml
{"operationId": "email_security_update_sending_domain_restriction", "summary": "Update a sending domain restriction", "description": "Updates an existing sending domain restriction. Only provided fields will be modified. Changes affect which domains require TLS and which subdomains are excluded.", "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/email-security_UpdateSendingDomainRestriction"}}}}, "responses": {"200": {"description": "Updated sending domain restriction.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/email-security_api-response-single"}, {"properties": {"result": {"$ref": "#/components/schemas/email-security_SendingDomainRestriction"}}, "type": "object"}]}}}}, "4XX": {"$ref": "#/components/responses/email-security_ClientError"}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Email Security Settings"], "x-api-token-group": ["Cloud Email Security: Write"], "x-stability": "beta"}
```
