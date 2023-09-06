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

A power sensor is a device that reports information about voltage, current and power.
It plays a crucial role in monitoring and managing power usage in different applications.
The power sensor component measures voltage, current, and power consumption.

## Configuration

For configuration information, click on your sensor’s model:

Model | Description <a name="model-table"></a>
----- | -----------
[`fake`](./fake/) | a digital power sensor for testing
[`ina219`](./ina219) | INA219 power sensor; current and power monitor
[`ina226`](./ina226) | INA219 power sensor; current and power monitor
[`renogy`](./renogy) | solar charge controller
