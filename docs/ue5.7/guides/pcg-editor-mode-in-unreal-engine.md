<!-- Source: https://dev.epicgames.com/documentation/en-us/unreal-engine/pcg-editor-mode-in-unreal-engine -->

Learn to use this **Experimental** feature, but use caution when shipping with it.

The **PCG editor tool mode** is a feature you can use to place PCG content in levels, including splines, surfaces, painting, and volumes, using a library of customizable tools to create presets that leverage the PCG framework, each with an associated PCG Graph and parameters.

To access the PCG editor mode, open the **Modes** dropdown menu and select **PCG**.

## PCG Mode Tools

When you select one of the tools, the results depend on whether you have an appropriate actor selected.

- If an actor is selected, it adds a PCG component to the actor if needed, and creates a new tool data asset if none exists.

- If no actor is selected, an actor is created to perform the operation, which changes based on the selected graph or preset.

Tool buttons are disabled when you select an actor that isn't the correct actor class or that doesn't have the right components for the graph to run properly on it. Similarly, any presets that aren't compatible with the selected actor aren't shown.

Click on a tool button to start using the tool. This displays the **Apply** and **Cancel** buttons and a secondary row of buttons for presets.

Presets are graphs and instances that are marked as tool presets. They are a quick way to select a graph without using the dropdown menu. Using a preset is functionally equivalent to selecting a graph from the dropdown menu.

## Tool Instance Settings

When you select a tool, the panel shows the instance settings that you can access on the PCG component directly, which provides a way for you to change them while interacting with the tool.

| Instance Setting | Description |
| --- | --- |
| **Tool Graph** | This graph is set on the PCG component; it drives what parameters are available, and you can use it to select the actor class when creating a new actor. |
| **Parameter Overrides** | All the graph parameters that are exposed on the graph are available here. |
| **Data Instance** | Defines which 'data instance' the tool writes to. This has limited use for splines or volumes, but for painting tools, it provides you with a way to write to different layers (and do different processing per layer). You can change layers with keyboard shortcuts (1, 2, ...) |
| **Actor Label** | The label of the spawned actor, if none were selected. Changing the label here renames the actor. The default value comes from the graph's tool settings. |
| **Component Name** | The name of the component added to the actor (if you are not using an existing one). |
| **Actor Class to Spawn** | This is the class of the actor spawned when starting the tool from no selection. This field is not visible when starting the tool on an existing actor. |

## PCG Mode Tools

### Draw Spline

You can use the **Draw Spline** mode to place objects "on a spline" projected on the environment. Examples include fences, roads, and similar, and works with open and closed splines. This is similar to other spline creation modes, except it is tailored for PCG. Graphs that support this tool have the **SplineTool** tool tag in their properties.

### Draw Spline Surface

You can use the **Draw Spline Surface** mode to define a spline-bound closed area, inside which a PCG graph populates the interior. Examples include fields, cornrows, grass, and similar. This tool uses the **SplineSurfaceTool** tool tag.

### Paint

The **Paint** tool provides a way for you to paint on the world (based on collisions), or on the selected actor and is similar to the Foliage mode.

This creates points at the locations where raycasts hit physical objects. You can also remove points when holding the shift key (the brush becomes red). This tool uses the **PaintTool** tool tag.

### Volume

The **Volume** tool provides a way for you to create new PCG volumes by first dragging out the footprint, then the height of the box. This tool is disabled unless the actor is a volume or has a box component. This tool uses the **VolumeTool** tool tag.

## Tool-specific Controls

### Spline Controls

Draw modes control how you interact with the tool and are similar to other spline tools.

### Raycast Rules

The **Raycast Rules** control how several tools interact with the world. When enabled, each rule defines a particular interaction with your project.

| Raycast Rule | Description |
| --- | --- |
| **Landscape** | Accepts interaction on the landscape. |
| **Meshes** | Accepts interaction with meshes (for example, actors with collisions). |
| **Ignore PCG Components** | Rejects interaction on PCG-created components. |
| **Allowed Classes** | Accepts interaction only on actor classes in the list (or derived from a parent class in the list). |
| **Constrain to Actor** | Accepts only interactions on the selected actor. |

## Setting Up a Tool Graph

To set up a tool graph, look in the PCG Graph settings in the **Tool Data** section, and then set the values as appropriate for your new tool graph.

| Tool Data Graph Settings | Description |
| --- | --- |
| **Display Name** | Defines the name shown on the tool preset buttons. |
| **Tooltip** | Defines the tooltip shown when hovering over the tool preset button. |
| **Compatible Tool Tags** | Lists the compatible tags you can use this graph with. You must set this for the graph to appear in the graph drop-down in the matching tool. Valid values: **SplineTool**, **SplineSurfaceTool**, **PaintTool**, **VolumeTool**. |
| **Initial Actor Class To Spawn** | Defines the actor class to spawn when starting the tool from no selection, and acts as a restriction on what actor classes match with this graph. |
| **New Actor Label** | Defines the default actor label used when spawning an actor. |
| **Is Preset** | Controls whether the graph will appear as a tool preset button or not. You can override this in instances. |

## Setting Up an Instance as a Preset

In a similar way to a tool graph, the graph instances have a **Tool Data Override** section.

| Tool Data Override | Description |
| --- | --- |
| **Display Name** | The same as for graphs. |
| **Tooltip** | The same as for graphs. |
| **Is Preset** | Defines whether this instance is a preset or not, regardless of the value on the original graph. |

## PCG Editor Mode Settings

The **PCG Editor Mode settings** control the behavior of the PCG tool mode and are found in **Editor Preferences > PCG Editor Mode Settings**.

| PCG Editor Mode Setting | Description |
| --- | --- |
| **Graph Refresh Rate** | Defines the rate at which changes are propagated for PCG to pick them up. If generation is very slow, you can increase this value. |
| **Hide Tool Buttons During Active Tool** | If enabled, when you enter a tool, the UI hides the tool row, and only shows the presets. |
| **Show Editor Toast on Tool Errors** | Controls whether errors will be shown in a toast or just on the tool window. |
| **Interactive Tool Settings** | Defines what tool controls are shown, and their defaults. Contains pairs of tool class settings and default graph to select. |
| **Default New Actor Name** | If an actor's name isn't provided in the graph, it uses this value instead. |
| **Default New PCG Component Name** | If a PCG component's name isn't provided in the graph, it uses this value instead. |
| **Default New Spline Component Name** | If a Spline component's name isn't provided in the graph, it uses this value instead. |
