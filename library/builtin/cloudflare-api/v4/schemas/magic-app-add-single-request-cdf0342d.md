---
title: magic_app_add_single_request
page_id: schema-magic-app-add-single-request-cdf0342d
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# magic_app_add_single_request

```yaml
{"type": "object", "properties": {"hostnames": {"$ref": "#/components/schemas/magic_app_hostnames"}, "ip_subnets": {"$ref": "#/components/schemas/magic_app_subnets"}, "name": {"$ref": "#/components/schemas/magic_app_name"}, "source_subnets": {"$ref": "#/components/schemas/magic_app_source_subnets"}, "type": {"$ref": "#/components/schemas/magic_app_type"}}, "anyOf": [{"required": ["hostnames"], "title": "Hostnames"}, {"required": ["ip_subnets"], "title": "Subnets"}, {"required": ["source_subnets"], "title": "Source Subnets"}], "required": ["name", "type"]}
```
