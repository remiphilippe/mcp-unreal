<!-- Source: https://dev.epicgames.com/documentation/en-us/unreal-engine/command-line-arguments-in-unreal-engine -->

In Unreal Engine, **Command-line Arguments**, also called **Additional Launch Parameters**, customize how the engine runs on startup. Similar to [console commands](https://dev.epicgames.com/documentation/en-us/unreal-engine/console-variables-cplusplus-in-unreal-engine), command-line arguments can be an invaluable tool for testing and optimizing your project. These settings range from high-level operations, such as forcing the **Unreal Editor** to run in game mode instead of full-editor mode, to more detailed options, such as choosing a specific map to run within your game at a particular resolution and framerate.

## Pass Command-Line Arguments

There are three common methods to pass command-line arguments to your Unreal Engine project or executable:

- From the command-line.
- From the Unreal Editor.
- From an executable shortcut.

### From the Command-Line

The general syntax for adding command-line arguments to an executable run from the command-line is:

```cpp
<EXECUTABLE> [URL_PARAMETERS] [ARGUMENTS]
```

where:

- `EXECUTABLE` is the name of your executable file (e.g. `UnrealEditor.exe`, `MyGame.exe`)
- `URL_PARAMETERS` are any optional URL parameters (e.g. `MyMap`, `/Game/Maps/BonusMaps/BonusMap.umap?game=MyGameMode`)
- `ARGUMENTS` include additional, optional command-line flags or key-value pairs (e.g. `-log`, `-game`, `-windowed`, `-ResX=400 -ResY=620`)

For example, the following input runs the `MyGame` project on the `BonusMap` in the `MyGameMode` game mode fullscreen on Windows:

```cpp
UnrealEditor.exe MyGame.uproject /Game/Maps/BonusMaps/BonusMap.umap?game=MyGameMode -game -fullscreen
```

### From the Editor

The Unreal Editor supports customizing standalone games with command-line arguments. In the Unreal Editor, command-line arguments are referred to as _additional launch parameters_. Additional launch parameters are only supported for the **Play in Standalone Game** mode.

The Unreal Editor also supports command-line arguments passed specifically to a separate, dedicated server for testing multiplayer games:

- Server Map Name Override: this is where you can pass map name as a URL parameter.
- Additional Server Game Options: this is where you can pass additional URL parameters.
- Additional Server Launch Parameters: this is where you can pass any other additional command-line flags or key-value pairs.

#### Game Arguments

To pass command-line arguments to a standalone game launched from within the Unreal Editor:

1. Navigate to **Edit > Editor Preferences**.
2. On the left-hand side, select **Level Editor > Play**.
3. On the right-hand side, find the section titled **Play in Standalone Game**.
4. In this section, there is a textbox for **Additional Launch Parameters**. Paste your command-line arguments here.

#### Server Arguments

If you have checked **Launch Separate Server** and disabled **Run Under One Process**, you can specify the **Server Map Name Override**, **Additional Server Game Options**, and **Additional Server Launch Parameters**:

1. Navigate to **Edit > Editor Preferences**.
2. On the left-hand side, select **Level Editor > Play**.
3. On the right-hand side, find the section titled **Multiplayer Options**.
4. Enable **Launch Separate Server** and disable **Run Under One Process**.
5. Navigate to **Multiplayer Options > Server**.
6. Use the three text boxes to specify different types of command-line arguments for your dedicated server.

Additional server launch parameters are only available if you choose to **Launch Separate Server** and disable **Run Under One Process**. When Run Under One Process is disabled, your clients run slower because each client spawns a separate instance of the Unreal Editor.

### From an Executable Shortcut

1. Create a shortcut to your executable.
2. Right-click on the shortcut and select **Properties**.
3. Under the **Shortcut** section, add your command-line arguments to the end of the **Target** field.
4. When you run this shortcut, the command-line arguments are passed to the original executable.

## Command-Line Arguments on Non-Desktop Platforms

To pass command-line arguments to non-desktop platforms such as consoles, mobile, and extended reality (XR); you can set the command-line by creating or editing a file titled `UECommandLine.txt`. UE automatically reads in command-line arguments from `UECommandLine.txt` upon launch. If this file does not already exist, create the file in your project's root directory and add your command-line arguments.

## Create Your Own Command-Line Arguments

Unreal Engine provides some helpful C++ functions for parsing the command-line. You can create your own command-line arguments by passing your desired flag or key-value pair to the command-line. To use command-line arguments that you have passed, you need to read them from the command-line within your code.

### Flags

Flags are switches that turn a setting on or off by their presence on the command-line. For example:

```cpp
UnrealEditor.exe MyGame.uproject -game
```

#### Parse Flags

To parse a flag from the command-line, use the `FParse::Param` function:

```cpp
bool bMyFlag = false;
if (FParse::Param(FCommandLine::Get(), TEXT("myflag")))
{
    bMyFlag = true;
}
```

### Key-Value Pairs

Key-value pairs are settings switches that specify a particular value for a switch:

```cpp
UnrealEditor.exe MyGame.uproject -game -windowed -ResX=1080 -ResY=1920
```

#### Parse Key-Value Pairs

To parse a key-value pair, use the `FParse::Value` function:

```cpp
int32 myKeyValue;
if (FParse::Value(FCommandLine::Get(), TEXT("mykey="), myKeyValue))
{
    // if the program enters this "if" statement, mykey was present on the command-line
    // myKeyValue now contains the value passed through the command-line
}
```

You can find more information about what functions are available to interact with the command-line in `CommandLine.h` located in `Engine\Source\Runtime\Core\Public\Misc`.

## Customize Engine Configuration from the Command-Line

Engine configuration is normally set in engine configuration `.ini` files. You can also customize engine configuration from the command-line. See the [Configuration Files](https://dev.epicgames.com/documentation/en-us/unreal-engine/configuration-files-in-unreal-engine#overrideconfigurationfromthecommand-line) documentation for more information.

## Customize Console Commands from the Command-Line

Console commands are normally executed from the console in the Unreal Editor. You can also customize console commands from the command-line. See the [Console Variables](https://dev.epicgames.com/documentation/en-us/unreal-engine/console-variables-cplusplus-in-unreal-engine#commandline) documentation for more information.

## Command-Line Arguments Reference

### URL Parameters

URL parameters force your game to load a specific map upon startup. URL parameters are optional, but if you do provide them, they must immediately follow the executable name or any mode flag if one is present.

URL parameters consist of two parts:

- A map name or server IP address.
- A series of optional additional parameters.

#### Map Name

A map name can refer to any map located within the Maps directory. You can optionally include the `.umap` file extension. To load a map not found in the Maps directory, you must use an absolute path or a path relative to the maps directory. In this case, the `.umap` file extension is required.

#### Server IP Address

You can use a server IP address as a URL parameter to connect a game client to a dedicated server.

#### Additional Parameters

You can specify additional parameter options by appending them to the map name or the server IP address. Each option is prefaced by a `?` (question mark) and set with `=` (equality or assignment). Prepending an option with `-` (dash) removes that option from the cached URL options.

#### Examples

##### Open Game with Map Located in Maps Directory

```cpp
MyGame.exe MyMap
```

##### Open Game with Map Located Outside Maps Directory

```cpp
MyGame.exe /Game/Maps/BonusMaps/BonusMap.umap
```

##### Open Game in Unreal Editor with Map Located Outside Maps Directory

```cpp
UnrealEditor.exe MyGame.uproject /Game/Maps/BonusMaps/BonusMap.umap?game=MyGameMode -game
```

##### Connect a Game Client to a Dedicated Server

```cpp
UnrealEditor.exe MyGame.uproject /Game/Maps/BonusMaps/BonusMap.umap -server -port=7777 -log
UnrealEditor.exe MyGame.uproject 127.0.0.1:7777 -game -log
```

### Flags and Key-Value Pairs

#### Read Command-Line Arguments from a File

You can store your command-line arguments in a text file and pass this file in the command-line for convenience:

```cpp
<EXECUTABLE> -CmdLineFile=ABSOLUTE\PATH\TO\FILE.txt
```

For a list of all available command-line arguments, see the [Command-Line Arguments Reference](https://dev.epicgames.com/documentation/en-us/unreal-engine/unreal-engine-command-line-arguments-reference).
