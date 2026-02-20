<!-- Source: https://dev.epicgames.com/documentation/en-us/unreal-engine/procedural-content-generation-framework-node-reference-in-unreal-engine -->

The Procedural Content Generation (PCG) Framework utilizes the Procedural Node Graph to generate procedural content both in the Editor and at Runtime. Using a format similar to the Material Editor, spatial data flows into the graph from a PCG Component in your Level and is used to generate points. The points are filtered and modified through a series of nodes, which are listed below:

## Blueprint

| Node | Description |
| --- | --- |
| **Execute Blueprint** | Executes a specified custom Blueprint Class with the Execute or Execute With Context method on a clean instance of a Blueprint Class deriving from `UPCGBlueprintElement`. |

## Control Flow

| Node | Description |
| --- | --- |
| **Branch** | Selects one of two outputs based on a Boolean attribute. Controls execution flow through the graph so that depending on specific circumstances some different parts of a graph are executed. The branch that does not have any data is culled from execution. |
| **Select** | Selects one of two inputs to be forwarded to a single output based on a Boolean attribute. |
| **Select (Multi)** | Multi-input version of the Select node, which can be made to be an integer, enum, or string-based. |
| **Switch** | Multi-output version of the Branch node, which can be made to pick an integer, string, or enum. |

## Debug

| Node | Description |
| --- | --- |
| **Debug** | Debugs the previous node in the graph but is not transient. Does not execute in non-editor builds. |
| **Sanity Check Point Data** | Validates that the input data point(s) have a value in the given range; outside of the range this node logs an error and cancels the generation. |
| **Print String** | Prints a message that outputs a prefixed message optionally to the log, node, and screen. Acts as a passthrough node in the shipping build. |

## Density

| Node | Description |
| --- | --- |
| **Curve Remap Density** | Remaps the density of each point in the point data to another density value according to the provided curve. Final Density = Curve Remap(Input Density) |
| **Density Remap** | Applies a linear transform to the point densities. Optionally, this can be set to not affect values outside the input range. |
| **Distance to Density** | Sets the point density according to the distance of each point from a reference point. For most intents and purposes, this node is superseded by the native Distance node. |

## Filter

| Node | Description |
| --- | --- |
| **Attribute Filter** | General purpose filtering on attributes and properties. Works on both Point Data and Attribute Sets. |
| **Attribute Filter Range** | Range-based version of the Attribute Filter where input data is separated into what's inside the range and what's outside the range. |
| **Density Filter** | Filters points based on density and the provided filter ranges. Fully superseded by the Attribute Filter node, but is more specialized and more efficient. |
| **Discard Points on Irregular Surface** | Tests multiple points around given source points to determine if they are on the same plane. |
| **Filter Data By Attribute** | Separates data based on whether they have a specified metadata attribute. |
| **Filter Data by Index** | Separates data based on their index and the filter provided in the settings. |
| **Filter Data By Tag** | Separates data according to their tags. You can specify a comma-separated list of Tags to filter by. |
| **Filter Data By Type** | Separates data based on their type, as dictated by the Target Type. |
| **Point Filter** | Applies a per-point filter on the In point data. |
| **Point Filter Range** | Applies a range-based filter on point data. |
| **Self Pruning** | Removes intersections between points in the same point data, prioritizing data based on the settings (Large to Small, etc.). |

## Generic

| Node | Description |
| --- | --- |
| **Add Tags** | Adds tags on the provided data based on comma-separated lists of tags. |
| **Apply On Actor** | Sets properties on an actor, driven by the properties provided in an Attribute Set. Changes on an actor are not revertable by PCG. |
| **Delete Tags** | Removes tags from the input data, either for all matches against a comma-separated list or if a tag is not in the provided list. |
| **Gather** | Takes all inputs and generates a single collection holding all the input data. Contains a **Dependency Only** pin to sequence execution. |
| **Get Data Count** | Returns an Attribute Set containing the number of data passed to the input pin. |
| **Get Entries Count** | Returns the number of entries in an Attribute Set. |
| **Get Loop Index** | Returns an attribute set containing the current loop index if this is executed inside of a loop subgraph. |
| **Proxy** | Placeholder node replacement that allows dynamic override during execution of the graph. |
| **Replace Tags** | Replaces tags on the input data by their matching counterpart. Supports 1:1, N:1, or N:N relationships. |
| **Sort Attributes** | Sorts the input data by a specified attribute in ascending or descending order. |
| **Sort Points** | Alias for Sort Attributes. |

