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
    `isDelete`   tinyint       default 0                                                                        not null comment '是否删除 0 - 未删除 1- 删除',
    `createTime` datetime      default CURRENT_TIMESTAMP                                                      not null comment '创建时间',
    `updateTime` datetime      default CURRENT_TIMESTAMP                                                      not null on update CURRENT_TIMESTAMP comment '更新时间',
    UNIQUE KEY `username` (`username`),
    UNIQUE KEY `phone` (`phone`),
    UNIQUE KEY `email` (`email`)
    ) comment '用户表';

-- 插入初始用户
insert into `user` (`id`, `username`, `password`, `avatarUrl`, `email`, `phone`, `userRole`) values (1, '刘智辉', 'au', 'LKi', 'sammie.wiegand@gmail.com', '15059363515', 0);
insert into `user` (`id`, `username`, `password`, `avatarUrl`, `email`, `phone`, `userRole`) values (2, '方胤祥', 'zu', 'rsJjp', 'clyde.watsica@gmail.com', '15229432233', 0);
insert into `user` (`id`, `username`, `password`, `avatarUrl`, `email`, `phone`, `userRole`) values (3, '姚风华', 'Oi4', 'filTI', 'mckenzie.bradtke@hotmail.com', '15711273230', 0);
insert into `user` (`id`, `username`, `password`, `avatarUrl`, `email`, `phone`, `userRole`) values (4, '朱苑博', 'mGr', 'wI0', 'francesco.morissette@yahoo.com', '13377703453', 0);
insert into `user` (`id`, `username`, `password`, `avatarUrl`, `email`, `phone`, `userRole`) values (5, '金笑愚', 'Gqi9', 'xM', 'long.kohler@gmail.com', '17589086736', 0);
insert into `user` (`id`, `username`, `password`, `avatarUrl`, `email`, `phone`, `userRole`) values (6, '邵锦程', '4nr', 'FS6O', 'willette.quigley@hotmail.com', '14544482997', 0);
insert into `user` (`id`, `username`, `password`, `avatarUrl`, `email`, `phone`, `userRole`) values (7, '江耀杰', 'Kjp', 'b9aph', 'keith.kovacek@hotmail.com', '15823058762', 0);
insert into `user` (`id`, `username`, `password`, `avatarUrl`, `email`, `phone`, `userRole`) values (8, '刘果', 'Kus', 'LjM', 'sadye.corwin@gmail.com', '17666134068', 0);
insert into `user` (`id`, `username`, `password`, `avatarUrl`, `email`, `phone`, `userRole`) values (9, '廖思聪', '7BN5', 'yVGQ', 'colby.bode@hotmail.com', '14777401876', 0);
insert into `user` (`id`, `username`, `password`, `avatarUrl`, `email`, `phone`, `userRole`) values (10, '马子轩', 'sscF', 'pHNvs', 'moises.frami@yahoo.com', '17519358052', 0);
insert into `user` (`id`, `username`, `password`, `avatarUrl`, `email`, `phone`, `userRole`) values (11, '唐建辉', 'ZU', '3U', 'loyd.marks@yahoo.com', '17535626868', 0);
insert into `user` (`id`, `username`, `password`, `avatarUrl`, `email`, `phone`, `userRole`) values (12, '贾子涵', 'Bs', 'A8ya', 'nenita.hodkiewicz@yahoo.com', '17735009690', 0);
insert into `user` (`id`, `username`, `password`, `avatarUrl`, `email`, `phone`, `userRole`) values (13, '龚驰', 'kA2', 'Gh', 'conception.prosacco@yahoo.com', '17695326656', 0);
insert into `user` (`id`, `username`, `password`, `avatarUrl`, `email`, `phone`, `userRole`) values (14, '熊旭尧', 'sXBX', '16gYh', 'rhett.hansen@hotmail.com', '15843123828', 0);
insert into `user` (`id`, `username`, `password`, `avatarUrl`, `email`, `phone`, `userRole`) values (15, '武峻熙', '6FAHa', 'cX9', 'kamala.kertzmann@yahoo.com', '13401224036', 0);
insert into `user` (`id`, `username`, `password`, `avatarUrl`, `email`, `phone`, `userRole`) values (16, '姚健雄', 'k2z', 'Edg7', 'randy.treutel@yahoo.com', '15196328078', 0);
insert into `user` (`id`, `username`, `password`, `avatarUrl`, `email`, `phone`, `userRole`) values (17, '萧君浩', 'Jd', 'DRn3A', 'andre.moen@hotmail.com', '15884421611', 0);
insert into `user` (`id`, `username`, `password`, `avatarUrl`, `email`, `phone`, `userRole`) values (18, '阎伟宸', 'dHG', 'G2BZr', 'shane.mosciski@hotmail.com', '15961620099', 0);
insert into `user` (`id`, `username`, `password`, `avatarUrl`, `email`, `phone`, `userRole`) values (19, '王展鹏', 'uM', 'a6', 'dulce.effertz@hotmail.com', '17377514203', 0);
insert into `user` (`id`, `username`, `password`, `avatarUrl`, `email`, `phone`, `userRole`) values (20, '贾擎苍', 'h3JkC', 'pHtWD', 'stacey.heidenreich@hotmail.com', '17144627886', 0);

# ALTER TABLE `user` DROP INDEX `phone`;
# ALTER TABLE `user` DROP INDEX `email`;


ALTER TABLE `user` ADD INDEX `phone` (`phone`);
ALTER TABLE `user` ADD INDEX `email` (`email`);
