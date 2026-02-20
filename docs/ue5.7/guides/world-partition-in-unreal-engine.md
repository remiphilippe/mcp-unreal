<!-- Source: https://dev.epicgames.com/documentation/en-us/unreal-engine/world-partition-in-unreal-engine -->

Building large maps used to require developers to manually divide maps into sublevels, then use the Level streaming system to load and unload them as the player traversed the landscape. This method often created issues sharing files between multiple users, and viewing the whole world in context became a difficult task.

World Partition is an automatic data management and distance-based level streaming system that provides a complete solution for large world management. The system removes the previous need to divide large levels into sublevels by storing your world in a single persistent level separated into grid cells, and provides you with an automatic streaming system to load and unload those cells based on distance from a streaming source.

World Partition works closely with the following features:

- [One File Per Actor](https://dev.epicgames.com/documentation/en-us/unreal-engine/one-file-per-actor-in-unreal-engine)
- [World Partition - Data Layers](https://dev.epicgames.com/documentation/en-us/unreal-engine/world-partition---data-layers-in-unreal-engine)
- [Level Instancing](https://dev.epicgames.com/documentation/en-us/unreal-engine/level-instancing-in-unreal-engine)
- [World Partition - Hierarchical Levels of Detail](https://dev.epicgames.com/documentation/en-us/unreal-engine/world-partition---hierarchical-level-of-detail-in-unreal-engine)

## Enabling World Partition

There are three ways to enable World Partition within Unreal Engine:

- Creating a new project from a template in the **Games** category.
- Creating a new Level using the Open World template.
- Converting existing levels to use World Partition.

### Creating Your Project using a Games Template

World Partition is enabled by default in many of the project templates found in the **Games** category.

To reduce complexity and provide a scalable solution when creating new projects, grid cell streaming can be enabled and disabled using the **Enable Streaming** option in the **World Settings**.

The following templates use World Partition but have **Enable Streaming** disabled by default:

- Blank
- First Person
- Third Person
- Top Down
- Vehicle Advanced

### Using the Open World Default Map

The Open World default map type is designed to be a starting point for creating large open-world maps and comes with the following features enabled by default:

- World Partition
- One File Per Actor
- Data Layers
- Hierarchical Levels of Detail

The map contains a sample 2 km x 2 km Landscape with a Landscape Material and lighting setup for an outdoor environment. This includes a sky atmosphere, skylight, directional light, exponential height fog, and volumetric clouds.

To use the Open World default map type within your project:

1. Open the **File** menu and select **New Level**.
2. Select the **Open World** map type.
3. Click the **Create** button to create a new map.

### Converting Existing Levels to use World Partition

You can add World Partition to any Level by converting it using the **Tools > Convert Level** menu option, or by using the World Partition Convert Commandlet.

To use the World Partition Convert commandlet, follow these steps:

Command: `UnrealEditor.exe QAGame  -run=WorldPartitionConvertCommandlet Playground.umap -AllowCommandletRendering`

To convert your existing Levels to World Partition:

1. In Windows, open a Command Prompt window.
2. At the prompt, begin by navigating to the location of the `UnrealEditor.exe` executable file. In the above example: `c:\Builds\Home_UE5_Engine\Engine\Binaries\Win64`.
3. Next, begin the command with the name of the .exe file that will run the commandlet, `UnrealEditor.exe`
4. Add the name of the project. Here, `QAGame`.
5. Continue with the name of the commandlet to run, `-run=WorldPartitionConvertCommandlet`.
6. Add the name of the map file that will be converted. In the above example `Playground.umap`.
7. Finish the command with the additional argument `-AllowCommandletRendering`.
8. Press **Enter** and the commandlet will convert the map to use World Partition.

The following optional arguments are available for this commandlet:

| Optional Argument | Description |
| --- | --- |
| **-SCCProvider=(None,Perforce...)** | Specifies which source control provider to use. To run without source control, specify `-SCCProvider=None`. |
| **-Verbose** | Displays verbose logging. |
| **-ConversionSuffix** | Appends the \_WP suffix to a converted map. This is useful when converting Levels for testing purposes while keeping the source Level intact. |
| **-DeleteSourceLevels** | Deletes source Levels after conversion. |
| **-ReportOnly** | Reports what would happen during the conversion. Does not do the conversion. |
| **-GenerateIni** | Generates a default `.ini` conversion file for this map. Does not do the conversion. |
| **-SkipStableGUIDValidation** | Skips the unstable actor GUIDs validation process. Levels with unstable actor GUIDs will result in different conversion output when converting several times. Resaving the Level fixes this. |
| **-OnlyMergeSubLevels** | Converts and merges Levels and Sublevels to One File Per Actor without World Partition. The converted Level can be used as a Level Instance in a World Partition Level. |
| **-FoliageTypePath=\[Path\]** | Extracts Foliage Types as Assets to the given path. Use if the Level contains embedded Foliage Types. |

If you want to alter the conversion settings, use a default conversion `.ini` file with the commandlet. The `.ini` file needs to be in the same folder as your map file and have the same filename as your map, but with the `.ini` extension. For example, an `.ini` file written for the `FirstPersonExampleMap.umap` would be named `FirstPersonExampleMap.ini`.

Here is an example of a default conversion `.ini` file:

```
[/Script/UnrealEd.WorldPartitionConvertCommandlet]
	EditorHashClass=Class'/Script/Engine.WorldPartitionEditorSpatialHash
	RuntimeHashClass=Class'/Script/Engine.WorldPartitionRuntimeSpatialHash
	LevelsGridPlacement=(("/Game/Maps/Highrise_Audio", Bounds),("/Game/Maps/Highrise_Collisions_Temp", Bounds),("/Game/Maps/Highrise_Gameplay", Bounds),("/Game/Maps/Highrise_Lights", Bounds),("/Game/Maps/Highrise_Vista", AlwaysLoaded))
	HLODLayerAssetsPath=
	DefaultHLODLayerName=

	[/Script/Engine.WorldPartitionEditorSpatialHash]
	CellSize=51200
	WorldImage=None
```

## Using World Partition

The World Partition system works by storing your world in a single persistent Level file and subdividing the space into streamable grid cells using a configurable runtime grid.
These cells are loaded and unloaded at runtime by the presence of streaming sources, such as the player. In this way, Unreal Engine only loads the parts of the Level that the player sees and interacts with at a given time.

### Actors in World Partition

When editing the world, Actors can be added anywhere and are automatically assigned to a grid cell based on their **Is Spatially Loaded** setting, found in their **Details** panel's **World Partition** section.

| Option | Description |
| --- | --- |
| Runtime Grid | Determines in which partition grid this Actor is placed. If **None**, the grid will be chosen by the partition system. |
| Is Spatially Loaded | Determines if the Actor is spatially loaded: If enabled, this Actor is loaded when in range of any streaming source when not assigned to a disabled Data Layer. If disabled, this Actor is loaded when not assigned to a disabled Data Layer. |

Since Actors are saved to their own individual files using the One File Per Actor feature, you do not need to check out the Level file from source control to make changes to the Actors in the world. This frees up the Level file for others on your team.

For more information on the One File Per Actor system and Unreal Engine's integrated source control, see the [One File Per Actor](https://dev.epicgames.com/documentation/en-us/unreal-engine/one-file-per-actor-in-unreal-engine) documentation.

Actors that reference other Actors in the Level will be bundled together and loaded at the same time.

### Streaming Sources

Streaming of grid cells within the grid at runtime is determined by two factors:

- Streaming Sources
- Runtime Grid Settings

The first is the position of streaming sources in the Level.

Streaming sources are components that define a position in the world and trigger the loading of cells around them. Player Controllers are a streaming source. Other streaming sources can be added to the world using the **World Partition Streaming Source** component. For example, a streaming source component can be activated at the location that a player will teleport to, so it can load the cells there. Once the grid cells are loaded, the player teleports to the location and the streaming source component is deactivated. Since there is no longer a streaming source at the player's previous location, those grid cells would be unloaded.

### Using the Player as a Streaming Source

Each Player Controller is used as a World Partition streaming source using the **Enable Streaming Source** option. This is enabled by default.

### Using the World Partition Streaming Source Component

World Partition streaming is also done using the World Partition Streaming Source component:

| Option | Description |
| --- | --- |
| **Default Visualizer Loading Range** | Determines the size of the debug visualizer grid when the visualizer is enabled. |
| **Target Grid** | Determines the streaming grid affected by this source. |
| **Debug Color** | Determines the color used for debugging. |
| **Target HLOD Layer** | Determines the HLOD Layer affected by the streaming source. |
| **Shapes** | Determines the shape list used to build a custom shape for this streaming source. If empty, will use a sphere with a radius equal to the grid loading range. |
| **Priority** | Determines the priority of the streaming source. If a grid cell intersects multiple streaming sources, its priority will be the highest priority amongst all streaming sources. |
| **Streaming Source Enabled** | Determines if this component is enabled. |
| **Target State** | Determines which state the intersecting grid cell should be in (either Loaded or Activated). If a grid cell intersects multiple streaming sources, the target state will be the highest target value (where activated is greater than loaded). |

The Blueprint functions **Enable Streaming Source** and **Disable Streaming Source** will enable and disable streaming with this component.

The Blueprint function **Is Streaming Completed** returns true when the component has finished streaming the grid cells that it intersects with.

### Runtime Grid Settings

The second factor that determines whether a grid cell is loaded or unloaded at runtime is the settings of the runtime grid itself. Runtime grid settings are located in the **World Settings** panel, in the **World Partition Setup** section.

A **2D Runtime Hash** grid is provided by default. Using more than one grid can negatively impact performance.

For more information on the recommended setup and settings for your 2D Runtime Hash grid, see the **Big City** map located in the [City Sample](https://www.fab.com/listings/4898e707-7855-404b-af0e-a505ee690e68) project.

| Option | Description |
| --- | --- |
| **Grid Name** | Contains the name of the runtime grid. |
| **Cell Size** | Determines the size of the grid cells that are used to generate the streaming Levels. In the example, the **Cell Size** is 256m x 256m x 256m. |
| **Loading Range** | Determines the range from a streaming source where cells are loaded. In the image above, the **Loading Range** is a 768 meter radius around a streaming source. |
| **Block on Slow Streaming** | Blocks loading in situations where grid cells are not loading fast enough. |
| **Priority** | Determines the priority of the streaming source. If a grid cell intersects multiple streaming sources, its priority will be the highest priority amongst all streaming sources. |
| **Debug Color** | Determines the color of the grid lines that are shown when **Preview Grids** is enabled. |
| **Preview Grids** | When enabled, displays the grid lines in the viewport. |

### Loading and Unloading Regions in the Editor

To support the development of large worlds, the world is initially unloaded. When the Level opens, the Editor only loads Actors that have their **Is Spatially Loaded** setting marked as **False**, such as environment backdrops and managers. This supports the development of large worlds where it is impossible to load the entire map in the Editor.

#### Load and Unload Regions Using the World Partition Window

In the World Partition window, you can manually select the region to work in. Open the window by selecting **Window > World Partition > World Partition Editor** from the main menu.

In the window, click and drag out a region. Then, right-click the selection to open the context menu and select **Load Region from Selection**.

#### Load and Unload Regions Using Location Volumes

Regions can also be loaded and unloaded using **Location Volumes**.

A Location Volume is an Editor-only volume placed into the level that represents a region of the map in the World Partition window.

To add a Location Volume to your level, follow the steps below:

1. Open the **Place Actors** window, if it is not already open.
2. Search for **Location Volume**. Click on it and drag one into your level.
3. Save your level.

Once the level is saved, a new region will appear in the World Partition window with the same name as the Location Volume.

Select the region and right-click to open the context menu. Click **Load Selected Region** to load the region. Several regions can be loaded or unloaded at the same time by drag selecting or ctrl+clicking multiple regions.

Location Volumes have the following supported settings:

**Brush Settings**

| Setting | Description |
| --- | --- |
| **Display Shaded Volume** | Display the brush with a shaded volume. |
| **Shaded Volume Opacity Value** | Sets the opacity of the shaded volume using a value from 0.0-1.0. |

Many of the options in the **Brush Settings** section depend on the **Brush Shape** you select.

**World Partition**

| Setting | Description |
| --- | --- |
| **Load and Unload** | Loads and unloads the region within the selected volume. |

### Generating a Minimap for the World Partition Window

You can generate a Minimap to make navigating in the World Partition window easier by using the **Build Minimap** option found in the World Partition section of the **Build** menu, or by using the World Partition Minimap Builder commandlet.

Running this commandlet creates a Minimap image of your World Partition world and places it in the World Partition window.

If you build your Minimap and it does not appear in the World Partition window, you will need to enable Virtual Textures support in your project. To enable Virtual Textures, from the main menu, go to **Edit > Project Settings**. Then, enable the **Enable virtual texture support** checkbox.

### Useful Shortcuts and Options in the World Partition Window

| Hotkey | Description |
| --- | --- |
| **Shift+Drag** | Snaps selection to the current runtime grid size. |
| **Double Click** | Moves the camera to the location in all viewports. |
| **Shift+ Double Click** | Starts a PIE session at the clicked location. |
| **Ctrl+Double Click** | Loads a region around the clicked location. |
| **MMB+Drag** | Shows the distance from the clicked point to the end point in Unreal units. |

### Generating Hierarchical Levels of Detail (HLODs)

HLODs are generated using the Build HLODs option found in the World Partition section of the Build menu, or by using the World Partition HLODs Builder commandlet.

Running this commandlet creates the HLOD Actors for your World Partition cells according to the generation settings you specified in your HLOD Layers.
For more information about using Hierarchical Levels of Detail (HLODs) in World Partition and using the World Partition HLODs Builder commandlet, see the [World Partition - Hierarchical Level of Detail](https://dev.epicgames.com/documentation/en-us/unreal-engine/world-partition---hierarchical-level-of-detail-in-unreal-engine) documentation.

### Cooking a World Partition world

Cooking a World Partition map requires the use of the Cook commandlet:

Command: `UnrealEditor.exe QAGame  -run=cook  -targetplatform=WindowsNoEditor -Unversioned -map=Playground`

To cook your World Partition map:

1. In Windows, open a Command Prompt window.
2. At the prompt, begin by navigating to the location of the `UnrealEditor.exe` executable file.
3. Next, begin the command with the name of the .exe file that will run the commandlet, `UnrealEditor.exe`.
4. Add the name of the project. Here, `QAGame`.
5. Continue with the name of the commandlet to run, `-run=cook`.
6. Finish the command with the following arguments:
   - `-targetplatform=WindowsNoEditor` cooks the project for Windows platforms.
   - `-UnVersioned` saves all of the cooked packages without versions. These are then assumed to be the current version on load.
   - `-map=Playground` specifies the name of the map, in this case `Playground.umap`.

For more information on using the Cook commandlet, see [Content Cooking](https://dev.epicgames.com/documentation/en-us/unreal-engine/sharing-and-releasing-projects-for-unreal-engine).

### Using World Partition with Blueprint

Both Blueprint Classes and Level Blueprints are supported in a World Partition world. However, Blueprint Classes are preferred, as any Actors referenced in the Level Blueprint will be marked as Always Loaded.

## Testing a Partitioned World

### Debugging and Runtime Overrides

There are several useful console commands for debugging a World Partition world during runtime.

| Console Command | Description |
| --- | --- |
| **wp.Runtime.ToggleDrawRuntimeHash2D** | Toggles 2D debug display of world partition runtime hash. |
| **wp.Runtime.ToggleDrawRuntimeHash3D** | Toggles 3D debug display of world partition runtime hash. |
| **wp.Runtime.ShowRuntimeSpatialHashGridLevel** | Chooses which grid level to display when showing world partition runtime hash. |
| **wp.Runtime.ShowRuntimeSpatialHashGridLevelCount** | Chooses how many grid levels to display when showing world partition runtime hash. |
| **wp.Runtime.ShowRuntimeSpatialHashGridIndex** | Shows a specific grid when showing world partition runtime hash. An invalid index will show all. |
| **wp.Runtime.RuntimeSpatialHashCellToSourceAngleContributionToCellImportance** | Takes a value between 0 and 1 that modulates the contribution of the angle between streaming source-to-cell vector and source-forward vector to the cell importance. The closer the value to 0, the less the angle will contribute to the cell importance. |
| **wp.Runtime.OverrideRuntimeSpatialHashLoadingRange** | Sets the runtime loading range. Takes the following arguments: `-grid=[index]`: Sets the runtime grid you would like to affect. `-range=[override_loading_range]`: Sets the new runtime loading range |
| **wp.Runtime.MaxLoadingLevelStreamingCells** | Limits the number of concurrent loading World Partition streaming cells. |
| **wp.Runtime.HLOD 0** | Shows the world without HLODs using `wp.Runtime.HLOD`. |

## World Partition Builder Commandlets

World Partition introduces a builder commandlet framework through the **UWorldPartitionBuilderCommandlet** and the **UWorldPartitionBuilder** base class.

These commandlets are used to automate batch processes and generate or modify data in World Partition Levels. Large worlds will not have to be loaded all at once to do things like generating HLODs, AI Nav Data, or resave a large number of Actors.

For more information on using World Partition Builder Commandlets, see [World Partition Builder Commandlets Reference](https://dev.epicgames.com/documentation/en-us/unreal-engine/world-partition-builder-commandlet-reference).