## Helpers

| Node | Description |
| --- | --- |
| **Spatial Data Bounds To Point** | Computes the bounds and returns a single point representing the bounds of a spatial data. |

## Hierarchical Generation

| Node | Description |
| --- | --- |
| **Grid Size** | Specifies at which grid size to execute downstream nodes. Used with Hierarchical Generation. |

## Input Output

| Node | Description |
| --- | --- |
| **Data Table Row to Attribute Set** | Extracts a single row from a data table to an Attribute Set. |
| **Load Alembic File** | Loads an Alembic file into PCG point data. Requires the **Procedural Content Generation (PCG) External Data Interop** plugin. |
| **Load Data Table** | Loads a UDataTable into PCG point data. Can either import the table as Point Data or as an Attribute Set. |
| **Load PCG Data Asset** | Loads, either synchronously or asynchronously, a PCG Data Asset object and passes its data downstream in the graph. |

## Metadata

| Node | Description |
| --- | --- |
| **Add Attribute** | Adds an attribute to point data or an attribute set. |
| **Attribute Noise** | Computes new values for a target attribute for each point. Works with Point Data and Attribute Sets. |
| **Attribute Partition** | Splits the input data in a partition according to the attributes selected. |
| **Attribute Rename** | Renames an existing attribute. |
| **Attribute Select** | Computes the **Min**, **Max**, or **Median** on a selected **Axis**. |
| **Attribute String Op** | Performs String-related attribute operations, such as appending strings. |
| **Break Transform Attribute** | Breaks down a Transform attribute into its components: **Translation**, **Rotation**, and **Scale**. |
| **Break Vector Attribute** | Breaks down a Vector attribute into its components: **X**, **Y**, **Z**, and **W**. |
| **Copy Attribute** | Copies an attribute either from the **Attribute** pin or from the input itself, to a new point data. |
| **Create Attribute** | Creates an Attribute Set with a single attribute. |
| **Delete Attributes** | Filters (keep or remove) comma-separated attributes from an Attribute Set or Spatial Data. |
| **Density Noise** | Alias for Attribute Noise. |
| **Filter Attributes by Name** | Filters (keep or remove) comma-separated attributes. Alias for Delete Attributes. |
| **Get Attribute from Point Index** | Retrieves a single point from point data and its attributes in a separate Attribute Set. |
| **Make Transform Attribute** | Creates a Transform attribute from three provided attributes: **Translation**, **Rotation**, and **Scale**. |
| **Make Vector Attribute** | Creates a Vector attribute from two to four attributes based on **Output Type**. |
| **Match And Set Attributes** | Selects an entry in the provided Attribute Set table (Match Data) and copies its values to the input data. |
| **Merge Attributes** | Merges multiple attribute sets (in order of connection) together. |
| **Point Match and Set** | Finds a match for each point based on the selection criteria, then applies the value to an attribute. |
| **Transfer Attribute** | Sets an attribute from an object of the same type with the same data set size. |

### Attribute Bitwise Op

| Node | Description |
| --- | --- |
| **And** | Computes the result of the bitwise AND between two attributes. |
| **Not** | Computes the result of the bitwise NOT between two attributes. |
| **Or** | Computes the result of the bitwise OR between two attributes. |
| **Xor** | Computes the result of the bitwise XOR (Exclusive OR) between two attributes. |

### Attribute Boolean Op

| Node | Description |
| --- | --- |
| **And** | Computes the result of the boolean AND between two attributes. |
| **Imply** | Computes the result of the boolean IMPLY between two attributes. |
| **Nand** | Computes the result of the boolean NAND between two attributes. |
| **Nimply** | Computes the result of the boolean NIMPLY between two attributes. |
| **Nor** | Computes the result of the boolean NOR between two attributes. |
| **Not** | Computes the result of the boolean NOT between two attributes. |
| **Or** | Computes the result of the boolean OR between two attributes. |
| **Xnor** | Computes the result of the boolean XNOR (Exclusive NOR) between two attributes. |
| **Xor** | Computes the result of the boolean XOR (Exclusive OR) between two attributes. |

### Attribute Compare Op

