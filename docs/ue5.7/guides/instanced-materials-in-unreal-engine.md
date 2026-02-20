<!-- Source: https://dev.epicgames.com/documentation/en-us/unreal-engine/instanced-materials-in-unreal-engine -->

**Material instancing** in Unreal Engine is used to change the appearance of a Material without incurring an expensive recompilation of the Material.

Whereas a typical Material cannot be changed without recompiling (something that must happen prior to gameplay), a parameterized Material can be edited in a Material instance without such recompilation. This has numerous workflow advantages, and can improve Material performance.

Certain types of instanced Materials can even change during gameplay in response to in-game events (such as a tree whose Material blackens and chars while it burns). This allows tremendous visual flexibility in your artistic elements.

## Material Inheritance

The relationship between Materials and Material instances is a hierarchical parent-child relationship. A Material instance inherits all of its attributes from the parent (or master) Material. For example, this is the Material graph for one of the chair props found in the starter content.

Any Material instances created from **M\_Chair** inherit all the attributes from the graph shown above.

Note the naming conventions used above. This is a good practice to adopt so that you can easily identify parent Materials and Material instances in the Content Browser.

1. The prefix **M\_** denotes a parent Material, as in **M\_Chair**.

2. The prefix **MI\_** denotes a Material instance, as in the two examples pictured right.

Because they inherit their attributes from the parent, newly created Material instances appear identical to the parent Material when applied to objects in the Level. In the image below, the chair on the far left has the parent Material applied while the center and right chairs use unaltered Material instances.

The key workflow benefit of Material instances is that you can very rapidly customize them the **Material Instance Editor** without ever editing the node graph or recompiling the Material.

## Material Parameterization

It is important to know that you cannot edit every single characteristic of a Material instance by default. To make Material attributes editable within an instance, you must designate them as parameters in the parent Material. This is called **parameterizing** your Material.

A parameter is created like any other data node in the Material Editor, and contains the same information as its non-parameterized counterpart.

For example, a **Constant** expression contains a single floating-point value, and is frequently used to control Material inputs like Roughness and Metallic. The parameterized version of this node is called a **Scalar Parameter**.

Note that the Scalar Parameter also becomes a named value which serves as a conduit to send data values into a Material instance. It is important that you give every parameter a unique, descriptive name in the **Details panel**. Select the node in the Material Graph to access its properties in the Details panel.

In the simple parameterized Material shown below, a **Vector Parameter** is connected to the Base Color input, while **Scalar Parameters** are plugged into Metallic and Roughness.

To further illustrate the idea of parameterization, there is also a constant value of 0.5 passed into the **Specular** input.

When opened in the Material Instance Editor, the three parameters are exposed and editable, whereas the constant is not. Values you want to expose to artists should be parameters, and values that you don't want anyone to change should remain as constants.

## Types of Parameters

Parameters can be used anywhere in your Material graph to drive a wide range of Material effects.

