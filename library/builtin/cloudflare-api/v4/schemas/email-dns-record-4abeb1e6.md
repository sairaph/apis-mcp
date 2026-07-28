---
title: email_dns_record
page_id: schema-email-dns-record-4abeb1e6
path: schemas
description: List of records needed to enable an Email Routing zone.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# email_dns_record

List of records needed to enable an Email Routing zone.

```yaml
{"description": "List of records needed to enable an Email Routing zone.", "type": "object", "properties": {"content": {"description": "DNS record content.", "example": "route1.mx.cloudflare.net", "type": "string"}, "name": {"description": "DNS record name (or @ for the zone apex).", "type": "string", "example": "example.com", "maxLength": 255, "x-auditable": true}, "priority": {"description": "Required for MX, SRV and URI records. Unused by other record types. Records with lower priorities are preferred.", "type": "number", "example": 12, "maximum": 65535, "minimum": 0, "x-auditable": true}, "ttl": {"description": "Time to live, in seconds, of the DNS record. Must be between 60 and 86400, or 1 for 'automatic'.", "type": "number", "example": 1, "anyOf": [{"example": 3600, "maximum": 86400, "minimum": 1, "type": "number"}, {"enum": [1], "type": "number"}], "x-auditable": true}, "type": {"description": "DNS record type.", "type": "string", "example": "NS", "enum": ["A", "AAAA", "CNAME", "HTTPS", "TXT", "SRV", "LOC", "MX", "NS", "CERT", "DNSKEY", "DS", "NAPTR", "SMIMEA", "SSHFP", "SVCB", "TLSA", "URI"], "readOnly": true, "x-auditable": true}}}
```
