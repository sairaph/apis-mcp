---
title: dns-custom-nameservers_CustomNS
page_id: schema-dns-custom-nameservers-customns-6cb5ab48
path: schemas
description: A single account custom nameserver.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# dns-custom-nameservers_CustomNS

A single account custom nameserver.

```yaml
{"description": "A single account custom nameserver.", "type": "object", "properties": {"dns_records": {"description": "A and AAAA records associated with the nameserver.", "type": "array", "items": {"properties": {"type": {"description": "DNS record type.", "type": "string", "example": "A", "enum": ["A", "AAAA"], "x-auditable": true}, "value": {"description": "DNS record contents (an IPv4 or IPv6 address).", "type": "string", "example": "1.1.1.1", "x-auditable": true}}, "type": "object"}, "x-stainless-collection-type": "set"}, "ns_name": {"$ref": "#/components/schemas/dns-custom-nameservers_ns_name"}, "ns_set": {"$ref": "#/components/schemas/dns-custom-nameservers_ns_set"}, "status": {"description": "Verification status of the nameserver.", "type": "string", "example": "verified", "enum": ["moved", "pending", "verified"], "deprecated": true, "x-auditable": true}, "zone_tag": {"$ref": "#/components/schemas/dns-custom-nameservers_identifier-2"}}, "required": ["dns_records", "ns_name", "status", "zone_tag"], "title": "Custom NS"}
```
