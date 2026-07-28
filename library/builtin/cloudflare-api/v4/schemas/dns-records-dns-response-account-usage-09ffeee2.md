---
title: dns-records_dns_response_account_usage
page_id: schema-dns-records-dns-response-account-usage-09ffeee2
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# dns-records_dns_response_account_usage

```yaml
{"type": "object", "allOf": [{"$ref": "#/components/schemas/dns-records_api-response-single"}, {"properties": {"result": {"type": "object", "properties": {"internal_record_quota": {"description": "Maximum number of DNS records allowed across all internal zones in the account. Only present if internal DNS is enabled.", "type": "integer", "example": 1000000, "minimum": 0}, "internal_record_usage": {"description": "Current number of DNS records across all internal zones in the account. Only present if internal DNS is enabled.", "type": "integer", "example": 5000, "minimum": 0}, "record_quota": {"description": "Maximum number of DNS records allowed across all public zones in the account. Null if using zone-level quota.", "type": "integer", "example": 1000000, "minimum": 0, "nullable": true}, "record_usage": {"description": "Current number of DNS records across all public zones in the account.", "type": "integer", "example": 5000, "minimum": 0}}, "required": ["record_usage", "record_quota"]}}, "type": "object"}]}
```
