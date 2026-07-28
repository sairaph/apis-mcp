---
title: access_oidc_claim_rule
page_id: schema-access-oidc-claim-rule-11975919
path: schemas
description: |-
    Matches an OIDC claim.
    Requires an OIDC identity provider.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# access_oidc_claim_rule

Matches an OIDC claim.
Requires an OIDC identity provider.

```yaml
{"description": "Matches an OIDC claim.\nRequires an OIDC identity provider.", "type": "object", "properties": {"oidc": {"type": "object", "properties": {"claim_name": {"description": "The name of the OIDC claim.", "type": "string", "example": "group"}, "claim_value": {"description": "The OIDC claim value to look for.", "type": "string", "example": "devs@cloudflare.com"}, "identity_provider_id": {"description": "The ID of your OIDC identity provider.", "type": "string", "example": "ea85612a-29c8-46c2-bacb-669d65136971"}}, "required": ["claim_name", "claim_value", "identity_provider_id"]}}, "required": ["oidc"], "title": "OIDC claim"}
```
