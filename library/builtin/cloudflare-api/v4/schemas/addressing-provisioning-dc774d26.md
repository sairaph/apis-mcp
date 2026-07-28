---
title: addressing_provisioning
page_id: schema-addressing-provisioning-dc774d26
path: schemas
description: Status of a Service Binding's deployment to the Cloudflare network
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# addressing_provisioning

Status of a Service Binding's deployment to the Cloudflare network

```yaml
{"description": "Status of a Service Binding's deployment to the Cloudflare network", "type": "object", "properties": {"state": {"description": "When a binding has been deployed to a majority of Cloudflare datacenters, the binding will become active and can be used with its associated service.\n", "type": "string", "example": "provisioning", "enum": ["provisioning", "active"], "x-auditable": true}}}
```
