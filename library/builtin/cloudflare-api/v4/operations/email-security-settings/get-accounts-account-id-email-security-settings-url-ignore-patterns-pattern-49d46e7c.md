---
title: Get a URL ignore pattern
page_id: operation-get-accounts-account-id-email-security-settings-url-ignore-patterns-patt-b2858e1e
path: operations/email-security-settings
description: Returns a single URL rewrite ignore pattern by its identifier.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/email-security/settings/url_ignore_patterns/{pattern_id}
operation_ids:
    - email_security_get_url_ignore_pattern
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get a URL ignore pattern

`GET /accounts/{account_id}/email-security/settings/url_ignore_patterns/{pattern_id}`

Operation ID: `email_security_get_url_ignore_pattern`

Returns a single URL rewrite ignore pattern by its identifier.

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/email-security_account_id"}, {"name": "pattern_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/email-security_UrlIgnorePatternId"}}]
```

## Definition

```yaml
{"operationId": "email_security_get_url_ignore_pattern", "summary": "Get a URL ignore pattern", "description": "Returns a single URL rewrite ignore pattern by its identifier.", "responses": {"200": {"description": "URL ignore pattern.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/email-security_api-response-single"}, {"properties": {"result": {"$ref": "#/components/schemas/email-security_UrlIgnorePattern"}}, "type": "object"}]}}}}, "4XX": {"$ref": "#/components/responses/email-security_ClientError"}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Email Security Settings"], "x-api-token-group": ["Cloud Email Security: Write", "Cloud Email Security: Read"], "x-stability": "beta"}
```
