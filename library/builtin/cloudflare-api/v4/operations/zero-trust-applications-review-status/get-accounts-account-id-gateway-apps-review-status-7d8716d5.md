---
title: List applications review statuses
page_id: operation-get-accounts-account-id-gateway-apps-review-status-90e54825
path: operations/zero-trust-applications-review-status
description: Retrieve the statuses of your applications.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/gateway/apps/review_status
operation_ids:
    - zero-trust-applications-review-status-list
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List applications review statuses

`GET /accounts/{account_id}/gateway/apps/review_status`

Operation ID: `zero-trust-applications-review-status-list`

Retrieve the statuses of your applications.

## Definition

```yaml
{"operationId": "zero-trust-applications-review-status-list", "summary": "List applications review statuses", "description": "Retrieve the statuses of your applications.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/zero-trust-gateway_identifier-3"}}], "responses": {"200": {"description": "List applications review status response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/zero-trust-gateway_applications_review_status_response"}}}}, "4XX": {"description": "List applications review status failure response.", "content": {"application/json": {"schema": {"type": "object", "allOf": [{"$ref": "#/components/schemas/zero-trust-gateway_applications_review_status_response"}, {"$ref": "#/components/schemas/zero-trust-gateway_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Zero Trust applications review status"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.apps.review.status", "x-fern-sdk-method-name": "list", "x-forge-hidden": true}
```
