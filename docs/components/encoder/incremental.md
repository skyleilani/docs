---
title: "Configure an incremental encoder"
linkTitle: "incremental"
type: "docs"
description: "Configure an incremental encoder."
tags: ["encoder", "components"]
# SMEs: Rand
---

Use the `incremental` encoder model to configure [a quadrature encoder](https://en.wikipedia.org/wiki/Incremental_encoder).

Configuring an `incremental` encoder requires configuring two pins on the board to which the encoder is wired.
These two pins provide the phase outputs used to measure the speed and direction of rotation in relation to a given reference point.

{{< tabs name="Configure an incremental encoder" >}}
{{% tab name="Config Builder" %}}

Navigate to the **config** tab of your robot's page in [the Viam app](https://app.viam.com).
Click on the **Components** subtab and navigate to the **Create component** menu.
Enter a name for your encoder, select the type `encoder`, and select the `incremental` model.

![Creation of an incremental encoder in the Viam app config builder.](../img/create-incremental.png)

Click **Create component**.
Fill in the attributes for your encoder:

![Configuration of an incremental encoder in the Viam app config builder.](../img/configure-incremental.png)

{{% /tab %}}
{{% tab name="JSON Template" %}}

```json {class="line-numbers linkable-line-numbers"}
{
    "name": "<encoder_name>",
    "type": "encoder",
    "model" : "incremental",
    "attributes": {
      "board": "<board_name>",
      "pins": {
        "a": <string>,
        "b": <string>
      }
    }
}
```

{{% /tab %}}
{{< /tabs >}}

The following attributes are available for `incremental` encoders:

| Name | Inclusion | Description |
| ---- | --------- | ----------- |
| `board` | *Required* | The name of the board to which the encoder is wired. |
| `pins` | *Required* | A struct holding the names of the pins wired to the encoder: <ul> <li> <code>a</code>: Pin number of one of the pins to which the encoder is wired. Use pin number, not GPIO number. </li> <li> <code>b</code>: Required for two phase encoder. Pin number for the second board pin to which the encoder is wired. </li> </ul> |

Viam also supports a model of encoder called [`"single"`](../single) which requires only one pin (`i`).

### Example Config

The following example shows the configuration of a board and an encoder:

```json {class="line-numbers linkable-line-numbers"}
{
  "components": [
    {
      "name": "local",
      "type": "board",
      "model": "pi",
      "attributes": {}
    },
    {
      "name": "myEncoder",
      "type": "encoder",
      "model": "incremental",
      "attributes": {
        "board": "local",
        "pins": {
          "a": "13",
          "b": "11"
        }
      }
    }
  ]
}
```