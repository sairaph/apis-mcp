---
title: registrar-api_domain_check_request
page_id: schema-registrar-api-domain-check-request-8f885c48
path: schemas
description: Request body for checking domain availability.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# registrar-api_domain_check_request

Request body for checking domain availability.

```yaml
{"description": "Request body for checking domain availability.", "type": "object", "properties": {"domains": {"description": "List of fully qualified domain names (FQDNs) to check for availability. Each domain must include the extension.\n- Minimum: 1 domain\n- Maximum: 20 domains per request\n- Domains on unsupported extensions are returned with `registrable: false` and a `reason` field\n- Malformed domain names (e.g., missing extension) may be omitted from the response\n", "type": "array", "items": {"type": "string"}, "example": ["example.com", "example.net"], "maxItems": 20, "minItems": 1}}, "required": ["domains"]}
```
