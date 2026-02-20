<!-- Source: https://dev.epicgames.com/documentation/en-us/unreal-engine/level-streaming-overview-in-unreal-engine -->

The Level Streaming feature makes it possible to load and unload map files into memory as well as toggle their visibility during play. This makes it possible to have worlds broken up
into smaller chunks so that only the relevant parts of the world are taking up resources and being rendered at any point. If done properly, this allows for the creation of very
large, seamless games that can make the player feel as if they are playing within a world that dwarfs them in size.

## The Persistent Level

The first step towards getting your Levels to blend together seamlessly is creating the persistent Level. This can be thought of as a master Level that is primarily used to govern
which Levels will be streamed in and out.

## Streaming Levels

Streaming Levels are managed through the [**Levels** window](https://dev.epicgames.com/documentation/en-us/unreal-engine/managing-multiple-levels-in-unreal-engine). They can overlap your persistent Level, or be offset to create larger worlds. Each
streaming Level can be set to **Always Loaded** or **Blueprint** for the streaming type. This setting can be toggled in the **Levels** window by right-clicking on the sublevel.

When a streaming Level is set to **Always Loaded**, it is loaded together with the persistent Level. It also becomes visible at the same time as the persistent Level.
It ignores streaming volumes assigned to it, as well as any loading or unloading requests sent from Blueprints or C++. This type of sublevel is often used for splitting
content that would ordinarily be in the persistent Level into multiple "layers", so artists can collaborate working in parallel without blocking each other. For more information on this workflow,
see [Managing Multiple Levels](https://dev.epicgames.com/documentation/en-us/unreal-engine/managing-multiple-levels-in-unreal-engine). In game, an **Always Loaded** sublevel is never streamed out, unless the game changes the persistent Level.

### Dynamic Streaming Methods

Streaming Levels with the **Blueprint** streaming type can actually be controlled with Level Streaming Volumes, Blueprints, or C++ code. These Levels can be loaded
or unloaded dynamically.

### Level Streaming Volumes

**Level Streaming Volumes** make controlling level streaming very easy. The idea is simple: load/unload requests for a streaming Level are issued based on whether the viewpoint is
inside any of the Level Streaming Volumes associated with a Level.

Specifically, Level Streaming Volumes can be used in two ways:

- In the **game**, Level Streaming Volumes cause Levels to load when the player's viewpoint is inside the volume and unload when the player's viewpoint is outside the volume.
- In the **editor**, Level Streaming Volumes can be used to preview Level streaming by automatically hiding/unhiding Levels based on the location of the perspective viewport's camera.

Volume-based level streaming is simple to use and does not require any scripting, making it an ideal way to control level streaming. Furthermore, volume-based level streaming is
much easier to maintain than scripted streaming: when demands on the loading system change, Level loading/unloading behavior can be modified by simply resizing the streaming volumes.

- [Level Streaming Volumes Reference](https://dev.epicgames.com/documentation/en-us/unreal-engine/level-streaming-volumes-reference-in-unreal-engine)
- [Level Streaming using Volumes](https://dev.epicgames.com/documentation/en-us/unreal-engine/level-streaming-using-volumes-in-unreal-engine)

### Scripted Level Streaming

If you want to set up more complex behavior for loading and unloading Levels, similar to the door unlocking example above, you can also set up Blueprint or C++ logic for Level
loading and unloading. The key functions are **Load Stream Level** and **Unload Stream Level**, and you designate the Level to load by its name.

In addition to **Load Stream Level** and **Unload Stream Level**, the **Get Streaming Level** function enables even more dynamic streaming behavior. Since it gives you access
to the actual level streaming object, you can modify and query its state. Using this function in combination with **Create Instance**, you can also create and stream in copies of a specific sublevel.
By applying transforms and sending parameters to these copies, you can create procedural worlds.

- [Loading and Unloading Levels using Blueprints](https://dev.epicgames.com/documentation/en-us/unreal-engine/loading-and-unloading-levels-using-blueprints-in-unreal-engine)
- [Loading and Unloading Levels using C++](https://dev.epicgames.com/documentation/en-us/unreal-engine/loading-and-unloading-levels-using-cplusplus-in-unreal-engine)
