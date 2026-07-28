---
title: magic_app_update_request
page_id: schema-magic-app-update-request-c2b82b7d
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# magic_app_update_request

```yaml
{"type": "object", "properties": {"hostnames": {"$ref": "#/components/schemas/magic_app_hostnames"}, "ip_subnets": {"$ref": "#/components/schemas/magic_app_subnets"}, "name": {"$ref": "#/components/schemas/magic_app_name"}, "source_subnets": {"$ref": "#/components/schemas/magic_app_source_subnets"}, "type": {"$ref": "#/components/schemas/magic_app_type"}}, "anyOf": [{"required": ["name"], "title": "Update App Name"}, {"required": ["type"], "title": "Update App Type"}, {"required": ["hostnames"], "title": "Update App Hostnames"}, {"required": ["ip_subnets"], "title": "Update App Subnets"}, {"required": ["source_subnets"], "title": "Update App Source Subnets"}]}
```
