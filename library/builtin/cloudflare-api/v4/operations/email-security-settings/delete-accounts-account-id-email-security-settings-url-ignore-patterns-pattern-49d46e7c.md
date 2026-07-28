---
title: Delete a URL ignore pattern
page_id: operation-delete-accounts-account-id-email-security-settings-url-ignore-patterns-p-9c908548
path: operations/email-security-settings
description: Removes a URL rewrite ignore pattern. After deletion, URLs matching this pattern will be rewritten again.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/email-security/settings/url_ignore_patterns/{pattern_id}
operation_ids:
    - email_security_delete_url_ignore_pattern
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete a URL ignore pattern

`DELETE /accounts/{account_id}/email-security/settings/url_ignore_patterns/{pattern_id}`

Operation ID: `email_security_delete_url_ignore_pattern`

Removes a URL rewrite ignore pattern. After deletion, URLs matching this pattern will be rewritten again.

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/email-security_account_id"}, {"name": "pattern_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/email-security_UrlIgnorePatternId"}}]
```

## Definition

```yaml
{"operationId": "email_security_delete_url_ignore_pattern", "summary": "Delete a URL ignore pattern", "description": "Removes a URL rewrite ignore pattern. After deletion, URLs matching this pattern will be rewritten again.", "responses": {"200": {"description": "Deleted URL ignore pattern.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/email-security_api-response-single"}, {"properties": {"result": {"$ref": "#/components/schemas/email-security_DeletedUrlIgnorePattern"}}, "type": "object"}]}}}}, "4XX": {"$ref": "#/components/responses/email-security_ClientError"}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Email Security Settings"], "x-api-token-group": ["Cloud Email Security: Write"], "x-stability": "beta"}
```
