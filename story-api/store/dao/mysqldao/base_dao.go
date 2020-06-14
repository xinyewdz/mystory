package mysqldao

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"go.uber.org/zap"
	"reflect"
	"story-api/global"
	"story-api/util/config"
)

var(
	log = global.MainLog
	sqlMap = make(map[string]string)
	conn *sql.DB

)

func init()  {
	host := config.Get("mysqldao.host")
	user := config.Get("mysqldao.user")
	password := config.Get("mysqldao.password")
	database :=config.Get("mysqldao.database")
	url := user+":"+password+"@("+host+")/"+database+"?charset=utf8&parseTime=True&loc=Local"
	var err error
	conn,err = sql.Open("mysqldao",url)
	//conn,err = gorm.Open("mysqldao",url)
	if err!=nil{
		log.Error("get mysqldao conn error.",zap.String("url",url),zap.Error(err))
	}
}

func update(conn *sql.DB,table string,v interface{})(error){
	sql := "update "+table+" set "
	rv := reflect.ValueOf(v).Elem()
	rt := reflect.TypeOf(v).Elem()
	var id interface{}
	vals := []interface{}{}
	for idx:=0;idx<rv.NumField();idx++{
		val:=rv.Field(idx)
		t := rt.Field(idx)
		name := t.Tag.Get("db")
		if name==""{
			name = t.Name
		}
		if name=="id"||name=="Id"{
			id = val.Interface()
			continue
		}
		if len(vals)>0{
			sql += ","
		}
		vals = append(vals,val.Interface())
		sql += name+"=?"
	}
	vals = append(vals,id)
	sql += " where id=?"
	_, err :=conn.Exec(sql,vals...)
	if err!=nil{
		return err
	}
	return nil
}

func insert(conn *sql.DB,table string,v interface{})(int64,error){
	sql := "insert into "+table+"("
	rv := reflect.ValueOf(v).Elem()
	rt := reflect.TypeOf(v).Elem()
	vals := []interface{}{}
	valTag := "values("
	for idx:=0;idx<rv.NumField();idx++{
		val:=rv.Field(idx)
		if val.IsZero(){
			continue
		}
		if len(vals)>0{
			sql += ","
			valTag+=","
		}
		t := rt.Field(idx)
		vals = append(vals,val.Interface())
		name := t.Tag.Get("db")
		if name==""{
			name = t.Name
		}

		valTag +="?"
		sql += name
	}
	valTag+=")"
	sql +=")"
	sql += valTag
	result, err :=conn.Exec(sql,vals...)
	if err!=nil{
		return 0,err
	}
	id,_ := result.LastInsertId()
	return id,nil
}

func fromRows(rows *sql.Rows,v interface{})error{
	cNames,err := rows.Columns()

	rv := reflect.ValueOf(v).Elem()
	rt := reflect.TypeOf(v).Elem()
	dest := make([]interface{},len(cNames))
	vMap := make(map[string]reflect.Value)
	for idx:=0;idx<rt.NumField();idx++{
		rtt := rt.Field(idx)
		name := rtt.Tag.Get("db")
		if name==""{
			name = rtt.Name
		}
		vMap[name] = rv.Field(idx)
	}
	for idx:=0;idx<len(cNames);idx++{
		cName := cNames[idx]
		val := vMap[cName]
		dest[idx] = val.Addr().Interface()
	}
	err = rows.Scan(dest...)
	return err
}
