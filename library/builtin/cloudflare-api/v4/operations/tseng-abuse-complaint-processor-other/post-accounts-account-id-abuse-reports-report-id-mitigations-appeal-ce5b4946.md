---
title: Request review on mitigations
page_id: operation-post-accounts-account-id-abuse-reports-report-id-mitigations-appeal-dccd28f0
path: operations/tseng-abuse-complaint-processor-other
description: Request a review for mitigations on an account.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/abuse-reports/{report_id}/mitigations/appeal
operation_ids:
    - RequestReview
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Request review on mitigations

`POST /accounts/{account_id}/abuse-reports/{report_id}/mitigations/appeal`

Operation ID: `RequestReview`

Request a review for mitigations on an account.

## Definition

```yaml
{"operationId": "RequestReview", "summary": "Request review on mitigations", "description": "Request a review for mitigations on an account.", "parameters": [{"name": "account_id", "in": "path", "description": "Cloudflare Account ID", "required": true, "schema": {"type": "string", "example": "023e105f4ecef8ad9ca31a8372d0c353", "maxLength": 32}}, {"name": "report_id", "in": "path", "description": "Abuse Report ID", "required": true, "schema": {"type": "string"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/abuse-reports_MitigationAppealRequest"}}}}, "responses": {"200": {"description": "Mitigation appeals received", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/abuse-reports_MitigationAppealResult"}}}}, "401": {"description": "The request is not authorized. Missing or invalid credentials may be\nrejected with the standard Cloudflare API authentication error. If\ncredentials are valid, the API token does not have the Abuse Reports\nedit permission required for this account.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/abuse-reports_UnauthorizedErrorResponse"}}}}, "500": {"description": "Failed to request review on delayed action.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"$ref": "#/components/schemas/abuse-reports_Message"}}, "messages": {"type": "array", "items": {"$ref": "#/components/schemas/abuse-reports_Message"}}, "success": {"type": "boolean"}}, "required": ["success"]}}}}}, "security": [{"api_token": []}], "tags": ["tseng-abuse-complaint-processor_other"], "x-api-token-group": ["Trust and Safety Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "abuse-reports.mitigations", "x-fern-sdk-method-name": "review", "x-forge-hidden": true}
```
