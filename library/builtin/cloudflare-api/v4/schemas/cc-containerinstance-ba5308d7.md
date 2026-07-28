---
title: cc_ContainerInstance
page_id: schema-cc-containerinstance-ba5308d7
path: schemas
description: Represents a container instance within an application, backed by a Durable Object.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# cc_ContainerInstance

Represents a container instance within an application, backed by a Durable Object.

```yaml
{"description": "Represents a container instance within an application, backed by a Durable Object.", "type": "object", "properties": {"application_id": {"$ref": "#/components/schemas/cc_ApplicationID"}, "created_at": {"$ref": "#/components/schemas/cc_ISO8601Timestamp"}, "deployment_id": {"$ref": "#/components/schemas/cc_DeploymentID"}, "id": {"$ref": "#/components/schemas/cc_ContainerInstanceID"}, "ingress_url": {"description": "Public ingress URL for the instance. Omitted when managed application ingress is not configured.\n", "type": "string", "format": "uri"}, "name": {"description": "The human-readable name of the instance.", "type": "string"}, "placement_id": {"$ref": "#/components/schemas/cc_PlacementID"}, "status": {"$ref": "#/components/schemas/cc_ContainerInstanceStatus"}}, "required": ["id", "name", "application_id", "created_at"]}
```
