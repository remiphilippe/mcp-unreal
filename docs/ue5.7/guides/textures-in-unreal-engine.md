<!-- Source: https://dev.epicgames.com/documentation/en-us/unreal-engine/textures-in-unreal-engine -->

**Textures** are image assets that are primarily used in Materials but can also be directly applied outside of Materials, like when using a texture for a heads up display (HUD).

For Materials, textures are mapped to surfaces which the Material is applied to. Textures can be used for a variety of calculations within a Material by being applied directly to an input (such as, Base Color), used as a mask, or using the RGBA values for other calculations.

Materials may make use of several textures that are all sampled and applied for different purposes. For instance, a simple material may have a Base Color texture, a Specular texture, and a Normal Map texture. In addition, there may be a map for Emissive and Roughness stored in the alpha channels of one or more of these same textures. Packing multiple values in a single texture allows them to be used more readily while saving draw calls for performance and reducing disk space.

## Importing Textures

Textures are imported into the engine through the **Content Browser** by using the **Import** button or by dragging-and-dropping images directly from your operating system's windows into the Content Browser.

A variety of image formats and file types are supported:

- .bmp
- .float
- .jpeg
- .jpg
- .pcx
- .png
- .psd
- .tga
- .dds (Cubemap or 2D)
- .exr (HDR)
- .tif (TIFF)
- .tiff (TIFF)

When importing your textures, consider the following suggestions for the dimensions of your textures:

- Use power of two sizes when possible, such as 32, 64, 128, 2048, and so on.
  - Power of two values can be mipmapped and streamed. Non-power of two sizes are never streamed and do not generate mipmaps.
- Some GPUs have hardware limits in the maximum size texture they can support. For example, some GPUs may not support texture sizes larger than 8192 pixels (8k).

## Texture Graph Editor

The **Texture Graph Editor** provides artists a node-based interface to procedurally create and edit textures in Unreal Engine.

You can combine texture graphs with Blueprints, materials, and material functions for unique workflows that are only possible within Unreal Engine. The editor works in conjunction with the Texture Asset Editor, which provides additional controls for managing the texture asset.

For more information, see [Getting started with Texture Graph](https://dev.epicgames.com/documentation/en-us/unreal-engine/getting-started-with-texture-graph-in-unreal-engine).

## Texture Asset Editor

The **Texture Asset Editor** is a standalone window where you can view and edit texture assets.

From this editor window, you can view the texture and its color channels. The **Details** panel provides additional information about the imported texture along with a set of properties to configure the texture. This includes being able to set the compression, adjust brightness and saturation, set level of detail, and much more.

For more information, see [Texture Asset Editor](https://dev.epicgames.com/documentation/en-us/unreal-engine/texture-asset-editor-in-unreal-engine).

## Texture Workflows and Optimizations

The following topics detail some common workflows and optimizations you can do with textures in your projects.

- [Getting started with Texture Graph](https://dev.epicgames.com/documentation/en-us/unreal-engine/getting-started-with-texture-graph-in-unreal-engine) - Fundamentals of the Texture Graph Editor and asset to procedurally create textures.
- [Texture Format Support and Settings](https://dev.epicgames.com/documentation/en-us/unreal-engine/texture-format-support-and-settings-in-unreal-engine) - Reference for supported texture formats and file types and their configuration.
- [Texture Streaming](https://dev.epicgames.com/documentation/en-us/unreal-engine/texture-streaming-in-unreal-engine) - System for loading and unloading textures into and out of memory during gameplay.
- [Virtual Texturing](https://dev.epicgames.com/documentation/en-us/unreal-engine/virtual-texturing-in-unreal-engine) - An overview of the available virtual texturing methods in Unreal Engine.
