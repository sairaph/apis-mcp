---
title: access_scim_config_mapping
page_id: schema-access-scim-config-mapping-f04e085c
path: schemas
description: Transformations and filters applied to resources before they are provisioned in the remote SCIM service.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# access_scim_config_mapping

Transformations and filters applied to resources before they are provisioned in the remote SCIM service.

```yaml
{"description": "Transformations and filters applied to resources before they are provisioned in the remote SCIM service.", "type": "object", "properties": {"enabled": {"description": "Whether or not this mapping is enabled.", "type": "boolean"}, "filter": {"description": "A [SCIM filter expression](https://datatracker.ietf.org/doc/html/rfc7644#section-3.4.2.2) that matches resources that should be provisioned to this application.", "type": "string", "example": "title pr or userType eq \"Intern\""}, "operations": {"description": "Whether or not this mapping applies to creates, updates, or deletes.", "type": "object", "properties": {"create": {"description": "Whether or not this mapping applies to create (POST) operations.", "type": "boolean"}, "delete": {"description": "Whether or not this mapping applies to DELETE operations.", "type": "boolean"}, "update": {"description": "Whether or not this mapping applies to update (PATCH/PUT) operations.", "type": "boolean"}}}, "schema": {"description": "Which SCIM resource type this mapping applies to.", "type": "string", "example": "urn:ietf:params:scim:schemas:core:2.0:User"}, "strictness": {"description": "The level of adherence to outbound resource schemas when provisioning to this mapping. ‘Strict’ removes unknown values, while ‘passthrough’ passes unknown values to the target.", "type": "string", "enum": ["strict", "passthrough"]}, "transform_jsonata": {"description": "A [JSONata](https://jsonata.org/) expression that transforms the resource before provisioning it in the application.", "type": "string", "example": "$merge([$, {'userName': $substringBefore($.userName, '@') & '+test@' & $substringAfter($.userName, '@')}])"}}, "required": ["schema"]}
```