| Node | Description |
| --- | --- |
| **Equal** | Writes the comparison Equal To result between two attributes to a boolean attribute. |
| **Greater** | Writes the comparison Greater Than result between two attributes to a boolean attribute. |
| **Greater or Equal** | Writes the comparison Greater Than Or Equal To result between two attributes to a boolean attribute. |
| **Less** | Writes the comparison Less Than result between two attributes to a boolean attribute. |
| **Less or Equal** | Writes the comparison Less Than or Equal To result between two attributes to a boolean attribute. |
| **Not Equal** | Writes the comparison Not Equal result between two attributes to a boolean attribute. |

### Attribute Maths Op

| Node | Description |
| --- | --- |
| **Abs** | Computes the Absolute Value of an input attribute. |
| **Add** | Adds two input values together. |
| **Ceil** | Rounds an input value up to the next integer. |
| **Clamp** | Constrains input values to a specific range. |
| **Clamp Max** | Provides a maximum value for the Clamp operation. |
| **Clamp Min** | Provides a minimum value for the Clamp operation. |
| **Divide** | Divides the first input by the second. |
| **Floor** | Rounds an input value down to the nearest integer. |
| **Frac** | Returns the fractional portion of an input value. |
| **Lerp** | Linear interpolation between two points using a Ratio value. |
| **Max** | Outputs the higher of two input values. |
| **Min** | Outputs the lower of two input values. |
| **Modulo** | Returns the remainder of dividing the first input by the second. |
| **Multiply** | Multiplies two input values. |
| **One Minus** | Computes 1 minus the input value. |
| **Pow** | Raises a base value to the exponent power. |
| **Round** | Rounds an input value to the nearest whole number. |
| **Set** | Sets the output attribute to the value of the provided attributes. |
| **Sign** | Returns -1 for negative, 0 for zero, or 1 for positive input. |
| **Sqrt** | Computes the Square Root of an input. |
| **Subtract** | Subtracts the second input from the first. |
| **Truncate** | Truncates a value by discarding the fractional part. |

### Attribute Reduce Op

| Node | Description |
| --- | --- |
| **Average** | Gathers ensemble average information. |
| **Max** | Gathers ensemble maximum information. |
| **Min** | Gathers ensemble minimum information. |

### Attribute Rotator Op

| Node | Description |
| --- | --- |
| **Combine** | Combines two rotation values, combining first A, then B. |
| **Inverse Transform Rotation** | Transforms a Rotator by the inverse of the supplied transform. |
| **Invert** | Finds the inverse of a provided Rotator. |
| **Lerp** | Linearly interpolates between two Rotator inputs A and B based on the Ratio. |
| **Normalize** | Clamps an angle to a range of -180 to 180. |
| **Transform Rotation** | Transforms a rotation by a given Transform. |

### Attribute Transform Op

| Node | Description |
| --- | --- |
| **Compose** | Composes two Transforms in order: A * B. |
| **Invert** | Inverts the input Transform. |
| **Lerp** | Linearly interpolates between two Transform inputs A and B based on the Ratio. |

### Attribute Trig Op

| Node | Description |
| --- | --- |
| **Acos** | Returns the inverse cosine (arccos) of an input. |
| **Asin** | Returns the inverse sine (arcsin) of an input. |
| **Atan** | Returns the inverse tangent (arctan) of an input. |
| **Atan 2** | Returns the inverse tangent (arctan2) of 2 inputs (B/A). |
| **Cos** | Returns the cosine (cos) of an input. |
| **Deg to Rad** | Converts degrees to radians. |
| **Rad to Deg** | Converts radians to degrees. |
| **Sin** | Returns the sine (sin) of an input. |
| **Tan** | Returns the tangent (tan) of an input. |

### Attribute Vector Op

| Node | Description |
| --- | --- |
| **Cross** | Outputs the Cross Product of two input vectors. |
| **Distance** | Calculates the distance between two Vector inputs. |
| **Dot** | Returns the Dot Product of two input Vectors. |
| **Inverse Transform Direction** | Transforms a direction Vector by the inverse of the input Transform. |
| **Inverse Transform Location** | Transforms a location by the inverse of the input Transform. |
| **Length** | Returns the length of a Vector. |
| **Normalize** | Outputs a normalized copy of the Vector. Returns a zero vector if vector length is too small. |
| **Rotate Around Axis** | Calculates Vector A rotated by Angle (Deg) around Axis. |
| **Transform Direction** | Transforms an input direction Vector by the supplied transform. |
| **Transform Rotation** | Transforms a rotator or quaternion by the input Transform. |

