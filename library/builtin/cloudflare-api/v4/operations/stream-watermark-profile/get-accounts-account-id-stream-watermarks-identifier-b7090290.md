---
title: Watermark profile details
page_id: operation-get-accounts-account-id-stream-watermarks-identifier-200bfaf0
path: operations/stream-watermark-profile
description: Retrieves details for a single watermark profile.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/stream/watermarks/{identifier}
operation_ids:
    - stream-watermark-profile-watermark-profile-details
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Watermark profile details

`GET /accounts/{account_id}/stream/watermarks/{identifier}`

Operation ID: `stream-watermark-profile-watermark-profile-details`

Retrieves details for a single watermark profile.

## Definition

```yaml
{"operationId": "stream-watermark-profile-watermark-profile-details", "summary": "Watermark profile details", "description": "Retrieves details for a single watermark profile.", "parameters": [{"name": "identifier", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/stream_watermark_identifier"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/stream_account_identifier"}}], "responses": {"200": {"description": "Watermark profile details response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/stream_watermark_response_single"}}}}, "4XX": {"description": "Watermark profile details response failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/stream_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Stream Watermark Profile"], "x-api-token-group": ["Stream Write", "Stream Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "stream.watermarks", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
