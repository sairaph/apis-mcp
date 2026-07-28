---
title: access_apps
page_id: schema-access-apps-2fb67317
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# access_apps

```yaml
{"type": "object", "anyOf": [{"allOf": [{"$ref": "#/components/schemas/access_basic_app_response_props-2"}, {"$ref": "#/components/schemas/access_self_hosted_props-2"}], "title": "Self Hosted Application", "type": "object"}, {"allOf": [{"$ref": "#/components/schemas/access_basic_app_response_props-2"}, {"$ref": "#/components/schemas/access_saas_props-2"}], "title": "SaaS Application", "type": "object"}, {"allOf": [{"$ref": "#/components/schemas/access_basic_app_response_props-2"}, {"$ref": "#/components/schemas/access_ssh_props-2"}], "title": "Browser SSH Application", "type": "object"}, {"allOf": [{"$ref": "#/components/schemas/access_basic_app_response_props-2"}, {"$ref": "#/components/schemas/access_vnc_props-2"}], "title": "Browser VNC Application", "type": "object"}, {"allOf": [{"$ref": "#/components/schemas/access_basic_app_response_props-2"}, {"$ref": "#/components/schemas/access_app_launcher_props-2"}], "title": "App Launcher Application", "type": "object"}, {"allOf": [{"$ref": "#/components/schemas/access_basic_app_response_props-2"}, {"$ref": "#/components/schemas/access_warp_props-2"}], "title": "Device Enrollment Permissions Application", "type": "object"}, {"allOf": [{"$ref": "#/components/schemas/access_basic_app_response_props-2"}, {"$ref": "#/components/schemas/access_biso_props-2"}], "title": "Browser Isolation Permissions Application", "type": "object"}, {"allOf": [{"$ref": "#/components/schemas/access_basic_app_response_props-2"}, {"$ref": "#/components/schemas/access_bookmark_props-2"}], "title": "Bookmark application", "type": "object"}]}
```
