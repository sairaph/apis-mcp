---
title: Create custom prompt topic
page_id: operation-post-accounts-account-id-dlp-custom-prompt-topics-0a79a06a
path: operations/dlp-custom-prompt-topics
description: Creates a DLP custom prompt topic entry.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/dlp/custom_prompt_topics
operation_ids:
    - dlp-custom-prompt-topics-create
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create custom prompt topic

`POST /accounts/{account_id}/dlp/custom_prompt_topics`

Operation ID: `dlp-custom-prompt-topics-create`

Creates a DLP custom prompt topic entry.

## Definition

```yaml
{"operationId": "dlp-custom-prompt-topics-create", "summary": "Create custom prompt topic", "description": "Creates a DLP custom prompt topic entry.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string"}}], "requestBody": {"description": "A new custom prompt topic to create.", "required": true, "content": {"application/json": {"schema": {"type": "object", "properties": {"description": {"type": "string", "nullable": true}, "enabled": {"type": "boolean"}, "name": {"type": "string"}, "profile_id": {"type": "string", "format": "uuid"}, "topic": {"type": "string", "maxLength": 50, "minLength": 2}}, "required": ["name", "enabled", "topic"]}}}}, "responses": {"200": {"description": "Create custom prompt topic response.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/dlp_api-response-single"}, {"properties": {"result": {"$ref": "#/components/schemas/dlp_CustomPromptTopic"}}, "type": "object"}]}}}}, "4XX": {"description": "Create custom prompt topic failure response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dlp_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["DLP Custom Prompt Topics"], "x-api-token-group": ["Zero Trust Write"], "x-stainless-skip": ["terraform"]}
```
