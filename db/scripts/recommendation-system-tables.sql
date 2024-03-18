CREATE TABLE IF NOT EXISTS `users` (
    `uid` bigint unsigned NOT NULL AUTO_INCREMENT ,
    `email` varchar(255),
    `password_hashed` varchar(255) ,
    `confirm` integer,
    `timestamp` bigint,
    PRIMARY KEY (`uid`)
);

create index users_email_index on users(email DESC);

CREATE TABLE IF NOT EXISTS `recommendations` (
    `uid` bigint unsigned NOT NULL AUTO_INCREMENT ,
    `promotion_messages` varchar(255),
    PRIMARY KEY (`uid`)
);

