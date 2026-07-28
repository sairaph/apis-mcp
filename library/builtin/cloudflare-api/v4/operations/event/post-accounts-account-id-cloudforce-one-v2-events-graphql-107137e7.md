---
title: GraphQL endpoint for event aggregation
page_id: operation-post-accounts-account-id-cloudforce-one-v2-events-graphql-2d02d946
path: operations/event
description: Execute GraphQL aggregations over threat events. Supports multi-dimensional group-bys, optional date range filtering, and multi-dataset aggregation.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/cloudforce-one/v2/events/graphql
operation_ids:
    - post_EventGraphQLV2
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# GraphQL endpoint for event aggregation

`POST /accounts/{account_id}/cloudforce-one/v2/events/graphql`

Operation ID: `post_EventGraphQLV2`

Execute GraphQL aggregations over threat events. Supports multi-dimensional group-bys, optional date range filtering, and multi-dataset aggregation.

## Definition

```yaml
{"operationId": "post_EventGraphQLV2", "summary": "GraphQL endpoint for event aggregation", "description": "Execute GraphQL aggregations over threat events. Supports multi-dimensional group-bys, optional date range filtering, and multi-dataset aggregation.", "parameters": [{"name": "account_id", "in": "path", "description": "Account ID.", "required": true, "schema": {"description": "Account ID.", "type": "string"}}], "responses": {"200": {"description": "GraphQL response payload (data and errors).", "content": {"application/json": {"schema": {"type": "object", "properties": {"data": {"type": "object", "nullable": true}, "errors": {"type": "array", "items": {"type": "object"}, "nullable": true}}}}}}, "400": {"description": "Bad Request.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"message": {"type": "string", "example": "An error occurred."}}, "required": ["message"], "type": "object"}}, "result": {"type": "object"}, "success": {"type": "boolean"}}, "required": ["result", "success", "errors"]}}}}}, "security": [{"api_token": []}], "tags": ["Event"], "x-api-token-group": ["Cloudforce One Write"]}
```
