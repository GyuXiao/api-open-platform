-- 创建库
create database if not exists api_db;

-- 切换库
use api_db;

-- 用户表
create table if not exists `user`
(
    `id`         bigint                                                                                         not null auto_increment comment '唯一 id' primary key,
    `username`   varchar(256)                                                                                   not null comment '用户昵称',
    `password`   varchar(512)                                                                                   not null comment '用户密码',
    `avatarUrl`  varchar(1024) default 'https://gyu-pic-bucket.oss-cn-shenzhen.aliyuncs.com/gyustudio_icon.jpg' not null comment '用户头像',
    `email`      varchar(256)                                                                                   null comment '用户邮箱',
    `phone`      varchar(256)                                                                                   null comment '手机号',
    `userRole`   tinyint       default 0                                                                        not null comment '用户角色 0 - 普通用户 1 - 管理员',
    `accessKey` varchar(512) not null comment 'accessKey',
    `secretKey` varchar(512) not null comment 'secretKey',
    `isDelete`   tinyint       default 0                                                                        not null comment '是否删除 0 - 未删除 1- 删除',
    `createTime` datetime      default CURRENT_TIMESTAMP                                                      not null comment '创建时间',
    `updateTime` datetime      default CURRENT_TIMESTAMP                                                      not null on update CURRENT_TIMESTAMP comment '更新时间',
    UNIQUE KEY `username` (`username`),
    UNIQUE KEY `phone` (`phone`),
    UNIQUE KEY `email` (`email`)
    ) comment '用户表';

-- 插入初始用户

# ALTER TABLE `user` DROP INDEX `phone`;
# ALTER TABLE `user` DROP INDEX `email`;


ALTER TABLE `user` ADD INDEX `phone` (`phone`);
ALTER TABLE `user` ADD INDEX `email` (`email`);
