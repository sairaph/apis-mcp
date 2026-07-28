---
title: Update custom prompt topic
page_id: operation-put-accounts-account-id-dlp-custom-prompt-topics-entry-id-ff5052b0
path: operations/dlp-custom-prompt-topics
description: Updates a DLP custom prompt topic entry.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /accounts/{account_id}/dlp/custom_prompt_topics/{entry_id}
operation_ids:
    - dlp-custom-prompt-topics-update
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update custom prompt topic

`PUT /accounts/{account_id}/dlp/custom_prompt_topics/{entry_id}`

Operation ID: `dlp-custom-prompt-topics-update`

Updates a DLP custom prompt topic entry.

## Definition

```yaml
{"operationId": "dlp-custom-prompt-topics-update", "summary": "Update custom prompt topic", "description": "Updates a DLP custom prompt topic entry.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string"}}, {"name": "entry_id", "in": "path", "required": true, "schema": {"type": "string", "format": "uuid"}}], "requestBody": {"description": "Update to be applied to the custom prompt topic.", "required": true, "content": {"application/json": {"schema": {"type": "object", "properties": {"description": {"type": "string", "nullable": true}, "enabled": {"type": "boolean"}, "name": {"type": "string"}, "topic": {"type": "string", "maxLength": 50, "minLength": 2}}, "required": ["name", "enabled", "topic"]}}}}, "responses": {"200": {"description": "Update custom prompt topic response.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/dlp_api-response-single"}, {"properties": {"result": {"$ref": "#/components/schemas/dlp_CustomPromptTopic"}}, "type": "object"}]}}}}, "4XX": {"description": "Update custom prompt topic failure response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dlp_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["DLP Custom Prompt Topics"], "x-api-token-group": ["Zero Trust Write"], "x-stainless-skip": ["terraform"]}
```
