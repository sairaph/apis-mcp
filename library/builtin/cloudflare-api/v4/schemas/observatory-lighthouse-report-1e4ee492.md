---
title: observatory_lighthouse_report
page_id: schema-observatory-lighthouse-report-1e4ee492
path: schemas
description: The Lighthouse report.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# observatory_lighthouse_report

The Lighthouse report.

```yaml
{"description": "The Lighthouse report.", "type": "object", "properties": {"cls": {"description": "Cumulative Layout Shift.", "type": "number", "example": 100, "x-auditable": true}, "deviceType": {"$ref": "#/components/schemas/observatory_device_type"}, "error": {"type": "object", "properties": {"code": {"$ref": "#/components/schemas/observatory_lighthouse_error_code"}, "detail": {"description": "Detailed error message.", "type": "string", "example": "Details: net::ERR_CONNECTION_CLOSED", "x-auditable": true}, "finalDisplayedUrl": {"description": "The final URL displayed to the user.", "type": "string", "example": "example.com", "x-auditable": true}}}, "fcp": {"description": "First Contentful Paint.", "type": "number", "example": 100, "x-auditable": true}, "jsonReportUrl": {"description": "The URL to the full Lighthouse JSON report.", "type": "string", "x-auditable": true}, "lcp": {"description": "Largest Contentful Paint.", "type": "number", "example": 100, "x-auditable": true}, "performanceScore": {"description": "The Lighthouse performance score.", "type": "number", "example": 90, "x-auditable": true}, "si": {"description": "Speed Index.", "type": "number", "example": 100, "x-auditable": true}, "state": {"$ref": "#/components/schemas/observatory_lighthouse_state"}, "tbt": {"description": "Total Blocking Time.", "type": "number", "example": 100, "x-auditable": true}, "ttfb": {"description": "Time To First Byte.", "type": "number", "example": 100, "x-auditable": true}, "tti": {"description": "Time To Interactive.", "type": "number", "example": 100, "x-auditable": true}}}
```
