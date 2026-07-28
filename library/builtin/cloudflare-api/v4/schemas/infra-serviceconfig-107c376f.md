---
title: infra_ServiceConfig
page_id: schema-infra-serviceconfig-107c376f
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# infra_ServiceConfig

```yaml
{"discriminator": {"mapping": {"http": "#/components/schemas/infra_HttpServiceConfig", "tcp": "#/components/schemas/infra_TcpServiceConfig"}, "propertyName": "type"}, "oneOf": [{"$ref": "#/components/schemas/infra_HttpServiceConfig"}, {"$ref": "#/components/schemas/infra_TcpServiceConfig"}]}
```
