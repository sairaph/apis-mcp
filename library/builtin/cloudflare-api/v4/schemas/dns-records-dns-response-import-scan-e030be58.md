---
title: dns-records_dns_response_import_scan
page_id: schema-dns-records-dns-response-import-scan-e030be58
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# dns-records_dns_response_import_scan

```yaml
{"type": "object", "allOf": [{"$ref": "#/components/schemas/dns-records_api-response-single"}, {"properties": {"result": {"type": "object", "properties": {"recs_added": {"description": "Number of DNS records added.", "type": "number", "example": 5}, "total_records_parsed": {"description": "Total number of DNS records parsed.", "type": "number", "example": 5}}}}, "type": "object"}]}
```
