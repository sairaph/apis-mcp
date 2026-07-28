---
title: tunnel_mesh_configuration_request_body
page_id: schema-tunnel-mesh-configuration-request-body-ca9cc851
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# tunnel_mesh_configuration_request_body

```yaml
{"type": "object", "properties": {"config": {"description": "Provider-specific configuration. Required shape depends on ha_mode. For `aws`, must contain `fnr_id`. For `local`, must contain `vips`. For `none` and `disabled`, must be empty or omitted.", "type": "object", "nullable": true, "oneOf": [{"$ref": "#/components/schemas/tunnel_mesh_aws_config"}, {"$ref": "#/components/schemas/tunnel_mesh_local_config"}, {"additionalProperties": false, "description": "Empty object for none/disabled modes.", "type": "object"}]}, "ha_mode": {"$ref": "#/components/schemas/tunnel_mesh_ha_mode"}}, "required": ["ha_mode"]}
```
