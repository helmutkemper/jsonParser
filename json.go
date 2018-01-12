package jsonParser

import "strconv"

// Ajuda a extrair dados de um JSon fora de padrão.
// No exemplo abaixo, dt deveria ser numérico e não uma string
//
// var data interface{}
// s := `{ "dt": "01092017" }`
//
// if err := json.Unmarshal([]byte(s), &data); err != nil {
//   panic(err)
// }
//        |interface
//        |           |tipo
//        |           |
// parser(data, "dt", int(0))
//              |
//              |chave
//
// int:
//   parser( data, key, int(0) )
//
// bool:
//   parser( data, key, false )
//
// float64:
//   parser( data, key, float64(0) )
//
// string:
//   parser( data, key, "")
func Parser( itf interface{}, key string, tp interface{} ) interface{} {
  var retInt int64
  var retFlt float64
  var retBoo bool
  switch tp.(type) {
  case string:

    switch itf.(map[string]interface{})[key].(type) {
    case string:
      return itf.(map[string]interface{})[key].(string)
    case int:
      return strconv.Itoa( itf.(map[string]interface{})[key].(int) )
    case int64:
      return strconv.FormatInt( itf.(map[string]interface{})[key].(int64), 10 )
    case float32:
      return strconv.FormatFloat( float64( itf.(map[string]interface{})[key].(float32) ), 'E', -1, 64 )
    case float64:
      return strconv.FormatFloat( itf.(map[string]interface{})[key].(float64), 'E', -1, 64 )
    case bool:
      return strconv.FormatBool( itf.(map[string]interface{})[key].(bool) )
    }

  case int:

    switch itf.(map[string]interface{})[key].(type) {
    case string:
      retInt, _ = strconv.ParseInt( itf.(map[string]interface{})[key].(string), 10, 64 )
      return int( retInt )
    case int:
      return itf.(map[string]interface{})[key].(int)
    case int64:
      return int( itf.(map[string]interface{})[key].(int64) )
    case float32:
      return int( itf.(map[string]interface{})[key].(float32) )
    case float64:
      return int( itf.(map[string]interface{})[key].(float64) )
    case bool:
      if itf.(map[string]interface{})[key].(bool) {
        return 1
      } else {
        return 0
      }
    }

  case int64:

    switch itf.(map[string]interface{})[key].(type) {
    case string:
      retInt, _ = strconv.ParseInt( itf.(map[string]interface{})[key].(string), 10, 64 )
      return retInt
    case int:
      return int64( itf.(map[string]interface{})[key].(int) )
    case int64:
      return itf.(map[string]interface{})[key].(int64)
    case float32:
      return int64( itf.(map[string]interface{})[key].(float32) )
    case float64:
      return int64( itf.(map[string]interface{})[key].(float64) )
    case bool:
      if itf.(map[string]interface{})[key].(bool) {
        return int64( 1 )
      } else {
        return int64( 0 )
      }
    }

  case float64:

    switch itf.(map[string]interface{})[key].(type) {
    case string:
      retFlt, _ = strconv.ParseFloat( itf.(map[string]interface{})[key].(string), 64 )
      return retFlt
    case int:
      return float64( itf.(map[string]interface{})[key].(int) )
    case int64:
      return float64( itf.(map[string]interface{})[key].(int64) )
    case float32:
      return float64( itf.(map[string]interface{})[key].(float32) )
    case float64:
      return itf.(map[string]interface{})[key].(float64)
    case bool:
      if itf.(map[string]interface{})[key].(bool) {
        return float64( 1.0 )
      } else {
        return float64( 0.0 )
      }
    }

  case bool:

    switch itf.(map[string]interface{})[key].(type) {
    case string:
      retBoo, _ = strconv.ParseBool( itf.(map[string]interface{})[key].(string) )
      return retBoo
    case int:
      if itf.(map[string]interface{})[key].(int) == 0 {
        return false
      } else {
        return true
      }
    case int64:
      if itf.(map[string]interface{})[key].(int64) == 0 {
        return false
      } else {
        return true
      }
    case float32:
      if itf.(map[string]interface{})[key].(float32) == 0 {
        return false
      } else {
        return true
      }
    case float64:
      if itf.(map[string]interface{})[key].(float64) == 0 {
        return false
      } else {
        return true
      }
    case bool:
      return itf.(map[string]interface{})[key].(bool)
    }

  }

  return nil
}
