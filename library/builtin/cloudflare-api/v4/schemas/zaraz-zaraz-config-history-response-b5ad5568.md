---
title: zaraz_zaraz-config-history-response
page_id: schema-zaraz-zaraz-config-history-response-b5ad5568
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# zaraz_zaraz-config-history-response

```yaml
{"allOf": [{"$ref": "#/components/schemas/zaraz_api-response-common"}, {"properties": {"result": {"description": "Object where keys are numeric configuration IDs.", "type": "object", "example": {"12345": {"config": {"consent": {"cookieName": "zaraz-consent", "customIntroDisclaimerDismissed": true, "enabled": false}, "dataLayer": true, "debugKey": "my-debug-key", "settings": {"autoInjectScript": true}, "tools": {"aJvt": {"component": "facebook-pixel", "defaultFields": {"testKey": "TEST123456"}, "enabled": true, "name": "Facebook Pixel", "neoEvents": [{"actionType": "pageview", "blockingTriggers": [], "data": {"__zaraz_setting_name": "Page view", "ev": "PageView"}, "firingTriggers": ["Pageview"]}], "permissions": ["access_client_kv"], "settings": {"accessToken": "ABcdEFg", "ecommerce": true, "property": "12345"}, "type": "component"}}, "triggers": {"ktBn": {"Pageview": {"clientRules": [], "description": "All page loads", "excludeRules": [], "loadRules": [{"match": "{{ client.__zarazTrack }}", "op": "EQUALS", "value": "Pageview"}], "name": "Pageview", "system": "pageload"}}}, "variables": {"Autd": {"name": "ip", "type": "string", "value": "{{ system.device.ip }}"}}, "zarazVersion": 43}, "createdAt": "2023-02-23T05:05:55.155273Z", "id": 12345, "updatedAt": "2023-02-23T05:05:55.155273Z", "userId": "278d0d0g123cd8e49d45ea64f12faa37"}, "23456": null}, "additionalProperties": {"allOf": [{"$ref": "#/components/schemas/zaraz_zaraz-config-row-base"}, {"description": "Configuration record corresponding to an ID provided in query params.", "nullable": true, "properties": {"config": {"$ref": "#/components/schemas/zaraz_zaraz-config-return"}}, "required": ["config"], "type": "object"}]}}}, "type": "object"}]}
```