### Make Rotator Nodes

| Node | Description |
| --- | --- |
| **Make Rot from Angles** | Returns a Rotator created using Roll, Pitch, and Yaw values. |
| **Make Rot from Axis** | Returns a Rotator using a reference frame created from Forward, Right, and Up axis. |
| **Make Rot from X** | Returns a Rotator created using only an X axis. |
| **Make Rot from XY** | Returns a Rotator created using X and Y axes. |
| **Make Rot from XZ** | Returns a Rotator created using X and Z axes. |
| **Make Rot from Y** | Returns a Rotator created using only a Y axis. |
| **Make Rot from YX** | Returns a Rotator created using Y and X axes. |
| **Make Rot from YZ** | Returns a Rotator created using Y and Z axes. |
| **Make Rot from Z** | Returns a Rotator created using only a Z axis. |
| **Make Rot from ZX** | Returns a Rotator created using Z and X axes. |
| **Make Rot from ZY** | Returns a Rotator created using Z and Y axes. |

## Param

| Node | Description |
| --- | --- |
| **Get Actor Property** | Retrieves the contents of a property from the actor holding the PCG component (or higher in the object hierarchy). By default, it looks at actor-level properties (useful for Blueprint variables), but it can look at component properties as well using the **Select Component** option. |
| **Get Property From Object Path** | Similar to Get Actor Property except that it can take actor references (soft object paths) through an Attribute Set. |
| **Point To Attribute Set** | Converts a Point Data to Attribute Set by dropping all of the point properties and keeping only the point attributes. |

## Point Ops

| Node | Description |
| --- | --- |
| **Apply Scale to Bounds** | For each point, the bounds min and max is multiplied by their scale and the scale will be reset to 1. |
| **Bounds Modifier** | Modifies the bounds property on points in the provided point data. |
| **Build Rotation From Up Vector** | Builds a rotation from an up vector. |
| **Combine Points** | For each input Point Data, outputs a new Point Data containing a single point that encompasses all points. |
| **Duplicate Point** | For each point, duplicate the point and move it along an axis defined by the Direction. Useful for building fractal-like patterns. |
| **Extents Modifier** | Modifies the extent of each point in the point data by manipulating the bounds. |
| **Split Points** | For each point, create two points split along the specified Split Axis and Split Position. |
| **Transform Points** | Changes the points transforms using basic random rules. Each component of the transform (translation, rotation, scale) can be set to Absolute instead of relative. Useful for generating spatial variation. |

## Sampler

| Node | Description |
| --- | --- |
| **Copy Points** | Copies an instance of all points in the Source per point in the Target input. Often used to instantiate Point Data generated in local space. |
| **Mesh Sampler** | Samples points on a specified static mesh. Requires the **PCG Geometry Script Interop** plugin and the **Geometry Script** plugin. |
| **Texture Sampler** | Samples the UV of a texture at each point. |
| **Select Points** | Selects a subset of points from the input Point Data using a probability. |
| **Spline Sampler** | Samples points using the spline as the source material. Sampling inside the spline requires the spline to be closed. |
| **Surface Sampler** | Samples points on a Surface data, in a regular grid pattern. Options include Points Per Square Meter, Point Extents, and Looseness. |
| **Volume Sampler** | Samples the provided spatial data on a regular 3D grid. |

## Spatial

