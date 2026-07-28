---
title: bill-subs-api_rate_plan
page_id: schema-bill-subs-api-rate-plan-152c738b
path: schemas
description: The rate plan applied to the subscription.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# bill-subs-api_rate_plan

The rate plan applied to the subscription.

```yaml
{"description": "The rate plan applied to the subscription.", "type": "object", "properties": {"currency": {"description": "The currency applied to the rate plan subscription.", "type": "string", "example": "USD", "x-auditable": true, "x-stainless-terraform-configurability": "computed_optional"}, "externally_managed": {"description": "Whether this rate plan is managed externally from Cloudflare.", "type": "boolean", "example": false, "x-auditable": true, "x-stainless-terraform-configurability": "computed_optional"}, "id": {"description": "The ID of the rate plan.", "type": "string", "example": "free", "x-auditable": true}, "is_contract": {"description": "Whether a rate plan is enterprise-based (or newly adopted term contract).", "type": "boolean", "example": false, "x-auditable": true, "x-stainless-terraform-configurability": "computed_optional"}, "public_name": {"description": "The full name of the rate plan.", "type": "string", "example": "Business Plan", "x-auditable": true, "x-stainless-terraform-configurability": "computed_optional"}, "scope": {"description": "The scope that this rate plan applies to.", "type": "string", "example": "zone", "x-auditable": true, "x-stainless-terraform-configurability": "computed_optional"}, "sets": {"description": "The list of sets this rate plan applies to. Returns array of strings.", "type": "array", "items": {"type": "string", "x-auditable": true}, "example": [], "x-stainless-terraform-configurability": "optional"}}}
```
