---
title: intel_passive-dns-by-ip
page_id: schema-intel-passive-dns-by-ip-fb91d8ee
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# intel_passive-dns-by-ip

```yaml
{"type": "object", "properties": {"count": {"description": "Total results returned based on your search parameters.", "type": "number", "example": 1}, "page": {"description": "Current page within paginated list of results.", "type": "number", "example": 1}, "per_page": {"description": "Number of results per page of results.", "type": "number", "example": 20}, "reverse_records": {"description": "Reverse DNS look-ups observed during the time period.", "type": "array", "items": {"properties": {"first_seen": {"description": "First seen date of the DNS record during the time period.", "type": "string", "format": "date", "example": "2021-04-01", "x-auditable": true}, "hostname": {"description": "Hostname that the IP was observed resolving to.", "type": "string", "x-auditable": true}, "last_seen": {"description": "Last seen date of the DNS record during the time period.", "type": "string", "format": "date", "example": "2021-04-30", "x-auditable": true}}, "type": "object"}}}}
```
