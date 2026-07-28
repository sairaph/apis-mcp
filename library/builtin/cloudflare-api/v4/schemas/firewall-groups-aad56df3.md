---
title: firewall_groups
page_id: schema-firewall-groups-aad56df3
path: schemas
description: An object that allows you to enable or disable WAF rule groups for the current WAF override. Each key of this object must be the ID of a WAF rule group, and each value must be a valid WAF action (usually `default` or `disable`). When creating a new URI-based WAF override, you must provide a `groups` object or a `rules` object.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# firewall_groups

An object that allows you to enable or disable WAF rule groups for the current WAF override. Each key of this object must be the ID of a WAF rule group, and each value must be a valid WAF action (usually `default` or `disable`). When creating a new URI-based WAF override, you must provide a `groups` object or a `rules` object.

```yaml
{"description": "An object that allows you to enable or disable WAF rule groups for the current WAF override. Each key of this object must be the ID of a WAF rule group, and each value must be a valid WAF action (usually `default` or `disable`). When creating a new URI-based WAF override, you must provide a `groups` object or a `rules` object.", "type": "object", "example": {"ea8687e59929c1fd05ba97574ad43f77": "default"}, "additionalProperties": true}
```
