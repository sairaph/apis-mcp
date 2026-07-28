---
title: firewall_zonelockdown
page_id: schema-firewall-zonelockdown-2fbc830c
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# firewall_zonelockdown

```yaml
{"type": "object", "properties": {"configurations": {"$ref": "#/components/schemas/firewall_configurations"}, "created_on": {"$ref": "#/components/schemas/firewall_created_on"}, "description": {"$ref": "#/components/schemas/firewall_lockdowns_components-schemas-description"}, "id": {"$ref": "#/components/schemas/firewall_lockdowns_components-schemas-id"}, "modified_on": {"$ref": "#/components/schemas/firewall_modified_on"}, "paused": {"$ref": "#/components/schemas/firewall_schemas-paused"}, "urls": {"$ref": "#/components/schemas/firewall_schemas-urls"}}, "additionalProperties": false, "required": ["id", "created_on", "modified_on", "paused", "description", "urls", "configurations"]}
```
