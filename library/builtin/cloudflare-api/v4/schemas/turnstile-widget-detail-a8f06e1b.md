---
title: turnstile_widget_detail
page_id: schema-turnstile-widget-detail-a8f06e1b
path: schemas
description: A Turnstile widget's detailed configuration
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# turnstile_widget_detail

A Turnstile widget's detailed configuration

```yaml
{"description": "A Turnstile widget's detailed configuration", "type": "object", "properties": {"bot_fight_mode": {"$ref": "#/components/schemas/turnstile_bot_fight_mode"}, "clearance_level": {"$ref": "#/components/schemas/turnstile_clearance_level"}, "created_on": {"$ref": "#/components/schemas/turnstile_created_on"}, "deployed_via": {"$ref": "#/components/schemas/turnstile_deployed_via"}, "domains": {"$ref": "#/components/schemas/turnstile_domains"}, "ephemeral_id": {"$ref": "#/components/schemas/turnstile_ephemeral_id"}, "last_modified_via": {"$ref": "#/components/schemas/turnstile_last_modified_via"}, "mode": {"$ref": "#/components/schemas/turnstile_widget_mode"}, "modified_on": {"$ref": "#/components/schemas/turnstile_modified_on"}, "name": {"$ref": "#/components/schemas/turnstile_name"}, "offlabel": {"$ref": "#/components/schemas/turnstile_offlabel"}, "region": {"$ref": "#/components/schemas/turnstile_region"}, "secret": {"$ref": "#/components/schemas/turnstile_secret"}, "sitekey": {"$ref": "#/components/schemas/turnstile_sitekey"}}, "required": ["sitekey", "secret", "created_on", "modified_on", "name", "domains", "mode", "region", "bot_fight_mode", "offlabel", "clearance_level", "ephemeral_id"]}
```
