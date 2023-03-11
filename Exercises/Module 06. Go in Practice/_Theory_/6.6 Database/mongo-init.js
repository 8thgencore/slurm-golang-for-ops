db.createUser(
    {
        user: "tester",
        pwd: "testerPwd",
        roles: [
            {
                role: "readWrite",
                db: "testing"
            }
        ]

    }
)