<!-- Source: https://dev.epicgames.com/documentation/en-us/unreal-engine/event-graph-in-unreal-engine -->

The **EventGraph** of a Blueprint contains a node graph that uses events and function calls to perform actions in response to gameplay events associated with the Blueprint. This is used to add functionality that is common to all instances of a Blueprint. This is where interactivity and dynamic responses are setup. For example, a light Blueprint could respond to a damage event by turning off its `LightComponent` and changing the material used by its mesh. This would automatically provide this behavior to all instances of the light Blueprint.

The **EventGraph** of a Level Blueprint contains a node graph that uses events and function calls to perform actions in response
to gameplay events. This is used to handle events for the level as a whole and to add functionality for specific instances of
Actors and Blueprints within the world.

In either case, an **EventGraph** is used by adding one or more events to act as entry points and then connecting Function
Calls, Flow Control nodes, and Variables to perform the desired actions.

## Working with Graphs

The **Graph** tab displays the visual representation of a particular graph of nodes as it shows all of the nodes contained in the graph as well as the connections between them. It provides editing capabilities for adding and removing nodes, arranging nodes, and creating links between nodes. Breakpoints can also be set in the **Graph** tab to aid in debugging Blueprints.

See the [Blueprint Editor Graph Editor](https://dev.epicgames.com/documentation/en-us/unreal-engine/graph-editor-for-the-blueprints-visual-scripting-editor-in-unreal-engine) for a detailed guide to working with the **EventGraph** and other **Graphs** with Blueprints.
