using System.Collections.Generic;
using System.IO;
using System.Text;
using UnityEngine;

public class TestTable
{
    public int ID;
    public string Name;
    public int HP;
    public string AP_Name;
    public  TestTable(string line)
    {
        string[] fileds = line.Split('	');
        ID =int.Parse(fileds[0]);
        Name =fileds[1];
        HP =int.Parse(fileds[2]);
        AP_Name =fileds[3];
     }
}

public class TestTableManager
{
    Dictionary<int, TestTable> data = new Dictionary<int, TestTable>();

    public void InitTable()
    {
        string tableDir = PathUtil.GetTableDataPath();
        string tableName = "TestTable";
        string tableDataPath = string.Format("{0}{1}.txt", tableDir, tableName);
        StreamReader sr = new StreamReader(tableDataPath, Encoding.UTF8);
        string line;
        while ((line = sr.ReadLine()) != null)
        {
            line = line.Trim();
            if (line.Length > 0)
            {
                TestTable rowData = new TestTable(line);
                data.Add(rowData.ID, rowData);
}
            else
            {
                continue;
            }

        }
    }

    public TestTable GetDataByID(int id)
    {
        TestTable rowData = null;
        data.TryGetValue(id, out rowData);
        return rowData;
    }

    private TestTableManager() { }

    public static readonly TestTableManager Instance = new TestTableManager();

} 