CREATE TABLE subscription (
  server varchar(64) NOT NULL,
  topic varchar(64) NOT NULL,
  online boolean DEFAULT 0 
);

CREATE TABLE user (
    name varchar(64) NOT NULL
);

CREATE TABLE user_subscription (
    user_id int NOT NULL REFERENCES user(rowid),
    subscription_id int NOT NULL REFERENCES subscription(rowid)
);

CREATE UNIQUE INDEX user_subscription ON user_subscription(user_id, subscription_id);


INSERT INTO user VALUES ("minh");
INSERT INTO subscription VALUES ("tcp://localhost:1883", "test/topic", False)