| Node | Description |
| --- | --- |
| **Attribute Set To Point** | Converts an Attribute Set to Data Point by creating one default point per entry. |
| **Clip Paths** | Used to intersect or difference splines with Polygon 2D shapes. |
| **Create Points** | Creates a point data containing points from a static description. |
| **Create Points Grid** | Creates a point data containing a simple grid of points. |
| **Create Polygon 2D** | Creates a Polygon 2D data from the input point data or spline data. |
| **Create Spline** | Creates a Spline from the input point data. Contains options for closed or linear splines and custom tangents. |
| **Create Surface From Polygon 2D** | Creates an implicit surface from a Polygon 2D. |
| **Create Surface From Spline** | Creates an implicit surface from a closed spline. |
| **Cull Points Outside Actor Bounds** | Culls points based on the current component bounds with additional control for bounds expansion. |
| **Difference** | Outputs the result of the difference of each source against the union of the differences. Density Function options: Minimum, Clamped Subtraction, Binary. Mode options: Inferred, Continuous, Discrete. |
| **Distance** | For each point in the first input, calculates the distance to the nearest point in the second input. |
| **Find Convex Hull 2D** | Computes a 2D convex hull from the input Point Data using the location only. |
| **Get Actor Data** | General version of the Get Data nodes. Reads data from an Actor using the Actor Filter and the Mode. Most common way to access data from PCG. |
| **Get Bounds** | Creates an attribute set containing the world space bounds (min & max) of any given Spatial data. |
| **Get Landscape Data** | Specialization of Get Actor Data that returns Landscape data. |
| **Get PCG Component Data** | Specialization of Get Actor Data that returns generated output from selected actor PCG components. |
| **Get Points Count** | Returns the number of points in the provided input Point Data. |
| **Get Primitive Data** | Specialization of Get Actor Data that returns Primitive data. |
| **Get Segment** | Returns a segment index point or spline from a point, spline, or Polygon 2D. |
| **Get Spline Data** | Specialization of Get Actor Data that returns Spline data. |
| **Get Texture Data** | Loads a texture to a surface data. Supports sampling of compressed textures, UTexture2DArrays, and CPU-available Textures. |
| **Get Volume Data** | Specialization of Get Actor Data that returns Volume data. |
| **Inner Intersection** | Computes the inner intersection between all data provided to the node. |
| **Intersection** | Computes an outer intersection for each data provided in the Primary Source pin against the union of data on each other Source pin. |
| **Make Concrete** | Collapses composite data types (intersection, difference, union) into point data. |
| **Merge Points** | Combines multiple input point data into a single point data. |
| **Mutate Seed** | Mutates the seed of every point according to its position, previous seed, this node's seed, and the component's seed. |
| **Normal To Density** | Computes point data density based on the point normal and the provided settings. |
| **Offset Polygon** | Applies an offset to a Polygon 2D shape to make it larger or smaller. |
| **Point Neighborhood** | Computes neighborhood-based values including distance to center, average neighborhood center, average density and average color. |
| **Point From Mesh** | Builds a point data containing one point with the bounds of the provided static mesh and a reference to that mesh. |
| **Polygon Operation** | Polygon-to-polygon operations, including intersection, union, and difference. |
| **Projection** | Creates a projection data from a source data to project onto a target. |
| **Spatial Noise** | Constructs a spatially-consistent noise pattern (such as Perlin noise) and writes it to a specified attribute. |
| **Spline Intersection** | Finds intersecting splines in 3D and adds control points at each intersection. |
| **Split Splines** | Divides splines into multiple splines based on alpha, distance, key, or control points predicate. |
| **To Point** | Casts the data to a point data or discretizes the spatial data to point data. |
| **Union** | Creates a logical union between data. Density Function options: Maximum, Clamped Addition, Binary. |
| **World Ray Hit Query** | Creates a surface-like data that performs ray casts in the physics world. |
| **World Volumetric Query** | Creates a volume-like data that gathers points from the physics world. |

## Spawner

| Node | Description |
| --- | --- |
| **Create Target Actor** | Creates an empty actor from a template that can be used as a target for writing PCG artifacts. |
| **Point from Player Pawn** | Creates a point at the current player pawn location. Used during runtime generation. |
| **Spawn Actor** | Spawns either the contents of an actor or an actor per point. Options include Template Actor Class, Collapse Actors, Merge PCG only, No Merging, and Attach modes. |
| **Static Mesh Spawner** | Spawns one static mesh per point. Static Mesh options are added to the Mesh Entries array and selected using each entry's Weight. Mesh Selector Types: PCG Mesh Selector Weighted, By Attribute, Weighted By Category. |

## Subgraph

| Node | Description |
| --- | --- |
| **Loop** | Executes another graph as a subgraph, once per data in the loop pins. Non-loop pins are passed as-is. Feedback pins have special behavior for iterative processing. |
| **Subgraph** | Executes another graph as a subgraph. A graph can call itself recursively. |

## Uncategorized

| Node | Description |
| --- | --- |
| **Add Comment** | Visual aid to categorize and organize a graph. |
| **Add Reroute Node** | Graph organizational tool to add control points on edges. |
| **Add Named Reroute Declaration Node** | Named reroute nodes without visual edges, used to remove long or spaghetti edges across large graphs. |
