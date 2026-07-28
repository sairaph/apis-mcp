---
title: Update Zone Bot Management Config
page_id: operation-put-zones-zone-id-bot-management-8c68de28
path: operations/bot-settings
description: |-
    Updates the Bot Management configuration for a zone.

    This API is used to update:
    - **Bot Fight Mode**
    - **Super Bot Fight Mode**
    - **Bot Management for Enterprise**

    See [Bot Plans](https://developers.cloudflare.com/bots/plans/) for more information on the different plans
    \
    If you recently upgraded or downgraded your plan, refer to the following examples to clean up old configurations.
    Copy and paste the example body to remove old zone configurations based on your current plan.
    #### Clean up configuration for Bot Fight Mode plan
    ```json
    {
      "sbfm_likely_automated": "allow",
      "sbfm_definitely_automated": "allow",
      "sbfm_verified_bots": "allow",
      "sbfm_static_resource_protection": false,
      "optimize_wordpress": false,
      "suppress_session_score": false
    }
    ```
    #### Clean up configuration for SBFM Pro plan
    ```json
    {
      "sbfm_likely_automated": "allow",
      "fight_mode": false
    }
    ```
    #### Clean up configuration for SBFM Biz plan
    ```json
    {
      "fight_mode": false
    }
    ```
    #### Clean up configuration for BM Enterprise Subscription plan
    It is strongly recommended that you ensure you have [custom rules](https://developers.cloudflare.com/waf/custom-rules/) in place to protect your zone before disabling the SBFM rules. Without these protections, your zone is vulnerable to attacks.
    ```json
    {
      "sbfm_likely_automated": "allow",
      "sbfm_definitely_automated": "allow",
      "sbfm_verified_bots": "allow",
      "sbfm_static_resource_protection": false,
      "optimize_wordpress": false,
      "fight_mode": false
    }
    ```
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /zones/{zone_id}/bot_management
operation_ids:
    - bot-management-for-a-zone-update-config
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update Zone Bot Management Config

`PUT /zones/{zone_id}/bot_management`

Operation ID: `bot-management-for-a-zone-update-config`

Updates the Bot Management configuration for a zone.

This API is used to update:
- **Bot Fight Mode**
- **Super Bot Fight Mode**
- **Bot Management for Enterprise**

See [Bot Plans](https://developers.cloudflare.com/bots/plans/) for more information on the different plans
\
If you recently upgraded or downgraded your plan, refer to the following examples to clean up old configurations.
Copy and paste the example body to remove old zone configurations based on your current plan.
#### Clean up configuration for Bot Fight Mode plan
```json
{
  "sbfm_likely_automated": "allow",
  "sbfm_definitely_automated": "allow",
  "sbfm_verified_bots": "allow",
  "sbfm_static_resource_protection": false,
  "optimize_wordpress": false,
  "suppress_session_score": false
}
```
#### Clean up configuration for SBFM Pro plan
```json
{
  "sbfm_likely_automated": "allow",
  "fight_mode": false
}
```
#### Clean up configuration for SBFM Biz plan
```json
{
  "fight_mode": false
}
```
#### Clean up configuration for BM Enterprise Subscription plan
It is strongly recommended that you ensure you have [custom rules](https://developers.cloudflare.com/waf/custom-rules/) in place to protect your zone before disabling the SBFM rules. Without these protections, your zone is vulnerable to attacks.
```json
{
  "sbfm_likely_automated": "allow",
  "sbfm_definitely_automated": "allow",
  "sbfm_verified_bots": "allow",
  "sbfm_static_resource_protection": false,
  "optimize_wordpress": false,
  "fight_mode": false
}
```

## Definition

```yaml
{"operationId": "bot-management-for-a-zone-update-config", "summary": "Update Zone Bot Management Config", "description": "Updates the Bot Management configuration for a zone.\n\nThis API is used to update:\n- **Bot Fight Mode**\n- **Super Bot Fight Mode**\n- **Bot Management for Enterprise**\n\nSee [Bot Plans](https://developers.cloudflare.com/bots/plans/) for more information on the different plans\n\\\nIf you recently upgraded or downgraded your plan, refer to the following examples to clean up old configurations.\nCopy and paste the example body to remove old zone configurations based on your current plan.\n#### Clean up configuration for Bot Fight Mode plan\n```json\n{\n  \"sbfm_likely_automated\": \"allow\",\n  \"sbfm_definitely_automated\": \"allow\",\n  \"sbfm_verified_bots\": \"allow\",\n  \"sbfm_static_resource_protection\": false,\n  \"optimize_wordpress\": false,\n  \"suppress_session_score\": false\n}\n```\n#### Clean up configuration for SBFM Pro plan\n```json\n{\n  \"sbfm_likely_automated\": \"allow\",\n  \"fight_mode\": false\n}\n```\n#### Clean up configuration for SBFM Biz plan\n```json\n{\n  \"fight_mode\": false\n}\n```\n#### Clean up configuration for BM Enterprise Subscription plan\nIt is strongly recommended that you ensure you have [custom rules](https://developers.cloudflare.com/waf/custom-rules/) in place to protect your zone before disabling the SBFM rules. Without these protections, your zone is vulnerable to attacks.\n```json\n{\n  \"sbfm_likely_automated\": \"allow\",\n  \"sbfm_definitely_automated\": \"allow\",\n  \"sbfm_verified_bots\": \"allow\",\n  \"sbfm_static_resource_protection\": false,\n  \"optimize_wordpress\": false,\n  \"fight_mode\": false\n}\n```\n", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/bot-management_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"examples": {"example1": {"summary": "Bot Fight Mode default example", "value": {"ai_bots_protection": "disabled", "cf_robots_variant": "off", "crawler_protection": "disabled", "enable_js": true, "fight_mode": true, "is_robots_txt_managed": false}}, "example2": {"summary": "SBFM Pro plan default example", "value": {"ai_bots_protection": "disabled", "cf_robots_variant": "off", "crawler_protection": "disabled", "enable_js": true, "is_robots_txt_managed": false, "sbfm_definitely_automated": "block", "sbfm_static_resource_protection": true, "sbfm_verified_bots": "block"}}, "example3": {"summary": "SBFM Biz plan default example", "value": {"ai_bots_protection": "disabled", "cf_robots_variant": "off", "crawler_protection": "disabled", "enable_js": true, "is_robots_txt_managed": false, "sbfm_definitely_automated": "block", "sbfm_likely_automated": "managed_challenge", "sbfm_static_resource_protection": true, "sbfm_verified_bots": "block"}}, "example4": {"summary": "BM Enterprise Subscription default example", "value": {"ai_bots_protection": "disabled", "auto_update_model": true, "bm_cookie_enabled": true, "cf_robots_variant": "off", "crawler_protection": "disabled", "enable_js": true, "is_robots_txt_managed": false}}, "example5": {"summary": "Clean up configuration for Bot Fight Mode plan", "value": {"bm_cookie_enabled": true, "optimize_wordpress": false, "sbfm_definitely_automated": "allow", "sbfm_likely_automated": "allow", "sbfm_static_resource_protection": false, "sbfm_verified_bots": "allow", "suppress_session_score": false}}, "example6": {"summary": "Clean up configuration for SBFM Pro plan", "value": {"fight_mode": false, "sbfm_likely_automated": "allow"}}, "example7": {"summary": "Clean up configuration for SBFM Biz plan", "value": {"fight_mode": false}}, "example8": {"summary": "Clean up configuration for BM Enterprise Subscription", "value": {"fight_mode": false, "optimize_wordpress": false, "sbfm_definitely_automated": "allow", "sbfm_likely_automated": "allow", "sbfm_static_resource_protection": false, "sbfm_verified_bots": "allow"}}, "example9": {"summary": "Block AI Bots on Ad pages only", "value": {"ai_bots_protection": "only_on_ad_pages"}}}, "schema": {"$ref": "#/components/schemas/bot-management_config_single"}}}}, "responses": {"200": {"description": "Update Bot Management response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/bot-management_bot_management_response_body"}}}}, "4XX": {"description": "Update Bot Management response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/bot-management_bot_management_response_body"}, {"$ref": "#/components/schemas/bot-management_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Bot Settings"], "x-api-token-group": ["Bot Management Write"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}}
```
