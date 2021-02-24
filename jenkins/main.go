import(
"fmt"
"os"
 
gocb"gopkg.in/couchbase/gocb.v1"
)
 
typeBucketInterfaceinterface{
Get(key string,valueinterface{})(gocb.Cas,error)
Insert(key string,valueinterface{},expiry uint32)(gocb.Cas,error)
}
 
typeDatabasestruct{
bucket BucketInterface
}
 
typePersonstruct{
Type      string`json:"type"`
Firstname string`json:"firstname"`
Lastname  string`json:"lastname"`
}
 
func(dDatabase)GetPersonDocument(key string)(interface{},error){
vardatainterface{}
_,err:=d.bucket.Get(key,&data)
iferr!=nil{
returnnil,err
}
returndata,nil
}
 
func(dDatabase)CreatePersonDocument(key string,datainterface{})(interface{},error){
_,err:=d.bucket.Insert(key,data,0)
iferr!=nil{
returnnil,err
}
returndata,nil
}
 
func main(){
fmt.Println("Starting the application...")
vardatabase Database
cluster,_:=gocb.Connect("couchbase://"+os.Getenv("DB_HOST"))
cluster.Authenticate(gocb.PasswordAuthenticator{Username:os.Getenv("DB_USER"),Password:os.Getenv("DB_PASS")})
database.bucket,_=cluster.OpenBucket(os.Getenv("DB_BUCKET"),"")
fmt.Println(database.GetPersonDocument("8eaf1065-5bc7-49b5-8f04-c6a33472d9d5"))
database.CreatePersonDocument("blawson",Person{Type:"person",Firstname:"Brett",Lastname:"Lawson"})
}
