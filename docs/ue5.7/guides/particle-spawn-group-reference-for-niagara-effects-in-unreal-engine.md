<!-- Source: https://dev.epicgames.com/documentation/en-us/unreal-engine/particle-spawn-group-reference-for-niagara-effects-in-unreal-engine -->

**Particle Spawn** modules occur once for each created particle. Modules in this section set up initial values for each particle. If **Use Interpolated Spawning** is set, some Particle Spawn modules will be updated in the Spawn stage instead of in the Particle Update stage. Modules are executed in order from the top to the bottom of the stack.

## Beam Modules

| Module | Description |
| --- | --- |
| **Beam Width** | Controls the width of the spawned beam and writes that width to the **Particles.RibbonWidth** parameter. |
| **Spawn Beam** | Places particles along a bezier spline, or simply along a line between two points. Useful for sprite facing along a beam-style path, or for using with the ribbon renderer. |

## Camera Modules

| Module | Description |
| --- | --- |
| **Camera Offset** | Offsets the particle along the vector between the particle and the camera. |
| **Maintain in Camera Particle Scale** | Retains the in-camera particle size by taking into account the camera's FOV, the particle's camera-relative depth, and the render target's size. |

## Chaos Modules

| Module | Description |
| --- | --- |
| **Apply Chaos Data** | Sets the particle position, velocity, and color from a Chaos solver. |

## Color Modules

| Module | Description |
| --- | --- |
| **Color** | Directly sets the **Particles.Color** parameter, with scale factors for the Float3 Color and Scalar Alpha components. |

## Event Modules

| Module | Description |
| --- | --- |
| **Generate Location Event** | Generates an event that contains the position of the particle. The event payload also contains particle velocity, a particle ID, the age, and a random number. |

## Forces Modules

| Module | Description |
| --- | --- |
| **Acceleration Force** | Adds to the **Physics.Force** parameter, which translates into acceleration within the solver. |
| **Apply Initial Forces** | Converts rotational and linear forces into rotational and linear velocity. |
| **Curl Noise Force** | Adds to **Physics.Force** using a curl noise field. |
| **Drag** | Applies Drag value directly to particle velocity and rotational velocity. Accumulates into **Physics.Drag** and **Physics.RotationalDrag**. |
| **Gravity Force** | Applies a gravitational force (in cm/s) to the **Physics.Force** parameter. |
| **Limit Force** | Scales **Physics.Force** down to the specified magnitude if it exceeds the **Force Limit**. |
| **Line Attraction Force** | Accumulates a pull toward the nearest position on a line segment. |
| **Linear Force** | Adds a force vector (in cm/s) to **Physics.Force** in a specific coordinate space. |
| **Mesh Rotation Force** | Adds a rotational force and accumulates to **Physics.RotationalForce**. |
| **Point Attraction Force** | Accumulates a pull toward **AttractorPosition** into **Physics.Force**. |
| **Point Force** | Adds force from an arbitrary point in space with optional falloff. |
| **Vector Noise Force** | Introduces random noise into **Physics.Force**. |
| **Vortex Force** | Takes a velocity around a vortex axis and injects it into **Physics.Force**. |
| **Wind Force** | Applies a wind force to particles, with an optional **Air Resistance** parameter. |

## Initialization Modules

| Module | Description |
| --- | --- |
| **Initialize Particle** | Contains several common particle parameters including Point Attributes (Lifetime, Position, Mass, Color), Sprite Attributes (Sprite Size, Sprite Rotation), and Mesh Attributes (Mesh Scale). Should be at the top of the stack. |
| **Initialize Ribbon** | Contains common parameters for ribbons including Ribbon Attributes (Ribbon Width, Ribbon Twist). Should be at the top of the stack. |

## Kill Modules

| Module | Description |
| --- | --- |
| **Kill Particles** | Kills all particles if set to True. Allows for particles to be dynamically killed at any point in the execution stack. |
| **Kill Particles in Volume** | Kill particles if they are inside of analytical shapes (box, plane, slab, or sphere). The result can be inverted. |

## Location Modules

| Module | Description |
| --- | --- |
| **Box Location** | Spawns particles in a rectangular box shape. |
| **Cone Location** | Spawns particles in a cone shape. |
| **Cylinder Location** | Spawns particles in a cylinder shape with lathe-style controls. |
| **Grid Location** | Spawns particles in an even distribution on a grid. |
| **Jitter Position** | Jitters a spawned particle in a random direction, on a delay timer. |
| **Rotate Around Point** | Finds a position on a forward vector-aligned circle around a user-defined center point. |
| **Skeletal Mesh Location** | Places particles on a bone, socket, triangle or vertex of a Skeletal Mesh. |
| **Sphere Location** | Spawns particles in a spherical shape. |
| **Static Mesh Location** | Spawns particles from the surface of a static mesh. |
| **System Location** | Spawns particles from the system's location. |
| **Torus Location** | Spawns particles in a torus shape. |

## Mass Modules

| Module | Description |
| --- | --- |
| **Calculate Mass and Rotational Inertia by Volume** | Calculates the mass and rotational inertia based on the particle's bounds and a density value (kg/m^3). |
| **Calculate Size and Rotational Inertia by Mass** | Calculates the particle's scale and rotational inertia based on user-driven mass and density values. |

## Materials Modules

| Module | Description |
| --- | --- |
| **Dynamic Material Parameters** | Writes to the Dynamic Parameter Vertex Interpolator node in the Material Editor. Supports up to four unique dynamic parameter nodes. |

