---
title: abuse-reports_SubmitErrorResponse
page_id: schema-abuse-reports-submiterrorresponse-713dc4f2
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# abuse-reports_SubmitErrorResponse

```yaml
{"type": "object", "properties": {"err_code": {"$ref": "#/components/schemas/abuse-reports_ErrorCode"}, "error_code": {"$ref": "#/components/schemas/abuse-reports_ErrorCode"}, "msg": {"description": "The error message for the error", "type": "string"}, "request": {"type": "object", "properties": {"act": {"$ref": "#/components/schemas/abuse-reports_SubmissionReportType"}}, "required": ["act"]}, "result": {"description": "The result should be 'error' for successful response", "type": "string"}}, "required": ["request", "result", "msg", "err_code", "error_code"]}
```
