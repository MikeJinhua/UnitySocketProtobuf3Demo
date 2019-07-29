from const import cs_table_file_dir


def GenCSTableManagerFile(tableName, fieldsIndex, table):
    filePath = cs_table_file_dir + tableName + ".cs"

    fileContent = ""
    fileContent += \
    'using System.Collections.Generic;\n'\
    'using System.IO;\n'\
    'using System.Text;\n'\
    'using UnityEngine;\n'\
    '\n'

    ##############TableClass#############

    fileContent += \
        'public class ' + tableName + "\n"
    fileContent += "{\n"
    for index in fieldsIndex:
        fieldtype = table.cell(2, index).value
        fieldtype = fieldtype.lower()
        if fieldtype == "Map[String]Int".lower():
            fieldtype = "Dictionary<string,int>"
        fieldName = table.cell(3, index).value
        fileContent += \
            "    public " + fieldtype + " " + fieldName + ";\n"

    fileContent += \
        "    public " + tableName + "(string line)\n"
    fileContent += \
        "    {\n"
    fileContent += \
        "        string[] fileds = line.Split('\t');\n"

    for index in fieldsIndex:
        fieldtype = table.cell(2, index).value
        fieldtype = fieldtype.lower()
        # if fieldtype == "Map[String]Int".lower():
        #     fieldtype = "public Dictionary<string,int>"
        fieldName = table.cell(3, index).value
        index = str(index)
        if fieldtype == "int[]".lower():
            fileContent += \
                "        " + fieldName + " = TableUtil.LoadInts(fileds[" + index + "]);\n"
        elif fieldtype == "Map[String]Int".lower():
            fileContent += \
                "        " + fieldName + " = TableUtil.LoadDictString2Int(fileds[" + index + "]);\n"
        elif fieldtype == "int" or fieldtype == "float":
            fileContent += \
                "        " + fieldName + " = " + fieldtype + ".Parse(fileds[" + index + "]);\n"
        elif fieldtype == "string":
            fileContent += \
                "        " + fieldName + " = fileds[" + index + "];\n"

    fileContent += \
        "     }\n"

    fileContent += \
        "}\n"
    fileContent += \
        "\n"
    ##############TableClass Manager#############

    fileContent += \
        "public class " + tableName + "Manager\n"
    fileContent += \
        '{\n' \
        '    Dictionary<int, ' + tableName + '> data = new Dictionary<int, ' + tableName + '>();\n' \
        "\n" \
        '    public void InitTable()\n' \
        '    {\n'\
        '        string tableDir = PathUtil.GetTableDataPath();\n'
    fileContent += \
        '        string tableName = \"' +tableName +"\";\n"

    fileContent += \
        '''        string tableDataPath = string.Format("{0}{1}.txt", tableDir, tableName);
        StreamReader sr = new StreamReader(tableDataPath, Encoding.UTF8);
        string line;
        while ((line = sr.ReadLine()) != null)
        {
            line = line.Trim();
            if (line.Length > 0)
            {
                ''' + tableName + ' rowData = new ' + tableName + '(line);\n'''
    fileContent += \
    "                data.Add(rowData."+ table.cell(3, 0).value + ", rowData);\n"
    fileContent += \
        '''            }
            else
            {
                continue;
            }

        }
    }\n\n'''

    keyType = table.cell(2, 0).value
    fileContent += \
        '''    public ''' + tableName + ''' GetDataByID(''' + keyType.lower() + ''' id)\n'''
    fileContent += \
'''    {
        ''' + tableName + ''' rowData = null;
        data.TryGetValue(id, out rowData);
        return rowData;
    }\n\n'''

    keyType = table.cell(2, 1).value
    fileContent += \
        '''    public ''' + tableName + ''' GetDataByName(''' + keyType.lower() + ''' name)\n'''
    fileContent += \
'''    {
        foreach(KeyValuePair<int, ''' + tableName + '''> kp in data)
        {
            if (kp.Value.Name == name) return kp.Value;
        }
        return null;
    }

    private ''' + tableName + '''Manager() { }

    public static readonly ''' + tableName + '''Manager Instance = new ''' + tableName + '''Manager();

} '''

    fo = open(filePath, "wb")
    fo.write(fileContent.encode('utf-8'))
    fo.close()






def genCSLoadTablesFile(files):
    filePath = cs_table_file_dir + "LoadTables.cs"
    fileContent = ""
    fileContent += \
'''
using System.Collections;
using System.Collections.Generic;
using UnityEngine;

public class LoadTables
{

    /// <summary>
    /// Init ALL tables
    /// </summary>
	public static void Init()
    {\n'''
    for file in files:
        fileContent +="        "+file+"Manager.Instance.InitTable();\n"

    fileContent += \
'''    }


}

'''
    fo = open(filePath, "wb")
    fo.write(fileContent.encode('utf-8'))
    fo.close()














