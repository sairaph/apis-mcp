---
title: vuln_scanner_scan
page_id: schema-vuln-scanner-scan-a7730159
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# vuln_scanner_scan

```yaml
{"type": "object", "properties": {"id": {"description": "Scan identifier.", "type": "string", "format": "uuid", "x-auditable": true}, "report": {"description": "Vulnerability report produced after the scan completes. The shape depends on the scan type. Present only for finished scans.", "type": "object", "allOf": [{"$ref": "#/components/schemas/vuln_scanner_bola-report"}], "nullable": true}, "scan_type": {"description": "The type of vulnerability scan.", "type": "string", "enum": ["bola"], "x-auditable": true}, "status": {"description": "Current lifecycle status of the scan.", "type": "string", "enum": ["created", "scheduled", "planning", "running", "finished", "failed"], "x-auditable": true}, "target_environment_id": {"description": "The target environment this scan runs against.", "type": "string", "format": "uuid", "x-auditable": true}}, "required": ["id", "target_environment_id", "scan_type", "status"]}
```
