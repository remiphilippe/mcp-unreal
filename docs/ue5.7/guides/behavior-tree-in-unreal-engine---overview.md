<!-- Source: https://dev.epicgames.com/documentation/en-us/unreal-engine/behavior-tree-in-unreal-engine---overview -->

## Behavior Tree in Unreal Engine - Overview

Creating artificial intelligence (AI) in Unreal Engine 5 (Unreal Engine) for characters can be done in many different ways. You can use Blueprint Visual Scripting to instruct a character to "do something" such as play an animation, move to a specific location, react when being hit by something, and more. When you want AI characters to do their thinking and make their own decisions, this is where **Behavior Trees** can help! An example of a Behavior Tree is illustrated below where an AI character switches between patrolling and chasing the player.

## Basics of Behavior Trees

Behavior Trees are created in a visual way similar to Blueprint by adding and connecting a series of nodes that have some functionality attached to them to a Behavior Tree Graph. While a Behavior Tree executes logic, a separate asset called a **Blackboard** is used to store information (called **Blackboard Keys**) the Behavior Tree needs to know about to make informed decisions. A typical workflow would be to create a Blackboard, add some Blackboard Keys, then create a Behavior Tree that uses the Blackboard asset (pictured below, a Blackboard is assigned to a Behavior Tree).

Behavior Trees in Unreal Engine execute their logic from left-to-right, and top-to-bottom. The numerical order of operation can be viewed in the upper-right corner of nodes placed in the graph.

