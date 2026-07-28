---
title: Delete watermark profiles
page_id: operation-delete-accounts-account-id-stream-watermarks-identifier-264b5414
path: operations/stream-watermark-profile
description: Deletes a watermark profile.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/stream/watermarks/{identifier}
operation_ids:
    - stream-watermark-profile-delete-watermark-profiles
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete watermark profiles

`DELETE /accounts/{account_id}/stream/watermarks/{identifier}`

Operation ID: `stream-watermark-profile-delete-watermark-profiles`

Deletes a watermark profile.

## Definition

```yaml
{"operationId": "stream-watermark-profile-delete-watermark-profiles", "summary": "Delete watermark profiles", "description": "Deletes a watermark profile.", "parameters": [{"name": "identifier", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/stream_watermark_identifier"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/stream_account_identifier"}}], "responses": {"200": {"description": "Delete watermark profiles response.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/stream_api-response-single"}, {"properties": {"result": {"type": "string", "example": "", "x-auditable": true}}, "type": "object"}]}}}}, "4XX": {"description": "Delete watermark profiles response failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/stream_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Stream Watermark Profile"], "x-api-token-group": ["Stream Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "stream.watermarks", "x-fern-sdk-method-name": "delete", "x-forge-hidden": true}
```
