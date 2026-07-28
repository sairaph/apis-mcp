---
title: bot-management_feedback_report
page_id: schema-bot-management-feedback-report-e22fb05a
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# bot-management_feedback_report

```yaml
{"type": "object", "properties": {"created_at": {"type": "string", "format": "date-time", "readOnly": true}, "description": {"type": "string"}, "expression": {"description": "Wirefilter expression describing the traffic being reported.", "type": "string"}, "first_request_seen_at": {"type": "string", "format": "date-time"}, "last_request_seen_at": {"type": "string", "format": "date-time"}, "requests": {"type": "integer", "format": "int64"}, "requests_by_attribute": {"$ref": "#/components/schemas/bot-management_requests_by_attribute"}, "requests_by_score": {"$ref": "#/components/schemas/bot-management_requests_by_score"}, "requests_by_score_src": {"$ref": "#/components/schemas/bot-management_requests_by_score_src"}, "subtype": {"type": "string"}, "type": {"$ref": "#/components/schemas/bot-management_feedback_type"}}, "required": ["type", "description", "expression", "first_request_seen_at", "last_request_seen_at", "requests", "requests_by_score", "requests_by_score_src", "requests_by_attribute"]}
```
