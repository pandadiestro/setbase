package db

import (
	"bufio"
	"io"
	"os"
	"path/filepath"
)

var data []Register
var dataHome string

func registerWalkFunc(path string, info os.DirEntry, err error) error {
    if err != nil {
        return err
    }

    if info.IsDir() || filepath.Ext(info.Name()) != ".META" {
        return nil
    }


    sets := []string{}

    metaFile, openErr := os.OpenFile(path, os.O_RDWR, 00666)
    if openErr != nil {
        return openErr
    }

    reader := bufio.NewReader(metaFile)

    origPath, _, nameErr := reader.ReadLine()
    if nameErr != nil {
        return nameErr
    }

    line, _, err := reader.ReadLine();
    for err != io.EOF {
        sets = append(sets, string(line))
        line, _, err = reader.ReadLine();
    }

    metaFile.Close()

    data = append(data, Register{
        Path: dataHome + "/" + string(origPath),
        Sets: sets,
    })

    return nil
}


func StartDb() (*DB, error) {
    dataHome = os.Getenv("setbase_data")
    traverseErr := filepath.WalkDir(dataHome, registerWalkFunc)

    if traverseErr != nil {
        return nil, traverseErr
    }

    newBase := &DB{
        ListedData: data,
    }

    return newBase, nil
}

