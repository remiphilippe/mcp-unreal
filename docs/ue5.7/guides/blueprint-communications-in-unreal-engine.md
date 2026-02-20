<!-- Source: https://dev.epicgames.com/documentation/en-us/unreal-engine/blueprint-communications-in-unreal-engine -->

On this page you will learn step-by-step how to set up and use different methods of **Blueprint Communication**.

### Direct Blueprint Communications

Below we have two **Blueprints** in our level that we want to have communicate with one another. Suppose we wanted the Cube Blueprint to communicate to the Sparks Blueprint to turn itself off when a player character enters the cube. This can easily be achieved with **Direct Blueprint Communications**.

- The cube above is a Blueprint created using the **Shape\_Cube** mesh and its collision has been set to **OverlapOnlyPawn** so it acts as a trigger. Also enable **Generate Overlap Events**
- The sparks above are the **Blueprint\_Effect\_Sparks** asset (included with starter content).

Using Direct Blueprint Communication, we would do the following:

1. In the **Shape\_Cube** Blueprint, under **My Blueprint** click the button on the Variables category.

2. In the **Details** panel, under **Variable Type** select the type of Blueprint you wish to access. Hover over the **Blueprint\_Effects\_Sparks** Variable Type and select **Reference** from the list. Here we are stating that we want to access a **Blueprint\_Effect\_Sparks** Blueprint.

3. In the **Details** panel, update the sections outlined below.
   1. **Variable Name** - give the variable a descriptive name such as **TargetBlueprint**.
   2. **Variable Type** - this should be the type of Blueprint that you want to access.
   3. **Editable** - make sure this is checked to expose the variable and make it public which will allow you to access it in the Level Editor.
   4. **Tool Tip** - add a short description of what the variable does or what it will reference.

4. With the **Shape\_Cube** Blueprint selected in the level, in the Level Editor under **Details** you should see the exposed variable created in the previous step.
   1. Click the **None** drop-down box to assign a Target Blueprint.
   2. All instances of the Blueprint placed in the level are displayed here, allowing you to specify which instance is the **Target Blueprint**.

Here we are stating which **Blueprint\_Effect\_Sparks** Blueprint Actor we want to affect that is placed in our level; this is considered the Instance Actor. If we had multiple sparks in our level and we only wanted to turn off one of them, we could choose which instance of the Blueprint we want to affect here by setting it as the **Target Blueprint**.

5. Instead of using the drop-down, you can click the eyedropper icon, then click on the object you want to reference that is placed in your level. You can only set the **Target Blueprint** to target the specified Variable Type (in our case **Blueprint\_Effect\_Sparks**).

6. In the **Shape\_Cube** Blueprint, while holding down the **Ctrl** key, drag the variable into the graph.
   1. This creates a Getter node which allows you to access Events, Variables, Functions, etc. from the **Target Blueprint**.
   2. Drag off the out pin to view the context menu.

Here we are searching for the spark effect and spark audio components from our **Target Blueprint** as we want to access those components.

7. The sample script below states that when the player enters the cube, deactivate the spark effect and spark audio.

#### Direct Blueprint Communication for Spawned Actors

There may be an instance when you want to communicate between two Blueprints, however one (or more) of your Blueprints is not placed in the level yet (for example, a magic effect that is spawned when the player presses a button). In this case neither the player character nor the magic effect has spawned in the level, so setting a **Target Blueprint** and instance like above cannot be performed.

When using the **Spawn Actor from Class** node, you can drag off its **Return Value** and assign it as a variable.

In the example below, inside our **MyCharacter** Blueprint, we have stated that when **F** is pressed we want to spawn an instance of the **Blueprint\_Effect\_Smoke** Blueprint at the player's location and assign it to a variable called **Target Blueprint**.

We can then access the **Blueprint\_Effect\_Smoke** Blueprint and get the Smoke Effect and Smoke Audio components from our **Target Blueprint** and **Deactivate** them when **F** is pressed a second time (which is what the **Flip/Flop** node does). Therefore we are accessing one Blueprint from inside another Blueprint through **Direct Blueprint Communication**.

