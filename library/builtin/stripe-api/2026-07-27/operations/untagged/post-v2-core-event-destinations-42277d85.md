---
title: Create an Event Destination
page_id: operation-post-v2-core-event-destinations-9ec0483d
path: operations/untagged
description: Create a new event destination.
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - POST
api_endpoints:
    - /v2/core/event_destinations
operation_ids:
    - PostV2CoreEventDestinations
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Create an Event Destination

`POST /v2/core/event_destinations`

Operation ID: `PostV2CoreEventDestinations`

Create a new event destination.

## Definition

```yaml
{"summary": "Create an Event Destination", "description": "Create a new event destination.", "operationId": "PostV2CoreEventDestinations", "requestBody": {"content": {"application/json": {"schema": {"required": ["enabled_events", "event_payload", "name", "type"], "type": "object", "properties": {"amazon_eventbridge": {"required": ["aws_account_id", "aws_region"], "type": "object", "properties": {"aws_account_id": {"type": "string", "description": "The AWS account ID."}, "aws_region": {"type": "string", "description": "The region of the AWS event source."}}, "description": "Amazon EventBridge configuration."}, "azure_event_grid": {"required": ["azure_region", "azure_resource_group_name", "azure_subscription_id"], "type": "object", "properties": {"azure_region": {"type": "string", "description": "The Azure region."}, "azure_resource_group_name": {"type": "string", "description": "The name of the Azure resource group."}, "azure_subscription_id": {"type": "string", "description": "The Azure subscription ID."}}, "description": "Azure Event Grid configuration."}, "description": {"type": "string", "description": "An optional description of what the event destination is used for."}, "enabled_events": {"type": "array", "description": "The list of events to enable for this endpoint.", "items": {"type": "string"}}, "event_payload": {"type": "string", "description": "Payload type of events being subscribed to.", "enum": ["snapshot", "thin"]}, "events_from": {"type": "array", "description": "Specifies which accounts' events route to this destination.\n`@self`: Receive events from the account that owns the event destination.\n`@accounts`: Receive events emitted from other accounts you manage which includes your v1 and v2 accounts.\n`@organization_members`: Receive events from accounts directly linked to the organization.\n`@organization_members/@accounts`: Receive events from all accounts connected to any platform accounts in the organization.", "items": {"type": "string"}}, "include": {"type": "array", "description": "Additional fields to include in the response.", "items": {"type": "string", "enum": ["webhook_endpoint.signing_secret", "webhook_endpoint.url"]}}, "metadata": {"type": "object", "additionalProperties": {"type": "string"}, "description": "Metadata."}, "name": {"type": "string", "description": "Event destination name."}, "snapshot_api_version": {"type": "string", "description": "If using the snapshot event payload, the API version events are rendered as."}, "type": {"type": "string", "description": "Event destination type.", "enum": ["amazon_eventbridge", "azure_event_grid", "webhook_endpoint"]}, "webhook_endpoint": {"required": ["url"], "type": "object", "properties": {"url": {"type": "string", "description": "The URL of the webhook endpoint."}}, "description": "Webhook endpoint configuration."}}}}}}, "responses": {"200": {"description": "Successful response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/v2.core.event_destination"}}}}, "default": {"description": "Error response", "content": {"application/json": {"schema": {"oneOf": [{"$ref": "#/components/schemas/v2.error.expired_azure_partner_authorization"}, {"$ref": "#/components/schemas/v2.error.idempotency_error"}, {"$ref": "#/components/schemas/v2.error.invalid_azure_partner_authorization"}, {"$ref": "#/components/schemas/v2.error.invalid_azure_partner_configuration"}, {"$ref": "#/components/schemas/v2.error"}]}}}}}}
```
