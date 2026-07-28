---
title: registrar-api-sandbox_domain_check_response
page_id: schema-registrar-api-sandbox-domain-check-response-ea6e6a29
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# registrar-api-sandbox_domain_check_response

```yaml
{"allOf": [{"$ref": "#/components/schemas/registrar-api-sandbox_api-response-common"}, {"properties": {"result": {"description": "Contains the availability check results.", "type": "object", "properties": {"domains": {"description": "Array of domain availability results. Domains on unsupported\nextensions are included with `registrable: false` and a `reason`\nfield. Malformed domain names may be omitted.\n", "type": "array", "items": {"$ref": "#/components/schemas/registrar-api-sandbox_domain_check_result"}}}, "required": ["domains"]}}, "required": ["result"], "type": "object"}]}
```
