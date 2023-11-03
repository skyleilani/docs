---
title: "Extend Viam with Modular Resources"
linkTitle: "Modular Resources"
weight: 50
type: "docs"
tags:
  [
    "server",
    "rdk",
    "extending viam",
    "modular resources",
    "components",
    "services",
  ]
description: "You can use hardware components or services that Viam does not natively support, by adding them through existing modular resources from the Viam Registry or by creating new modular resources."
no_list: true
aliases:
  - "/program/extend/modular-resources/"
  - "/extend/modular-resources/"
  - "/extend/"
  - "/program/extend/modular-resources/key-concepts/"
  - "/modular-resources/key-concepts/"
---

Viam provides built-in support for a variety of {{< glossary_tooltip term_id="resource" text="resources" >}}:

- Various types of hardware {{< glossary_tooltip term_id="component" text="components" >}}.
- High-level functionality exposed as {{< glossary_tooltip term_id="service" text="services" >}}.

If the {{< glossary_tooltip term_id="component" text="component" >}} or {{< glossary_tooltip term_id="service" text="service" >}} you want to use for your project is not natively supported, you can use _{{< glossary_tooltip term_id="modular-resource" text="modular resources" >}}_ from [the Viam Registry](#the-viam-registry) or [create your own modular resources](#create-your-own-modular-resources).
You configure and manage modular resources the same way as built-in resources and use the same [standardized APIs](/program/apis/).

## The Viam Registry

The [Viam registry](https://app.viam.com/registry) is the central place where you can browse modular resources that add capabilities to your smart machine beyond what is built-in to `viam-server`.

You can search the available modular resources from the Viam Registry here:

<div id="searchbox"></div>
<p>
<div id="searchstats"></div></p>
<div class="mr-component" id="">
  <div class="modellistheader">
    <div class="type">API</div>
    <div class="name">Model</div>
    <div>Description</div>
  </div>
<div id="hits" class="modellist">
</div>
<div id="pagination"></div>
</div>

You can see details about each module in the [Viam registry](https://app.viam.com/registry) on its dedicated module page.

To use a modular resource from the registry, add it from your robot's **Config** tab in [the Viam app](https://app.viam.com/), using the **Create component** button.

After adding a module to your robot, you can choose to configure it for automatic version updates from the Viam registry, or update your module manually.
By default, newly added modules will remain at the version they were when you installed them, and will not update automatically.

## Create your own modular resources

If none of the existing modular resources in the Viam registry support your use case, you can create your own modular resources:

- **Implement a custom component**: You can write a driver for an unsupported {{< glossary_tooltip term_id="component" text="component" >}} by implementing the corresponding component API.

- **Implement a custom service**: You can implement your own algorithm or {{< glossary_tooltip term_id="model" text="model" >}} against a corresponding service API or use custom algorithms or data models when working with services such as [SLAM](/services/slam/), [vision](/services/vision/), or [motion planning](/services/motion/).

You can write modules in a variety of programming languages, such as, Go, Python, C++, Rust, while implementing the same [APIs](/program/apis/).
To create a new modular resource:

1. [Code a new resource model](/modular-resources/create/) by implementing all methods for the component's or service's [standardized API](/program/apis/).
1. Package your modular resource or modular resources as a {{< glossary_tooltip term_id="module" text="module" >}} and [upload the module to the Viam registry](/modular-resources/upload/) to make it available for deployment to robots.
   You can upload _private_ modules for your [organization](/manage/fleet/organizations/) or _public_ modules.
1. Once you have uploaded your module to the registry, you can [deploy and configure the module](/modular-resources/configure/) from [the Viam app](https://app.viam.com/).
   You can test your added resource using the [**Control** tab](/manage/fleet/#remote-control) and [program](/program/) it with Viam's Go or Python SDKs.

### Unique cases

If you are using unique hardware that does not already have an [appropriate API](/program/apis/#component-apis) defined to support it, you can use the [generic API](/components/generic/) to add support for that unique hardware type to your smart machine.

Some use cases may require you to define a new API, or to deploy custom components using a server on a remote part.
For more information, see [Advanced Modular Resources](/modular-resources/advanced/).

## Related tutorials

{{< cards >}}
{{% card link="/tutorials/projects/make-a-plant-watering-robot/" %}}
{{% card link="/tutorials/custom/custom-base-dog/" %}}
{{% card link="/modular-resources/examples/custom-arm/" %}}
{{< /cards >}}