Refer to the [Blueprint Communication Usage](https://dev.epicgames.com/documentation/en-us/unreal-engine/blueprint-communication-usage-in-unreal-engine) documentation for more information.

### Blueprint Casting

For this example, we have a fire effect Blueprint in our level (which is an Actor) and we want it to communicate with the playable Character Blueprint the player is using. When the player enters the fire, we want to send a signal to the Character Blueprint that the player has entered the fire and that they should now take damage. By using the Return Value of an OverlapEvent, we can Cast To our Character Blueprint and access the Events, Functions, or Variables within it.

- The fire effect above is the **Blueprint\_Effect\_Fire** asset (included with starter content).
- A sphere component named **Trigger** was added to the Blueprint and was set to **OverlapOnlyPawn** for its collision.

Using Blueprint Casting, we would do the following:

1. The Character Blueprint assigned to the **Default Pawn Class** (the playable character) is our Target Blueprint we want to access. You can view the **Default Pawn Class** from the **Edit** menu under **Project Settings** in the **Maps & Modes** section.

2. Now that we know our target is the **MyCharacter** Blueprint, inside it we create a **Bool** variable that states if the player **Is on Fire**. Above the **Event Tick** feeds a **Branch** where if **True**, we print **Apply Damage** to the screen (off True is where you would have your apply damage script).

3. Inside the **Blueprint\_Effect\_Fire** Blueprint, we add two events for the **Trigger**: **OnComponentBeginOverlap** and **OnComponentEndOverlap**.

4. With the Events added, we drag off the **Other Actor** pin and enter **Cast To My** in the search field. Here we check/assign the Actor (MyCharacter Blueprint) we want to trigger the event and **Cast To** it so that we may access it within the fire Blueprint.

5. Select the **Cast To MyCharacter** option.

6. With the node added, we can drag off the **As My Character C** pin and access the Events, Variables, Functions, etc. within it (in this case **Set Is on Fire**).

7. Both Events in the **Blueprint\_Effect\_Fire** Blueprint would then look like this. When overlapping the fire, we are setting the **IsOnFire** variable in the **MyCharacter** Blueprint to **True** and setting it to **False** when not overlapping it. Inside the **MyCharacter** Blueprint, when **IsOnFire** is set to **True** via the fire Blueprint, we print to the screen **Apply Damage** (or if you have created a Health/Damage system, you could apply damage and reduce player's health here).

#### Other Types of Casting

There are some special functions that can be used to **Cast To** different classifications of **Blueprints**.

| Blueprint | Description |
| --- | --- |
| **Character** (1) | Here the **Get Player Character** node is used and we are casting to a Character Blueprint called **MyCharacter**. |
| **PlayerController** (2) | Here the **Get Player Controller** node is used and we are casting to a Player Controller Blueprint called **MyController**. |
| **Game Mode** (3) | Here the **Get Game Mode** node is used and we are casting to a Game Mode Blueprint called **MyGame**. |
| **Pawn** (4) | Here the **Get Controlled Pawn** and **Get Player Controller** nodes are used to cast to a Pawn Blueprint called **MyPawn**. |
| **HUD** (5) | Here the **Get HUD** and **Get Player Controller** nodes are used to cast to a HUD Blueprint called **MyHUD**. |

In each of the examples above, dragging off the **As My X** (where X is the type of Blueprint) node will allow you to access the Events, Variables, Functions, etc. from their respective Blueprints.

Also of note, the **Player Index** value in the **Get Player Character** and **Get Player Controller** nodes can be used to specify different players in a multiplayer scenario. For a single player scenario, leaving these as 0 is fine.

#### Target Blueprint Casting

There are also instances where you can use a variable to **Cast To** a specific type of Blueprint in order to access it.

In the image above for number 1, a **Save Game Object** is created and assigned to a **SaveGameObject** variable. That variable is then used to **Cast To** a Save Game Blueprint called **MySaveGame** - which could be used to pass off or retrieve save game information such as a high score, best lap time, etc.

In the image above for number 2, a **Widget Blueprint** is created and assigned to a **UserWidget** variable. That variable is then used to **Cast To** a Widget Blueprint called **MyWidgetBlueprint** - which could be used to update or retrieve information from the Widget Blueprint (which could be a HUD or other UI element you want to access).

### Event Dispatchers

Below we have a bush **Actor Blueprint** in our level that we want to receive communication from the **Character Blueprint** when the player presses a button to set the bush on fire, then destroy the fire and the bush after a few seconds. We are wanting to communicate from the **Character Blueprint** to the **Level Blueprint** which can be done through using an **Event Dispatcher**.

- The bush above is the **SM\_Bush** asset (included with starter content).

Using an **Event Dispatcher** we would do the following:

1. Inside the **MyCharacter** Blueprint, click the **Event Dispatcher** icon or the >> arrows if the icon is hidden and name the Event Dispatcher **StartFire**.

2. **Right-click** in the graph and add the **F** key event, then off the **Pressed** search for and add the **Call StartFire** Event Dispatcher.

3. **Compile** and **Save** then close the **MyBlueprint**.

4. Click on the bush in the level to select it, then open the **Level Blueprint**.

5. **Right-click** in the graph and add the reference to the bush from the level.

6. **Right-click** and add an **Event Begin Play** node and a **Get Player Character** node, then off the **Get Player Character**, **Cast To MyCharacter**.

7. Off the **As My Character C**, add the **Assign Start Fire** Event Dispatcher (a new Binded Event will be created).

8. Off the **StartFire\_Event**, add a **SpawnActorFromClass** node (class set to **Blueprint\_Effect\_Fire**) and for the **Transform**, get the **SM\_Bush** Transform.

9. Off the **Return Value** of the Spawn Actor node, add the **Assign On Destroyed** node.

10. Off the **OnDestroyed\_Event**, add a **Destroy Actor** node with target being **SM\_Bush**.

11. Off the **Bind Event to OnDestroyed**, add a **Delay** (3 seconds) then add a **Destroy Actor** with the target the **Return Value** from the Spawn Actor node.

If you compile and save then play in the editor, you should see that when you press the **F** key a fire effect is spawned inside the bush. After 3 seconds both the fire and bush should then be removed from the level.

This is of course a simple example and you would probably want more checking to occur (if the player is near the bush allow them to set it on fire) as well only allowing the player to start the fire once, but it illustrates how you can execute events inside the **Level Blueprint** by way of a **Character Blueprint** using an **Event Dispatcher**. The example above also shows how you can assign an **Event Dispatcher** to a Spawned Actor and execute Events when something occurs to that Actor (in this case, when it is destroyed).

Refer to the **[Event Dispatchers](https://dev.epicgames.com/documentation/en-us/unreal-engine/event-dispatchers-in-unreal-engine)** documentation for more information.

### Blueprint Interfaces

Below we have four **Blueprints** in our level: a cube which acts as a trigger, a fire effect, a spark effect and a hanging light. We want the fire, light and sparks to each do something differently when the player enters the cube. We also want to increase the movement speed of the character each time they enter the cube.

- The cube above is a **Blueprint** created using the **Shape\_Cube** mesh and its collision has been set to **OverlapOnlyPawn** so it acts as a trigger.
- The sparks above are the **Blueprint\_Effect\_Sparks** asset (included with starter content).
- The fire above is the **Blueprint\_Effect\_Fire** asset (included with starter content).
- The light above is the **Blueprint\_CeilingLight** asset (included with starter content).

Using a **Blueprint Interface** would allow us to communicate with the three different Blueprints as well as the player character Blueprint.

To communicate with each of them, we would do the following:

1. In the **Content Browser**, in an empty space, **Right-click** and select **Blueprints** then **Blueprint Interface**.

2. Name the Interface, **CubeInterface** (or some other name) then **Double-click** on it to open it up and click the **Add Function** button.

3. Name the new function **MagicCube** or any name you want, then **Compile**, **Save**, and close the Interface.

4. Open the Cube Blueprint, then **Right-click** on the **StaticMesh** and add a **OnComponentBeginOverlap** Event.

5. Create a new **Actor** variable called **Targets** and click the box next to Variable Type to make it an **Array**, then check the **Editable** box. This will store the Actors that are affected by the **Blueprint Interface**.

6. **Right-click** in the graph and under **Interface Messages**, click the **MagicCube** (or whatever you called it) function.

7. Set up your graph: drag in the **Targets** variable then drag off it to get the **Add** node. Plug **Targets** into the **MagicCube** node and plug a **Get Player Character** node into the **Add** node.

8. Select the cube in the level, then in the **Details** panel, click the + sign under targets and add the fire, light and sparks from the level.

9. Open the **Blueprint\_Effect\_Fire** Blueprint, then click the **Blueprint Props** button from the toolbar.

10. In the **Details** panel under **Interfaces**, click the **Add** button then select your Interface (**CubeInterface\_C** in our example).

11. **Right-click** in the graph and under **Add Event**, select the **Event Magic Cube** Event.

12. Anything following the **Event Magic Cube** will now be executed when the player enters the cube. Here we are increasing the size of the fire when the player enters the cube, then resetting it when they enter it a second time.

13. Open the **Blueprint\_CeilingLight** Blueprint, click **Blueprint Props** then add the Interface from the **Details** panel as before.

14. **Right-click** in the graph and add the **Event Magic Cube** Event so that anything following the Event is executed when the player enters the cube. Here we are turning the light off by setting its Brightness to 0, then turning it on by setting the Brightness to 1500.

15. Repeat the process of adding **Blueprint Props** to the **Blueprint\_Effect\_Sparks** Blueprint, then add the **CubeInterface\_C**. Here we are moving the spark effect up when entering the cube, then down when entering the cube a second time.

16. Repeat the process of adding **Blueprint Props** to the **MyCharacter** Blueprint, then add the **CubeInterface\_C**. Here we increase the character's movement speed each time they enter the cube by 100.

As you can see from the examples above, by using a **Blueprint Interface** you can communicate with several different types of **Blueprints** at once where each can perform a different function all stemming from the same singular source (in this case a trigger).

This example is good for having an Event execute functionality in multiple Blueprints, however it is not the only manner in which **Blueprint Interfaces** can be used. The next section discusses how Variables can be passed between Blueprints using **Blueprint Interfaces**.

#### Passing Variables through Blueprint Interfaces

Below we have the **Blueprint\_Effect\_Fire** Blueprint which will represent the player character's life force. This Blueprint will check what the player's health is and once it is 0, will dissolve and disappear.

- The fire above is the **Blueprint\_Effect\_Fire** asset (included with starter content).

Using a **Blueprint Interface** and passing two variables through (the player's health and whether the player is dead or not) we can tell the fire when to disappear.

Here is how we would set up passing through those variables:

1. In the **Content Browser**, in an empty space, **Right-click** and select **Blueprints** then **Blueprint Interface**.

2. Name the Interface, **BP\_Interface** (or some other name) then **Double-click** on it to open it up and click the **Add Function** button.

3. Name the new function **GetHealth**, then in the **Details** panel, add two **Outputs** by clicking the **New** button.

4. Make one of the new outputs a **Bool** called **playerIsDead** and the other a **Float** called **playerHealth**, then **Compile** and **Save** and close the Interface.

5. Open the **MyCharacter** Blueprint, then click the **Blueprint Props** button from the toolbar.

6. In the **Details** panel under **Interfaces**, click the **Add** button then select your Interface (**BP\_Interface\_C** in our example).

7. Create a **Bool** variable called **OutOfHealth** and a **Float** variable called **PlayerHealthValue**, **Compile** then set the **PlayerHealthValue** to **100**.

8. Under the **Interfaces** section of **MyBlueprint**, **Double-click** on the **GetHealth** function to open it up.

9. In the graph, drag in the **OutOfHealth** and **PlayerHealthValues** and plug them in to the **ReturnNode**. This will pass the values that are stored in the **MyCharacter** Blueprint to the Interface.

10. Return to the **EventGraph** of the **MyCharacter** Blueprint and re-create the setup: when the player's health is greater than 0 and the player presses **F**, subtract 25 from the current health value and call an **Event Dispatcher** called **TakeDamage**. When health is less than or equal to 0, set a Bool variable called OutOfHealth to _true_ and call the **Event Dispatcher** called **TakeDamage**. We use an **Event Dispatcher** here to signal another Blueprint that the player has taken damage, rather than the other Blueprint checking the player's health each tick using an Event Tick.

11. Open the Blueprint you want to pass the variables to (**Blueprint\_Effect\_Fire**), and click the **Blueprint Props** button, then **Add** the interface via the **Details** panel.

12. In the **EventGraph**, bind an Event to the **TakeDamage** event from the **MyCharacter** Blueprint. Drag off the **Get Player Character** node and **Cast To MyCharacter**, then off the **As My Character C**, **Assign Take Damage** to create a Binded Event.

13. Off the Binded Event **TakeDamage\_Event**, add the **GetHealth** Interface Message. Be sure to implement the **Interface Message** and not the **Call Function**.

14. The **GetHealth** Interface is connected to a series of **Branch** nodes where the first checks if the **PlayerIsDead** (defined in the **MyCharacter** Blueprint) and if so, text is printed to the screen and the fire effect/audio is deactivated. The second **Branch** node checks if the **PlayerHealth** value is 0 and if it is, text is printed to the screen that states "1 More Hit" before the character is "dead". This is by no means a perfect Health/Damage setup, but it illustrates how you can pass two variables through an Interface and how you could then use those variables in another Blueprint. The "PlayerHealth" value in this example could be passed to a HUD and updated to reflect current Health for example.

Refer to the **[Blueprint Interface](https://dev.epicgames.com/documentation/en-us/unreal-engine/blueprint-interface-in-unreal-engine)** documentation for more information.
