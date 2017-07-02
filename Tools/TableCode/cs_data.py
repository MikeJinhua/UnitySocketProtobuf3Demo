from const import cs_table_data_dir


def GenCSTableData(tableName, fieldsIndex, table):
    filePath = cs_table_data_dir + tableName + ".txt"

    fileContent = ""


    for row in range(4,table.nrows):
        for ncols in range(table.ncols):
            if ncols in fieldsIndex:
                fieldType = table.cell(2, ncols).value
                if fieldType == "Int":
                    value = table.cell(row, ncols).value
                    fileContent += ("{0}\t").format(int(table.cell(row, ncols).value))
                else:
                    fileContent += ("{0}\t").format(table.cell(row, ncols).value)

        fileContent += "\n"

    fo = open(filePath, "wb")
    fo.write(fileContent.encode('utf-8'))
    fo.close()







