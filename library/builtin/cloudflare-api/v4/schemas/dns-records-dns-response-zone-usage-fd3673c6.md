---
title: dns-records_dns_response_zone_usage
page_id: schema-dns-records-dns-response-zone-usage-fd3673c6
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# dns-records_dns_response_zone_usage

```yaml
{"type": "object", "allOf": [{"$ref": "#/components/schemas/dns-records_api-response-single"}, {"properties": {"result": {"type": "object", "properties": {"record_quota": {"description": "Maximum number of DNS records allowed for the zone. Null if using account-level quota.", "type": "integer", "example": 200, "minimum": 0, "nullable": true}, "record_usage": {"description": "Current number of DNS records in the zone.", "type": "integer", "example": 150, "minimum": 0}}, "required": ["record_usage", "record_quota"]}}, "type": "object"}]}
```
