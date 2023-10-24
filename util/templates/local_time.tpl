package define

import (
    "bytes"
	"database/sql/driver"
	"fmt"
	"time"
)

type LocalTime struct {
   time.Time
}

var ExTimeUnmarshalTimeFormat = "2006-01-02 15:04:05"
var ExTimeMarshalTimeFormat = "2006-01-02 15:04:05"

func (t LocalTime) UnmarshalTimeFormat() string {
   return ExTimeUnmarshalTimeFormat
}

func (t LocalTime) MarshalTimeFormat() string {
   return ExTimeMarshalTimeFormat
}

func (t *LocalTime) UnmarshalJSON(b []byte) error {
   b = bytes.Trim(b, "\"")   // 此除需要去掉传入的数据的两端的 ""
   ext, err := time.Parse(t.UnmarshalTimeFormat(), string(b))
   if err != nil {
      return err
   }
   *t = LocalTime{ext}
   return nil
}

// MarshalJSON 2. 为 Xtime 重写 MarshaJSON 方法，在此方法中实现自定义格式的转换；
func (t LocalTime) MarshalJSON() ([]byte, error) {
   output := fmt.Sprintf("\"%s\"", t.Format("2006-01-02 15:04:05"))
   return []byte(output), nil
}

// Value 3. 为 Xtime 实现 Value 方法，写入数据库时会调用该方法将自定义时间类型转换并写入数据库；
func (t LocalTime) Value() (driver.Value, error) {
   var zeroTime time.Time
   if t.Time.UnixNano() == zeroTime.UnixNano() {
      return nil, nil
   }
   return t.Time, nil
}

// Scan 4. 为 Xtime 实现 Scan 方法，读取数据库时会调用该方法将时间数据转换成自定义时间类型；
func (t *LocalTime) Scan(v interface{}) error {
   value, ok := v.(time.Time)
   if ok {
      *t = LocalTime{Time: value}
      return nil
   }
   return fmt.Errorf("can not convert %v to timestamp", v)
}