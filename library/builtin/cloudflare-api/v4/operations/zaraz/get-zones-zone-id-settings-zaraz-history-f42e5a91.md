---
title: List Zaraz historical configuration records
page_id: operation-get-zones-zone-id-settings-zaraz-history-51733926
path: operations/zaraz
description: Lists a history of published Zaraz configuration records for a zone.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/settings/zaraz/history
operation_ids:
    - get-zones-zone_identifier-zaraz-history
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List Zaraz historical configuration records

`GET /zones/{zone_id}/settings/zaraz/history`

Operation ID: `get-zones-zone_identifier-zaraz-history`

Lists a history of published Zaraz configuration records for a zone.

## Definition

```yaml
{"operationId": "get-zones-zone_identifier-zaraz-history", "summary": "List Zaraz historical configuration records", "description": "Lists a history of published Zaraz configuration records for a zone.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/zaraz_identifier"}}, {"name": "offset", "in": "query", "description": "Ordinal number to start listing the results with. Default value is 0.", "schema": {"type": "integer", "minimum": 0}, "example": 0}, {"name": "limit", "in": "query", "description": "Maximum amount of results to list. Default value is 10.", "schema": {"type": "integer", "minimum": 1}, "example": 10}, {"name": "sortField", "in": "query", "description": "The field to sort by. Default is updated_at.", "schema": {"type": "string", "enum": ["id", "user_id", "description", "created_at", "updated_at"]}, "example": "updated_at"}, {"name": "sortOrder", "in": "query", "description": "Sorting order. Default is DESC.", "schema": {"type": "string", "enum": ["DESC", "ASC"]}, "example": "DESC"}], "responses": {"200": {"description": "List Zaraz historical configuration records response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/zaraz_zaraz-history-response"}}}}, "4XX": {"description": "List Zaraz historical configuration records failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/zaraz_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Zaraz"], "x-api-token-group": ["Zaraz Edit", "Zaraz Read", "Zaraz Admin"]}
```
