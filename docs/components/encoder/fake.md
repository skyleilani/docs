---
title: "Configure a fake encoder"
linkTitle: "fake"
type: "docs"
description: "Configure a fake encoder."
images: ["/components/img/components/encoder.svg"]
tags: ["encoder", "components"]
# SMEs: Rand
---

The `fake` encoder is an encoder model for testing code without any hardware.

{{< tabs name="Configure an fake encoder" >}}
{{% tab name="Config Builder" %}}

Navigate to the **config** tab of your robot's page in [the Viam app](https://app.viam.com).
Click on the **Components** subtab and navigate to the **Create component** menu.
Enter a name for your encoder, select the type `encoder`, and select the `fake` model.

![Creation of a fake encoder in the Viam app config builder.](../img/create-fake.png)

Click **Create component**.
Fill in the attributes for your encoder:

![Configuration of a fake encoder in the Viam app config builder.](../img/configure-fake.png)

{{% /tab %}}
{{% tab name="JSON Template" %}}

```json {class="line-numbers linkable-line-numbers"}
{
    "name": "<encoder_name>",
    "type": "encoder",
    "model" : "fake",
    "attributes": {}
}
```

{{% /tab %}}
{{< /tabs >}}