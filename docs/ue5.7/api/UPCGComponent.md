<!-- Source: https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/UPCGComponent -->

|  |  |
| --- | --- |
| _Name_ | UPCGComponent |
| _Type_ | class |
| _Header File_ | /Engine/Plugins/PCG/Source/PCG/Public/PCGComponent.h |
| _Include Path_ | #include "PCGComponent.h" |

## Syntax

```cpp

UCLASS (MinimalAPI, BlueprintType, ClassGroup=(Procedural),

       Meta=(BlueprintSpawnableComponent, prioritizeCategories="PCG"))

class UPCGComponent :

    public UActorComponent ,

    public IPCGGraphExecutionSource
```

## Inheritance Hierarchy

- [UObjectBase](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/CoreUObject/UObjectBase) → [UObjectBaseUtility](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/CoreUObject/UObjectBaseUtility) → [UObject](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/CoreUObject/UObject) → [UActorComponent](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Engine/UActorComponent) → **UPCGComponent**

## Implements Interfaces

- [IAsyncPhysicsStateProcessor](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Engine/IAsyncPhysicsStateProcessor)
- [IInterface\_AssetUserData](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Engine/IInterface_AssetUserData)
- [IPCGGraphExecutionSource](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/IPCGGraphExecutionSource)

## Constructors

