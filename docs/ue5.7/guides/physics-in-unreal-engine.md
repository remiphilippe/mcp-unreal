<!-- Source: https://dev.epicgames.com/documentation/en-us/unreal-engine/physics-in-unreal-engine -->

## Physics in Unreal Engine

**Chaos Physics** is a light-weight physics simulation solution available in Unreal Engine, built from the ground up to meet the needs of next-generation games. The system includes the following major features:

- Destruction
- Networked Physics
- Chaos Visual Debugger
- Rigid Body Dynamics
- Rigid Body Animation Nodes and Physical Animation
- Cloth Physics and Machine Learning Cloth Simulation
- Ragdoll Physics
- Vehicles
- Physics Fields
- Fluid Simulation
- Hair Physics
- Flesh Simulation

### Destruction

The **Chaos Destruction** system is a collection of tools within Unreal Engine that can be used to achieve cinematic-quality levels of destruction in real time. The system utilizes **Geometry Collections**, a type of asset built from one or more Static Meshes. These Geometry Collections can be fractured to achieve real-time destruction.

The system provides control over the fracturing process using an intuitive non-linear workflow. Users can create multiple levels of fracturing, as well as selective fracturing on parts of the Geometry Collection. Users can also define **Damage Thresholds** per cluster that will trigger a fracture.

**Connection Graphs** can be used to manipulate how a structure collapses as it takes damage. Chaos Destruction comes with World Support, which allows certain parts of the Geometry Collection to be set to **Kinematic**.

Chaos Destruction has deep integration with the **Niagara particle system**. Niagara can read Chaos Destruction **Break Events** and **Collision Events** to spawn particles or modify an existing particle system at runtime.

Chaos Destruction integrates with **Physics Fields**, a system that affects Chaos Physics simulations at runtime in a specified region of space.

