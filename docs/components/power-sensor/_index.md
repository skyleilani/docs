---
title: "Power Sensor Component"
linkTitle: "Power Sensor"
childTitleEndOverwrite: "Power Sensor"
weight: 70
no_list: true
type: "docs"
description: "A device that provides information about a robot's systems, including voltage, current, and power consumption."
tags: ["sensor", "components", "power sensor", "ina219", "ina226", "renogy"]
icon: "/icons/components/sensor.svg"
images: ["/icons/components/sensor.svg"]
# SME: #team-bucket
---

A power sensor is a device that reports measurements of the voltage, current and power consumption in your robot's system.
It plays a crucial role in monitoring and managing power usage in different applications.

## Configuration

For configuration information, click on your sensor’s model:

Model | Description <a name="model-table"></a>
----- | -----------
[`fake`](./fake/) | Digital power sensor for testing
[`ina219`](./ina219) | [INA219](https://pdf1.alldatasheet.com/datasheet-pdf/view/249609/TI/INA219.html) power sensor; current and power monitor
[`ina226`](./ina226) | [INA226](https://www.ti.com/lit/ds/symlink/ina226.pdf?ts=1688994548364&ref_url=https%253A%252F%252Fwww.ti.com%252Fproduct%252Fde-de%252FINA226) power sensor; current and power monitor
[`renogy`](./renogy) | [Renogy](https://www.renogy.com/content/RSP200D/RSP200D%20G2%20Datasheet.pdf) solar charge controller

## Control your power sensor with Viam’s client SDK libraries

To get started using Viam's SDKs to connect to and control your robot, go to your robot's page on [the Viam app](https://app.viam.com), navigate to the **Code sample** tab, select your preferred programming language, and copy the sample code generated.

{{% snippet "show-secret.md" %}}

When executed, this sample code will create a connection to your robot as a client.
Then control your robot programmatically by adding API method calls as shown in the following examples.

These examples assume you have a power sensor called `"my_power_sensor"` configured as a component of your robot.
If your power sensor has a different name, change the `name` in the code.

Be sure to import the power sensor package for the SDK you are using:

{{< tabs >}}
{{% tab name="Python" %}}

```python
from viam.components.powersensor import powersensor
```

{{% /tab %}}
{{% tab name="Go" %}}

```go
import (
  "go.viam.com/rdk/components/powersensor"
)
```

{{% /tab %}}
{{< /tabs >}}
