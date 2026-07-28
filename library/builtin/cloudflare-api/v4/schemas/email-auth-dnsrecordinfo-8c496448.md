---
title: email-auth_DnsRecordInfo
page_id: schema-email-auth-dnsrecordinfo-8c496448
path: schemas
description: Summary of a single DNS record
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# email-auth_DnsRecordInfo

Summary of a single DNS record

```yaml
{"description": "Summary of a single DNS record", "type": "object", "properties": {"content": {"description": "Record content", "example": "v=DMARC1; p=none; rua=mailto:rua@dmarc-reports.cloudflare.net", "type": "string"}, "id": {"description": "DNS record ID", "type": "string", "example": "e5bb46707a802688812d5d1c9f7977d4"}, "name": {"description": "DNS record name", "type": "string", "example": "_dmarc.example.com"}, "ttl": {"description": "Time to live in seconds", "type": "integer", "example": 300}, "type": {"description": "Record type", "type": "string", "example": "TXT"}}}
```
