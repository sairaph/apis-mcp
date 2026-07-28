---
title: teams-devices_dns_search_suffix
page_id: schema-teams-devices-dns-search-suffix-f2b7d581
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# teams-devices_dns_search_suffix

```yaml
{"type": "object", "properties": {"description": {"description": "A description of the DNS search suffix.", "type": "string", "example": "Example internal domains", "x-auditable": true, "x-stainless-terraform-configurability": "computed_optional"}, "suffix": {"description": "The DNS search suffix to append when resolving short hostnames.", "type": "string", "example": "internal.corp", "x-auditable": true}}, "required": ["suffix"]}
```
