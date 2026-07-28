---
title: GET /paas/v4/async-result/{id}
page_id: operation-get-paas-v4-async-result-id-f1020ab6
path: operations/untagged
description: This endpoint is used to query the result of an asynchronous request.
source: https://docs.z.ai/openapi.json
http_methods:
    - GET
api_endpoints:
    - /paas/v4/async-result/{id}
source_type: openapi
imported_from: https://docs.z.ai/openapi.json
---

# GET /paas/v4/async-result/{id}

`GET /paas/v4/async-result/{id}`

This endpoint is used to query the result of an asynchronous request.

## Definition

```yaml
{"description": "This endpoint is used to query the result of an asynchronous request.", "parameters": [{"$ref": "#/components/parameters/AcceptLanguage"}, {"name": "id", "in": "path", "required": true, "schema": {"type": "string", "description": "Task id."}}], "responses": {"200": {"description": "Processing successful", "content": {"application/json": {"schema": {"oneOf": [{"title": "Video Generation", "$ref": "#/components/schemas/AsyncVideoGenerationResponse"}, {"title": "Image Generation", "$ref": "#/components/schemas/AsyncImageGenerationResponse"}]}}}}, "default": {"description": "The request has failed.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/Error"}}}}}}
```
