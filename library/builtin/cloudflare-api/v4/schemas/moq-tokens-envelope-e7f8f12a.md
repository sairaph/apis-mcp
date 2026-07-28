---
title: moq_tokens_envelope
page_id: schema-moq-tokens-envelope-e7f8f12a
path: schemas
description: |-
    A relay's token collection, keyed on issuer `type` (a discriminated
    union). V1 ships exactly one arm (`cloudflare_jwt`). Clients iterate
    `issuers`, switch on `type`, and ignore unknown types — that contract is
    what makes adding or removing an arm non-breaking.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# moq_tokens_envelope

A relay's token collection, keyed on issuer `type` (a discriminated
union). V1 ships exactly one arm (`cloudflare_jwt`). Clients iterate
`issuers`, switch on `type`, and ignore unknown types — that contract is
what makes adding or removing an arm non-breaking.

```yaml
{"description": "A relay's token collection, keyed on issuer `type` (a discriminated\nunion). V1 ships exactly one arm (`cloudflare_jwt`). Clients iterate\n`issuers`, switch on `type`, and ignore unknown types — that contract is\nwhat makes adding or removing an arm non-breaking.\n", "type": "object", "properties": {"issuers": {"type": "array", "items": {"$ref": "#/components/schemas/moq_issuer"}}}, "required": ["issuers"]}
```
