---
title: intel-sinkholes_ingress_create_params
page_id: schema-intel-sinkholes-ingress-create-params-1f5a81f9
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# intel-sinkholes_ingress_create_params

```yaml
{"type": "object", "properties": {"cidr": {"description": "The CIDR block for the ingress rule in IPv4 or IPv6 notation (e.g., 192.0.2.0/24). Must be a Cloudflare BYOIP associated with your account.", "type": "string"}}, "required": ["cidr"]}
```
