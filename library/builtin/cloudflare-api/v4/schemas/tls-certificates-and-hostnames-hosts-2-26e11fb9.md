---
title: tls-certificates-and-hostnames_hosts-2
page_id: schema-tls-certificates-and-hostnames-hosts-2-26e11fb9
path: schemas
description: Comma separated list of valid host names for the certificate packs. Must contain the zone apex, may not contain more than 50 hosts, and may not be empty.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# tls-certificates-and-hostnames_hosts-2

Comma separated list of valid host names for the certificate packs. Must contain the zone apex, may not contain more than 50 hosts, and may not be empty.

```yaml
{"description": "Comma separated list of valid host names for the certificate packs. Must contain the zone apex, may not contain more than 50 hosts, and may not be empty.", "type": "array", "items": {"type": "string", "x-auditable": true}, "example": ["example.com", "*.example.com", "www.example.com"], "x-stainless-collection-type": "set", "x-stainless-terraform-configurability": "computed_optional"}
```
