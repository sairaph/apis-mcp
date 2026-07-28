---
title: vuln_scanner_bola-report
page_id: schema-vuln-scanner-bola-report-924c800a
path: schemas
description: A BOLA vulnerability scan report, versioned for future evolution.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# vuln_scanner_bola-report

A BOLA vulnerability scan report, versioned for future evolution.

```yaml
{"description": "A BOLA vulnerability scan report, versioned for future evolution.", "type": "object", "properties": {"report": {"allOf": [{"$ref": "#/components/schemas/vuln_scanner_bola-report-v1"}]}, "report_schema_version": {"description": "Version of the report schema.", "type": "string", "enum": ["v1"]}}, "required": ["report_schema_version", "report"]}
```
