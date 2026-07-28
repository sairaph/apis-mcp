---
title: email-auth_ZoneDnsRecords
page_id: schema-email-auth-zonednsrecords-ed521d33
path: schemas
description: Live DNS records for the zone, grouped by type
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# email-auth_ZoneDnsRecords

Live DNS records for the zone, grouped by type

```yaml
{"description": "Live DNS records for the zone, grouped by type", "type": "object", "properties": {"bimi_records": {"description": "BIMI TXT records", "type": "array", "items": {"$ref": "#/components/schemas/email-auth_DnsRecordInfo"}}, "cname_dkim_records": {"description": "CNAME records for DKIM", "type": "array", "items": {"$ref": "#/components/schemas/email-auth_DnsRecordInfo"}}, "cname_dmarc_records": {"description": "CNAME records at _dmarc (problematic)", "type": "array", "items": {"$ref": "#/components/schemas/email-auth_DnsRecordInfo"}}, "cname_spf_records": {"description": "CNAME records for SPF", "type": "array", "items": {"$ref": "#/components/schemas/email-auth_DnsRecordInfo"}}, "dkim_records": {"description": "DKIM TXT records", "type": "array", "items": {"$ref": "#/components/schemas/email-auth_DnsRecordInfo"}}, "dmarc_records": {"description": "DMARC TXT records", "type": "array", "items": {"$ref": "#/components/schemas/email-auth_DnsRecordInfo"}}, "spf_records": {"description": "SPF TXT records", "type": "array", "items": {"$ref": "#/components/schemas/email-auth_DnsRecordInfo"}}}}
```
