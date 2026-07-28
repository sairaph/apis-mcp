---
title: abuse-reports_AbuseReport
page_id: schema-abuse-reports-abusereport-25e3577c
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# abuse-reports_AbuseReport

```yaml
{"type": "object", "properties": {"cdate": {"description": "Creation date of report. Time in RFC 3339 format (https://www.rfc-editor.org/rfc/rfc3339.html)", "type": "string", "example": "2009-11-10T23:00:00Z"}, "domain": {"description": "Domain that relates to the report.", "type": "string"}, "id": {"description": "Public facing ID of abuse report, aka abuse_rand.", "type": "string"}, "justification": {"description": "Justification for the report.", "type": "string"}, "mitigation_summary": {"$ref": "#/components/schemas/abuse-reports_MitigationSummary"}, "original_work": {"description": "Original work / Targeted brand in the alleged abuse.", "type": "string"}, "status": {"$ref": "#/components/schemas/abuse-reports_ReportStatus"}, "submitter": {"$ref": "#/components/schemas/abuse-reports_SubmitterDetails"}, "type": {"$ref": "#/components/schemas/abuse-reports_ReportType"}, "urls": {"type": "array", "items": {"type": "string"}}}, "required": ["id", "cdate", "domain", "type", "status", "mitigation_summary"]}
```
