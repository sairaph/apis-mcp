---
title: access_app_req_embedded_policies
page_id: schema-access-app-req-embedded-policies-203526d0
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# access_app_req_embedded_policies

```yaml
{"type": "object", "properties": {"policies": {"description": "The policies that Access applies to the application, in ascending order of precedence. Items can reference existing policies or create new policies exclusive to the application. Reusable and inline policies are mutually exclusive.", "type": "array", "items": {"oneOf": [{"$ref": "#/components/schemas/access_app_policy_link"}, {"allOf": [{"description": "A policy UID to link to this application."}, {"$ref": "#/components/schemas/access_uuid-2"}]}, {"allOf": [{"type": "object"}, {"description": "An application-scoped policy JSON. If the policy does not yet exist, it will be created.", "properties": {"id": {"$ref": "#/components/schemas/access_uuid-2"}}}, {"$ref": "#/components/schemas/access_app_policy_request"}]}]}}}}
```
