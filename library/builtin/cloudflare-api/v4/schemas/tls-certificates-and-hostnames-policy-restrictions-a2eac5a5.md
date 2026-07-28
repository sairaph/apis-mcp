---
title: tls-certificates-and-hostnames_policy_restrictions
page_id: schema-tls-certificates-and-hostnames-policy-restrictions-a2eac5a5
path: schemas
description: |-
    The policy restrictions returned by the API. This field is returned in responses
    when a policy has been set. The API accepts the "policy" field in requests but
    returns this field as "policy_restrictions" in responses.

    Specifies the region(s) where your private key can be held locally for optimal
    TLS performance. Format is a boolean expression, for example:
    "(country: US) or (region: EU)"
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# tls-certificates-and-hostnames_policy_restrictions

The policy restrictions returned by the API. This field is returned in responses
when a policy has been set. The API accepts the "policy" field in requests but
returns this field as "policy_restrictions" in responses.

Specifies the region(s) where your private key can be held locally for optimal
TLS performance. Format is a boolean expression, for example:
"(country: US) or (region: EU)"

```yaml
{"description": "The policy restrictions returned by the API. This field is returned in responses\nwhen a policy has been set. The API accepts the \"policy\" field in requests but\nreturns this field as \"policy_restrictions\" in responses.\n\nSpecifies the region(s) where your private key can be held locally for optimal\nTLS performance. Format is a boolean expression, for example:\n\"(country: US) or (region: EU)\"\n", "type": "string", "example": "(country: US) or (region: EU)", "readOnly": true}
```
