///Turret Alarm V1
///Created by Foltast
///Version: 2021.1.1
/// <summary>
/// Can be changed by users
/// </summary>
string beaconName = "Beacon"; // Name of the beacon
string lcdTag = "[TA LCD]"; //Tag for LCD panels intended for displaying information

// Set 'false' for disable auto refreshing or 'true' to enable
bool refreshEnabled = true;

//Override script mode. Can be set to 'AUTO' ('0'), 'VANILLA' ('1'), 'WEAPONCORE' ('2'). By default set as '0'
int forcedModeOverride = 0;

//LCD Panels settings
string lcdFontFamily = "Monospace";
float lcdFontSize = 0.8f;
Color lcdFontColor = new Color(255,130,0);
TextAlignment lcdFontAlignment = TextAlignment.CENTER;

/// <summary>
/// Do not touch below this line  
/// </summary>
/// -------------------------------------------------------------------- ///
IMyBeacon beacon;
static WcPbApi api;

List<IMyTerminalBlock> moddedTurrets = new List<IMyTerminalBlock>();
List<IMyLargeTurretBase> vanillaTurrets = new List<IMyLargeTurretBase>();
List<MyDefinitionId> tempIds = new List<MyDefinitionId>();

IMyTextSurface[] lcds = null;

AlarmStatus currentStatus;
bool isWCUsed = false; 
bool isHaveTarget = false;

string version = "2021.1.1";
string refreshPass;

int currentMode;
int checkRate = 300;
int currentCheckPass = 0;

int maxAttempts = 150;
int currentAttempt = 0;

public Program()
{ 
    Runtime.UpdateFrequency = UpdateFrequency.Update1; 
    api = new WcPbApi(); 
}

private void Initialize()
{
    isWCUsed = api.Activate(Me);

    CheckBeacon();
    SearchLCDs();

    if (isWCUsed)
        SearchModdedTurrets();
    else
        SearchVanillaTurrets();
}

public void Main()
{
    Echo($"Turret Alarm v1 by Foltast\nVersion: {version}\n");

    if (currentAttempt <= maxAttempts)
    {
        currentAttempt++;

        if (isWCUsed == false && currentAttempt < maxAttempts)
            Initialize();
    }

    if (refreshEnabled)
    {
        refreshPass = (currentCheckPass / 100 + 1).ToString();
        Echo($"Next refresh in: {refreshPass}");
        currentCheckPass--;
        CheckBeacon();
        SearchLCDs();
    }

    isHaveTarget = false;

    switch(forcedModeOverride)
    {
        case 0:
            currentStatus = isWCUsed ? GetStatusWC() : GetStatusVanilla();
            break;
        case 1:
            Echo("CURRENT MODE IS FORCED");
            currentStatus = GetStatusVanilla();
            break;
        case 2:
            Echo("CURRENT MODE IS FORCED");
            currentStatus = GetStatusWC();
            break;
    }

    if(beacon != null)
        beacon.Enabled = isHaveTarget;

    UpdateLCDs();
}

private AlarmStatus GetStatusVanilla() 
{
    Echo("Vanilla mode\n");

    currentMode = 1;

    if(refreshEnabled)
        SearchVanillaTurrets();
 
    if (beacon == null) 
    { 
        Echo("ERROR: No beacon detected");
        return AlarmStatus.idle; 
    } 
 
    if (vanillaTurrets.Count < 1) 
    { 
        Echo("ERROR: No turrets detected");
        return AlarmStatus.idle; 
    } 
 
    foreach (var turret in vanillaTurrets) 
    { 
        if (turret.IsShooting) 
        { 
            isHaveTarget = true; 
            Echo("Status: Target detected"); 
            return AlarmStatus.detected; 
        } 
    }

    Echo("Status: Waiting targets");
    return AlarmStatus.idle; 
}

private AlarmStatus GetStatusWC() 
{
    Echo("WeaponCore mode\n");

    currentMode = 2;

    if(refreshEnabled)
        SearchModdedTurrets();

    if (beacon == null || moddedTurrets.Count < 1) 
    { 
        Echo("ERROR: No turrets or beacon detected"); 
        return AlarmStatus.idle; 
    } 
 
    foreach (var turret in moddedTurrets) 
    {
        MyDetectedEntityInfo? entity = api.GetWeaponTarget(turret, 0);

        if (!entity.Value.IsEmpty())
        {
            isHaveTarget = true;
            Echo("Status: Target detected");
            return AlarmStatus.detected;
        }
        //if(api.GetHeatLevel(turret) > 0)
        //{
        //    isHaveTarget = true; 
        //    Echo("Target detected"); 
        //    return AlarmStatus.detected; 
        //}
    }

    Echo("Status: Waiting targets");
    return AlarmStatus.idle; 
}

public List<IMyTextPanel> GetLCDsWithTag(string tag)
{
    List<IMyTextPanel> textPanels = new List<IMyTextPanel>();
    
    return textPanels;
}

private void CheckBeacon()
{
    if (currentCheckPass <= 0)
    {
        beacon = GridTerminalSystem.GetBlockWithName(beaconName) as IMyBeacon;
    }
}

private void SearchModdedTurrets()
{
    if (currentCheckPass <= 0)
    {
        api.GetAllCoreTurrets(tempIds);
        List<string> defSubIds = new List<string>();
        tempIds.ForEach(x => defSubIds.Add(x.SubtypeName));
        GridTerminalSystem.GetBlocksOfType<IMyTerminalBlock>(moddedTurrets, b => b.CubeGrid == Me.CubeGrid && defSubIds.Contains(b.BlockDefinition.SubtypeName));

        currentCheckPass = checkRate;
    }
}

