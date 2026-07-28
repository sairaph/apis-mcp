---
title: firewall_package_definition
page_id: schema-firewall-package-definition-1e62bcd0
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# firewall_package_definition

```yaml
{"type": "object", "properties": {"description": {"$ref": "#/components/schemas/firewall_schemas-description"}, "detection_mode": {"$ref": "#/components/schemas/firewall_detection_mode"}, "id": {"$ref": "#/components/schemas/firewall_identifier"}, "name": {"$ref": "#/components/schemas/firewall_name"}, "status": {"$ref": "#/components/schemas/firewall_status"}, "zone_id": {"$ref": "#/components/schemas/firewall_identifier"}}, "required": ["id", "name", "description", "detection_mode", "zone_id"], "title": "Traditional WAF package"}
```
