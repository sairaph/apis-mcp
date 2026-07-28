---
title: workers_binding_kind_mtls_certificate
page_id: schema-workers-binding-kind-mtls-certificate-906969bc
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# workers_binding_kind_mtls_certificate

```yaml
{"type": "object", "properties": {"certificate_id": {"description": "Identifier of the certificate to bind to.", "type": "string", "example": "efwu2n6s-q69d-2kr9-184j-4913e8h391k6", "x-auditable": true}, "name": {"$ref": "#/components/schemas/workers_binding_name"}, "type": {"description": "The kind of resource that the binding provides.", "type": "string", "enum": ["mtls_certificate"], "x-auditable": true}}, "required": ["name", "type", "certificate_id"]}
```
