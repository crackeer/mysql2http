package container

import (
    "gorm.io/gorm"
    "gorm.io/driver/mysql"
    "sync"
)

type ModelParse func() interface{} 
type ModelListParse func() interface{} 


var databaseConnecttion map[string]*gorm.DB = map[string]*gorm.DB{}

var modelParseFuncMap map[string]map[string]ModelParse = map[string]map[string]ModelParse{}

var modelListParseFuncMap map[string]map[string]ModelListParse = map[string]map[string]ModelListParse{}

var locker *sync.Mutex

{% autoescape off %}
var databaseDSN map[string]string = map[string]string{
{% for item in databases %} 
    "{{item.database}}" : "{{item.dsn}}",
{% endfor %}
}
{% endautoescape %}


func init() {
    databaseConnecttion = map[string]*gorm.DB{}
    locker = &sync.Mutex{}
    for name, dsn := range databaseDSN {
        db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
        if err != nil {
            panic(err)
        }
        databaseConnecttion[name] = db
    }
}

func GetGormDB(name string) *gorm.DB {
    if value, ok := databaseConnecttion[name]; ok {
        return value
    }
    return nil
}

func RegisterModelParseFunc(database, table string, someFunc ModelParse) {
    locker.Lock()
    defer locker.Unlock()
    if _, ok := modelParseFuncMap[database]; !ok {
        modelParseFuncMap[database] = map[string]ModelParse{}
    }
    modelParseFuncMap[database][table] = someFunc
}

func RegisterModelListParseFunc(database, table string, someFunc ModelListParse) {
    locker.Lock()
    defer locker.Unlock()
    if _, ok := modelListParseFuncMap[database]; !ok {
        modelListParseFuncMap[database] = map[string]ModelListParse{}
    }
    modelListParseFuncMap[database][table] = someFunc
}

func GetModelObject(database, table string) interface{} {
    if tmpMap, ok := modelParseFuncMap[database]; ok {
        if someFunc, ok := tmpMap[table]; ok {
            return someFunc()
        }
    }
    return nil
}

func GetModelListObject(database, table string) interface{} {
    if tmpMap, ok := modelListParseFuncMap[database]; ok {
        if someFunc, ok := tmpMap[table]; ok {
            return someFunc()
        }
    }
    return nil
}