private void SearchVanillaTurrets()
{
    if (currentCheckPass <= 0)
    {
        GridTerminalSystem.GetBlocksOfType(vanillaTurrets, b => b.CubeGrid == Me.CubeGrid);
        currentCheckPass = checkRate;
    }
}

private void UpdateLCDs()
{
    foreach (var lcd in lcds)
    {
        lcd.WriteText($"TURRET ALARM INFO PANEL\n\nV1 version: {version}\n\nCurrent mode: {(forcedModeOverride != 0 ? "(OVRD)" : "")}{(currentMode < 2 ? "Vanilla" : "WeaponCore")}\n\nCurrent Status: {currentStatus.ToString().ToUpper()}{(refreshEnabled ? "\n\nRefresh in: " + (currentCheckPass > 50 ? refreshPass : "progress") : "\n\nRefresh is DISABLED")}", false);
    }
}

private void SearchLCDs()
{
    List<IMyTerminalBlock> tmp_lcds = new List<IMyTerminalBlock>();
    GridTerminalSystem.GetBlocksOfType<IMyTerminalBlock>(tmp_lcds, b => b.CubeGrid == Me.CubeGrid && ((b is IMyTextSurfaceProvider && (b as IMyTextSurfaceProvider).SurfaceCount > 0) || b is IMyTextSurface) && b.CustomName.StartsWith(lcdTag));

    lcds = new IMyTextSurface[tmp_lcds.Count];

    for (int i = tmp_lcds.Count; i-- > 0;)
    {
        if (tmp_lcds[i] is IMyTextSurfaceProvider)
        {
            bool cust_si = false;
            if (tmp_lcds[i].CustomName[lcdTag.Length] == '[' && tmp_lcds[i].CustomName[lcdTag.Length + 2] == ']')
            {
                int srf_idx = (int)tmp_lcds[i].CustomName[lcdTag.Length + 1] - 48;
                if ((cust_si = srf_idx > 0 && srf_idx < 10 && (tmp_lcds[i] as IMyTextSurfaceProvider).SurfaceCount > srf_idx)) lcds[i] = ((IMyTextSurfaceProvider)tmp_lcds[i]).GetSurface(srf_idx);
            }
            if (!cust_si) lcds[i] = ((IMyTextSurfaceProvider)tmp_lcds[i]).GetSurface(0);
        }
        else lcds[i] = (IMyTextSurface)tmp_lcds[i];

        lcds[i].ContentType = (ContentType)1;
        lcds[i].Font = lcdFontFamily;
        lcds[i].FontSize = lcdFontSize;
        lcds[i].FontColor = lcdFontColor;
        lcds[i].Alignment = lcdFontAlignment;
        lcds[i].ContentType = ContentType.TEXT_AND_IMAGE;
    }
}

public class WcPbApi
{
    private Action<ICollection<MyDefinitionId>> _getCoreWeapons;
    private Func<Sandbox.ModAPI.Ingame.IMyTerminalBlock, float> _getHeatLevel;
    private Action<ICollection<MyDefinitionId>> _getCoreTurrets;

    private Func<Sandbox.ModAPI.Ingame.IMyTerminalBlock, int, Sandbox.ModAPI.Ingame.MyDetectedEntityInfo> _getWeaponTarget;

    public bool Activate(IMyTerminalBlock pbBlock)
    {
        var dict = pbBlock.GetProperty("WcPbAPI")?.As<Dictionary<string, Delegate>>().GetValue(pbBlock);
        if (dict == null) return false;
        return ApiAssign(dict);
    }

    public bool ApiAssign(IReadOnlyDictionary<string, Delegate> delegates)
    {
        if (delegates == null)
            return false;

        AssignMethod(delegates, "GetHeatLevel", ref _getHeatLevel);
        AssignMethod(delegates, "GetCoreWeapons", ref _getCoreWeapons);
        AssignMethod(delegates, "GetCoreTurrets", ref _getCoreTurrets);
        AssignMethod(delegates, "GetWeaponTarget", ref _getWeaponTarget);
        return true;
    }

    private void AssignMethod<T>(IReadOnlyDictionary<string, Delegate> delegates, string name, ref T field) where T : class
    {
        if (delegates == null)
        {
            field = null;
            return;
        }
        Delegate del;
        if (!delegates.TryGetValue(name, out del))
            throw new Exception($"{GetType().Name} :: Couldn't find {name} delegate of type {typeof(T)}");
        field = del as T;
        if (field == null)
            throw new Exception(
                $"{GetType().Name} :: Delegate {name} is not type {typeof(T)}, instead it's: {del.GetType()}");
    }

    public MyDetectedEntityInfo? GetWeaponTarget(Sandbox.ModAPI.Ingame.IMyTerminalBlock weapon, int weaponId = 0) =>
    _getWeaponTarget?.Invoke(weapon, weaponId);

    public float GetHeatLevel(Sandbox.ModAPI.Ingame.IMyTerminalBlock weapon) => _getHeatLevel?.Invoke(weapon) ?? 0f;

    public void GetAllCoreWeapons(ICollection<MyDefinitionId> collection) => _getCoreWeapons?.Invoke(collection);

    public void GetAllCoreTurrets(ICollection<MyDefinitionId> collection) => _getCoreTurrets?.Invoke(collection);

}
 
public enum AlarmStatus 
{ 
    idle, 
    detected, 
    lost 
}