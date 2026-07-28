---
title: resource-tagging_set_tags_request_zone_level
page_id: schema-resource-tagging-set-tags-request-zone-level-691dbfdb
path: schemas
description: Request body schema for setting tags on zone-level resources.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# resource-tagging_set_tags_request_zone_level

Request body schema for setting tags on zone-level resources.

```yaml
{"description": "Request body schema for setting tags on zone-level resources.", "discriminator": {"mapping": {"access_application_policy": "#/components/schemas/resource-tagging_set_tags_request_zone_level_access_application_policy", "api_gateway_operation": "#/components/schemas/resource-tagging_set_tags_request_zone_level_base", "custom_certificate": "#/components/schemas/resource-tagging_set_tags_request_zone_level_base", "custom_hostname": "#/components/schemas/resource-tagging_set_tags_request_zone_level_base", "dns_record": "#/components/schemas/resource-tagging_set_tags_request_zone_level_base", "healthcheck": "#/components/schemas/resource-tagging_set_tags_request_zone_level_base", "load_balancer": "#/components/schemas/resource-tagging_set_tags_request_zone_level_base", "managed_client_certificate": "#/components/schemas/resource-tagging_set_tags_request_zone_level_base", "worker_route": "#/components/schemas/resource-tagging_set_tags_request_zone_level_base", "zone": "#/components/schemas/resource-tagging_set_tags_request_zone_level_base", "zone_ruleset": "#/components/schemas/resource-tagging_set_tags_request_zone_level_base"}, "propertyName": "resource_type"}, "oneOf": [{"$ref": "#/components/schemas/resource-tagging_set_tags_request_zone_level_base"}, {"$ref": "#/components/schemas/resource-tagging_set_tags_request_zone_level_access_application_policy"}]}
```
