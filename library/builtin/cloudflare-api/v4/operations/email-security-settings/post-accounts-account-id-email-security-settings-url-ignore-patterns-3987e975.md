---
title: Create a URL ignore pattern
page_id: operation-post-accounts-account-id-email-security-settings-url-ignore-patterns-b1bc0720
path: operations/email-security-settings
description: Creates a new URL rewrite ignore pattern. URLs matching this pattern will not be rewritten.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/email-security/settings/url_ignore_patterns
operation_ids:
    - email_security_create_url_ignore_pattern
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create a URL ignore pattern

`POST /accounts/{account_id}/email-security/settings/url_ignore_patterns`

Operation ID: `email_security_create_url_ignore_pattern`

Creates a new URL rewrite ignore pattern. URLs matching this pattern will not be rewritten.

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/email-security_account_id"}]
```

## Definition

```yaml
{"operationId": "email_security_create_url_ignore_pattern", "summary": "Create a URL ignore pattern", "description": "Creates a new URL rewrite ignore pattern. URLs matching this pattern will not be rewritten.", "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/email-security_CreateUrlIgnorePattern"}}}}, "responses": {"201": {"description": "Created URL ignore pattern.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/email-security_api-response-single"}, {"properties": {"result": {"$ref": "#/components/schemas/email-security_UrlIgnorePattern"}}, "type": "object"}]}}}}, "4XX": {"$ref": "#/components/responses/email-security_ClientError"}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Email Security Settings"], "x-api-token-group": ["Cloud Email Security: Write"], "x-stability": "beta"}
```
