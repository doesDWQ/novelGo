CREATE TABLE `user`
(
    `id` varchar(255) default "" NOT NULL COMMENT '唯一主键',
    `name` varchar(255) NOT NULL default "" COMMENT '姓名',
    `gender` int(11) NOT NULL default 0 COMMENT '性别',
    `email` varchar(255) NOT NULL default "" COMMENT '邮箱',
    PRIMARY KEY(`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;