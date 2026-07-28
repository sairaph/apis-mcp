---
title: intel_whois
page_id: schema-intel-whois-d7b13d6d
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# intel_whois

```yaml
{"type": "object", "properties": {"created_date": {"type": "string", "format": "date", "example": "2009-02-17", "x-auditable": true}, "domain": {"$ref": "#/components/schemas/intel_domain_name"}, "nameservers": {"type": "array", "items": {"type": "string", "x-auditable": true}, "example": ["ns3.cloudflare.com", "ns4.cloudflare.com", "ns5.cloudflare.com", "ns6.cloudflare.com", "ns7.cloudflare.com"]}, "registrant": {"type": "string", "example": "DATA REDACTED", "x-auditable": true}, "registrant_country": {"type": "string", "example": "United States", "x-auditable": true}, "registrant_email": {"type": "string", "example": "https://domaincontact.cloudflareregistrar.com/cloudflare.com", "x-auditable": true}, "registrant_org": {"type": "string", "example": "DATA REDACTED", "x-auditable": true}, "registrar": {"type": "string", "example": "Cloudflare, Inc.", "x-auditable": true}, "updated_date": {"type": "string", "format": "date", "example": "2017-05-24", "x-auditable": true}}}
```
