# manager-sql-struct
## How to use?

      //package origins
      type People struct {
  	    Id int64 `json:"id" db:"id"`
        Idxxx sql.NullInt64 `json:"id_xxx" db:"id_xxx"`
        Name sql.NullString `json:"name" db:"name"`
        UpdatedAt mysql.NullTime `json:"updated_at" db:"updated_at"`
        //etc...
      }
      
      //package destinations
      type People struct {
        Id int64 `json:"id"`
        Idxxx int64 `json:"id_xxx"`
        Name string `json:"name"`
        UpdatedAt time.Time `json:"updated_at"`
        //etc...
      }
      
      //package main
      origin := &origins.People{}
      destiny := &destinations.People{}
      err = sqlstruct.Marshall(*origin,destiny)
      
      if err != nil {
        log.Println(err)
        return response, err
      }
