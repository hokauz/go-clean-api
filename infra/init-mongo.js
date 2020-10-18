db.createUser({
    user: "cleanADMIN",
    pwd: "asd123A5q3@#4@#45",
    roles: [
        { role: "readWrite", db: "clean-api" },
    ]
}
)