---
title: Delete a plan
page_id: operation-delete-v1-plans-plan-3b5eb70b
path: operations/untagged
description: <p>Deleting plans means new subscribers can’t be added. Existing subscribers aren’t affected.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - DELETE
api_endpoints:
    - /v1/plans/{plan}
operation_ids:
    - DeletePlansPlan
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Delete a plan

`DELETE /v1/plans/{plan}`

Operation ID: `DeletePlansPlan`

<p>Deleting plans means new subscribers can’t be added. Existing subscribers aren’t affected.</p>

## Definition

```yaml
{"summary": "Delete a plan", "description": "<p>Deleting plans means new subscribers can’t be added. Existing subscribers aren’t affected.</p>", "operationId": "DeletePlansPlan", "parameters": [{"name": "plan", "in": "path", "required": true, "style": "simple", "explode": false, "schema": {"maxLength": 5000, "type": "string"}}], "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"type": "object", "properties": {}, "additionalProperties": false}, "encoding": {}}}, "required": false}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/deleted_plan"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
