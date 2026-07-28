---
title: Get custom prompt topic
page_id: operation-get-accounts-account-id-dlp-custom-prompt-topics-entry-id-55d78f12
path: operations/dlp-custom-prompt-topics
description: Fetches a DLP custom prompt topic entry by ID.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/dlp/custom_prompt_topics/{entry_id}
operation_ids:
    - dlp-custom-prompt-topics-get
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get custom prompt topic

`GET /accounts/{account_id}/dlp/custom_prompt_topics/{entry_id}`

Operation ID: `dlp-custom-prompt-topics-get`

Fetches a DLP custom prompt topic entry by ID.

## Definition

```yaml
{"operationId": "dlp-custom-prompt-topics-get", "summary": "Get custom prompt topic", "description": "Fetches a DLP custom prompt topic entry by ID.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string"}}, {"name": "entry_id", "in": "path", "required": true, "schema": {"type": "string", "format": "uuid"}}], "responses": {"200": {"description": "Get custom prompt topic response.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/dlp_api-response-single"}, {"properties": {"result": {"$ref": "#/components/schemas/dlp_CustomPromptTopic"}}, "type": "object"}]}}}}, "4XX": {"description": "Get custom prompt topic failure response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dlp_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["DLP Custom Prompt Topics"], "x-api-token-group": ["Zero Trust Read", "Zero Trust Write"]}
```
