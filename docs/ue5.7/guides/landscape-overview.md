<!-- Source: https://dev.epicgames.com/documentation/en-us/unreal-engine/landscape-overview -->

Using the **Landscape** system, you can create terrain for your world. Mountains, valleys, uneven or sloped ground, and even openings for caves are possible. Using the collection of tools in the Landscape system, you can modify your terrain's shape and appearance.

For information about opening and using the Landscape tool, refer to the Landscape Quick Start Guide.

## Landscape Tool Modes

The Landscape tool has three modes, accessible by their buttons at the top of the Landscape tool's window.

| Icon | Mode | Description |
| --- | --- | --- |
| | **Manage mode** | Create new Landscapes, and modify Landscape components. Manage mode is also where you work with [Landscape Copy Tool](https://dev.epicgames.com/documentation/en-us/unreal-engine/landscape-copy-tool-in-unreal-engine) to copy, paste, import, and export parts of your Landscape. For more information about Manage mode, refer to [Landscape Manage Mode](https://dev.epicgames.com/documentation/en-us/unreal-engine/landscape-manage-mode-in-unreal-engine). |
| | **Sculpt mode** | Modify the shape of your Landscape by selecting and using specific tools. For more information about Sculpt mode, refer to [Landscape Sculpt Mode](https://dev.epicgames.com/documentation/en-us/unreal-engine/landscape-sculpt-mode-in-unreal-engine). |
| | **Paint mode** | Modify the appearance of parts of your Landscape by painting textures on it, based on the layers defined in the Landscape's Material. For more information about Paint mode, refer to [Landscape Paint Mode](https://dev.epicgames.com/documentation/en-us/unreal-engine/landscape-paint-mode-in-unreal-engine). |

Creating a Landscape means creating a Landscape Actor. As with other Actors, you can edit many of its properties, including its assigned Material, in the Level Editor's **Details** panel. For more information about **Details** panels, refer to [Level Editor Details Panel](https://dev.epicgames.com/documentation/en-us/unreal-engine/level-editor-details-panel-in-unreal-engine).

## Landscape Features

Below are the main features and techniques employed by the Landscape terrain system.

### Large Terrain Sizes

The Landscape system paves the way for terrains that are orders of magnitude larger than what has been possible in Unreal Engine previously. Because of its powerful Level of Detail (LOD) system and the way it makes efficient use of memory, heightmaps up to 8192x8192 are now legitimately possible and feasible. Unreal Engine now supports expansive outdoor worlds, which means quick game creation without modifying the stock engine or tools.

### World Partition

Landscape integrates natively with [World Partition](https://dev.epicgames.com/documentation/en-us/unreal-engine/world-partition-in-unreal-engine) to subdivide a landscape into separately streamable parts. The landscape is edited as a single unified object, which is internally subdivided into separate Actors (called **Landscape Proxies**). The Proxies can be dynamically loaded and unloaded in the editor and by the client runtime.

This allows overall landscape sizes that would be too slow to usefully render or edit all at once. It also enables multiple users to edit different areas of the landscape at the same time without file contention conflicts.

### Landscape Memory Usage

Landscapes are generally a better choice for creating large terrains than **Static Meshes**.

Landscapes use 4 bytes per vertex for the vertex data. Static Meshes store position as a 12-byte vector, and tangent X and Z vectors packed into 4 bytes each, and either 16-bit or 32-bit float UVs for a total of either 24 or 28 bytes per vertex.

This means Static Meshes use 6 or 7 times the memory Landscapes use for the same vertex density. Landscapes also store their data as **Textures** and can stream out unused LOD levels for distant areas and load them from disk in the background as you approach them. Landscapes use a regular heightfield that efficiently stores collision data compared to the collision data for Static Meshes.

### Static Render Data Stored as Textures in GPU Memory

On most platforms, the Landscape system stores the render data for the terrain in Textures in GPU memory. This storage can be used for data look-up in the vertex shader. The render data uses a 32-bit Texture for storage, with the height occupying 16-bits in the form of the R and G channels and the normal stored as 28-bit values for X and Y, occupying the B and A channels, respectively.

### Continuous Geo-MipMap LOD

Standard Texture mipmaps handle LODs for Landscape terrains. Each mipmap is a level of detail, and the mipmap to sample can be specified using the `text2Dlod` HLSL instruction. Your Landscape can have a large number of LODs, yet maintain smooth LOD transitions, because mip levels for both LODs involved in a transition can be sampled, and then the heights and X and Y offsets can be interpolated in the vertex shader creating a clean morphing effect.

### Heightmap and Weight Data Streaming

With Textures storing data, the standard Texture streaming system in Unreal Engine handles streaming mipmaps in and out as needed. This applies to the heightmap data and the weights for Texture layers. Only requiring the mipmaps needed for each LOD minimizes the amount of memory used at any time, which means you can create a more extensive terrain.

### High Resolution LOD-Independent Lighting

The entire high-resolution (non-LOD'd) normal data is available for lighting calculations due to the storage of the X and Y slopes of the Landscape.

This means you can always use the highest resolution of the terrain for per-pixel lighting, even on distant components that have been LOD'd out.

When this high-resolution normal data combines with detailed normal maps, Landscape terrains can achieve highly detailed lighting with very little overhead.

### Collision

Landscape uses a heightfield object for its collision. Each target layer can specify a [Physical Material](https://dev.epicgames.com/documentation/en-us/unreal-engine/physical-materials-in-unreal-engine). The collision system will use the dominant layer at each position to determine which Physical Material to use. It is possible to use a reduced resolution collision heightfield (for example, 0.5x render resolution) to save on memory requirements for large Landscape terrains. The collision and render components for distant Landscapes can also be streamed out using the level streaming system.

## Landscape Project Settings

| Option | Description |
| --- | --- |
| **Max Number of Layers** | Defines the maximum number of edit layers that can be added to the landscape. |
| **Show Dialog for Automatic Layer Creation** | When true, automatic edit layer creation pops up a dialog where the new layer can be reordered relative to other layers. |
| **Show Update Edit Layers During Interactive Changes** | For landscape layers-affecting changes, allows the landscape to be updated when performing an interactive change (e.g. when changing an edit layer's alpha). Set to false if the performance when editing gets too bad (the landscape will be properly updated when the dragging operation is done). |
| **Max Components** | Defines the maximum number of components in a Landscape. |
| **Max Image Import Cache Size Mega Bytes** | Defines the maximum size of the import image cache in MB. |
| **Paint Strength Gamma** | Defines the exponent used for adjusting the strength of the **Paint** tool. |
| **Disable Painting Startup Slowdown** | (Will be decomm'ed in 5.8) Enabling this feature creates a reduced brush strength at the beginning of a brush stroke. You need to hover in one place until the brush output is at the desired strength. |
| **Landscape Dirtying Mode** | Defines when the engine requires the Landscape to be resaved: **In Landscape Mode and User Triggered Changes** (default), **In Landscape Mode Only**, or **Auto**. |
| **Brush Size UIMax** | Maximum size that can be set via the slider for the landscape sculpt/paint brushes. |
| **Brush Size Clamp Max** | Maximum size that can be set manually for the landscape sculpt/paint brushes. |
| **Disable Temporal Anti Aliasing in Landscape Mode** | When true, temporal anti-aliasing will be inactive while in landscape mode. This avoids the ghosting effect on the landscape brush but can lead to aliasing or shimmering on other parts of the image. |
| **Default Landscape Material** | Defines which landscape material is assigned to new landscapes by default. |
| **Default Layer Info Object** | Defines which **Layer Info Object** is added to a new landscape by default. |
| **Display Target Layer Thumbnails** | When true, each target layer will have a representative thumbnail in landscape mode. Setting this to false will skip needlessly rendering landscape layer thumbnails, which can improve the editing experience. |
| **Target Layer Default Blend Method** | Target layer blend method to use for newly created Landscape Layer Info assets. This is only used when `DefaultLayerInfoObject` isn't set. |
| **HLOD Max Texture Size** | Maximum size of the textures generated for Landscape HLODs. |
| **Spline Icon World ZOffset** | Offset in Z for the landscape spline icon in world-space. |
| **Spline Icon Scale** | Size of the landscape spline control point icon in the viewport. |
