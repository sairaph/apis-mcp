---
title: mcn_provider_initial_setup_response
page_id: schema-mcn-provider-initial-setup-response-0c81377c
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# mcn_provider_initial_setup_response

```yaml
{"type": "object", "allOf": [{"$ref": "#/components/schemas/mcn_good_response"}, {"properties": {"result": {"type": "object", "discriminator": {"propertyName": "item_type"}, "oneOf": [{"$ref": "#/components/schemas/mcn_aws_trust_policy"}, {"$ref": "#/components/schemas/mcn_azure_setup"}, {"$ref": "#/components/schemas/mcn_gcp_setup"}]}}, "type": "object"}]}
```
