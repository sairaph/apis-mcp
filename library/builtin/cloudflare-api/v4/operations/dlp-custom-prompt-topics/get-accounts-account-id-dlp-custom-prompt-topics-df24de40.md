---
title: List custom prompt topics
page_id: operation-get-accounts-account-id-dlp-custom-prompt-topics-a2db5663
path: operations/dlp-custom-prompt-topics
description: Lists all DLP custom prompt topic entries in an account.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/dlp/custom_prompt_topics
operation_ids:
    - dlp-custom-prompt-topics-list
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List custom prompt topics

`GET /accounts/{account_id}/dlp/custom_prompt_topics`

Operation ID: `dlp-custom-prompt-topics-list`

Lists all DLP custom prompt topic entries in an account.

## Definition

```yaml
{"operationId": "dlp-custom-prompt-topics-list", "summary": "List custom prompt topics", "description": "Lists all DLP custom prompt topic entries in an account.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string"}}], "responses": {"200": {"description": "List custom prompt topics response.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/dlp_api-response-single"}, {"properties": {"result": {"type": "array", "items": {"$ref": "#/components/schemas/dlp_CustomPromptTopic"}}}, "type": "object"}]}}}}, "4XX": {"description": "List custom prompt topics failure response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dlp_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["DLP Custom Prompt Topics"], "x-api-token-group": ["Zero Trust Read", "Zero Trust Write"]}
```
