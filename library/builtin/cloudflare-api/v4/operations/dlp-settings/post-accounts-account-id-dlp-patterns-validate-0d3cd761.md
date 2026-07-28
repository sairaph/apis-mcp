---
title: Validate a DLP regex pattern
page_id: operation-post-accounts-account-id-dlp-patterns-validate-0d956e10
path: operations/dlp-settings
description: |-
    Validates whether this pattern is a valid regular expression. Rejects it if
    the regular expression is too complex or can match an unbounded-length
    string. The regex will be rejected if it uses `*` or `+`. Bound the maximum
    number of characters that can be matched using a range, e.g. `{1,100}`.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/dlp/patterns/validate
operation_ids:
    - dlp-pattern-validate
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Validate a DLP regex pattern

`POST /accounts/{account_id}/dlp/patterns/validate`

Operation ID: `dlp-pattern-validate`

Validates whether this pattern is a valid regular expression. Rejects it if
the regular expression is too complex or can match an unbounded-length
string. The regex will be rejected if it uses `*` or `+`. Bound the maximum
number of characters that can be matched using a range, e.g. `{1,100}`.

## Definition

```yaml
{"operationId": "dlp-pattern-validate", "summary": "Validate a DLP regex pattern", "description": "Validates whether this pattern is a valid regular expression. Rejects it if\nthe regular expression is too complex or can match an unbounded-length\nstring. The regex will be rejected if it uses `*` or `+`. Bound the maximum\nnumber of characters that can be matched using a range, e.g. `{1,100}`.", "parameters": [{"name": "account_id", "in": "path", "description": "Account ID.", "required": true, "schema": {"type": "string"}}], "requestBody": {"description": "Validation query.", "required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dlp_RegexValidationQuery"}}}}, "responses": {"200": {"description": "Validation response.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/dlp_api-response-single"}, {"properties": {"result": {"$ref": "#/components/schemas/dlp_RegexValidationResult"}}, "type": "object"}]}}}}, "4XX": {"description": "Failed to validate.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dlp_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["DLP Settings"], "x-api-token-group": ["Zero Trust Write"]}
```
