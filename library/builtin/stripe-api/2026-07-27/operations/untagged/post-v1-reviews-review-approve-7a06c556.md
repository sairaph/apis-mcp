---
title: Approve a review
page_id: operation-post-v1-reviews-review-approve-9eb16b9c
path: operations/untagged
description: <p>Approves a <code>Review</code> object, closing it and removing it from the list of reviews.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - POST
api_endpoints:
    - /v1/reviews/{review}/approve
operation_ids:
    - PostReviewsReviewApprove
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Approve a review

`POST /v1/reviews/{review}/approve`

Operation ID: `PostReviewsReviewApprove`

<p>Approves a <code>Review</code> object, closing it and removing it from the list of reviews.</p>

## Definition

```yaml
{"summary": "Approve a review", "description": "<p>Approves a <code>Review</code> object, closing it and removing it from the list of reviews.</p>", "operationId": "PostReviewsReviewApprove", "parameters": [{"name": "review", "in": "path", "required": true, "style": "simple", "explode": false, "schema": {"maxLength": 5000, "type": "string"}}], "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"type": "object", "properties": {"expand": {"type": "array", "description": "Specifies which fields in the response should be expanded.", "items": {"maxLength": 5000, "type": "string"}}}, "additionalProperties": false}, "encoding": {"expand": {"style": "deepObject", "explode": true}}}}, "required": false}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/review"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
