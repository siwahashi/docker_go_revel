package models

import "time"

type User struct {
  Id                int64
  Age               int16
  Name              string
  Gmail             string
  PuMail            string
  Password          string
  CreatedAt         time.Time
  UpdatedAt         time.Time
  DeletedAt         time.Time
  LastLoginAt       time.Time
  PasswordUpdatedAt time.Time
}