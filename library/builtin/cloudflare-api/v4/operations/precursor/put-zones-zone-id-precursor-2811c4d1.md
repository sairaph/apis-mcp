---
title: Update Zone Precursor Config
page_id: operation-put-zones-zone-id-precursor-ae064c63
path: operations/precursor
description: |-
    Updates the Precursor configuration for a zone.

    `default_mode` sets the zone-level enforcement mode. `enforcement_rules`
    is the ordered list of rules that override enforcement for matching
    requests.

    This is a partial update: only the fields present in the request body
    are changed.

    - Sending an empty array (`[]`) clears all enforcement rules.
    - At least one of `default_mode` or `enforcement_rules` must be present;
      an empty body (`{}`) is rejected with `400`.
    - Rule `id` is read-only (assigned by Cloudflare) and ignored on input.
    - Rule `mode` must be `min-friction` or `max-security` (`off` is not a
      valid rule mode; use `default_mode` to disable enforcement).
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /zones/{zone_id}/precursor
operation_ids:
    - precursor-for-a-zone-update-config
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update Zone Precursor Config

`PUT /zones/{zone_id}/precursor`

Operation ID: `precursor-for-a-zone-update-config`

Updates the Precursor configuration for a zone.

`default_mode` sets the zone-level enforcement mode. `enforcement_rules`
is the ordered list of rules that override enforcement for matching
requests.

This is a partial update: only the fields present in the request body
are changed.

- Sending an empty array (`[]`) clears all enforcement rules.
- At least one of `default_mode` or `enforcement_rules` must be present;
  an empty body (`{}`) is rejected with `400`.
- Rule `id` is read-only (assigned by Cloudflare) and ignored on input.
- Rule `mode` must be `min-friction` or `max-security` (`off` is not a
  valid rule mode; use `default_mode` to disable enforcement).

## Definition

```yaml
{"operationId": "precursor-for-a-zone-update-config", "summary": "Update Zone Precursor Config", "description": "Updates the Precursor configuration for a zone.\n\n`default_mode` sets the zone-level enforcement mode. `enforcement_rules`\nis the ordered list of rules that override enforcement for matching\nrequests.\n\nThis is a partial update: only the fields present in the request body\nare changed.\n\n- Sending an empty array (`[]`) clears all enforcement rules.\n- At least one of `default_mode` or `enforcement_rules` must be present;\n  an empty body (`{}`) is rejected with `400`.\n- Rule `id` is read-only (assigned by Cloudflare) and ignored on input.\n- Rule `mode` must be `min-friction` or `max-security` (`off` is not a\n  valid rule mode; use `default_mode` to disable enforcement).\n", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/precursor_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"examples": {"min-friction": {"summary": "Zone-level min-friction enforcement", "value": {"default_mode": "min-friction"}}, "off": {"summary": "Disable precursor for the zone", "value": {"default_mode": "off"}}, "rules-only": {"summary": "Replace rules only, leaving default_mode unchanged", "value": {"enforcement_rules": [{"description": "Ease friction on the login path", "enabled": true, "expression": "http.request.uri.path eq \"/login\"", "mode": "min-friction"}]}}, "with-rules": {"summary": "max-security default with a min-friction override rule", "value": {"default_mode": "max-security", "enforcement_rules": [{"description": "Ease friction on the login path", "enabled": true, "expression": "http.request.uri.path eq \"/login\"", "mode": "min-friction"}]}}}, "schema": {"$ref": "#/components/schemas/precursor_precursor_config"}}}}, "responses": {"200": {"description": "Update Precursor response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/precursor_precursor_config_response_body"}}}}, "4XX": {"description": "Update Precursor response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/precursor_precursor_config_response_body"}, {"$ref": "#/components/schemas/precursor_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Precursor"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}}
```
