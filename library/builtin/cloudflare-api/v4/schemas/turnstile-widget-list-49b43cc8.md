---
title: turnstile_widget_list
page_id: schema-turnstile-widget-list-49b43cc8
path: schemas
description: A Turnstile Widgets configuration as it appears in listings
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# turnstile_widget_list

A Turnstile Widgets configuration as it appears in listings

```yaml
{"description": "A Turnstile Widgets configuration as it appears in listings", "type": "object", "properties": {"bot_fight_mode": {"$ref": "#/components/schemas/turnstile_bot_fight_mode"}, "clearance_level": {"$ref": "#/components/schemas/turnstile_clearance_level"}, "created_on": {"$ref": "#/components/schemas/turnstile_created_on"}, "deployed_via": {"$ref": "#/components/schemas/turnstile_deployed_via"}, "domains": {"$ref": "#/components/schemas/turnstile_domains"}, "ephemeral_id": {"$ref": "#/components/schemas/turnstile_ephemeral_id"}, "last_modified_via": {"$ref": "#/components/schemas/turnstile_last_modified_via"}, "mode": {"$ref": "#/components/schemas/turnstile_widget_mode"}, "modified_on": {"$ref": "#/components/schemas/turnstile_modified_on"}, "name": {"$ref": "#/components/schemas/turnstile_name"}, "offlabel": {"$ref": "#/components/schemas/turnstile_offlabel"}, "region": {"$ref": "#/components/schemas/turnstile_region"}, "sitekey": {"$ref": "#/components/schemas/turnstile_sitekey"}}, "required": ["sitekey", "created_on", "modified_on", "name", "domains", "mode", "region", "bot_fight_mode", "offlabel", "clearance_level", "ephemeral_id"]}
```
