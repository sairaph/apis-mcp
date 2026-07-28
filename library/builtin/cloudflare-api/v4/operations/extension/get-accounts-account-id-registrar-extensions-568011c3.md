---
title: List extensions
page_id: operation-get-accounts-account-id-registrar-extensions-9e377467
path: operations/extension
description: |-
    Returns metadata and JSON Schema documents describing the expected input
    structure for registration operations on each supported
    extension (TLD).

    This endpoint uses cursor-based pagination. Results are ordered by
    extension name by default. To fetch the next page, pass the `cursor`
    value from the `result_info` object in the response as the `cursor`
    query parameter in your next request. An empty `cursor` string
    indicates there are no more pages.

    Supports HTTP conditional GET via `ETag`. Include the `ETag` value
    from a previous response in an `If-None-Match` header to receive a
    `304 Not Modified` when the data has not changed.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/registrar/extensions
operation_ids:
    - registrar-extension-list
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List extensions

`GET /accounts/{account_id}/registrar/extensions`

Operation ID: `registrar-extension-list`

Returns metadata and JSON Schema documents describing the expected input
structure for registration operations on each supported
extension (TLD).

This endpoint uses cursor-based pagination. Results are ordered by
extension name by default. To fetch the next page, pass the `cursor`
value from the `result_info` object in the response as the `cursor`
query parameter in your next request. An empty `cursor` string
indicates there are no more pages.

Supports HTTP conditional GET via `ETag`. Include the `ETag` value
from a previous response in an `If-None-Match` header to receive a
`304 Not Modified` when the data has not changed.

## Definition

```yaml
{"operationId": "registrar-extension-list", "summary": "List extensions", "description": "Returns metadata and JSON Schema documents describing the expected input\nstructure for registration operations on each supported\nextension (TLD).\n\nThis endpoint uses cursor-based pagination. Results are ordered by\nextension name by default. To fetch the next page, pass the `cursor`\nvalue from the `result_info` object in the response as the `cursor`\nquery parameter in your next request. An empty `cursor` string\nindicates there are no more pages.\n\nSupports HTTP conditional GET via `ETag`. Include the `ETag` value\nfrom a previous response in an `If-None-Match` header to receive a\n`304 Not Modified` when the data has not changed.\n", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/registrar-api_identifier"}}, {"name": "name", "in": "query", "description": "Filter extensions by exact name match.\nFor example, `name=com` returns only the `com` extension.\n", "schema": {"type": "string"}, "example": "com"}, {"$ref": "#/components/parameters/registrar-api_cursor_pagination_cursor"}, {"$ref": "#/components/parameters/registrar-api_cursor_pagination_per_page"}, {"$ref": "#/components/parameters/registrar-api_cursor_pagination_direction"}, {"name": "sort_by", "in": "query", "description": "Column to sort results by. Defaults to `name` when omitted.\n", "schema": {"type": "string", "default": "name", "enum": ["name", "created_at", "updated_at"]}}], "responses": {"200": {"description": "Successfully returned extensions.", "headers": {"ETag": {"description": "Opaque identifier for the current version of the response. Send this value in the `If-None-Match` header on subsequent requests to enable conditional GET.", "schema": {"type": "string"}}}, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/registrar-api_extension-response-collection"}}}}, "304": {"description": "Not Modified — client cache is current. Returned when the request includes an `If-None-Match` header matching the current ETag."}, "4XX": {"description": "List extensions failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/registrar-api_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Extension"]}
```