The blue node in the image above is referred to as a [Decorator](https://dev.epicgames.com/documentation/en-us/unreal-engine/unreal-engine-behavior-tree-node-reference-decorators) (or a _Conditional_ in other Behavior Tree systems). It is attached to a [Composite](https://dev.epicgames.com/documentation/en-us/unreal-engine/unreal-engine-behavior-tree-node-reference-composites) node and used to validate whether the Blackboard Key is true or not. This determines whether the rest of the branch can be executed. The purple nodes are [Task](https://dev.epicgames.com/documentation/en-us/unreal-engine/unreal-engine-behavior-tree-node-reference-tasks) nodes and are the actions the AI can perform.

Please see [Behavior Tree User Guide](https://dev.epicgames.com/documentation/en-us/unreal-engine/behavior-tree-in-unreal-engine---user-guide) for more information on creating and editing Behavior Trees.

Once you have created your Behavior Tree and logic, you will need to run the Behavior Tree during gameplay. Usually, you will have a **Pawn** (or "body" for the AI that is either a character or some other entity) and that Pawn will have an associated **AI Controller** which will be used to take control of and direct the Pawn in performing actions (one of which is to run the Behavior Tree).

For your Pawn to navigate around the environment, you will need to add a **NavMeshBoundsVolume** to the Level.

Please see the [Behavior Tree Quick Start](https://dev.epicgames.com/documentation/en-us/unreal-engine/behavior-tree-in-unreal-engine---quick-start-guide) for a more detailed step-by-step guide in setting up AI using a Behavior Tree.

## Differences in Unreal Engine Behavior Trees

This section outlines some of the differences in Unreal Engine Behavior Trees when compared to traditional Behavior Tree systems.

### Behavior Trees are Event-Driven

One of the ways Unreal Engine Behavior Trees differ from other Behavior Tree systems is that Unreal Engine Behavior Trees are event-driven to avoid doing unnecessary work every frame. Instead of constantly checking whether any relevant change has occurred, the Behavior Tree passively listens for "events" that can be used to trigger changes in the tree.

Having an event-driven architecture grants improvements to both performance and debugging. However, to take the most advantage of these improvements, you will need to understand the other differences to Unreal Engine Behavior Trees and structure your Behavior Trees appropriately. Since the code does not have to iterate through the entire tree every tick, performance is much better. Conceptually, instead of constantly asking "Are we there yet?", we can just rest until we are prodded and told, "We are there!"

When stepping forward and backward through a Behavior Tree's execution history to visually debug the behavior, it is ideal to make the history show relevant changes and not show irrelevant ones. In the event-driven implementation in Unreal Engine, it is not necessary to filter out irrelevant steps that iterated over the tree and chose the same behavior as before, because that additional iteration never had to happen in the first place. Instead, only changes to the location of execution in the tree or Blackboard values matter, and it is those differences that are shown.

### Conditionals are not Leaf Nodes

In the standard model for Behavior Trees, conditionals are Task leaf nodes, which simply do not do anything other than succeed or fail. Although nothing prevents you from making traditional conditional tasks, it is highly recommended that you use [Decorators](https://dev.epicgames.com/documentation/en-us/unreal-engine/unreal-engine-behavior-tree-node-reference-decorators) for conditionals instead.

Making conditionals Decorators rather than Tasks has a couple of significant advantages:

1. Conditional Decorators make the Behavior Tree UI more intuitive and easier to read.
2. Since all leaves are action Tasks, it is easier to see what actual actions are being ordered by the tree.

Since conditionals are at the root of the sub-tree they are controlling, you can immediately see what part of the tree is "closed off" if the conditionals are not met. Also, since all leaves are action Tasks, it is easier to see what the actions are. In a traditional model, conditionals would be among the leaves, so you would have to spend more time trying to figure out which leaves are conditionals and which leaves are actions.

Another advantage of conditional Decorators is that it is easy to make those Decorators act as observers (waiting for events) at critical nodes in the tree. This feature is critical to gaining full advantage from the event-driven nature of the trees.

### Concurrent Behaviors

Standard Behavior trees often use a parallel composite node to handle concurrent behaviors and the parallel node begins execution on all of its children simultaneously. Special rules determine how to act if one or more of those child trees finish (depending on the desired behavior).

Parallel nodes are not necessarily multi-threading (executing tasks at the same time). They are just a way to conceptually perform several tasks at once. Often they still run on the same thread and begin in some sequence. That sequence should be irrelevant since they will all happen in the same frame, but it is still sometimes important.

Instead of complex parallel nodes, Unreal Engine Behavior Trees use **Simple Parallel** nodes, a special node type called [Services](https://dev.epicgames.com/documentation/en-us/unreal-engine/unreal-engine-behavior-tree-node-reference-services), and the property **Observer Aborts** on [Decorators](https://dev.epicgames.com/documentation/en-us/unreal-engine/unreal-engine-behavior-tree-node-reference-decorators) to accomplish the same sorts of behaviors.

#### Simple Parallel Nodes

Simple Parallel nodes have only two children: one which must be a single Task node (with optional Decorators), and the other of which can be a complete sub-tree. Think of the Simple Parallel node as "While doing A, do B as well." For example, "While attacking the enemy, move toward the enemy." A is a primary task, and B is a secondary or filler task while waiting for A to complete.

While there are some options as to how to handle the secondary task (Task B), the node is relatively simple in concept compared to traditional parallel nodes. Nonetheless, it supports much of the most common usage of parallel nodes. Simple Parallel nodes allow easy usage of our event-driven optimizations while full parallel nodes would be much more complex to optimize.

#### Services

[Services](https://dev.epicgames.com/documentation/en-us/unreal-engine/unreal-engine-behavior-tree-node-reference-services) are special nodes associated with any Composite node (Selector, Sequence, or Simple Parallel), which can register for callbacks every specified amount of seconds and perform updates of various sorts that need to occur periodically.

For example, a service can be used to determine which enemy is the best choice for the AI Pawn to pursue while the Pawn continues to act normally in its Behavior Tree toward its current enemy.

Services are active as long as execution remains in the sub-tree of the Composite node where the service has been added.

#### Observer Aborts

One common usage case for standard parallel nodes is to constantly check conditions so that a task can abort if the conditions it requires becomes false.

For example, if you have a cat that performs a sequence, such as "Hiss" and "Pounce", you may want to give up immediately if the mouse escapes into its mouse hole. With parallel nodes, you would have a child that checks if the mouse can be pounced on, and then another child that the sequence would perform. Since Unreal Engine Behavior Trees are event-driven, we instead handle this by having our conditional Decorators observe their values and abort when necessary. In this example, you would have a "Mouse Can Be Pounced On?" Decorator on the Sequence, with "Observer Aborts" set to "Self".

#### Advantages to Unreal Engine's Approach to Concurrent Behaviors

There are three main advantages to the way Unreal Engine handles concurrent behaviors:

1. **Clarity** - Using Services and Simple Parallel nodes creates simple trees that are easier to read and understand.
2. **Ease of Debugging** - Clearer graphs are easier to debug. Also, having fewer simultaneous execution paths makes it easier to see what is being executed.
3. **Easier Optimization** - Event-driven graphs are easier to optimize if they do not have a lot of sub-trees simultaneously executing.
