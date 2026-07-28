---
title: Submit an abuse report
page_id: operation-post-accounts-account-id-abuse-reports-report-param-024c34e6
path: operations/tseng-abuse-complaint-processor-other
description: |-
    Submit an abuse report of a particular type.

    Requires the abuse-reports entitlement on the account (Enterprise
    accounts have it by default; other accounts must request access) and an
    API token with the `Account > Abuse Reports > Edit` permission. If the
    account is not entitled, the request is rejected with an HTTP `401`
    response (see below).
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/abuse-reports/{report_param}
operation_ids:
    - SubmitAbuseReport
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Submit an abuse report

`POST /accounts/{account_id}/abuse-reports/{report_param}`

Operation ID: `SubmitAbuseReport`

Submit an abuse report of a particular type.

Requires the abuse-reports entitlement on the account (Enterprise
accounts have it by default; other accounts must request access) and an
API token with the `Account > Abuse Reports > Edit` permission. If the
account is not entitled, the request is rejected with an HTTP `401`
response (see below).

## Definition

```yaml
{"operationId": "SubmitAbuseReport", "summary": "Submit an abuse report", "description": "Submit an abuse report of a particular type.\n\nRequires the abuse-reports entitlement on the account (Enterprise\naccounts have it by default; other accounts must request access) and an\nAPI token with the `Account > Abuse Reports > Edit` permission. If the\naccount is not entitled, the request is rejected with an HTTP `401`\nresponse (see below).", "parameters": [{"name": "account_id", "in": "path", "description": "Cloudflare Account ID", "required": true, "schema": {"type": "string", "example": "023e105f4ecef8ad9ca31a8372d0c353", "maxLength": 32}}, {"name": "report_param", "in": "path", "description": "The report type to be submitted. Example: abuse_general", "required": true, "schema": {"$ref": "#/components/schemas/abuse-reports_SubmissionReportType"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/abuse-reports_SubmitReportRequest"}}}}, "responses": {"200": {"description": "Report submitted successfully", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/abuse-reports_SubmitReportResponse"}}}}, "400": {"description": "Report submitted with an error", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/abuse-reports_SubmitErrorResponse"}}}}, "401": {"description": "The request is not authorized to submit abuse reports. Missing or\ninvalid credentials may be rejected with the standard Cloudflare API\nauthentication error. If credentials are valid, this is most commonly\nbecause the account does not have the abuse-reports entitlement\n(Enterprise accounts have it by default; other accounts must request\naccess), or because the API token is missing the\n`Account > Abuse Reports > Edit` permission. Entitlement failures\nreturn the message `Not entitled to use feature: Abuse Report API`.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/abuse-reports_UnauthorizedErrorResponse"}}}}, "500": {"description": "Report submitted with an error", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/abuse-reports_SubmitErrorResponse"}}}}}, "security": [{"api_token": []}], "tags": ["tseng-abuse-complaint-processor_other"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "abuse-reports", "x-fern-sdk-method-name": "create", "x-forge-hidden": true}
```
