---
title: Delete captions or subtitles
page_id: operation-delete-accounts-account-id-stream-identifier-captions-language-55f2b1fb
path: operations/stream-subtitles-captions
description: Removes the captions or subtitles from a video.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/stream/{identifier}/captions/{language}
operation_ids:
    - stream-subtitles/-captions-delete-captions-or-subtitles
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete captions or subtitles

`DELETE /accounts/{account_id}/stream/{identifier}/captions/{language}`

Operation ID: `stream-subtitles/-captions-delete-captions-or-subtitles`

Removes the captions or subtitles from a video.

## Definition

```yaml
{"operationId": "stream-subtitles/-captions-delete-captions-or-subtitles", "summary": "Delete captions or subtitles", "description": "Removes the captions or subtitles from a video.", "parameters": [{"name": "language", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/stream_language"}}, {"name": "identifier", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/stream_identifier"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/stream_identifier-2"}}], "responses": {"200": {"description": "Delete captions or subtitles response.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/stream_api-response-common"}, {"properties": {"result": {"type": "string", "example": "", "x-auditable": true}}, "type": "object"}]}}}}, "4XX": {"description": "Delete captions or subtitles response failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/stream_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Stream Subtitles/Captions"], "x-api-token-group": ["Stream Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "stream.captions.language", "x-fern-sdk-method-name": "delete", "x-forge-hidden": true}
```
