---
title: workers_Domain
page_id: schema-workers-domain-76f63947
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# workers_Domain

```yaml
{"type": "object", "properties": {"cert_id": {"description": "ID of the TLS certificate issued for the domain.", "type": "string", "format": "uuid", "example": "9fdf92c8-64c2-4a3d-b1af-e15304961145", "readOnly": true, "x-auditable": true}, "environment": {"description": "Worker environment associated with the domain.", "type": "string", "example": "production", "deprecated": true, "readOnly": true, "x-auditable": true}, "hostname": {"description": "Hostname of the domain. Can be either the zone apex or a subdomain of the zone. Requests to this hostname will be routed to the configured Worker.", "type": "string", "example": "app.example.com", "x-auditable": true}, "id": {"description": "Immutable ID of the domain.", "type": "string", "example": "dbe10b4bc17c295377eabd600e1787fd", "readOnly": true, "x-auditable": true}, "service": {"description": "Name of the Worker associated with the domain. Requests to the configured hostname will be routed to this Worker.", "type": "string", "example": "my-worker", "x-auditable": true}, "zone_id": {"description": "ID of the zone containing the domain hostname.", "type": "string", "example": "593c9c94de529bbbfaac7c53ced0447d", "x-auditable": true, "x-stainless-terraform-configurability": "computed_optional"}, "zone_name": {"description": "Name of the zone containing the domain hostname.", "type": "string", "example": "example.com", "x-auditable": true, "x-stainless-terraform-configurability": "computed_optional"}}, "required": ["id", "cert_id", "zone_id", "zone_name", "hostname", "service"]}
```
