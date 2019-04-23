### tool
    Gorm两表查询
    type User struct{
        ID          uint
        UserTypeId  uint
        UserType    UserType   `gorm:"ForeignKey:UserTypeId"`
    }
    type UserType struct{
        ID          uint
        Name        string
    }
    var userSlice []User
    db.Preload("UserType").Find(&userSlice)