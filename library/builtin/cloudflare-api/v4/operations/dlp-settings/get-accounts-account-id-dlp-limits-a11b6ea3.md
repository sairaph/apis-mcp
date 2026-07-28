---
title: Fetch limits associated with DLP for account
page_id: operation-get-accounts-account-id-dlp-limits-21933df2
path: operations/dlp-settings
description: |-
    Retrieves current DLP usage limits and quotas for the account, including
    maximum allowed counts and current usage for custom entries, dataset cells,
    and document fingerprints.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/dlp/limits
operation_ids:
    - dlp-limits-get
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Fetch limits associated with DLP for account

`GET /accounts/{account_id}/dlp/limits`

Operation ID: `dlp-limits-get`

Retrieves current DLP usage limits and quotas for the account, including
maximum allowed counts and current usage for custom entries, dataset cells,
and document fingerprints.

## Definition

```yaml
{"operationId": "dlp-limits-get", "summary": "Fetch limits associated with DLP for account", "description": "Retrieves current DLP usage limits and quotas for the account, including\nmaximum allowed counts and current usage for custom entries, dataset cells,\nand document fingerprints.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string"}}], "responses": {"200": {"description": "Limits retrieved successfully.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/dlp_api-response-single"}, {"properties": {"result": {"$ref": "#/components/schemas/dlp_Limits"}}, "type": "object"}]}}}}, "4XX": {"description": "Limits get failed.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dlp_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["DLP Settings"], "x-api-token-group": ["Zero Trust Read", "Zero Trust Write"]}
```
