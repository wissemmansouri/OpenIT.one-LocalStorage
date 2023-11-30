package google_drive

import (
         "github.com/wissemmansouri/OpenIT.one-LocalStorage/internal/driver"
         "github.com/wissemmansouri/OpenIT.one-LocalStorage/internal/op"
)

const ICONURL = "./img/driver/GoogleDrive.svg"
const CLIENTID = ""            // Must put the real ID
const CLIENTSECRET = ""        // Must put the real Secret 

type Addition struct {
        driver.RootID  
        RefreshToken    string `json:"refresh_token" required:"true" omit:"true"`
        OrderBy         string `json:"order_by" type:"string" help:"such as: folder,name,modifiedTime" omi:"true"`
        OrderDirection  string `json:"order_direction" type:"select" options:"asc,desc" omit:"true"`
        ClientID        string `json:"client_id" required:"true" default:"54545454545454554544urhud-yrteetehjd5555jdj.apps.googleusercontent.com" omit:"true"` //must change the real clientID
        ClientSecret    string `json:"client_secret" required:"true" default:"GOCSPX-v-bJkdjdjfhfudj4MDLD5" omit:"true"` //must change the real Client Secret
        ChunkSize       int64  `json:"chunk_size" type:"number" help:"chunk size while uploading (unit: MB)" omit:"true"`
        AuthUrl         string `json:"auth_url" type:"string" default:"https://accounts.google.com/o/auth2"` // must change the real default 
        Icon            string `json:"icon" type:"string" default:"./img/driver/GoogleDrive.svg"`
        Code            string `json:"code" type:"string" help:"code from auth_url" omit:"true"`

}

var config = driver.Config{
        Name:        "GoogleDrive",
        OnlyProxy:   true,
        DefaultRoot: "root",
}

func init() {
        op.RegisterDriver(func() driver.Driver {
                return &GoogleDrive{}
        })
}
         
