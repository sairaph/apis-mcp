---
title: firewall_filters
page_id: schema-firewall-filters-8dba116d
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# firewall_filters

```yaml
{"type": "object", "properties": {"configuration.target": {"description": "The target to search in existing rules.", "type": "string", "example": "ip", "enum": ["ip", "ip_range", "asn", "country"], "x-auditable": true}, "configuration.value": {"description": "The target value to search for in existing rules: an IP address, an IP address range, or a country code, depending on the provided `configuration.target`.\nNotes: You can search for a single IPv4 address, an IP address range with a subnet of '/16' or '/24', or a two-letter ISO-3166-1 alpha-2 country code.", "type": "string", "example": "198.51.100.4"}, "match": {"description": "When set to `all`, all the search requirements must match. When set to `any`, only one of the search requirements has to match.", "type": "string", "default": "all", "enum": ["any", "all"], "x-auditable": true}, "mode": {"$ref": "#/components/schemas/firewall_schemas-mode"}, "notes": {"description": "The string to search for in the notes of existing IP Access rules.\nNotes: For example, the string 'attack' would match IP Access rules with notes 'Attack 26/02' and 'Attack 27/02'. The search is case insensitive.", "type": "string", "example": "my note", "x-auditable": true}}}
```
