---
title: Update a URL ignore pattern
page_id: operation-patch-accounts-account-id-email-security-settings-url-ignore-patterns-pa-36d26a88
path: operations/email-security-settings
description: Updates an existing URL rewrite ignore pattern. Only provided fields will be modified.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PATCH
api_endpoints:
    - /accounts/{account_id}/email-security/settings/url_ignore_patterns/{pattern_id}
operation_ids:
    - email_security_update_url_ignore_pattern
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update a URL ignore pattern

`PATCH /accounts/{account_id}/email-security/settings/url_ignore_patterns/{pattern_id}`

Operation ID: `email_security_update_url_ignore_pattern`

Updates an existing URL rewrite ignore pattern. Only provided fields will be modified.

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/email-security_account_id"}, {"name": "pattern_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/email-security_UrlIgnorePatternId"}}]
```

## Definition

```yaml
{"operationId": "email_security_update_url_ignore_pattern", "summary": "Update a URL ignore pattern", "description": "Updates an existing URL rewrite ignore pattern. Only provided fields will be modified.", "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/email-security_UpdateUrlIgnorePattern"}}}}, "responses": {"200": {"description": "Updated URL ignore pattern.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/email-security_api-response-single"}, {"properties": {"result": {"$ref": "#/components/schemas/email-security_UrlIgnorePattern"}}, "type": "object"}]}}}}, "4XX": {"$ref": "#/components/responses/email-security_ClientError"}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Email Security Settings"], "x-api-token-group": ["Cloud Email Security: Write"], "x-stability": "beta"}
```
