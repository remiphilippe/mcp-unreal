<!-- Source: https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/UPCGGraph -->

|  |  |
| --- | --- |
| _Name_ | UPCGGraph |
| _Type_ | class |
| _Header File_ | /Engine/Plugins/PCG/Source/PCG/Public/PCGGraph.h |
| _Include Path_ | #include "PCGGraph.h" |

## Syntax

```cpp

UCLASS (MinimalAPI, BlueprintType, ClassGroup=(Procedural), HideCategories=(Object))

class UPCGGraph : public UPCGGraphInterface

Copy full snippet
```

## Inheritance Hierarchy

- [UObjectBase](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/CoreUObject/UObjectBase) → [UObjectBaseUtility](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/CoreUObject/UObjectBaseUtility) → [UObject](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/CoreUObject/UObject) → [UPCGGraphInterface](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/UPCGGraphInterface) → **UPCGGraph**

## Derived Classes

- [UProceduralVegetationGraph](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/ProceduralVegetation/UProceduralVegetationGraph)

## Constructors

| Name | Remarks | Include Path | [Unreal Specifiers](https://dev.epicgames.com/documentation/unreal-engine/ufunctions-in-unreal-engine#functionspecifiers) |
| --- | --- | --- | --- |
| UPCGGraph<br>(<br>const [FObjectInitializer](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/CoreUObject/FObjectInitializer)& ObjectInitializer<br>) |  | PCGGraph.h |  |

## Structs

| Name | Remarks |
| --- | --- |
| [FGridInfo](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/UPCGGraph/FGridInfo) |  |

## Typedefs

| Name | Type | Remarks | Include Path |
| --- | --- | --- | --- |
| FComputeGraphInstanceKey | TPair< uint32, int32 > |  | PCGGraph.h |
| FComputeGraphInstancePool | [TMap](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/TMap) < FComputeGraphInstanceKey, [TArray](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/TArray) < [TSharedPtr](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/TSharedPtr) < [FComputeGraphInstance](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/ComputeFramework/FComputeGraphInstance) > \> > |  | PCGGraph.h |

## Variables

### Public

| Name | Type | Remarks | Include Path | [Unreal Specifiers](https://dev.epicgames.com/documentation/unreal-engine/unreal-engine-uproperties#propertyspecifiers) |
| --- | --- | --- | --- | --- |
| bIgnoreLandscapeTracking | bool | Marks the graph to be not refreshed automatically when the landscape changes, even if it is used. | PCGGraph.h | - BlueprintReadWrite<br>- EditAnywhere<br>- Category="Settings\|Advanced"<br>- Meta=(PCGNoHash, EditCondition="!IsStandaloneGraph()", EditConditionHides) |
| bIsStandaloneGraph | bool | When enabled, this graph can be executed outside of the world using an editor execution source. | PCGGraph.h | - BlueprintReadWrite<br>- EditAnywhere<br>- Category=AssetInfo<br>- AssetRegistrySearchable<br>- Meta=(PCGNoHash, EditCondition="CanToggleStandaloneGraph()", EditConditionHides) |
| bLandscapeUsesMetadata | bool |  | PCGGraph.h | - EditAnywhere<br>- Category="Settings\|Advanced"<br>- Meta=(EditCondition="!IsStandaloneGraph()", EditConditionHides) |
| Category | [FText](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/FText) |  | PCGGraph.h | - BlueprintReadWrite<br>- EditAnywhere<br>- Category=AssetInfo<br>- AssetRegistrySearchable<br>- Meta=(PCGNoHash) |
| Description | [FText](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/FText) |  | PCGGraph.h | - BlueprintReadWrite<br>- EditAnywhere<br>- Category=AssetInfo<br>- AssetRegistrySearchable<br>- Meta=(PCGNoHash) |
| GenerationRadii | [FPCGRuntimeGenerationRadii](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/FPCGRuntimeGenerationRadii) |  | PCGGraph.h | - BlueprintReadWrite<br>- EditAnywhere<br>- Category="Runtime Generation"<br>- Meta=(EditCondition="!IsStandaloneGraph()", EditConditionHides) |
| GraphCustomization | [FPCGGraphEditorCustomization](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/FPCGGraphEditorCustomization) |  | PCGGraph.h | - EditAnywhere<br>- Category=Customization<br>- Meta=(PCGNoHash, EditCondition="ShowGraphCustomization()", EditConditionHides) |
| ToolData | [FPCGGraphToolData](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/FPCGGraphToolData) | Contains the data relevant for PCG Editor Mode usage. | PCGGraph.h | - BlueprintReadWrite<br>- EditAnywhere<br>- Category=AssetInfo<br>- AssetRegistrySearchable<br>- Meta=(PCGNoHash) |

### Protected

| Name | Type | Remarks | Include Path | [Unreal Specifiers](https://dev.epicgames.com/documentation/unreal-engine/unreal-engine-uproperties#propertyspecifiers) |
| --- | --- | --- | --- | --- |
| AllComputeGraphInstances | FComputeGraphInstancePool | Used to track all valid compute graph instances that are alive for this graph. | PCGGraph.h |  |
| AvailableComputeGraphInstances | FComputeGraphInstancePool | Pool of compute graph instances available for use. | PCGGraph.h |  |
| bDelayedChangeNotification | bool |  | PCGGraph.h |  |
| bIsInspecting | bool |  | PCGGraph.h |  |
| bIsNotifying | bool |  | PCGGraph.h |  |
| bUserPausedNotificationsInGraphEditor | bool |  | PCGGraph.h |  |
| DelayedChangeType | [EPCGChangeType](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/EPCGChangeType) |  | PCGGraph.h |  |
| GraphChangeNotificationsDisableCounter | int32 |  | PCGGraph.h |  |
| InspectedStack | [FPCGStack](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/FPCGStack) |  | PCGGraph.h |  |
| PreviousPropertyBag | [TObjectPtr](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/TObjectPtr) < const [UPropertyBag](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/CoreUObject/UPropertyBag) > | Keep track of the previous PropertyBag, to see if we had a change in the number of properties, or if it is a rename/move. | PCGGraph.h |  |

## Functions

### Public

| Name | Remarks | Include Path | [Unreal Specifiers](https://dev.epicgames.com/documentation/unreal-engine/ufunctions-in-unreal-engine#functionspecifiers) |
| --- | --- | --- | --- |
| [UPCGNode](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/UPCGNode) \\* AddEdge<br>(<br>[UPCGNode](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/UPCGNode)\\* From,<br>const [FName](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/FName)& FromPinLabel,<br>[UPCGNode](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/UPCGNode)\\* To,<br>const [FName](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/FName)& ToPinLabel<br>) | Adds a directed edge in the graph. Returns the "To" node for easy chaining | PCGGraph.h | - BlueprintCallable<br>- Category=Graph |
| bool [AddLabeledEdge](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/UPCGGraph/AddLabeledEdge)<br>(<br>[UPCGNode](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/UPCGNode)\\* From,<br>const [FName](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/FName)& InboundLabel,<br>[UPCGNode](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/UPCGNode)\\* To,<br>const [FName](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/FName)& OutboundLabel<br>) | Creates an edge between two nodes/pins based on the labels. | PCGGraph.h |  |
| void [AddNode](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/UPCGGraph/AddNode)<br>(<br>[UPCGNode](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/UPCGNode)\\* InNode<br>) |  | PCGGraph.h |  |
| [UPCGNode](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/UPCGNode) \\* [AddNode](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/UPCGGraph/AddNode)<br>(<br>[UPCGSettingsInterface](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/UPCGSettingsInterface)\\* InSettings<br>) | Creates a node using the given settings interface. | PCGGraph.h |  |
| [UPCGNode](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/UPCGNode) \\* AddNodeCopy<br>(<br>const [UPCGSettings](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/UPCGSettings)\\* InSettings,<br>[UPCGSettings](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/UPCGSettings)\*& DefaultNodeSettings<br>) | Creates a node and copies the input settings. Returns the created node. | PCGGraph.h | - BlueprintCallable<br>- Category=Graph<br>- Meta=(DeterminesOutputType="InSettings", DynamicOutputParam="OutCopiedSettings") |
| [UPCGNode](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/UPCGNode) \\* AddNodeInstance<br>(<br>[UPCGSettings](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/UPCGSettings)\\* InSettings<br>) | Creates a node containing an instance to the given settings. Returns the created node. | PCGGraph.h | - BlueprintCallable<br>- Category=Graph |
| [UPCGNode](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/UPCGNode) \\* [AddNodeOfType](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/UPCGGraph/AddNodeOfType)<br>(<br>[TSubclassOf](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/CoreUObject/TSubclassOf) < class [UPCGSettings](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/UPCGSettings) \> InSettingsClass,<br>[UPCGSettings](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/UPCGSettings)\*& DefaultNodeSettings<br>) | Creates a default node based on the settings class wanted. Returns the newly created node. | PCGGraph.h | - BlueprintCallable<br>- Category=Graph<br>- Meta=(DeterminesOutputType="InSettingsClass", DynamicOutputParam="DefaultNodeSettings") |
| [UPCGNode](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/UPCGNode) \\* [AddNodeOfType](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/UPCGGraph/AddNodeOfType)<br>(<br>T\*& DefaultNodeSettings<br>) |  | PCGGraph.h |  |
| void AddNodes<br>(<br>[TArray](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/TArray) < [UPCGNode](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/UPCGNode)\\* >& InNodes<br>) |  | PCGGraph.h |  |
| [EPropertyBagAlterationResult](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/CoreUObject/EPropertyBagAlterationResult) [AddUserParameters](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/UPCGGraph/AddUserParameters)<br>(<br>const [TArray](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/TArray) < [FPropertyBagPropertyDesc](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/CoreUObject/FPropertyBagPropertyDesc) >& InDescs,<br>const [UPCGGraph](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/UPCGGraph)\\* InOptionalOriginalGraph<br>) | Add new user parameters using an array of descriptors. | PCGGraph.h |  |
| bool [Contains](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/UPCGGraph/Contains)<br>(<br>const [UPCGGraph](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/UPCGGraph)\\* InGraph<br>) const | Returns true if the current graph contains a subgraph node using statically the specified graph, recursively. | PCGGraph.h |  |
| bool [Contains](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/UPCGGraph/Contains)<br>(<br>[UPCGNode](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/UPCGNode)\\* Node<br>) const | Returns true if the current graph contains directly the specified node. | PCGGraph.h |  |
| bool DebugFlagAppliesToIndividualComponents() |  | PCGGraph.h |  |
| void DisableInspection() |  | PCGGraph.h |  |
| void DisableNotificationsForEditor() |  | PCGGraph.h |  |
| void EnableInspection<br>(<br>const [FPCGStack](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/FPCGStack)& InInspectedStack<br>) |  | PCGGraph.h |  |
| void EnableNotificationsForEditor() |  | PCGGraph.h |  |
| [UPCGNode](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/UPCGNode) \\* FindNodeByTitleName<br>(<br>[FName](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/FName) NodeTitle,<br>bool bRecursive,<br>[TSubclassOf](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/CoreUObject/TSubclassOf) < const [UPCGSettings](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/UPCGSettings) \> OptionalClass<br>) const | Returns the first node that matches the given name in the graph, if any. | PCGGraph.h |  |
| [TArray](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/TArray) < [UPCGNode](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/UPCGNode) \\* \> FindNodesWithSettings<br>(<br>const [TSubclassOf](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/CoreUObject/TSubclassOf) < [UPCGSettingsInterface](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/UPCGSettingsInterface) \> InSettingsClass,<br>bool bRecursive<br>) const |  | PCGGraph.h |  |
| [UPCGNode](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/UPCGNode) \\* FindNodeWithSettings<br>(<br>const [UPCGSettingsInterface](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/UPCGSettingsInterface)\\* InSettings,<br>bool bRecursive<br>) const | Returns the node with the given settings in the graph, if any | PCGGraph.h |  |
| void ForceNotificationForEditor<br>(<br>[EPCGChangeType](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/EPCGChangeType) ChangeType<br>) |  | PCGGraph.h | - BlueprintCallable<br>- Category="Graph\|Advanced" |
| bool ForEachNode<br>(<br>[TFunctionRef](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/TFunctionRef) < bool( [UPCGNode](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/UPCGNode)\*)\> Action<br>) const | Calls the lambda on every node in the graph or until the Action call returns false | PCGGraph.h |  |
| bool ForEachNodeRecursively<br>(<br>[TFunctionRef](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/TFunctionRef) < bool( [UPCGNode](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/UPCGNode)\*)\> Action<br>) const | Calls the lambda on every node (going through subgraphs too) or until the Action call returns false | PCGGraph.h |  |
| const [TArray](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/TArray) < [FPCGGraphCommentNodeData](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/FPCGGraphCommentNodeData) \> & GetCommentNodes() |  | PCGGraph.h |  |
| [TObjectPtr](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/TObjectPtr) < UPCGGraphCompilationData > [GetCookedCompilationData](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/UPCGGraph/GetCookedCompilationData) () |  | PCGGraph.h |  |
| const [TObjectPtr](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/TObjectPtr) < UPCGGraphCompilationData > [GetCookedCompilationData](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/UPCGGraph/GetCookedCompilationData) () |  | PCGGraph.h |  |
| [EPCGHiGenGrid](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/EPCGHiGenGrid) GetDefaultGrid() | Default grid size for generation. For hierarchical generation, nodes outside of grid size graph ranges will generate on this grid. | PCGGraph.h |  |
| uint32 GetDefaultGridSize() |  | PCGGraph.h |  |
| const [TArray](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/TArray) < [TObjectPtr](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/TObjectPtr) < [UObject](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/CoreUObject/UObject) > \> & GetExtraEditorNodes() |  | PCGGraph.h |  |
| double GetGridCleanupRadiusFromGrid<br>(<br>[EPCGHiGenGrid](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/EPCGHiGenGrid) Grid<br>) const | Gets cleanup radius from grid, considering grid exponential. | PCGGraph.h |  |
| uint32 GetGridExponential() | Returns exponential on grid size, which represents a shift in the grid | PCGGraph.h |  |
| double GetGridGenerationRadiusFromGrid<br>(<br>[EPCGHiGenGrid](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/EPCGHiGenGrid) Grid<br>) const | Gets generation radius from grid, considering grid exponential. | PCGGraph.h |  |
| void GetGridSizes<br>(<br>PCGHiGenGrid::FSizeArray& OutGridSizes,<br>bool& bOutHasUnbounded<br>) const | Determine the relevant grid sizes by inspecting all HiGenGridSize nodes. | PCGGraph.h |  |
| [UPCGNode](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/UPCGNode) \\* GetInputNode() | Returns the graph input node | PCGGraph.h | - BlueprintCallable<br>- Category=Graph |
| const [FPCGStack](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/FPCGStack) & GetInspectedStack() |  | PCGGraph.h |  |
| uint32 [GetNodeGenerationGridSize](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/UPCGGraph/GetNodeGenerationGridSize)<br>(<br>const [UPCGNode](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/UPCGNode)\\* InNode,<br>uint32 InDefaultGridSize<br>) const | Size of grid on which this node should be executed. | PCGGraph.h |  |
| const [TArray](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/TArray) < [UPCGNode](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/UPCGNode) \\* \> & GetNodes() |  | PCGGraph.h |  |
| [UPCGNode](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/UPCGNode) \\* GetOutputNode() | Returns the graph output node | PCGGraph.h | - BlueprintCallable<br>- Category=Graph |
| void GetParentGridSizes<br>(<br>const uint32 InChildGridSize,<br>PCGHiGenGrid::FSizeArray& OutParentGridSizes<br>) const | Returns all parent grid sizes for the given child grid size, calculated by inspecting nodes. | PCGGraph.h |  |
| void [GetTrackedActorKeysToSettings](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/UPCGGraph/GetTrackedActorKeysToSettings)<br>(<br>FPCGSelectionKeyToSettingsMap& OutKeysToSettings,<br>[TArray](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/TArray) < [TObjectPtr](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/TObjectPtr) < const [UPCGGraph](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/UPCGGraph) \> >& OutVisitedGraphs<br>) const |  | PCGGraph.h |  |
| FPCGSelectionKeyToSettingsMap [GetTrackedActorKeysToSettings](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/UPCGGraph/GetTrackedActorKeysToSettings) () |  | PCGGraph.h |  |
| bool HasDefaultConstructedInputs() |  | PCGGraph.h |  |
| bool IsHierarchicalGenerationEnabled() |  | PCGGraph.h |  |
| bool IsInspecting() |  | PCGGraph.h |  |
| bool NotificationsForEditorArePausedByUser() |  | PCGGraph.h |  |
| void OnPCGQualityLevelChanged() |  | PCGGraph.h |  |
| void PostNodeUndo<br>(<br>[UPCGNode](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/UPCGNode)\\* InPCGNode<br>) |  | PCGGraph.h |  |
| void PreNodeUndo<br>(<br>[UPCGNode](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/UPCGNode)\\* InPCGNode<br>) |  | PCGGraph.h |  |
| bool PrimeGraphCompilationCache() | Instruct the graph compiler to cache the relevant permutations of this graph. | PCGGraph.h |  |
| bool Recompile() | Trigger a recompilation of the relevant permutations of this graph and check for change in the compiled tasks. | PCGGraph.h |  |
| [TObjectPtr](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/TObjectPtr) < [UPCGNode](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/UPCGNode) \> ReconstructNewNode<br>(<br>const [UPCGNode](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/UPCGNode)\\* InNode<br>) | Duplicate a given node by creating a new node with the same settings and properties, but without any edges and add it to the graph | PCGGraph.h |  |
| void RemoveCommentNode<br>(<br>const [FGuid](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/FGuid)& InNodeGUID<br>) |  | PCGGraph.h |  |
| bool RemoveEdge<br>(<br>[UPCGNode](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/UPCGNode)\\* From,<br>const [FName](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/FName)& FromLabel,<br>[UPCGNode](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/UPCGNode)\\* To,<br>const [FName](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/FName)& ToLabel<br>) | Removes an edge in the graph. Returns true if an edge was removed. | PCGGraph.h | - BlueprintCallable<br>- Category=Graph |
| void RemoveExtraEditorNode<br>(<br>const [UObject](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/CoreUObject/UObject)\\* InNode<br>) |  | PCGGraph.h |  |
| bool RemoveInboundEdges<br>(<br>[UPCGNode](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/UPCGNode)\\* InNode,<br>const [FName](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/FName)& InboundLabel<br>) |  | PCGGraph.h |  |
| void RemoveNode<br>(<br>[UPCGNode](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/UPCGNode)\\* InNode<br>) | Removes a node from the graph. | PCGGraph.h | - BlueprintCallable<br>- Category=Graph |
| void RemoveNodes<br>(<br>[TArray](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/TArray) < [UPCGNode](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/UPCGNode)\\* >& InNodes<br>) | Bulk removal of nodes, to avoid notifying the world everytime. | PCGGraph.h | - BlueprintCallable<br>- Category=Graph |
| bool RemoveOutboundEdges<br>(<br>[UPCGNode](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/UPCGNode)\\* InNode,<br>const [FName](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/FName)& OutboundLabel<br>) |  | PCGGraph.h |  |
| [TSharedPtr](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/TSharedPtr) < [FComputeGraphInstance](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/ComputeFramework/FComputeGraphInstance) \> [RetrieveComputeGraphInstanceFromPool](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/UPCGGraph/RetrieveComputeGraphInstanceFrom-)<br>(<br>const FComputeGraphInstanceKey& InKey,<br>bool& bOutNewInstance<br>) const | Attempt to retrieve a pooled compute graph instance for the given key. | PCGGraph.h |  |
| void ReturnComputeGraphInstanceToPool<br>(<br>const FComputeGraphInstanceKey& InKey,<br>[TSharedPtr](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/TSharedPtr) < [FComputeGraphInstance](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/ComputeFramework/FComputeGraphInstance) \> InInstance<br>) const | Places given compute graph instance into pool. No-ops if pooling is disabled. | PCGGraph.h |  |
| virtual bool [SanitizeNodeTitle](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/UPCGGraph/SanitizeNodeTitle)<br>(<br>[FName](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/FName)& InOutTitle,<br>const [UPCGNode](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/UPCGNode)\\* InNode<br>) const | Overridable function for child classes to have a graph-wide node title sanitization when the title of a node changes. | PCGGraph.h |  |
| void SetCommentNodes<br>(<br>[TArray](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/TArray) < [FPCGGraphCommentNodeData](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/FPCGGraphCommentNodeData) \> InNodes<br>) |  | PCGGraph.h |  |
| void SetExtraEditorNodes<br>(<br>const [TArray](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/TArray) < [TObjectPtr](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/TObjectPtr) < const [UObject](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/CoreUObject/UObject) \> >& InNodes<br>) |  | PCGGraph.h |  |
| virtual bool ShouldDisplayDebuggingProperties() | Can be overriden by child class to disable debug globally on all settings. | PCGGraph.h |  |
| void ToggleUserPausedNotificationsForEditor() |  | PCGGraph.h |  |
| void UpdateUserParametersStruct<br>(<br>[TFunctionRef](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/TFunctionRef) < void( [FInstancedPropertyBag](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/CoreUObject/FInstancedPropertyBag)&)\> Callback<br>) | Will call the callback function with a mutable property bag and will trigger the updates when it's done. | PCGGraph.h |  |
| bool Use2DGrid() |  | PCGGraph.h |  |

#### Overridden from [UPCGGraphInterface](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/UPCGGraphInterface)

| Name | Remarks | Include Path | [Unreal Specifiers](https://dev.epicgames.com/documentation/unreal-engine/ufunctions-in-unreal-engine#functionspecifiers) |
| --- | --- | --- | --- |
| virtual [UPCGGraph](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/UPCGGraph) \\* [GetGraph](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/UPCGGraph/GetGraph) () | ~End [UObject](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/CoreUObject/UObject) interface ~Begin [UPCGGraphInterface](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/UPCGGraphInterface) interface | PCGGraph.h |  |
| virtual const [UPCGGraph](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/UPCGGraph) \\* [GetGraph](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/UPCGGraph/GetGraph) () |  | PCGGraph.h |  |
| virtual [TOptional](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/TOptional) < [FPCGGraphToolData](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/FPCGGraphToolData) \> GetGraphToolData() |  | PCGGraph.h |  |
| virtual const [FInstancedPropertyBag](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/CoreUObject/FInstancedPropertyBag) \\* GetUserParametersStruct() |  | PCGGraph.h |  |
| virtual void OnGraphParametersChanged<br>(<br>[EPCGGraphParameterEvent](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/EPCGGraphParameterEvent) InChangeType,<br>[FName](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/FName) InChangedPropertyName<br>) |  | PCGGraph.h |  |

#### Overridden from [UObject](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/CoreUObject/UObject)

| Name | Remarks | Include Path | [Unreal Specifiers](https://dev.epicgames.com/documentation/unreal-engine/ufunctions-in-unreal-engine#functionspecifiers) |
| --- | --- | --- | --- |
| virtual void BeginDestroy() |  | PCGGraph.h |  |
| virtual bool IsEditorOnly() |  | PCGGraph.h |  |
| virtual void PostEditChangeChainProperty<br>(<br>[FPropertyChangedChainEvent](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/CoreUObject/FPropertyChangedChainEvent)& PropertyChangedEvent<br>) |  | PCGGraph.h |  |
| virtual void PostEditUndo() |  | PCGGraph.h |  |
| virtual void PostLoad() | ~Begin [UObject](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/CoreUObject/UObject) interface | PCGGraph.h |  |
| virtual void PreEditChange<br>(<br>[FProperty](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/CoreUObject/FProperty)\\* InProperty<br>) |  | PCGGraph.h |  |
| virtual void PreSave<br>(<br>[FObjectPreSaveContext](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/CoreUObject/FObjectPreSaveContext) ObjectSaveContext<br>) |  | PCGGraph.h |  |

### Protected

| Name | Remarks | Include Path | [Unreal Specifiers](https://dev.epicgames.com/documentation/unreal-engine/ufunctions-in-unreal-engine#functionspecifiers) |
| --- | --- | --- | --- |
| void AddNodes\_Internal<br>(<br>[TArrayView](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/TArrayView) < [UPCGNode](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/UPCGNode)\\* \> InNodes<br>) |  | PCGGraph.h |  |
| void CacheGridSizesInternalNoLock() |  | PCGGraph.h |  |
| uint32 CalculateNodeGridSizeRecursive\_Unsafe<br>(<br>const [UPCGNode](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/UPCGNode)\\* InNode,<br>uint32 InDefaultGridSize<br>) const | Calculates node grid size. Not thread safe, must be called within write lock. | PCGGraph.h |  |
| PCGHiGenGrid::FSizeArray [CalculateNodeGridSizesRecursiveNoLock](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/UPCGGraph/CalculateNodeGridSizesRecursiveN-)<br>(<br>const [UPCGNode](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/UPCGNode)\\* InNode,<br>uint32 InDefaultGridSize<br>) const | Calculates all parent grid sizes that a given node depends on. | PCGGraph.h |  |
| virtual bool CanToggleStandaloneGraph() |  | PCGGraph.h |  |
| bool ForEachNodeRecursively\_Internal<br>(<br>[TFunctionRef](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/TFunctionRef) < bool( [UPCGNode](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/UPCGNode)\*)\> Action,<br>TSet< const [UPCGGraph](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/UPCGGraph)\\* >& VisitedGraphs<br>) const |  | PCGGraph.h |  |
| bool IsEditorOnly\_Internal() |  | PCGGraph.h |  |
| void [OnNodeAdded](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/UPCGGraph/OnNodeAdded)<br>(<br>[UPCGNode](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/UPCGNode)\\* InNode,<br>bool bNotify<br>) | Internal function to react to add/remove nodes. | PCGGraph.h |  |
| void OnNodeRemoved<br>(<br>[UPCGNode](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/UPCGNode)\\* InNode,<br>bool bNotify<br>) |  | PCGGraph.h |  |
| void OnNodesAdded<br>(<br>[TArrayView](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/TArrayView) < [UPCGNode](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/UPCGNode)\\* \> InNodes,<br>bool bNotify<br>) |  | PCGGraph.h |  |
| void OnNodesRemoved<br>(<br>[TArrayView](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/TArrayView) < [UPCGNode](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/UPCGNode)\\* \> InNodes,<br>bool bNotify<br>) |  | PCGGraph.h |  |
| void RemoveNodes\_Internal<br>(<br>[TArrayView](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/TArrayView) < [UPCGNode](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/UPCGNode)\\* \> InNodes<br>) |  | PCGGraph.h |  |
| void SetHiddenFlagInputNode<br>(<br>bool bHidden<br>) | Mark the input node hidden/not hidden. | PCGGraph.h |  |
| void SetHiddenFlagOutputNode<br>(<br>bool bHidden<br>) | Mark the output node hidden/not hidden. | PCGGraph.h |  |
| virtual bool ShowGraphCustomization() |  | PCGGraph.h |  |
| virtual bool SupportHierarchicalGeneration() |  | PCGGraph.h |  |
| bool UserParametersCanRemoveProperty<br>(<br>[FGuid](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/FGuid) InPropertyID,<br>[FName](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/FName) InPropertyName<br>) |  | PCGGraph.h | - BlueprintInternalUseOnly |
| bool UserParametersIsPinTypeAccepted<br>(<br>[FEdGraphPinType](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Engine/FEdGraphPinType) InPinType,<br>bool bIsChild<br>) |  | PCGGraph.h | - BlueprintInternalUseOnly |

#### Overridden from [UPCGGraphInterface](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/UPCGGraphInterface)

| Name | Remarks | Include Path | [Unreal Specifiers](https://dev.epicgames.com/documentation/unreal-engine/ufunctions-in-unreal-engine#functionspecifiers) |
| --- | --- | --- | --- |
| virtual [FInstancedPropertyBag](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/CoreUObject/FInstancedPropertyBag) \\* GetMutableUserParametersStruct() |  | PCGGraph.h |  |
| virtual bool IsTemplatePropertyEnabled() |  | PCGGraph.h |  |

### Static

| Name | Remarks | Include Path | [Unreal Specifiers](https://dev.epicgames.com/documentation/unreal-engine/ufunctions-in-unreal-engine#functionspecifiers) |
| --- | --- | --- | --- |
| static void AddReferencedObjects<br>(<br>[UObject](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/CoreUObject/UObject)\\* InThis,<br>[FReferenceCollector](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/CoreUObject/FReferenceCollector)& Collector<br>) |  | PCGGraph.h |  |
| static void DeclareConstructClasses<br>(<br>[TArray](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/TArray) < [FTopLevelAssetPath](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/CoreUObject/FTopLevelAssetPath) >& OutConstructClasses,<br>const [UClass](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/CoreUObject/UClass)\\* SpecificSubclass<br>) |  | PCGGraph.h |  |

* * *