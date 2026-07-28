---
title: Update applications review statuses
page_id: operation-put-accounts-account-id-gateway-apps-review-status-d478e8e9
path: operations/zero-trust-applications-review-status
description: Update the statuses of your applications.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /accounts/{account_id}/gateway/apps/review_status
operation_ids:
    - zero-trust-applications-review-status-update
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update applications review statuses

`PUT /accounts/{account_id}/gateway/apps/review_status`

Operation ID: `zero-trust-applications-review-status-update`

Update the statuses of your applications.

## Definition

```yaml
{"operationId": "zero-trust-applications-review-status-update", "summary": "Update applications review statuses", "description": "Update the statuses of your applications.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/zero-trust-gateway_identifier-3"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"type": "object", "properties": {"approved_apps": {"$ref": "#/components/schemas/zero-trust-gateway_approved_apps"}, "in_review_apps": {"$ref": "#/components/schemas/zero-trust-gateway_in_review_apps"}, "unapproved_apps": {"$ref": "#/components/schemas/zero-trust-gateway_unapproved_apps"}}, "required": ["approved_apps", "unapproved_apps", "in_review_apps"]}}}}, "responses": {"200": {"description": "Update applications review status response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/zero-trust-gateway_applications_review_status_response"}}}}, "4XX": {"description": "Update applications review status failure response.", "content": {"application/json": {"schema": {"type": "object", "allOf": [{"$ref": "#/components/schemas/zero-trust-gateway_applications_review_status_response"}, {"$ref": "#/components/schemas/zero-trust-gateway_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Zero Trust applications review status"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.apps.review.status", "x-fern-sdk-method-name": "update", "x-forge-hidden": true}
```
