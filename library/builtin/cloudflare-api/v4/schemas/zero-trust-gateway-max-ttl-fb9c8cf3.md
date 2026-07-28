---
title: zero-trust-gateway_max-ttl
page_id: schema-zero-trust-gateway-max-ttl-fb9c8cf3
path: schemas
description: Controls how DNS response TTLs are capped for this location relative to the account `max_ttl_secs` setting. Omitting `max_ttl` on update resets it to `inherit`.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# zero-trust-gateway_max-ttl

Controls how DNS response TTLs are capped for this location relative to the account `max_ttl_secs` setting. Omitting `max_ttl` on update resets it to `inherit`.

```yaml
{"description": "Controls how DNS response TTLs are capped for this location relative to the account `max_ttl_secs` setting. Omitting `max_ttl` on update resets it to `inherit`.", "type": "object", "properties": {"mode": {"description": "`inherit` uses the account `max_ttl_secs`. `override` uses this location's `ttl_secs`. `disabled` leaves returned TTLs unchanged.", "type": "string", "example": "override", "enum": ["inherit", "override", "disabled"], "x-auditable": true}, "ttl_secs": {"description": "Location-specific cap on DNS response TTLs, in seconds. Required when `mode` is `override`. Must be omitted when `mode` is `inherit` or `disabled`.", "type": "integer", "example": 3600, "maximum": 36000, "minimum": 60, "nullable": true, "x-auditable": true}}, "default": {"mode": "inherit"}, "nullable": true, "required": ["mode"], "x-stainless-terraform-configurability": "optional"}
```
