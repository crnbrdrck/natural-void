package main

import (
    "./go"
)

// Simple seed script to prepopulate the DB
func main() {
    // Generate a DAO
    dao := naturalvoid.DAO{}
    err := dao.New()
    if err != nil {
        panic(err)
    }
    // Get the prelim data and put it in the DB
    // Migrate the models
    dao.DB.AutoMigrate(&naturalvoid.User{})
    dao.DB.AutoMigrate(&naturalvoid.Story{})
    dao.DB.AutoMigrate(&naturalvoid.Episode{})

    // Create initial user to reflect me from the LDAP server
    user := naturalvoid.User{
        UserName: "crnbrdrck",
    }
    dao.DB.Create(&user)

    // Create a story record for Dread
    dread := naturalvoid.Story{
        Name:      "A Simple Trip to Waterdeep: Dread",
        ShortName: "ASTTW: Dread",
        Slug:      "asttw-dread",
        Description: []string{
            "Chapter 1 of our 5 chapter epic which follows our heroes Bran, Gundham, Lyra, Jake, and Viper on their respective trips to Waterdeep.",
            "During one night of particularly heavy fog, these five people and their coach driver get taken to the land of Barovia, ruled by Strahd von Zarovich.",
            "Will this group of 5 strangers be able to band together, overcome this situation and rescue their coach driver? Only time will tell.",
        },
        UserID: user.ID,
    }
    dao.DB.Create(&dread)
}
