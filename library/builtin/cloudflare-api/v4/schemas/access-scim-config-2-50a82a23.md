---
title: access_scim_config-2
page_id: schema-access-scim-config-2-50a82a23
path: schemas
description: Configuration for provisioning to this application via SCIM. This is currently in closed beta.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# access_scim_config-2

Configuration for provisioning to this application via SCIM. This is currently in closed beta.

```yaml
{"description": "Configuration for provisioning to this application via SCIM. This is currently in closed beta.", "type": "object", "properties": {"authentication": {"oneOf": [{"$ref": "#/components/schemas/access_scim_config_single_authentication-2"}, {"$ref": "#/components/schemas/access_scim_config_multi_authentication-2"}]}, "deactivate_on_delete": {"description": "If false, we propagate DELETE requests to the target application for SCIM resources. If true, we only set `active` to false on the SCIM resource. This is useful because some targets do not support DELETE operations.", "type": "boolean"}, "enabled": {"description": "Whether SCIM provisioning is turned on for this application.", "type": "boolean"}, "idp_uid": {"description": "The UID of the IdP to use as the source for SCIM resources to provision to this application.", "type": "string"}, "mappings": {"description": "A list of mappings to apply to SCIM resources before provisioning them in this application. These can transform or filter the resources to be provisioned.", "type": "array", "items": {"$ref": "#/components/schemas/access_scim_config_mapping"}}, "remote_uri": {"description": "The base URI for the application's SCIM-compatible API.", "type": "string"}}, "required": ["remote_uri", "idp_uid"]}
```