## Math/Blend Modules

| Module | Description |
| --- | --- |
| **Cone Mask** | Defines a cone in 3D space and checks if the position input lies inside it. Returns 1 if inside, 0 otherwise. |
| **Lerp Particle Attributes** | Enables linear interpolation of all default particle parameters. |
| **Recreate Camera Projection** | Recreates the camera-relative world position of a 2D scene capture's pixel. |
| **Temporal Lerp Float** | Performs a slow linear interpolation of a float over time. |
| **Temporal Lerp Vector** | Performs a slow linear interpolation of a vector over time. |

## Mesh Modules

| Module | Description |
| --- | --- |
| **Initialize Mesh Reproduction Sprite** | Chooses a random location on a skeletal mesh and calculates ideal particle size, UV scale, sprite alignment. |
| **Sample Skeletal Mesh Skeleton** | Samples bone or socket positions of a skeletal mesh. |
| **Sample Skeletal Mesh Surface** | Samples the surface of a skeletal mesh. |
| **Sample Static Mesh** | Samples a static mesh and writes values to particle parameters. |
| **Update Mesh Reproduction Sprite** | Used along with Initialize Mesh Reproduction Sprite for animated mesh reproduction. |

## Orientation Modules

| Module | Description |
| --- | --- |
| **Align Sprite to Mesh Orientation** | Aligns sprites to a mesh particle's orientation. |
| **Initial Mesh Orientation** | Aligns a mesh to a vector, or rotates it in place. |
| **Orient Mesh to Vector** | Aligns a mesh to an input vector. |

## Physics Modules

| Module | Description |
| --- | --- |
| **Add Rotational Velocity** | Adds to the Rotational Velocity value in a user-defined space. |
| **Find Kinetic and Potential Energy** | Returns a particle's kinetic energy, potential energy, and their sum. |

## SubUV Modules

| Module | Description |
| --- | --- |
| **SubUVAnimation** | Accepts the total number of sprites in a grid and plots them along a curve for smooth animation. |

## Texture Modules

| Module | Description |
| --- | --- |
| **Sample Pseudo Volume Texture** | Samples the color of a pseudo volume texture based on UVW coordinates. |
| **Sample Texture** | Samples a texture at a specific UV location (GPU simulations only). |
| **Sub UV Texture Sample** | Samples a single texture pixel in a row-by-column fashion. |
| **World Aligned Texture Sample** | Samples a texture's color based on particle position. |

## Utility Modules

| Module | Description |
| --- | --- |
| **Do Once** | Tracks whether its trigger condition has ever been true in a previous frame. |
| **Increment Over Time** | Increases a value each frame using tick delta and a user-specified rate. |
| **Update MS Vertex Animation Tools Morph Targets** | Reads morph target texture data and outputs positions and normal vectors. |

## Vector Field Modules

| Module | Description |
| --- | --- |
| **Apply Vector Field** | Takes vector samples from a vector field sampler and applies as force or velocity. |
| **Sample Vector Field** | Samples a vector field with per-particle intensity and optional falloff. |

## Velocity Modules

| Module | Description |
| --- | --- |
| **Add Velocity** | Assigns a velocity to spawned particles. |
| **Add Velocity from Point** | Adds velocity from an arbitrary point in space with optional falloff. |
| **Add Velocity in Cone** | Adds velocity in a cone shape with parameters for cone angle and velocity distribution. |
| **Inherit Velocity** | Adds inherited velocity from another source. |
| **Scale Velocity** | Multiplies **Particles.Velocity** by a separate vector in a specific coordinate space. |
| **Static Mesh Velocity** | Adds velocity based on normals from a static mesh plus inherited velocity. |
| **Vortex Velocity** | Calculates angular velocity around a vortex axis and injects into **Particles.Velocity**. |

## New Scratch Pad Module

Opens the **Scratch Pad** panel and places a **Scratch Pad module** in the **Selection** panel. Any modules or dynamic inputs you create in the Scratch Pad are automatically connected to your script.

## Set New or Existing Value Directly

Places a **Set Parameter** module in the **Selection** panel. Click the Plus sign (+) icon to select **Add Parameter** or **Create New Parameter**.

### Key Particle Parameters

| Parameter | Description |
| --- | --- |
| **Particles.Age** | The age of a particle. |
| **Particles.Color** | Directly sets the color of the particle. |
| **Particles.ID** | Engine-managed persistent ID for each spawned particle. |
| **Particles.Lifetime** | Lifetime of a spawned particle in seconds. |
| **Particles.Mass** | Mass of a spawned particle. |
| **Particles.MeshOrientation** | Axis-angle rotation applied to a spawned mesh particle. |
| **Particles.NormalizedAge** | Age divided by Lifetime, value between 0 and 1. |
| **Particles.Position** | Position of a spawned particle. |
| **Particles.RibbonID** | Assigns a Ribbon ID. Particles with the same ID are connected into a ribbon. |
| **Particles.RibbonWidth** | Width of a ribbon particle in UE units. |
| **Particles.Scale** | XYZ scale of a non-sprite particle. |
| **Particles.SpriteAlignment** | Makes the texture point toward the sprite's selected alignment axis. |
| **Particles.SpriteFacing** | Makes the surface of the sprite face towards a custom vector. |
| **Particles.SpriteRotation** | Screen-aligned roll of the particle in degrees. |
| **Particles.SpriteSize** | Size of the sprite particle's quad. |
| **Particles.Velocity** | Velocity of a particle in cm/s. |
