---
title: load-balancing_rules
page_id: schema-load-balancing-rules-923af969
path: schemas
description: 'BETA Field Not General Access: A list of rules for this load balancer to execute.'
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# load-balancing_rules

BETA Field Not General Access: A list of rules for this load balancer to execute.

```yaml
{"description": "BETA Field Not General Access: A list of rules for this load balancer to execute.", "type": "array", "items": {"additionalProperties": false, "description": "A rule object containing conditions and overrides for this load balancer to evaluate.", "properties": {"condition": {"description": "The condition expressions to evaluate. If the condition evaluates to true, the overrides or fixed_response in this rule will be applied. An empty condition is always true. For more details on condition expressions, please see https://developers.cloudflare.com/load-balancing/understand-basics/load-balancing-rules/expressions.", "type": "string", "example": "http.request.uri.path contains \"/testing\"", "x-auditable": true}, "disabled": {"description": "Disable this specific rule. It will no longer be evaluated by this load balancer.", "type": "boolean", "default": false, "x-auditable": true}, "fixed_response": {"$ref": "#/components/schemas/load-balancing_fixed_response"}, "name": {"description": "Name of this rule. Only used for human readability.", "type": "string", "example": "route the path /testing to testing datacenter.", "maxLength": 200, "x-auditable": true}, "overrides": {"$ref": "#/components/schemas/load-balancing_rule_overrides"}, "priority": {"description": "The order in which rules should be executed in relation to each other. Lower values are executed first. Values do not need to be sequential. If no value is provided for any rule the array order of the rules field will be used to assign a priority.", "type": "integer", "default": 0, "minimum": 0, "x-auditable": true}, "terminates": {"description": "If this rule's condition is true, this causes rule evaluation to stop after processing this rule.", "type": "boolean", "x-auditable": true}}, "type": "object"}}
```