| Name | Remarks | Include Path | [Unreal Specifiers](https://dev.epicgames.com/documentation/unreal-engine/ufunctions-in-unreal-engine#functionspecifiers) |
| --- | --- | --- | --- |
| UPCGComponent<br>(<br>const [FObjectInitializer](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/CoreUObject/FObjectInitializer)& InObjectInitializer<br>) |  | PCGComponent.h |  |

## Variables

### Public

| Name | Type | Remarks | Include Path | [Unreal Specifiers](https://dev.epicgames.com/documentation/unreal-engine/unreal-engine-uproperties#propertyspecifiers) |
| --- | --- | --- | --- | --- |
| bActivated | bool |  | PCGComponent.h | - BlueprintReadWrite<br>- EditAnywhere<br>- Category=Settings<br>- Meta=(DisplayPriority=100) |
| bDirtyGenerated | bool |  | PCGComponent.h | - BlueprintReadOnly<br>- VisibleAnywhere<br>- Transient<br>- Category=Debug<br>- Meta=(NoResetToDefault) |
| bForceGenerateOnBPAddedToWorld | bool | Property that will automatically be set on BP templates, to allow for "Generate on add to world" in editor. | PCGComponent.h |  |
| bGenerated | bool | Flag to indicate whether this component has run in the editor. | PCGComponent.h | - BlueprintReadOnly<br>- VisibleAnywhere<br>- Category=Debug<br>- NonTransactional<br>- Meta=(NoResetToDefault) |
| bGenerateOnDropWhenTriggerOnDemand | bool | When Generation Trigger is OnDemand, we can still force the component to generate on drop. | PCGComponent.h | - BlueprintReadOnly<br>- EditAnywhere<br>- Category="Settings\|Advanced"<br>- Meta=(EditCondition="GenerationTrigger == EPCGComponentGenerationTrigger::GenerateOnDemand") |
| bIgnoreLandscapeTracking | bool | Marks the component to be not refreshed automatically when the landscape changes, even if it is used. | PCGComponent.h | - BlueprintReadWrite<br>- EditAnywhere<br>- Category="Editing Settings"<br>- Meta=(DisplayPriority=451) |
| bIsComponentPartitioned | bool | Will partition the component in a grid, dispatching the generation to multiple local components. | PCGComponent.h | - BlueprintReadWrite<br>- EditAnywhere<br>- Category=Settings<br>- Meta=(EditCondition="!bIsComponentLocal", DisplayName="Is Partitioned", DisplayPriority=500) |
| bOnlyTrackItself | bool | Even if the graph has external dependencies, the component won't react to them. | PCGComponent.h | - BlueprintReadWrite<br>- EditAnywhere<br>- Category="Editing Settings"<br>- Meta=(DisplayPriority=450) |
| bOverrideGenerationRadii | bool | Manual overrides for the graph generation radii and cleanup radius multiplier. | PCGComponent.h | - BlueprintReadWrite<br>- EditAnywhere<br>- Category=RuntimeGeneration<br>- Meta=(EditCondition="GenerationTrigger == EPCGComponentGenerationTrigger::GenerateAtRuntime", EditConditionHides) |
| bParseActorComponents | bool |  | PCGComponent.h | - BlueprintReadWrite<br>- EditAnywhere<br>- Category="Input Node Settings (Deprecated)"<br>- Meta=(DisplayPriority=900) |
| bRegenerateInEditor | bool |  | PCGComponent.h | - BlueprintReadWrite<br>- EditAnywhere<br>- Category="Editing Settings"<br>- Meta=(DisplayName="Regenerate PCG Volume In Editor", DisplayPriority=400) |
| bRuntimeGenerated | bool |  | PCGComponent.h | - NonPIEDuplicateTransient |
| ExtraCapture | [PCGUtils::FExtraCapture](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/FExtraCapture) |  | PCGComponent.h |  |
| GenerationRadii | [FPCGRuntimeGenerationRadii](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/FPCGRuntimeGenerationRadii) |  | PCGComponent.h | - BlueprintReadWrite<br>- EditAnywhere<br>- Category=RuntimeGeneration<br>- Meta=(EditCondition="GenerationTrigger == EPCGComponentGenerationTrigger::GenerateAtRuntime && bOverrideGenerationRadii", EditConditionHides) |
| GenerationTrigger | [EPCGComponentGenerationTrigger](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/EPCGComponentGenerationTrigger) |  | PCGComponent.h | - BlueprintReadOnly<br>- EditAnywhere<br>- Category=Settings<br>- Meta=(EditCondition="!bIsComponentLocal", EditConditionHides, DisplayPriority=200) |
| InputType | [EPCGComponentInput](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/EPCGComponentInput) |  | PCGComponent.h | - BlueprintReadWrite<br>- EditAnywhere<br>- Category="Input Node Settings (Deprecated)"<br>- Meta=(DisplayPriority=800) |
| OnPCGGraphCancelledDelegate | FOnPCGGraphCancelled |  | PCGComponent.h |  |
| OnPCGGraphCancelledExternal | [FOnPCGGraphCancelledExternal](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/FOnPCGGraphCancelledExternal) | Event dispatched when a graph cancels generation on this component. | PCGComponent.h | - BlueprintAssignable<br>- Meta=(DisplayName="On Graph Cancelled") |
| OnPCGGraphCleanedDelegate | FOnPCGGraphCleaned |  | PCGComponent.h |  |
| OnPCGGraphCleanedExternal | [FOnPCGGraphCleanedExternal](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/FOnPCGGraphCleanedExternal) | Event dispatched when a graph cleans on this component. | PCGComponent.h | - BlueprintAssignable<br>- Meta=(DisplayName="On Graph Cleaned") |
| OnPCGGraphGeneratedDelegate | FOnPCGGraphGenerated |  | PCGComponent.h |  |
| OnPCGGraphGeneratedExternal | [FOnPCGGraphGeneratedExternal](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/FOnPCGGraphGeneratedExternal) | Event dispatched when a graph completes its generation on this component. | PCGComponent.h | - BlueprintAssignable<br>- Meta=(DisplayName="On Graph Generated") |
| OnPCGGraphStartGeneratingDelegate | FOnPCGGraphStartGenerating |  | PCGComponent.h |  |
| OnPCGGraphStartGeneratingExternal | [FOnPCGGraphStartGeneratingExternal](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/FOnPCGGraphStartGeneratingExtern-) | Event dispatched when a graph begins generation on this component. | PCGComponent.h | - BlueprintAssignable<br>- Meta=(DisplayName="On Graph Started Generating") |
| PostGenerateFunctionNames | [TArray](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/TArray) < [FName](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/FName) > | Can specify a list of functions from the owner of this component to be called when generation is done, in order. | PCGComponent.h | - BlueprintReadWrite<br>- EditAnywhere<br>- Category=Settings<br>- Meta=(DisplayPriority=700) |
| SchedulingPolicy | [TObjectPtr](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/TObjectPtr) < [UPCGSchedulingPolicyBase](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/UPCGSchedulingPolicyBase) > | This is the instanced UPCGSchedulingPolicy object which holds scheduling parameters and calculates priorities. | PCGComponent.h | - BlueprintReadOnly<br>- VisibleAnywhere<br>- Instanced<br>- Category=RuntimeGeneration<br>- Meta=(EditCondition="GenerationTrigger == EPCGComponentGenerationTrigger::GenerateAtRuntime", EditConditionHides) |
| SchedulingPolicyClass | [TSubclassOf](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/CoreUObject/TSubclassOf) < [UPCGSchedulingPolicyBase](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/UPCGSchedulingPolicyBase) > | A Scheduling Policy dictates the order in which instances of this component will be scheduled. | PCGComponent.h | - BlueprintReadOnly<br>- EditAnywhere<br>- NoClear<br>- Category=RuntimeGeneration<br>- Meta=(EditCondition="GenerationTrigger == EPCGComponentGenerationTrigger::GenerateAtRuntime", EditConditionHides) |
| Seed | int |  | PCGComponent.h | - BlueprintReadWrite<br>- EditAnywhere<br>- Category=Settings<br>- Meta=(DisplayPriority=600) |
| Timer | [PCGUtils::FCallTime](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/FCallTime) |  | PCGComponent.h |  |
| ToolDataContainer | [FPCGInteractiveToolDataContainer](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/FPCGInteractiveToolDataContainer) | This stores working data of a tool; runtime property as the PCG Component will rely on the tool's working data to generate properly. | PCGComponent.h | - BlueprintReadOnly<br>- EditAnywhere<br>- Category="Tool Data" |
| TrackingPriority | double | Tracking priority used to solve tracking dependencies between PCG Components. | PCGComponent.h | - BlueprintReadWrite<br>- EditAnywhere<br>- Category="Editing Settings"<br>- Meta=(DisplayPriority=452, ClampMin="-10000.0", ClampMax="10000.0", UIMin="-10000.0", UIMax="10000.0", Delta="1") |

### Protected

| Name | Type | Remarks | Include Path | [Unreal Specifiers](https://dev.epicgames.com/documentation/unreal-engine/unreal-engine-uproperties#propertyspecifiers) |
| --- | --- | --- | --- | --- |
| bDisableIsComponentPartitionedOnLoad | bool | Track if component should disable 'bIsComponentPartitioned'. | PCGComponent.h |  |
| bGenerationInProgress | bool |  | PCGComponent.h | - Transient<br>- VisibleInstanceOnly<br>- Category=Debug |
| bIsComponentLocal | bool |  | PCGComponent.h | - VisibleAnywhere<br>- Transient<br>- Category=Debug<br>- Meta=(EditCondition=false, EditConditionHides) |
| bProceduralInstancesInUse | bool | Whether procedural ISM components were used/generated in the last execution. | PCGComponent.h |  |
| bUnregisteredThroughLoading | bool | Track if component was unregistered while in a loading scope. | PCGComponent.h |  |
| bWasGeneratedThisSession | bool |  | PCGComponent.h |  |
| CachedActorData | [TObjectPtr](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/TObjectPtr) < [UPCGData](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/UPCGData) > |  | PCGComponent.h | - Transient<br>- NonPIEDuplicateTransient |
| CachedInputData | [TObjectPtr](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/TObjectPtr) < [UPCGData](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/UPCGData) > |  | PCGComponent.h | - Transient<br>- NonPIEDuplicateTransient |
| CachedLandscapeData | [TObjectPtr](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/TObjectPtr) < [UPCGData](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/UPCGData) > |  | PCGComponent.h | - Transient<br>- NonPIEDuplicateTransient |
| CachedLandscapeHeightData | [TObjectPtr](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/TObjectPtr) < [UPCGData](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/UPCGData) > |  | PCGComponent.h | - Transient<br>- NonPIEDuplicateTransient |
| CachedPCGData | [TObjectPtr](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/TObjectPtr) < [UPCGData](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/UPCGData) > |  | PCGComponent.h | - Transient<br>- NonPIEDuplicateTransient |
| CurrentCleanupTask | FPCGTaskId |  | PCGComponent.h |  |
| CurrentExecutionDynamicTracking | FPCGSelectionKeyToSettingsMap | Temporary storage for dynamic tracking that will be filled during component execution. | PCGComponent.h |  |
| CurrentExecutionDynamicTrackingLock | FTransactionallySafeCriticalSection |  | PCGComponent.h |  |
| CurrentExecutionDynamicTrackingSettings | TSet< const [UPCGSettings](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/UPCGSettings) \\* > | Temporary storage for dynamic tracking that will keep all settings that could have dynamic tracking, in order to detect changes. | PCGComponent.h |  |
| CurrentGenerationTask | FPCGTaskId |  | PCGComponent.h |  |
| CurrentRefreshTask | FPCGTaskId |  | PCGComponent.h |  |
| DynamicallyTrackedKeysToSettings | FPCGSelectionKeyToSettingsMap | Need to keep a reference to all tracked settings to still react to changes after a map load (since the component won't have been executed). | PCGComponent.h |  |
| ExecutionInspection | [FPCGGraphExecutionInspection](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/FPCGGraphExecutionInspection) |  | PCGComponent.h |  |
| ExecutionState | [FPCGComponentExecutionState](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/FPCGComponentExecutionState) |  | PCGComponent.h |  |
| GeneratedGraphOutput | [FPCGDataCollection](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/FPCGDataCollection) |  | PCGComponent.h | - VisibleInstanceOnly<br>- Category=Debug |
| GeneratedResources | [TArray](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/TArray) < [TObjectPtr](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/TObjectPtr) < [UPCGManagedResource](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/UPCGManagedResource) \> > | NOTE: This should not be made visible or editable because it will change the way the BP actors are duplicated/setup and might trigger an ensure in the resources. | PCGComponent.h |  |
| GeneratedResourcesInaccessible | bool | When doing a cleanup, locking resource modification. Used as sentinel. | PCGComponent.h |  |
| GeneratedResourcesLock | FTransactionallySafeCriticalSection |  | PCGComponent.h |  |
| IgnoredChangeOriginsLock | FTransactionallySafeRWLock |  | PCGComponent.h |  |
| IgnoredChangeOriginsToCounters | [TMap](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/TMap) < [TObjectKey](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/CoreUObject/TObjectKey) < [UObject](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/CoreUObject/UObject) >, int32 > | The tracking system will not trigger a generation on this component for these change origins. | PCGComponent.h |  |
| LastGeneratedBounds | FBox |  | PCGComponent.h | - VisibleInstanceOnly<br>- Category=Debug |
| LastGeneratedBoundsPriorToUndo | FBox |  | PCGComponent.h |  |
| LoadedPreviewGeneratedGraphOutput | [FPCGDataCollection](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/FPCGDataCollection) |  | PCGComponent.h | - Transient |
| LoadedPreviewResources | [TArray](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/TArray) < [TObjectPtr](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/TObjectPtr) < [UPCGManagedResource](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/UPCGManagedResource) \> > |  | PCGComponent.h | - Transient |
| PerPinGeneratedOutput | [TMap](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/TMap) < [FString](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/FString), [FPCGDataCollection](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/FPCGDataCollection) > | If any graph edges cross execution grid sizes, data on the edge is stored / retrieved from this map. | PCGComponent.h | - Transient<br>- NonTransactional<br>- VisibleAnywhere<br>- Category=Debug |
| PerPinGeneratedOutputLock | FTransactionallySafeRWLock |  | PCGComponent.h |  |
| RuntimeGridDescriptorHash | uint32 |  | PCGComponent.h |  |
| StaticallyTrackedKeysToSettings | FPCGSelectionKeyToSettingsMap |  | PCGComponent.h |  |

## Functions

### Public

| Name | Remarks | Include Path | [Unreal Specifiers](https://dev.epicgames.com/documentation/unreal-engine/ufunctions-in-unreal-engine#functionspecifiers) |
| --- | --- | --- | --- |
| void [AddActorsToManagedResources](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/UPCGComponent/AddActorsToManagedResources)<br>(<br>const [TArray](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/TArray) < [AActor](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Engine/AActor)\\* >& InActors<br>) | Creates a managed actors resource and adds it to the current component. | PCGComponent.h | - BlueprintCallable<br>- Category=PCG |
| void [AddComponentsToManagedResources](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/UPCGComponent/AddComponentsToManagedResources)<br>(<br>const [TArray](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/TArray) < [UActorComponent](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Engine/UActorComponent)\\* >& InComponents<br>) | Creates a managed component resource and adds it to the current component. | PCGComponent.h | - BlueprintCallable<br>- Category=PCG |
| void AddToManagedResources<br>(<br>[UPCGManagedResource](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/UPCGManagedResource)\\* InResource<br>) | Registers some managed resource to the current component | PCGComponent.h | - BlueprintCallable<br>- Category=PCG |
| bool AreManagedResourcesAccessible() |  | PCGComponent.h |  |
| bool AreProceduralInstancesInUse() | Whether this component created one or more procedural ISM components when last generated. | PCGComponent.h |  |
| void CancelGeneration() | Cancels in-progress generation | PCGComponent.h |  |
| bool CanPartition() |  | PCGComponent.h |  |
| void ChangeTransientState<br>(<br>[EPCGEditorDirtyMode](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/EPCGEditorDirtyMode) NewEditingMode<br>) | Changes the transient state (preview, normal, load on preview) - public only because it needs to be accessed by [APCGPartitionActor](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/APCGPartitionActor) | PCGComponent.h |  |
| void [Cleanup](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/UPCGComponent/Cleanup)<br>(<br>bool bRemoveComponents<br>) | Networked cleanup call | PCGComponent.h | - BlueprintCallable<br>- NetMulticast<br>- Reliable<br>- Category=PCG |
| void [Cleanup](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/UPCGComponent/Cleanup) () |  | PCGComponent.h |  |
| void [CleanupLocal](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/UPCGComponent/CleanupLocal)<br>(<br>bool bRemoveComponents,<br>bool bSave,<br>const [TArray](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/TArray) < FPCGTaskId >& Dependencies<br>) |  | PCGComponent.h |  |
| void [CleanupLocal](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/UPCGComponent/CleanupLocal)<br>(<br>bool bRemoveComponents,<br>bool bSave<br>) |  | PCGComponent.h |  |
| FPCGTaskId [CleanupLocal](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/UPCGComponent/CleanupLocal)<br>(<br>bool bRemoveComponents,<br>const [TArray](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/TArray) < FPCGTaskId >& Dependencies<br>) |  | PCGComponent.h |  |
| void [CleanupLocal](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/UPCGComponent/CleanupLocal)<br>(<br>bool bRemoveComponents<br>) | Cleans up the generation from a local (vs. remote) standpoint. | PCGComponent.h | - BlueprintCallable<br>- Category=PCG |
| void CleanupLocalDeleteAllGeneratedObjects<br>(<br>const [TArray](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/TArray) < FPCGTaskId >& Dependencies<br>) | Cleanup the generation while purging Actors and Components tagged as generated by PCG but are no longer managed by this or any other actor | PCGComponent.h |  |
| void [CleanupLocalImmediate](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/UPCGComponent/CleanupLocalImmediate)<br>(<br>bool bRemoveComponents,<br>bool bCleanupLocalComponents<br>) | Same as CleanupLocal, but without any delayed tasks. | PCGComponent.h |  |
| void ClearInspectionData<br>(<br>bool bClearPerNodeExecutionData<br>) |  | PCGComponent.h |  |
| [AActor](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Engine/AActor) \\* [ClearPCGLink](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/UPCGComponent/ClearPCGLink)<br>(<br>[UClass](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/CoreUObject/UClass)\\* TemplateActor<br>) | Move all generated resources under a new actor, following a template ( [AActor](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Engine/AActor) if not provided), clearing all link to this PCG component. | PCGComponent.h | - BlueprintCallable<br>- Category=PCG |
| void ClearPerPinGeneratedOutput() | Clear any data stored for any pins. | PCGComponent.h |  |
| void [DirtyGenerated](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/UPCGComponent/DirtyGenerated)<br>(<br>[EPCGComponentDirtyFlag](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/EPCGComponentDirtyFlag) DataToDirtyFlag,<br>const bool bDispatchToLocalComponents<br>) | Dirty generated data depending on the flag. | PCGComponent.h |  |
| void DisableInspection() |  | PCGComponent.h |  |
| bool DoesGridDependOnWorldStreaming<br>(<br>uint32 InGridSize<br>) const |  | PCGComponent.h |  |
| void EnableInspection() |  | PCGComponent.h |  |
| void ForEachConstManagedResource<br>(<br>[TFunctionRef](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/TFunctionRef) < void(const [UPCGManagedResource](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/UPCGManagedResource)\*)\> InFunction<br>) const |  | PCGComponent.h |  |
| void ForEachManagedResource<br>(<br>[TFunctionRef](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/TFunctionRef) < void( [UPCGManagedResource](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/UPCGManagedResource)\*)\> InFunction<br>) |  | PCGComponent.h |  |
| void [Generate](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/UPCGComponent/Generate) () |  | PCGComponent.h |  |
| void [Generate](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/UPCGComponent/Generate)<br>(<br>bool bForce<br>) | Networked generation call that also activates the component as needed | PCGComponent.h | - BlueprintCallable<br>- NetMulticast<br>- Reliable<br>- Category=PCG |
| void [GenerateLocal](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/UPCGComponent/GenerateLocal)<br>(<br>[EPCGComponentGenerationTrigger](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/EPCGComponentGenerationTrigger) RequestedGenerationTrigger,<br>bool bForce,<br>[EPCGHiGenGrid](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/EPCGHiGenGrid) Grid,<br>const [TArray](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/TArray) < FPCGTaskId >& Dependencies<br>) | Requests the component to generate only on the specified grid level (all grid levels if [EPCGHiGenGrid::Uninitialized](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/EPCGHiGenGrid)). | PCGComponent.h |  |
| void [GenerateLocal](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/UPCGComponent/GenerateLocal)<br>(<br>bool bForce<br>) | Starts generation from a local (vs. remote) standpoint. Will not be replicated. Will be delayed. | PCGComponent.h | - BlueprintCallable<br>- Category=PCG |
| FPCGTaskId [GenerateLocalGetTaskId](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/UPCGComponent/GenerateLocalGetTaskId)<br>(<br>bool bForce<br>) |  | PCGComponent.h |  |
| FPCGTaskId [GenerateLocalGetTaskId](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/UPCGComponent/GenerateLocalGetTaskId)<br>(<br>[EPCGComponentGenerationTrigger](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/EPCGComponentGenerationTrigger) RequestedGenerationTrigger,<br>bool bForce,<br>[EPCGHiGenGrid](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/EPCGHiGenGrid) Grid,<br>const [TArray](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/TArray) < FPCGTaskId >& Dependencies<br>) |  | PCGComponent.h |  |
| FPCGTaskId [GenerateLocalGetTaskId](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/UPCGComponent/GenerateLocalGetTaskId)<br>(<br>[EPCGComponentGenerationTrigger](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/EPCGComponentGenerationTrigger) RequestedGenerationTrigger,<br>bool bForce,<br>[EPCGHiGenGrid](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/EPCGHiGenGrid) Grid<br>) |  | PCGComponent.h |  |
| [UPCGData](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/UPCGData) \\* GetActorPCGData() |  | PCGComponent.h |  |
| double GetCleanupRadiusFromGrid<br>(<br>[EPCGHiGenGrid](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/EPCGHiGenGrid) Grid<br>) const | Compute the runtime cleanup radius for the given grid size. | PCGComponent.h |  |
| const [UPCGComponent](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/UPCGComponent) \\* GetConstOriginalComponent() |  | PCGComponent.h |  |
| [EPCGEditorDirtyMode](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/EPCGEditorDirtyMode) GetEditingMode() | Returns the current editing mode | PCGComponent.h | - BlueprintCallable<br>- Category="PCG\|Advanced" |
| PRAGMA\_DISABLE\_DEPRECATION\_WARNINGS [TMap](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/TMap) < [TObjectKey](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/CoreUObject/TObjectKey) < const [UPCGNode](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/UPCGNode) >, TSet< NodeExecutedNotificationData > > GetExecutedNodeStacks() |  | PCGComponent.h |  |
| const [FPCGDataCollection](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/FPCGDataCollection) & GetGeneratedGraphOutput() | Retrieves generated data | PCGComponent.h | - BlueprintCallable<br>- Category=PCG |
| [EPCGHiGenGrid](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/EPCGHiGenGrid) GetGenerationGrid() |  | PCGComponent.h |  |
| uint32 GetGenerationGridSize() |  | PCGComponent.h |  |
| const [FPCGRuntimeGenerationRadii](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/FPCGRuntimeGenerationRadii) & GetGenerationRadii() | Get the generation radii that are currently active for this component. | PCGComponent.h |  |
| double GetGenerationRadiusFromGrid<br>(<br>[EPCGHiGenGrid](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/EPCGHiGenGrid) Grid<br>) const | Get the runtime generation radius for the given grid size. | PCGComponent.h |  |
| FPCGTaskId GetGenerationTaskId() | Returns task ids to do internal chaining | PCGComponent.h |  |
| [UPCGGraph](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/UPCGGraph) \\* GetGraph() |  | PCGComponent.h |  |
| [UPCGGraphInstance](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/UPCGGraphInstance) \\* GetGraphInstance() |  | PCGComponent.h |  |
| FBox [GetGridBounds](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/UPCGComponent/GetGridBounds) () |  | PCGComponent.h |  |
| [FPCGGridDescriptor](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/FPCGGridDescriptor) GetGridDescriptor<br>(<br>uint32 GridSize<br>) const | Returns a GridDescriptor based on this component for the specified grid size | PCGComponent.h |  |
| [UPCGData](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/UPCGData) \\* GetInputPCGData() |  | PCGComponent.h |  |
| const [FPCGDataCollection](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/FPCGDataCollection) \\* GetInspectionData<br>(<br>const [FPCGStack](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/FPCGStack)& InStack<br>) const |  | PCGComponent.h |  |
| [UPCGData](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/UPCGData) \\* GetLandscapeHeightPCGData() |  | PCGComponent.h |  |
| [UPCGData](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/UPCGData) \\* GetLandscapePCGData() |  | PCGComponent.h |  |
| FBox GetLastGeneratedBounds() |  | PCGComponent.h |  |
| FBox GetLocalSpaceBounds() |  | PCGComponent.h |  |
| uint64 GetNodeInactivePinMask<br>(<br>const [UPCGNode](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/UPCGNode)\\* InNode,<br>const [FPCGStack](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/FPCGStack)& Stack<br>) const | Retrieve the inactive pin bitmask for the given node and stack in the last execution. | PCGComponent.h |  |
| [UPCGData](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/UPCGData) \\* GetOriginalActorPCGData() |  | PCGComponent.h |  |
| [UPCGComponent](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/UPCGComponent) \\* GetOriginalComponent() | If this is a local component returns the original component, otherwise returns self. | PCGComponent.h |  |
| FBox GetOriginalGridBounds() |  | PCGComponent.h |  |
| FBox GetOriginalLocalSpaceBounds() |  | PCGComponent.h |  |
| [UPCGData](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/UPCGData) \\* GetPCGData() |  | PCGComponent.h |  |
| [UPCGSchedulingPolicyBase](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/UPCGSchedulingPolicyBase) \\* GetRuntimeGenSchedulingPolicy() |  | PCGComponent.h |  |
| [EPCGEditorDirtyMode](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/EPCGEditorDirtyMode) GetSerializedEditingMode() |  | PCGComponent.h | - BlueprintCallable<br>- Category="PCG\|Advanced" |
| bool GetStackContext<br>(<br>[FPCGStackContext](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/FPCGStackContext)& OutStackContext<br>) const | Get execution stack information. | PCGComponent.h |  |
| [UPCGSubsystem](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/UPCGSubsystem) \\* GetSubsystem() |  | PCGComponent.h |  |
| bool HasNodeProducedData<br>(<br>const [UPCGNode](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/UPCGNode)\\* InNode,<br>const [FPCGStack](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/FPCGStack)& Stack<br>) const | Did the given node produce one or more data items in the given stack during the last execution. | PCGComponent.h |  |
| void IgnoreChangeOriginDuringGenerationWithScope<br>(<br>const [UObject](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/CoreUObject/UObject)\\* InChangeOriginToIgnore,<br>Func InFunc<br>) |  | PCGComponent.h |  |
| void [IgnoreChangeOriginsDuringGenerationWithScope](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/UPCGComponent/IgnoreChangeOriginsDuringGenerat-)<br>(<br>const [TArrayView](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/TArrayView) < const [UObject](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/CoreUObject/UObject)\\* \> InChangeOriginsToIgnore,<br>Func InFunc<br>) | Allow for the function to be defined outside of editor, so it's lighter at the calling site and will just call the InFunc. | PCGComponent.h |  |
| bool IsAnyObjectManagedByResource<br>(<br>const [TArrayView](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/TArrayView) < const [UObject](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/CoreUObject/UObject)\\* \> InObjects<br>) const | Will scan the managed resources to check if any resource manage one of the objects. | PCGComponent.h |  |
| bool IsCleaningUp() |  | PCGComponent.h |  |
| bool IsGenerating() | Return if we are currently generating the graph for this component | PCGComponent.h |  |
| bool IsGenerationInProgress() |  | PCGComponent.h |  |
| bool IsIgnoringAnyChangeOrigins<br>(<br>const [TArrayView](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/TArrayView) < const [UObject](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/CoreUObject/UObject)\\* \> InChangeOrigins,<br>const [UObject](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/CoreUObject/UObject)\*& OutFirstObjectFound<br>) const |  | PCGComponent.h |  |
| bool IsIgnoringChangeOrigin<br>(<br>const [UObject](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/CoreUObject/UObject)\\* InChangeOrigin<br>) const |  | PCGComponent.h |  |
| bool IsInPreviewMode() | Returns whether the component (or resources) should be marked as dirty following interaction/refresh based on the current editing mode | PCGComponent.h |  |
| bool IsInspecting() |  | PCGComponent.h |  |
| bool IsLocalComponent() |  | PCGComponent.h |  |
| bool [IsManagedByRuntimeGenSystem](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/UPCGComponent/IsManagedByRuntimeGenSystem) () | Returns true if the component is managed by the runtime generation system. | PCGComponent.h |  |
| bool IsObjectTracked<br>(<br>const [UObject](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/CoreUObject/UObject)\\* InObject,<br>bool& bOutIsCulled<br>) const |  | PCGComponent.h |  |
| bool IsPartitioned() |  | PCGComponent.h |  |
| bool IsRefreshInProgress() | Returns current refresh task ID. | PCGComponent.h |  |
| void MarkAsLocalComponent() | Responsibility of the PCG Partition Actor to mark is local | PCGComponent.h |  |
| void NotifyNodeDynamicInactivePins<br>(<br>const [UPCGNode](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/UPCGNode)\\* InNode,<br>const [FPCGStack](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/FPCGStack)\\* InStack,<br>uint64 InactivePinBitmask<br>) const | Whether the given node was culled by a dynamic branch in the given stack. | PCGComponent.h |  |
| void NotifyNodeExecuted<br>(<br>const [UPCGNode](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/UPCGNode)\\* InNode,<br>const [FPCGStack](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/FPCGStack)\\* InStack,<br>const [PCGUtils::FCallTime](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/FCallTime)\\* InTimer,<br>bool bNodeUsedCache<br>) |  | PCGComponent.h |  |
| void NotifyProceduralInstancesInUse() | Called during execution if one or more procedural ISM components are in use. | PCGComponent.h |  |
| void NotifyPropertiesChangedFromBlueprint() | Notify properties changed, used in runtime cases, will dirty & trigger a regeneration if needed | PCGComponent.h | - BlueprintCallable<br>- Category=PCG |
| void OnRefresh<br>(<br>bool bForceRefresh<br>) |  | PCGComponent.h |  |
| void [Refresh](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/UPCGComponent/Refresh)<br>(<br>[EPCGChangeType](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/EPCGChangeType) ChangeType,<br>bool bCancelExistingRefresh<br>) | Schedules refresh of the component. | PCGComponent.h |  |
| void [RegisterDynamicTracking](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/UPCGComponent/RegisterDynamicTracking)<br>(<br>const FPCGSelectionKeyToSettingsMap& InKeysToSettings<br>) | To be called to notify the component that this list of keys have a dynamic dependency. | PCGComponent.h |  |
| void [RegisterDynamicTracking](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/UPCGComponent/RegisterDynamicTracking)<br>(<br>const [UPCGSettings](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/UPCGSettings)\\* InSettings,<br>const [TArrayView](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/TArrayView) < TPair< [FPCGSelectionKey](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/FPCGSelectionKey), bool > >& InDynamicKeysAndCulling<br>) | To be called by an element to notify the component that this settings have a dynamic dependency. | PCGComponent.h |  |
| void ResetIgnoredChangeOrigins<br>(<br>bool bLogIfAnyPresent<br>) |  | PCGComponent.h |  |
| void ResetLastGeneratedBounds() | Reset last generated bounds to force PCGPartitionActor creation on next refresh | PCGComponent.h |  |
| const [FPCGDataCollection](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/FPCGDataCollection) \\* RetrieveOutputDataForPin<br>(<br>const [FString](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/FString)& InResourceKey<br>) | Lookup data using a resource key that identifies the pin. | PCGComponent.h |  |
| void SetEditingMode<br>(<br>[EPCGEditorDirtyMode](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/EPCGEditorDirtyMode) InEditingMode,<br>[EPCGEditorDirtyMode](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/EPCGEditorDirtyMode) InSerializedEditingMode<br>) |  | PCGComponent.h | - BlueprintCallable<br>- Category="PCG\|Advanced" |
| void SetGenerationGridSize<br>(<br>uint32 InGenerationGridSize<br>) |  | PCGComponent.h |  |
| void SetGraph<br>(<br>[UPCGGraphInterface](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/UPCGGraphInterface)\\* InGraph<br>) |  | PCGComponent.h | - BlueprintCallable<br>- NetMulticast<br>- Reliable<br>- Category=PCG |
| void SetGraphLocal<br>(<br>[UPCGGraphInterface](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/UPCGGraphInterface)\\* InGraph<br>) |  | PCGComponent.h |  |
| void [SetIsPartitioned](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/UPCGComponent/SetIsPartitioned)<br>(<br>bool bIsNowPartitioned<br>) | Utility function (mostly for tests) to properly set the value of bIsComponentPartitioned. | PCGComponent.h |  |
| void SetPropertiesFromOriginal<br>(<br>const [UPCGComponent](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/UPCGComponent)\\* Original<br>) | Updates internal properties from other component, dirties as required but does not trigger Refresh | PCGComponent.h |  |
| void SetSchedulingPolicyClass<br>(<br>[TSubclassOf](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/CoreUObject/TSubclassOf) < [UPCGSchedulingPolicyBase](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/UPCGSchedulingPolicyBase) \> InSchedulingPolicyClass<br>) | Set the runtime generation scheduling policy type. | PCGComponent.h |  |
| bool ShouldGenerateBPPCGAddedToWorld() | Know if we need to force a generation, in case of BP added to the world in editor | PCGComponent.h |  |
| void StartGenerationInProgress() |  | PCGComponent.h |  |
| void StartIgnoringChangeOriginDuringGeneration<br>(<br>const [UObject](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/CoreUObject/UObject)\\* InChangeOriginToIgnore<br>) | For duration of the current/next generation, any change triggers from this change origin will be discarded. | PCGComponent.h |  |
| void StartIgnoringChangeOriginsDuringGeneration<br>(<br>const [TArrayView](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/TArrayView) < const [UObject](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/CoreUObject/UObject)\\* \> InChangeOriginsToIgnore<br>) |  | PCGComponent.h |  |
| void StopGenerationInProgress() |  | PCGComponent.h |  |
| void StopIgnoringChangeOriginDuringGeneration<br>(<br>const [UObject](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/CoreUObject/UObject)\\* InChangeOriginToIgnore<br>) |  | PCGComponent.h |  |
| void StopIgnoringChangeOriginsDuringGeneration<br>(<br>const [TArrayView](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/TArrayView) < const [UObject](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/CoreUObject/UObject)\\* \> InChangeOriginsToIgnore<br>) |  | PCGComponent.h |  |
| void StoreInspectionData<br>(<br>const [FPCGStack](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/FPCGStack)\\* InStack,<br>const [UPCGNode](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/UPCGNode)\\* InNode,<br>const [PCGUtils::FCallTime](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/FCallTime)\\* InTimer,<br>const [FPCGDataCollection](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/FPCGDataCollection)& InInputData,<br>const [FPCGDataCollection](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/FPCGDataCollection)& InOutputData,<br>bool bUsedCache<br>) |  | PCGComponent.h |  |
| void StoreOutputDataForPin<br>(<br>const [FString](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/FString)& InResourceKey,<br>const [FPCGDataCollection](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/FPCGDataCollection)& InData<br>) | Store data with a resource key that identifies the pin. | PCGComponent.h |  |
| bool Use2DGrid() | Returns true if component should output on a 2D Grid | PCGComponent.h |  |
| bool WasGeneratedThisSession() | Functions for managing the node inspection cache | PCGComponent.h |  |
| bool WasNodeExecuted<br>(<br>const [UPCGNode](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/UPCGNode)\\* InNode,<br>const [FPCGStack](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/FPCGStack)& Stack<br>) const |  | PCGComponent.h |  |

#### Overridden from [UActorComponent](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Engine/UActorComponent)

| Name | Remarks | Include Path | [Unreal Specifiers](https://dev.epicgames.com/documentation/unreal-engine/ufunctions-in-unreal-engine#functionspecifiers) |
| --- | --- | --- | --- |
| virtual void BeginPlay() | ~End [UObject](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/CoreUObject/UObject) interface | PCGComponent.h |  |
| virtual void EndPlay<br>(<br>const [EEndPlayReason::Type](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Engine/EEndPlayReason__Type) EndPlayReason<br>) |  | PCGComponent.h |  |
| virtual void OnComponentDestroyed<br>(<br>bool bDestroyingHierarchy<br>) |  | PCGComponent.h |  |
| virtual void OnRegister() |  | PCGComponent.h |  |
| virtual void OnUnregister() |  | PCGComponent.h |  |

#### Overridden from [UObject](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/CoreUObject/UObject)

| Name | Remarks | Include Path | [Unreal Specifiers](https://dev.epicgames.com/documentation/unreal-engine/ufunctions-in-unreal-engine#functionspecifiers) |
| --- | --- | --- | --- |
| virtual void BeginDestroy() |  | PCGComponent.h |  |
| virtual bool CanEditChange<br>(<br>const [FProperty](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/CoreUObject/FProperty)\\* InProperty<br>) const |  | PCGComponent.h |  |
| virtual bool IsEditorOnly() |  | PCGComponent.h |  |
| virtual void PostEditImport() |  | PCGComponent.h |  |
| virtual void PostInitProperties() |  | PCGComponent.h |  |
| virtual void PostLoad() |  | PCGComponent.h |  |
| virtual void PreSave<br>(<br>[FObjectPreSaveContext](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/CoreUObject/FObjectPreSaveContext) ObjectSaveContext<br>) |  | PCGComponent.h |  |
| virtual void Serialize<br>(<br>[FArchive](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/FArchive)& Ar<br>) | ~End [IPCGGraphExecutionSource](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/IPCGGraphExecutionSource) interface ~Begin [UObject](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/CoreUObject/UObject) interface | PCGComponent.h |  |

#### Overridden from [IPCGGraphExecutionSource](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/IPCGGraphExecutionSource)

| Name | Remarks | Include Path | [Unreal Specifiers](https://dev.epicgames.com/documentation/unreal-engine/ufunctions-in-unreal-engine#functionspecifiers) |
| --- | --- | --- | --- |
| virtual [IPCGGraphExecutionState](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/IPCGGraphExecutionState) & [GetExecutionState](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/UPCGComponent/GetExecutionState) () | ~Begin [IPCGGraphExecutionSource](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/IPCGGraphExecutionSource) interface | PCGComponent.h |  |
| virtual const [IPCGGraphExecutionState](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/IPCGGraphExecutionState) & [GetExecutionState](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/UPCGComponent/GetExecutionState) () |  | PCGComponent.h |  |

### Protected

| Name | Remarks | Include Path | [Unreal Specifiers](https://dev.epicgames.com/documentation/unreal-engine/ufunctions-in-unreal-engine#functionspecifiers) |
| --- | --- | --- | --- |
| [FPCGGridDescriptor](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/FPCGGridDescriptor) GetGridDescriptorInternal<br>(<br>uint32 GridSize,<br>bool bRuntimeHashUpdate<br>) const |  | PCGComponent.h |  |
| void MarkSubObjectsAsGarbage() |  | PCGComponent.h |  |
| void RefreshSchedulingPolicy() |  | PCGComponent.h |  |

#### Overridden from [UActorComponent](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Engine/UActorComponent)

| Name | Remarks | Include Path | [Unreal Specifiers](https://dev.epicgames.com/documentation/unreal-engine/ufunctions-in-unreal-engine#functionspecifiers) |
| --- | --- | --- | --- |
| virtual [TStructOnScope](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/CoreUObject/TStructOnScope) < [FActorComponentInstanceData](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Engine/FActorComponentInstanceData) \> GetComponentInstanceData() |  | PCGComponent.h |  |

### Static

| Name | Remarks | Include Path | [Unreal Specifiers](https://dev.epicgames.com/documentation/unreal-engine/ufunctions-in-unreal-engine#functionspecifiers) |
| --- | --- | --- | --- |
| static void AddReferencedObjects<br>(<br>[UObject](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/CoreUObject/UObject)\\* InThis,<br>[FReferenceCollector](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/CoreUObject/FReferenceCollector)& Collector<br>) |  | PCGComponent.h |  |
| static [UPCGData](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/UPCGData) \\* [CreateActorPCGData](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/UPCGComponent/CreateActorPCGData)<br>(<br>[AActor](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Engine/AActor)\\* Actor,<br>const [UPCGComponent](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/UPCGComponent)\\* Component,<br>bool bParseActor<br>) | Builds the canonical PCG data from a given actor and its PCG component if any. | PCGComponent.h |  |
| static [FPCGDataCollection](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/FPCGDataCollection) CreateActorPCGDataCollection<br>(<br>[AActor](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Engine/AActor)\\* Actor,<br>const [UPCGComponent](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/UPCGComponent)\\* Component,<br>[EPCGDataType](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/EPCGDataType) InDataFilter,<br>bool bParseActor,<br>bool\* bOutOptionalSanitizedTagAttributeName<br>) | Builds the PCG data from a given actor and its PCG component, and places it in a data collection with appropriate tags | PCGComponent.h |  |
| static [TArray](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/TArray) < [TSoftObjectPtr](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/CoreUObject/TSoftObjectPtr) < [AActor](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Engine/AActor) > \> GetManagedActorPaths<br>(<br>[AActor](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Engine/AActor)\\* InActor<br>) |  | PCGComponent.h |  |
| static void PurgeUnlinkedResources<br>(<br>const [AActor](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Engine/AActor)\\* InActor<br>) | Purges Actors and Components generated by PCG but are no longer managed by any PCG Component | PCGComponent.h |  |

## Deprecated Variables

| Name | Type | Remarks | Include Path | [Unreal Specifiers](https://dev.epicgames.com/documentation/unreal-engine/unreal-engine-uproperties#propertyspecifiers) |
| --- | --- | --- | --- | --- |
| Graph\_DEPRECATED | [TObjectPtr](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/TObjectPtr) < [UPCGGraph](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/PCG/UPCGGraph) > |  | PCGComponent.h |  |

* * *