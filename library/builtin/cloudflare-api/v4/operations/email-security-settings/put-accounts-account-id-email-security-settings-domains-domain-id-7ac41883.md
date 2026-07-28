---
title: Replace an email domain
page_id: operation-put-accounts-account-id-email-security-settings-domains-domain-id-66e4bfb7
path: operations/email-security-settings
description: Replaces all mutable fields of a protected email domain in a single atomic operation. Unlike PATCH, all non-computed fields are required.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /accounts/{account_id}/email-security/settings/domains/{domain_id}
operation_ids:
    - email_security_replace_domain
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Replace an email domain

`PUT /accounts/{account_id}/email-security/settings/domains/{domain_id}`

Operation ID: `email_security_replace_domain`

Replaces all mutable fields of a protected email domain in a single atomic operation. Unlike PATCH, all non-computed fields are required.

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/email-security_account_id"}, {"name": "domain_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/email-security_DomainId"}}]
```

## Definition

```yaml
{"operationId": "email_security_replace_domain", "summary": "Replace an email domain", "description": "Replaces all mutable fields of a protected email domain in a single atomic operation. Unlike PATCH, all non-computed fields are required.", "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/email-security_ReplaceDomain"}}}}, "responses": {"200": {"description": "Replaced domain.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/email-security_api-response-single"}, {"properties": {"result": {"$ref": "#/components/schemas/email-security_Domain"}}, "type": "object"}]}}}}, "4XX": {"$ref": "#/components/responses/email-security_ClientError"}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Email Security Settings"], "x-stability": "beta"}
```