Some of the key parameter types are documented below, and a full list of [Parameter Expressions](https://dev.epicgames.com/documentation/assets/designing-visuals-rendering-and-graphics/materials/material-expressions/parameters) is found here.

### Scalar Parameters

A **ScalarParameter** is a parameter that contains a single floating-point value. Scalar parameters can drive any effect based on single values, as seen in the Roughness and Metallic examples above.

Scalar parameters are also frequently used to control the multiplication factor of an attribute. In this graph, a Scalar parameter is multiplied by a solid color, and the result is plugged into the Emissive Color input. The value in the Scalar Parameter controls the strength of the emissive effect. Higher values increase the emission brightness.

### Vector Parameters

A **VectorParameter** is a parameter that contains a 4-channel vector value, or four floating-point values.

These are generally used to provide configurable colors, but could also be used to represent positional data or drive any effect that requires multiple values.

### Texture Parameters

The most commonly used texture parameter is the TextureSampleParameter2D, which allows you to change textures within a Material instance.

There are several additional types of texture parameters available. Each one is specific to the type of texture that it accepts or the manner in which it is being used. For example:

- TextureSampleParameterCube accepts a TextureCube or cubemap.
- TextureSampleParameterFlipbook accepts a FlipbookTexture.
- TextureSampleParameterMeshSubUV accepts a Texture2D that is used for sub-uv effects with a mesh emitter.
- TextureSampleParameterMeshSubUV accepts a Texture2D that is used for sub-uv blending effects with a mesh emitter.

See the [Material Expression Reference](https://dev.epicgames.com/documentation/en-us/unreal-engine/unreal-engine-material-expressions-reference) for a complete list of texture parameters.

### Static Parameters

**Static** parameters are applied at compile-time, so they can be edited in the Material Instance Editor but not from script or at runtime.

They can be used to mask out branches of a Material. For example, a **StaticSwitch** parameter takes two inputs. It outputs the first value if the parameter value is true, and the outputs the second if false. This produces more optimal code as the branch that was masked out by a static parameter is not executed at runtime.

See [Static Switch Parameter](https://dev.epicgames.com/documentation/assets/designing-visuals-rendering-and-graphics/materials/material-expressions/parameters#StaticSwitchParameter) and [Static Component Mask Parameter](https://dev.epicgames.com/documentation/assets/designing-visuals-rendering-and-graphics/materials/material-expressions/parameters#StaticComponentMaskParameter) for information on the specific static parameter types.

A new Material is compiled for every combination of static parameters in the base Material that are used by instances.

This can lead to an excessive number of shaders that must compile. Try to minimize the number of static parameters in the Material and the number of permutations of those static parameters that are actually used.

## Constant and Dynamic Instances

There are two types of Material instances available in Unreal Engine:

- **Material Instance Constant** -- Only calculated prior to runtime.
- **Material Instance Dynamic** -- Can be calculated (and edited) at runtime.

### Material Instance Constant

A **Material Instance Constant** is an instanced Material that calculates only once, prior to runtime. This means that it cannot change during gameplay. Although they remain constant throughout your game, they still have the performance advantage of not requiring compilation.

For instance, if your game has a variety of cars with different paint jobs whose colors will not change during gameplay, the best practice approach is to create a master Material that represented the base aspects of a generic car paint. Then create **Material Instance Constants** to represent the variations for different types of car, such as different colors, varying levels of roughness, and so on. This approach was demonstrated with the chair example earlier on this page.

Material Instance Constants are created within the [Content Browser](https://dev.epicgames.com/documentation/en-us/unreal-engine/content-browser-in-unreal-engine) and are edited from the [Material Instance Editor](https://dev.epicgames.com/documentation/en-us/unreal-engine/unreal-engine-material-instance-editor-ui).

### Material Instance Dynamic

A **Material Instance Dynamic** (MID) is an instanced Material that can be calculated during gameplay (at runtime). This means that as you play, you can use script (either compiled code or Blueprint visual script) to change the parameters of your Material, thereby altering your Material throughout the game. The possible applications for this are endless, from showing different levels of damage to changing paint colors in an architectural visualization.

MIDs are created within script, either from a parameterized Material or a Material Instance Constant. In Blueprint, one would take a given Material that had parameterized properties, and feed it through a **Create Dynamic Material Instance** node. The result of that node is then applied to the object in question with a **Set Material** node. This produces a new Material that can be changed during gameplay.

## Creating and Using Material Instances

Creating and using Material instances is a two step process. First you must create a parent Material that uses parameter expressions for the properties you want to be able to override in a Material instance. Then you can create a Material Instance and customize the properties in the Material Instance Editor.

To learn how to create a parameterized Material and use it in a Material instance, read here: [Creating and Using Material Instances](https://dev.epicgames.com/documentation/en-us/unreal-engine/creating-and-using-material-instances-in-unreal-engine).
