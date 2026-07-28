---
title: firewall_anomaly_package
page_id: schema-firewall-anomaly-package-23eb0334
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# firewall_anomaly_package

```yaml
{"allOf": [{"$ref": "#/components/schemas/firewall_package_definition"}, {"properties": {"action_mode": {"$ref": "#/components/schemas/firewall_action_mode"}, "description": {"$ref": "#/components/schemas/firewall_anomaly_description"}, "detection_mode": {"$ref": "#/components/schemas/firewall_anomaly_detection_mode"}, "name": {"$ref": "#/components/schemas/firewall_anomaly_name"}, "sensitivity": {"$ref": "#/components/schemas/firewall_sensitivity"}}, "type": "object"}], "required": ["id", "name", "description", "zone_id", "detection_mode", "sensitivity", "action_mode"], "title": "Anomaly detection WAF package (OWASP)"}
```
