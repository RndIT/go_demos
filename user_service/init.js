// admin
db = db.getSiblingDB('admin');
// move to the admin db - always created in Mongo
db.auth("mongoadm", "mongoadm")

// user
userdb = db.getSiblingDB("notification_system")
userdb.createUser({
  "user": "nsuser",
  "pwd" : "nsuser",
  "roles": [
    { "role" : "readWrite", "db" : "notification_system"}
  ],
  "mechanisms": [ "SCRAM-SHA-1" ],
  "passwordDigestor": "client"
})
userdb.auth("nsuser", "nsuser")

userdb.log.insertOne({"message": "Database created."});

userdb.createCollection("collection123123")