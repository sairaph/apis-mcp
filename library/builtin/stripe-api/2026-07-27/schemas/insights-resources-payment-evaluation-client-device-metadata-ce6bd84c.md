---
title: insights_resources_payment_evaluation_client_device_metadata
page_id: schema-insights-resources-payment-evaluation-client-device-metadata-ce6bd84c
path: schemas
description: Client device metadata attached to this payment evaluation.
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# insights_resources_payment_evaluation_client_device_metadata

Client device metadata attached to this payment evaluation.

```yaml
{"title": "InsightsResourcesPaymentEvaluationClientDeviceMetadata", "required": ["radar_session"], "type": "object", "properties": {"radar_session": {"maxLength": 5000, "type": "string", "description": "ID for the Radar Session associated with the payment evaluation. A [Radar Session](https://docs.stripe.com/radar/radar-session) is a snapshot of the browser metadata and device details that help Radar make more accurate predictions on your payments."}}, "description": "Client device metadata attached to this payment evaluation.", "x-expandableFields": []}
```
