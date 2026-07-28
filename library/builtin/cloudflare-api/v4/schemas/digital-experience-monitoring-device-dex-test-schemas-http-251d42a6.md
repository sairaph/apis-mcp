---
title: digital-experience-monitoring_device-dex-test-schemas-http
page_id: schema-digital-experience-monitoring-device-dex-test-schemas-http-251d42a6
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# digital-experience-monitoring_device-dex-test-schemas-http

```yaml
{"type": "object", "properties": {"created": {"description": "Date the test was created, in RFC 3339 format.", "type": "string", "format": "date-time", "example": "2023-10-11T00:00:00Z", "readOnly": true, "x-stainless-terraform-configurability": "computed"}, "data": {"$ref": "#/components/schemas/digital-experience-monitoring_device-dex-test-schemas-data"}, "description": {"$ref": "#/components/schemas/digital-experience-monitoring_device-dex-test-schemas-description"}, "enabled": {"$ref": "#/components/schemas/digital-experience-monitoring_device-dex-test-schemas-enabled"}, "interval": {"$ref": "#/components/schemas/digital-experience-monitoring_device-dex-test-schemas-interval"}, "name": {"$ref": "#/components/schemas/digital-experience-monitoring_device-dex-test-schemas-name"}, "target_policies": {"allOf": [{"$ref": "#/components/schemas/digital-experience-monitoring_device-dex-test-target-policies"}], "x-stainless-terraform-configurability": "computed_optional"}, "targeted": {"type": "boolean", "x-stainless-terraform-configurability": "computed"}, "test_id": {"$ref": "#/components/schemas/digital-experience-monitoring_schemas-test-id"}, "updated": {"description": "Date the test was last updated, in RFC 3339 format.", "type": "string", "format": "date-time", "example": "2023-10-11T00:00:00Z", "readOnly": true, "x-stainless-terraform-configurability": "computed"}}, "required": ["name", "interval", "enabled", "data"]}
```
