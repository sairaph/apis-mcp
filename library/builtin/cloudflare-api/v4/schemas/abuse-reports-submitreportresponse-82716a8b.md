---
title: abuse-reports_SubmitReportResponse
page_id: schema-abuse-reports-submitreportresponse-82716a8b
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# abuse-reports_SubmitReportResponse

```yaml
{"type": "object", "properties": {"abuse_rand": {"description": "The identifier for the submitted abuse report.", "type": "string"}, "request": {"type": "object", "properties": {"act": {"$ref": "#/components/schemas/abuse-reports_SubmissionReportType"}}, "required": ["act"]}, "result": {"description": "The result should be 'success' for successful response", "type": "string"}}, "required": ["request", "result", "abuse_rand"]}
```
