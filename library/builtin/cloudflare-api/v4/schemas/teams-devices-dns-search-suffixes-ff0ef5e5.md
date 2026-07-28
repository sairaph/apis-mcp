---
title: teams-devices_dns_search_suffixes
page_id: schema-teams-devices-dns-search-suffixes-ff0ef5e5
path: schemas
description: List of DNS search suffixes to apply to clients. Suffixes are evaluated in order. Use an empty array to clear.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# teams-devices_dns_search_suffixes

List of DNS search suffixes to apply to clients. Suffixes are evaluated in order. Use an empty array to clear.

```yaml
{"description": "List of DNS search suffixes to apply to clients. Suffixes are evaluated in order. Use an empty array to clear.", "type": "array", "items": {"$ref": "#/components/schemas/teams-devices_dns_search_suffix"}, "default": [], "x-stainless-terraform-configurability": "computed_optional"}
```
