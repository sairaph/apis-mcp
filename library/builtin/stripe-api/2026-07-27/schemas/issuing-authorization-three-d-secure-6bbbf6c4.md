---
title: issuing_authorization_three_d_secure
page_id: schema-issuing-authorization-three-d-secure-6bbbf6c4
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# issuing_authorization_three_d_secure

```yaml
{"title": "IssuingAuthorizationThreeDSecure", "required": ["result"], "type": "object", "properties": {"result": {"type": "string", "description": "The outcome of the 3D Secure authentication request.", "enum": ["attempt_acknowledged", "authenticated", "failed", "required"], "x-stripeBypassValidation": true}}, "description": "", "x-expandableFields": []}
```
