using UnrealBuildTool;

public class MCPTestProjectTarget : TargetRules
{
	public MCPTestProjectTarget(TargetInfo Target) : base(Target)
	{
		Type = TargetType.Game;
		DefaultBuildSettings = BuildSettingsVersion.V6;
		IncludeOrderVersion = EngineIncludeOrderVersion.Latest;
		ExtraModuleNames.Add("MCPTestProject");
	}
}
