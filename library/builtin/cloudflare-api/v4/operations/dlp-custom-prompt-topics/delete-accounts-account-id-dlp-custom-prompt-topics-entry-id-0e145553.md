---
title: Delete custom prompt topic
page_id: operation-delete-accounts-account-id-dlp-custom-prompt-topics-entry-id-400aa619
path: operations/dlp-custom-prompt-topics
description: Deletes a DLP custom prompt topic entry.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/dlp/custom_prompt_topics/{entry_id}
operation_ids:
    - dlp-custom-prompt-topics-delete
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete custom prompt topic

`DELETE /accounts/{account_id}/dlp/custom_prompt_topics/{entry_id}`

Operation ID: `dlp-custom-prompt-topics-delete`

Deletes a DLP custom prompt topic entry.

## Definition

```yaml
{"operationId": "dlp-custom-prompt-topics-delete", "summary": "Delete custom prompt topic", "description": "Deletes a DLP custom prompt topic entry.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string"}}, {"name": "entry_id", "in": "path", "required": true, "schema": {"type": "string", "format": "uuid"}}], "responses": {"200": {"description": "Delete custom prompt topic response.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/dlp_api-response-single"}, {"properties": {"result": {"$ref": "#/components/schemas/dlp_Empty"}}, "type": "object"}]}}}}, "4XX": {"description": "Delete custom prompt topic failure response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dlp_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["DLP Custom Prompt Topics"], "x-api-token-group": ["Zero Trust Write"], "x-stainless-skip": ["terraform"]}
```
