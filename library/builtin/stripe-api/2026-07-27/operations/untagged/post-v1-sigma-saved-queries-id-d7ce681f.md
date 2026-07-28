---
title: Update an existing Sigma Query
page_id: operation-post-v1-sigma-saved-queries-id-f27e8ae9
path: operations/untagged
description: <p>Update an existing Sigma query that previously exists</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - POST
api_endpoints:
    - /v1/sigma/saved_queries/{id}
operation_ids:
    - PostSigmaSavedQueriesId
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Update an existing Sigma Query

`POST /v1/sigma/saved_queries/{id}`

Operation ID: `PostSigmaSavedQueriesId`

<p>Update an existing Sigma query that previously exists</p>

## Definition

```yaml
{"summary": "Update an existing Sigma Query", "description": "<p>Update an existing Sigma query that previously exists</p>", "operationId": "PostSigmaSavedQueriesId", "parameters": [{"name": "id", "in": "path", "description": "The `id` of the saved query to update. This should be a valid `id` that was previously created.", "required": true, "style": "simple", "explode": false, "schema": {"maxLength": 5000, "type": "string"}}], "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"type": "object", "properties": {"expand": {"type": "array", "description": "Specifies which fields in the response should be expanded.", "items": {"maxLength": 5000, "type": "string"}}, "name": {"maxLength": 5000, "type": "string", "description": "The name of the query to update."}, "sql": {"maxLength": 100000, "type": "string", "description": "The sql statement to update the specified query statement with. This should be a valid Trino SQL statement that can be run in Sigma."}}, "additionalProperties": false}, "encoding": {"expand": {"style": "deepObject", "explode": true}}}}, "required": false}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/sigma.sigma_api_query"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
