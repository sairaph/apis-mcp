---
title: registrar-api_pricing
page_id: schema-registrar-api-pricing-2ae67c03
path: schemas
description: |-
    Annual pricing information for a registrable domain. This object is only
    present when `registrable` is `true`. All prices are per year and returned
    as strings to preserve decimal precision.

    `registration_cost` and `renewal_cost` are frequently the same value, but
    may differ — especially for premium domains where registries set different
    rates for initial registration vs. renewal. For a multi-year registration
    (e.g., 4 years), the first year is charged at `registration_cost` and each
    subsequent year at `renewal_cost`. Registry pricing may change over time;
    the values returned here reflect the current registry rate. Premium pricing
    may be surfaced by Search and Check, but premium registration is not currently
    supported by this API.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# registrar-api_pricing

Annual pricing information for a registrable domain. This object is only
present when `registrable` is `true`. All prices are per year and returned
as strings to preserve decimal precision.

`registration_cost` and `renewal_cost` are frequently the same value, but
may differ — especially for premium domains where registries set different
rates for initial registration vs. renewal. For a multi-year registration
(e.g., 4 years), the first year is charged at `registration_cost` and each
subsequent year at `renewal_cost`. Registry pricing may change over time;
the values returned here reflect the current registry rate. Premium pricing
may be surfaced by Search and Check, but premium registration is not currently
supported by this API.

```yaml
{"description": "Annual pricing information for a registrable domain. This object is only\npresent when `registrable` is `true`. All prices are per year and returned\nas strings to preserve decimal precision.\n\n`registration_cost` and `renewal_cost` are frequently the same value, but\nmay differ — especially for premium domains where registries set different\nrates for initial registration vs. renewal. For a multi-year registration\n(e.g., 4 years), the first year is charged at `registration_cost` and each\nsubsequent year at `renewal_cost`. Registry pricing may change over time;\nthe values returned here reflect the current registry rate. Premium pricing\nmay be surfaced by Search and Check, but premium registration is not currently\nsupported by this API.\n", "type": "object", "properties": {"currency": {"description": "ISO-4217 currency code for the prices (e.g., \"USD\", \"EUR\", \"GBP\").", "type": "string", "example": "USD"}, "registration_cost": {"description": "The first-year cost to register this domain. For premium domains\n(`tier: premium`), this price is set by the registry and may be\nsignificantly higher than standard pricing. For multi-year\nregistrations, this cost applies to the first year only; subsequent\nyears are charged at `renewal_cost`.\n", "type": "string", "example": "8.57"}, "renewal_cost": {"description": "Per-year renewal cost for this domain. Applied to each year beyond\nthe first year of a multi-year registration, and to each annual\nauto-renewal thereafter. May differ from `registration_cost`,\nespecially for premium domains where initial registration often\ncosts more than renewals.\n", "type": "string", "example": "8.57"}}, "required": ["currency", "registration_cost", "renewal_cost"]}
```
