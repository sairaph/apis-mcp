---
title: Get Organization Usage (Version 2, Alpha, Restricted)
page_id: operation-get-organizations-organization-id-billable-usage-44b0da14
path: operations/billable-usage-v2
description: |-
    Returns cost and usage data for all accounts within an organization,
    aligned with the [FinOps FOCUS v1.3](https://focus.finops.org/focus-specification/v1-3/)
    Cost and Usage dataset specification.

    Each record represents one billable metric for one account on one day.
    This includes all metered usage, including usage that falls within
    free-tier allowances and may result in zero cost. The response
    includes usage for every account belonging to the specified
    organization.

    **Note:** Cost and pricing fields are not yet populated and
    will be absent from responses until billing integration is complete.

    When `from` and `to` are omitted, defaults to the start of the current
    month through today. The maximum date range is 31 days.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /organizations/{organization_id}/billable/usage
operation_ids:
    - billable-usage-v2-get-organization-usage
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get Organization Usage (Version 2, Alpha, Restricted)

`GET /organizations/{organization_id}/billable/usage`

Operation ID: `billable-usage-v2-get-organization-usage`

Returns cost and usage data for all accounts within an organization,
aligned with the [FinOps FOCUS v1.3](https://focus.finops.org/focus-specification/v1-3/)
Cost and Usage dataset specification.

Each record represents one billable metric for one account on one day.
This includes all metered usage, including usage that falls within
free-tier allowances and may result in zero cost. The response
includes usage for every account belonging to the specified
organization.

**Note:** Cost and pricing fields are not yet populated and
will be absent from responses until billing integration is complete.

When `from` and `to` are omitted, defaults to the start of the current
month through today. The maximum date range is 31 days.

## Definition

```yaml
{"operationId": "billable-usage-v2-get-organization-usage", "summary": "Get Organization Usage (Version 2, Alpha, Restricted)", "description": "Returns cost and usage data for all accounts within an organization,\naligned with the [FinOps FOCUS v1.3](https://focus.finops.org/focus-specification/v1-3/)\nCost and Usage dataset specification.\n\nEach record represents one billable metric for one account on one day.\nThis includes all metered usage, including usage that falls within\nfree-tier allowances and may result in zero cost. The response\nincludes usage for every account belonging to the specified\norganization.\n\n**Note:** Cost and pricing fields are not yet populated and\nwill be absent from responses until billing integration is complete.\n\nWhen `from` and `to` are omitted, defaults to the start of the current\nmonth through today. The maximum date range is 31 days.\n", "parameters": [{"$ref": "#/components/parameters/billable-usage-api_organization_id"}, {"$ref": "#/components/parameters/billable-usage-api_v2_from"}, {"$ref": "#/components/parameters/billable-usage-api_v2_to"}], "responses": {"200": {"description": "Organization usage data was successfully retrieved.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/billable-usage-api_v2_usage_response"}}}}, "4XX": {"description": "Indicates the request failed.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/billable-usage-api_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Billable Usage V2"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}}
```
