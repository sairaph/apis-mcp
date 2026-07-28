---
title: Create a Meter Event Stream Authentication Session
page_id: operation-post-v2-billing-meter-event-session-40c35471
path: operations/untagged
description: Creates a meter event session to send usage on the high-throughput meter event stream. Authentication tokens are only valid for 15 minutes, so you need to create a new meter event session when your token expires.
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - POST
api_endpoints:
    - /v2/billing/meter_event_session
operation_ids:
    - PostV2BillingMeterEventSession
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Create a Meter Event Stream Authentication Session

`POST /v2/billing/meter_event_session`

Operation ID: `PostV2BillingMeterEventSession`

Creates a meter event session to send usage on the high-throughput meter event stream. Authentication tokens are only valid for 15 minutes, so you need to create a new meter event session when your token expires.

## Definition

```yaml
{"summary": "Create a Meter Event Stream Authentication Session", "description": "Creates a meter event session to send usage on the high-throughput meter event stream. Authentication tokens are only valid for 15 minutes, so you need to create a new meter event session when your token expires.", "operationId": "PostV2BillingMeterEventSession", "responses": {"200": {"description": "Successful response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/v2.billing.meter_event_session"}}}}, "default": {"description": "Error response", "content": {"application/json": {"schema": {"oneOf": [{"$ref": "#/components/schemas/v2.error"}]}}}}}}
```
