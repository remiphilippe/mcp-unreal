<!-- Source: https://dev.epicgames.com/documentation/en-us/unreal-engine/events-in-unreal-engine -->

**Events** are nodes that are called from gameplay code to begin execution of an individual network within the **EventGraph**. They enable Blueprints to perform a series of actions in response to certain events that occur within the game, such as when the game starts, when a level resets, or when a player takes damage.

Events can be accessed within Blueprints to implement new functionality or to override or augment the default functionality. Any number of **Events** can be used within a single **EventGraph**; though only one of each type may be used.

An event can only execute a single object. If you want to trigger multiple actions from one event, you will need to string them together linearly.

## Event Level Reset

This Blueprint Event node is only available in the Level Blueprint. The **Level Reset** event sends out an execution signal when the level restarts. This is useful when you need something to be triggered once a level has reloaded, such as in a gaming situation when the player has died but the level does not need to reload.

## Event Actor Begin Overlap

This event will execute when a number of conditions are met at the same time:

- Collision response between the actors must allow Overlaps.
- Both Actors that are to execute the event have **Generate Overlap Events** set to true.
- And finally, both Actors' collision starts overlapping; moving together or one is created overlapping the other.

For more information on collision, see: [Collision Responses](https://dev.epicgames.com/documentation/en-us/unreal-engine/collision-in-unreal-engine).

| Output Pin | Description |
| --- | --- |
| **Other Actor** | Actor - This is the Actor that is overlapping this Blueprint. |

## Event Actor End Overlap

This event will execute when a number of conditions are met at the same time:

- Collision response between the actors must allow Overlaps.
- Both Actors that are to execute the event have **Generate Overlap Events** set to true.
- And finally, both Actors' collision stop overlapping; moving apart or if one is destroyed.

| Output Pin | Description |
| --- | --- |
| **Other Actor** | Actor - This is the Actor that is overlapping this Blueprint. |

## Event Hit

This event will execute as long as the collision settings on one of the Actors involved have **Simulation Generates Hit Events** set to true.

If you are creating movement using Sweeps, you will get this event even if you don't have the flag selected. This occurs as long as the Sweep stops you from moving past the blocking object.

| Output Pin | Description |
| --- | --- |
| **My Comp** | PrimitiveComponent - The Component on the executing Actor that was hit. |
| **Other** | Actor - The other Actor involved in the collision. |
| **Other Comp** | PrimitiveComponent - The component on the other Actor involved in the collision that was hit. |
| **Self Moved** | Boolean - Used when receiving a hit from another object's movement. |
| **Hit Location** | Vector - The location of contact between the two colliding Actors. |
| **Hit Normal** | Vector - The direction of the collision. |
| **Normal Impulse** | Vector - The force that the Actors collided with. |
| **Hit** | Struct HitResult - All the data collected in a Hit. |

## Event Any Damage

This Blueprint Event node executes only on the server. For single player games, the local client is considered the server. This event is passed along when general damage is to be dealt. Like drowning or environmental damage, not specifically point damage or radial damage.

| Output Pin | Description |
| --- | --- |
| **Damage** | Float - The amount of damage being passed into the Actor. |
| **Damage Type** | Object DamageType - Contains additional data on the Damage being dealt. |
| **Instigated By** | Controller - The Controller of the Object that is responsible for the damage. |
| **Damage Causer** | Actor - The Actor that caused the damage (e.g., a bullet or explosion). |

## Event Point Damage

This Blueprint Event node executes only on the server. **Point Damage** is meant to represent damage dealt by projectiles, hit scan weapons, or even melee weaponry.

| Output Pin | Description |
| --- | --- |
| **Damage** | Float - The amount of damage being passed into the Actor. |
| **Damage Type** | Object DamageType - Contains additional data on the Damage being dealt. |
| **Hit Location** | Vector - The location of where the damage is being applied. |
| **Hit Normal** | Vector - The direction of the collision. |
| **Hit Component** | PrimitiveComponent - The Component on the executing Actor that was hit. |
| **Bone Name** | Name - The name of the bone that was hit. |
| **Shot from Direction** | Vector - The direction the damage came from. |
| **Instigated By** | Actor - The Actor that is responsible for the damage. |
| **Damage Causer** | Actor - The Actor that caused the damage (e.g., a bullet or explosion). |

## Event Radial Damage

This Blueprint Event node executes only on the server. The **Radial Damage** Event is called whenever the parent Actor for this sequence receives radial damage. This is useful for handling events based on explosion damage, or damage caused indirectly.

| Output Pin | Description |
| --- | --- |
| **Damage Received** | Float - The amount of damage received from the event. |
| **Damage Type** | Object DamageType - Contains additional data on the Damage being dealt. |
| **Origin** | Vector - The location in 3D space where the damage originated. |
| **Hit Info** | Struct HitResult - All the data collected in a Hit. |
| **Instigated By** | Controller - The Controller (AI or Player) that instigated the damage. |
| **Damage Causer** | Actor - The Actor that caused the damage. |

## Event Actor Begin Cursor Over

When using the mouse interface, when the mouse cursor is moved over an Actor, this event will execute.

## Event Actor End Cursor Over

When using the mouse interface, when the mouse cursor is moved off an Actor, this event will execute.

## Event Begin Play

This event is triggered for all Actors when the game is started, any Actors spawned after the game is started will have this called immediately.

## Event End Play

This event is executed when the Actor ceases to be in the World.

| Output Pin | Description |
| --- | --- |
| **End Play Reason** | enum EEndPlayReason - An enum indicating the reason for why Event End Play is being called. |

## Event Destroyed

This event is executed when the Actor is Destroyed.

The Destroyed Event will be deprecated in a future release! The functionality of the Destroyed function has been incorporated into the EndPlay function.

## Event Tick

This is a simple event that is called on every frame of gameplay.

| Output Pin | Description |
| --- | --- |
| **Delta Seconds** | Float - The amount of time between frames. |

## Event Receive Draw HUD

This Event is only available to Blueprint Classes that inherit from the HUD class. This is a specialized event that enables Blueprints to draw to the HUD. The HUD draw nodes require this event to be the one that creates them.

| Output Pin | Description |
| --- | --- |
| **Size X** | Int - The width of the render window in pixels. |
| **Size Y** | Int - The height of the render window in pixels. |

## Custom Event

The Custom Event node is a specialized node with its own workflow.

To learn more:

- [Custom Events](https://dev.epicgames.com/documentation/en-us/unreal-engine/custom-events-in-unreal-engine): Create your own events that can be called at any point in your Blueprint's sequence.
- [Calling Events through Sequencer](https://dev.epicgames.com/documentation/en-us/unreal-engine/fire-blueprint-events-during-cinematics-in-unreal-engine): Create events to be triggered at specific times during the playback of a cinematic sequence.
