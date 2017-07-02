using System.Collections;
using System.Collections.Generic;
using UnityEngine;

public class PathUtil     {



    static string  GetGameDataPath()
    {
        string path = "";
#if UNITY_EDITOR
        path = Application.dataPath + "/GameData/";
#else
         
        //or sandbox dir.
        path = Application.streamingAssetsPath + "/GameData/";
#endif
        return path;
    }

    public static string GetTableDataPath()
    {
        string tableDataPath = GetGameDataPath() + "TableData/"; 
        return tableDataPath;
    }

    public static string GetBundleDataPath()
    {
        string tableDataPath = GetGameDataPath() + "Bundles/";
        return tableDataPath;
    }

}