- [Chaos Destruction Documentation](https://dev.epicgames.com/documentation/en-us/unreal-engine/chaos-destruction-in-unreal-engine)

### Networked Physics

**Networked physics** enables physics-driven simulations to work in a multiplayer environment. Physics replication refers to Actors with replicated movement that simulate physics.

Unreal Engine comes with three replication modes:

- **Default** - Legacy physics replication mode. Active on Actors that replicate their movement and their root component is set to simulate physics.
- **Predictive Interpolation** - Designed for server-authoritative Actors. Alters each object's velocity on the client to match the server's velocity, with better handling of interactions and local physics alterations.
- **Resimulation** - Designed for server-authoritative Pawns and Actors. Uses client prediction of physics and handles interactions with better accuracy. Triggers physics resimulation when server state differs from cached client state.

- [Networked Physics Documentation](https://dev.epicgames.com/documentation/en-us/unreal-engine/networked-physics)

### Chaos Visual Debugger

The **Chaos Visual Debugger (CVD)** is a visual debugging tool for Chaos Physics simulations. It provides a graphical view of the Chaos Physics scene and tools to visualize data and analyze simulation results.

CVD is included in Unreal Engine as an editor tool and runtime system that records the state of physics simulations during gameplay. It can replay those simulations and inspect data for any given frame or sub-step.

- [Chaos Visual Debugger Documentation](https://dev.epicgames.com/documentation/en-us/unreal-engine/chaos-visual-debugger-in-unreal-engine)

### Rigid Body Dynamics

Chaos Physics provides many features for **rigid-body dynamics** including collision responses, physics constraints, damping, and friction. It also provides asynchronous physics simulation and networked physics.

- [Collision](https://dev.epicgames.com/documentation/en-us/unreal-engine/collision-in-unreal-engine)
- [Traces with Raycasts](https://dev.epicgames.com/documentation/en-us/unreal-engine/traces-with-raycasts-in-unreal-engine)
- [Physics Constraints](https://dev.epicgames.com/documentation/en-us/unreal-engine/physics-constraints-in-unreal-engine)
- [Physics Components](https://dev.epicgames.com/documentation/en-us/unreal-engine/physics-components-in-unreal-engine)

### Rigid Body Animation Nodes and Physical Animation

Chaos Physics provides **Rigid Body simulation** and **physical animations** for characters at runtime. The system uses the **Physical Asset Editor** to set up rigid bodies attached to the character's Skeletal Mesh that can be simulated along the character's animations.

- [Physics Asset Editor](https://dev.epicgames.com/documentation/en-us/unreal-engine/physics-asset-editor-in-unreal-engine)

### Cloth Physics and Machine Learning Cloth Simulation

**Chaos Cloth** provides accurate and performant cloth simulation for games and real-time experiences. The system comes with physical reactions such as wind, and a powerful **Animation Drive system** which deforms a cloth Mesh to match its parent's animated Skeletal Mesh.

Chaos Cloth parameters are exposed to Blueprints for control over the cloth simulation at runtime. Users can modify simulation parameters based on gameplay conditions.

Chaos Cloth also provides **machine learning cloth simulation** using a trained dataset for higher fidelity results in real-time.

The **Chaos Cloth Panel node editor** (introduced in UE 5.3) provides a non-destructive way of authoring Chaos Cloth in-engine using a **cloth asset** that can be used with any Skeletal or Static Mesh via the Chaos Cloth component.

- [Cloth Simulation Documentation](https://dev.epicgames.com/documentation/en-us/unreal-engine/cloth-simulation-in-unreal-engine)

### Ragdoll Physics

Chaos Physics comes with **ragdoll physics**, where rigid bodies connected to a Skeletal Mesh are simulated in real-time. Commonly used to animate a humanoid character falling.

### Chaos Vehicles

**Chaos Vehicles** is Unreal Engine's lightweight system for vehicle physics simulations. It provides flexibility by simulating any number of wheels per vehicle, configurable forward and reverse gears, aerofoil surfaces for downforce or uplift, and thrust forces at specific locations.

The system supports the **Asynchronous Physics** mode in Unreal Engine 5 for improved determinism.

- [Vehicles Documentation](https://dev.epicgames.com/documentation/en-us/unreal-engine/vehicles-in-unreal-engine)

### Physics Fields

The **Physics Field System** provides a tool to affect **Chaos Physics** simulations at runtime in a specified region of space. Fields can exert force on rigid bodies, break Geometry Collection Clusters, and anchor or disable fractured rigid bodies.

The Physics Field System can also communicate with Niagara and Materials. Fields are set up by creating a **Field System Component** Blueprint configured as a **World Field** (for Materials and Niagara) or **Chaos Field** (for Geometry Collections and physics objects).

- [Physics Fields Documentation](https://dev.epicgames.com/documentation/en-us/unreal-engine/physics-fields-in-unreal-engine)

### Fluid Simulation

Unreal Engine 5 includes tools for simulating **2D and 3D fluid effects in real time**. These systems use physically-based simulation methods for fire, smoke, clouds, rivers, splashes, and waves.

- [Fluid Simulation Documentation](https://dev.epicgames.com/documentation/en-us/unreal-engine/fluid-simulation-in-unreal-engine)

### Hair Physics

The **Hair rendering and simulation system** uses a strand-based workflow to render each individual strand of hair with physically accurate motion.

- [Hair Physics Documentation](https://dev.epicgames.com/documentation/en-us/unreal-engine/hair-physics-in-unreal-engine)

### Chaos Flesh

The **Chaos Flesh** system provides high-quality, real-time simulation of deformable (soft) bodies. Unlike rigid body simulation, the shape of soft bodies can change during simulation. Primarily focused on character muscle deformation during skeletal animation.

- [Chaos Flesh Documentation](https://dev.epicgames.com/documentation/en-us/unreal-engine/chaos-flesh)

### Dataflow Graph System

The **Dataflow Graph** system is a node-based procedural asset generation environment inside the Unreal Engine editor. It improves iteration times for physics asset types such as Chaos Cloth, Chaos Flesh, and Geometry Collection fracturing. The system is extensible via C++.

- [Dataflow Graph Documentation](https://dev.epicgames.com/documentation/en-us/unreal-engine/dataflow-graph)

### Other Documentation

- [Physical Materials](https://dev.epicgames.com/documentation/en-us/unreal-engine/physical-materials-in-unreal-engine)
- [Physics Sub-Stepping](https://dev.epicgames.com/documentation/en-us/unreal-engine/physics-sub-stepping-in-unreal-engine)
- [Walkable Slope](https://dev.epicgames.com/documentation/en-us/unreal-engine/walkable-slope-in-unreal-engine)
