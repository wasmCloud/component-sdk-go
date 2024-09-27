# Custom WIT example

This example shows how to combine the component-sdk with your custom interfaces.

Build the example:

```bash
wash build
```

Deploy

```bash
wash app deploy ./wadm.yaml
```

Call the custom interface:

```bash
wash call invoke_example-invoker example:invoker/invoker.call
```
