---
title: Get extension
page_id: operation-get-accounts-account-id-registrar-extensions-extension-15e7ec40
path: operations/extension
description: |-
    Returns metadata and JSON Schema documents describing the expected input
    structure for registration operations on a specific
    extension (TLD).

    Supports HTTP conditional GET via `ETag`. Include the `ETag` value
    from a previous response in an `If-None-Match` header to receive a
    `304 Not Modified` when the data has not changed.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/registrar/extensions/{extension}
operation_ids:
    - registrar-extension-get
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get extension

`GET /accounts/{account_id}/registrar/extensions/{extension}`

Operation ID: `registrar-extension-get`

Returns metadata and JSON Schema documents describing the expected input
structure for registration operations on a specific
extension (TLD).

Supports HTTP conditional GET via `ETag`. Include the `ETag` value
from a previous response in an `If-None-Match` header to receive a
`304 Not Modified` when the data has not changed.

## Definition

```yaml
{"operationId": "registrar-extension-get", "summary": "Get extension", "description": "Returns metadata and JSON Schema documents describing the expected input\nstructure for registration operations on a specific\nextension (TLD).\n\nSupports HTTP conditional GET via `ETag`. Include the `ETag` value\nfrom a previous response in an `If-None-Match` header to receive a\n`304 Not Modified` when the data has not changed.\n", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/registrar-api_identifier"}}, {"name": "extension", "in": "path", "description": "The extension name (e.g., `com`, `co.uk`).", "required": true, "schema": {"type": "string"}, "example": "com"}], "responses": {"200": {"description": "Successfully returned extension.", "headers": {"ETag": {"description": "Opaque identifier for the current version of the response. Send this value in the `If-None-Match` header on subsequent requests to enable conditional GET.", "schema": {"type": "string"}}}, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/registrar-api_extension-response-single"}}}}, "304": {"description": "Not Modified — client cache is current. Returned when the request includes an `If-None-Match` header matching the current ETag."}, "4XX": {"description": "Get extension failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/registrar-api_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Extension"]}
```
