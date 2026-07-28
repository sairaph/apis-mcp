---
title: alexandria_application
page_id: schema-alexandria-application-9b2e4b11
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# alexandria_application

```yaml
{"type": "object", "properties": {"application_confidence_score": {"$ref": "#/components/schemas/alexandria_application_confidence_score"}, "application_score_composition": {"$ref": "#/components/schemas/alexandria_application_score_composition"}, "application_source": {"$ref": "#/components/schemas/alexandria_application_source"}, "application_type": {"$ref": "#/components/schemas/alexandria_application_type"}, "application_type_description": {"$ref": "#/components/schemas/alexandria_application_type_description"}, "created_at": {"$ref": "#/components/schemas/alexandria_application_created_at"}, "gen_ai_score": {"$ref": "#/components/schemas/alexandria_application_gen_ai_score"}, "hostnames": {"$ref": "#/components/schemas/alexandria_application_hostnames"}, "human_id": {"$ref": "#/components/schemas/alexandria_application_human_id"}, "id": {"$ref": "#/components/schemas/alexandria_application_id"}, "intel_id": {"$ref": "#/components/schemas/alexandria_application_intel_id"}, "ip_subnets": {"$ref": "#/components/schemas/alexandria_application_ip_subnets"}, "name": {"$ref": "#/components/schemas/alexandria_application_name"}, "port_protocols": {"$ref": "#/components/schemas/alexandria_application_port_protocols"}, "support_domains": {"$ref": "#/components/schemas/alexandria_application_support_domains"}, "supported": {"$ref": "#/components/schemas/alexandria_application_supported"}, "updated_at": {"$ref": "#/components/schemas/alexandria_application_updated_at"}, "version": {"$ref": "#/components/schemas/alexandria_application_version"}}, "required": ["application_source", "application_type", "application_type_description", "gen_ai_score", "application_confidence_score", "created_at", "supported", "hostnames", "human_id", "id", "ip_subnets", "name", "port_protocols", "support_domains", "updated_at", "version"]}
```
